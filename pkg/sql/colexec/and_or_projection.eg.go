// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
)

type andProjOp struct {
	allocator *Allocator
	input     Operator

	leftProjOpChain  Operator
	rightProjOpChain Operator
	leftFeedOp       *feedOperator
	rightFeedOp      *feedOperator

	leftIdx   int
	rightIdx  int
	outputIdx int

	// origSel is a buffer used to keep track of the original selection vector of
	// the input batch. We need to do this because we're going to modify the
	// selection vector in order to do the short-circuiting of logical operators.
	origSel []uint16
}

// NewAndProjOp returns a new projection operator that logical-And's
// the boolean columns at leftIdx and rightIdx, returning the result in
// outputIdx.
func NewAndProjOp(
	allocator *Allocator,
	input, leftProjOpChain, rightProjOpChain Operator,
	leftFeedOp, rightFeedOp *feedOperator,
	leftIdx, rightIdx, outputIdx int,
) Operator {
	return &andProjOp{
		allocator:        allocator,
		input:            input,
		leftProjOpChain:  leftProjOpChain,
		rightProjOpChain: rightProjOpChain,
		leftFeedOp:       leftFeedOp,
		rightFeedOp:      rightFeedOp,
		leftIdx:          leftIdx,
		rightIdx:         rightIdx,
		outputIdx:        outputIdx,
		origSel:          make([]uint16, coldata.BatchSize()),
	}
}

func (o *andProjOp) ChildCount(verbose bool) int {
	return 3
}

func (o *andProjOp) Child(nth int, verbose bool) execinfra.OpNode {
	switch nth {
	case 0:
		return o.input
	case 1:
		return o.leftProjOpChain
	case 2:
		return o.rightProjOpChain
	default:
		execerror.VectorizedInternalPanic(fmt.Sprintf("invalid idx %d", nth))
		// This code is unreachable, but the compiler cannot infer that.
		return nil
	}
}

func (o *andProjOp) Init() {
	o.input.Init()
}

// Next is part of the Operator interface.
// The idea to handle the short-circuiting logic is similar to what caseOp
// does: a logical operator has an input and two projection chains. First,
// it runs the left chain on the input batch. Then, it "subtracts" the
// tuples for which we know the result of logical operation based only on
// the left side projection (e.g. if the left side is false and we're
// doing AND operation, then the result is also false) and runs the right
// side projection only on the remaining tuples (i.e. those that were not
// "subtracted"). Next, it restores the original selection vector and
// populates the result of the logical operation.
func (o *andProjOp) Next(ctx context.Context) coldata.Batch {
	batch := o.input.Next(ctx)
	origLen := batch.Length()
	if origLen == 0 {
		return coldata.ZeroBatch
	}
	o.allocator.MaybeAddColumn(batch, coltypes.Bool, o.outputIdx)
	usesSel := false
	if sel := batch.Selection(); sel != nil {
		copy(o.origSel[:origLen], sel[:origLen])
		usesSel = true
	}

	// In order to support the short-circuiting logic, we need to be quite tricky
	// here. First, we set the input batch for the left projection to run and
	// actually run the projection.
	o.leftFeedOp.batch = batch
	batch = o.leftProjOpChain.Next(ctx)

	// Now we need to populate a selection vector on the batch in such a way that
	// those tuples that we already know the result of logical operation for do
	// not get the projection for the right side.
	// knownResult indicates the boolean value which if present on the left side
	// fully determines the result of the logical operation.
	knownResult := false
	leftCol := batch.ColVec(o.leftIdx)
	leftColVals := leftCol.Bool()
	var curIdx uint16
	if usesSel {
		sel := batch.Selection()
		origSel := o.origSel[:origLen]
		if leftCol.MaybeHasNulls() {
			leftNulls := leftCol.Nulls()
			for _, i := range origSel {
				isLeftNull := leftNulls.NullAt(i)
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		} else {
			for _, i := range origSel {
				isLeftNull := false
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		}
	} else {
		batch.SetSelection(true)
		sel := batch.Selection()
		if leftCol.MaybeHasNulls() {
			leftNulls := leftCol.Nulls()
			for i := uint16(0); i < origLen; i++ {
				isLeftNull := leftNulls.NullAt(i)
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		} else {
			for i := uint16(0); i < origLen; i++ {
				isLeftNull := false
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		}
	}

	var ranRightSide bool
	if curIdx > 0 {
		// We only run the right-side projection if there are non-zero number of
		// remaining tuples.
		batch.SetLength(curIdx)
		o.rightFeedOp.batch = batch
		batch = o.rightProjOpChain.Next(ctx)
		ranRightSide = true
	}

	// Now we need to restore the original selection vector and length.
	if usesSel {
		sel := batch.Selection()
		copy(sel[:origLen], o.origSel[:origLen])
	} else {
		batch.SetSelection(false)
	}
	batch.SetLength(origLen)

	var (
		rightCol     coldata.Vec
		rightColVals []bool
	)
	if ranRightSide {
		rightCol = batch.ColVec(o.rightIdx)
		rightColVals = rightCol.Bool()
	}
	outputCol := batch.ColVec(o.outputIdx)
	outputColVals := outputCol.Bool()
	outputNulls := outputCol.Nulls()
	// This is where we populate the output - do the actual evaluation of the
	// logical operation.
	if leftCol.MaybeHasNulls() {
		leftNulls := leftCol.Nulls()
		if rightCol != nil && rightCol.MaybeHasNulls() {
			rightNulls := rightCol.Nulls()
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		}
	} else {
		if rightCol != nil && rightCol.MaybeHasNulls() {
			rightNulls := rightCol.Nulls()
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for AND'ing two booleans are:
						// 1. if at least one of the values is FALSE, then the result is also FALSE
						// 2. if both values are TRUE, then the result is also TRUE
						// 3. in all other cases (one is TRUE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
							// Rule 1: at least one boolean is FALSE.
							outputColVals[idx] = false
						} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
							// Rule 2: both booleans are TRUE.
							outputColVals[idx] = true
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		}
	}

	return batch
}

type orProjOp struct {
	allocator *Allocator
	input     Operator

	leftProjOpChain  Operator
	rightProjOpChain Operator
	leftFeedOp       *feedOperator
	rightFeedOp      *feedOperator

	leftIdx   int
	rightIdx  int
	outputIdx int

	// origSel is a buffer used to keep track of the original selection vector of
	// the input batch. We need to do this because we're going to modify the
	// selection vector in order to do the short-circuiting of logical operators.
	origSel []uint16
}

// NewOrProjOp returns a new projection operator that logical-Or's
// the boolean columns at leftIdx and rightIdx, returning the result in
// outputIdx.
func NewOrProjOp(
	allocator *Allocator,
	input, leftProjOpChain, rightProjOpChain Operator,
	leftFeedOp, rightFeedOp *feedOperator,
	leftIdx, rightIdx, outputIdx int,
) Operator {
	return &orProjOp{
		allocator:        allocator,
		input:            input,
		leftProjOpChain:  leftProjOpChain,
		rightProjOpChain: rightProjOpChain,
		leftFeedOp:       leftFeedOp,
		rightFeedOp:      rightFeedOp,
		leftIdx:          leftIdx,
		rightIdx:         rightIdx,
		outputIdx:        outputIdx,
		origSel:          make([]uint16, coldata.BatchSize()),
	}
}

func (o *orProjOp) ChildCount(verbose bool) int {
	return 3
}

func (o *orProjOp) Child(nth int, verbose bool) execinfra.OpNode {
	switch nth {
	case 0:
		return o.input
	case 1:
		return o.leftProjOpChain
	case 2:
		return o.rightProjOpChain
	default:
		execerror.VectorizedInternalPanic(fmt.Sprintf("invalid idx %d", nth))
		// This code is unreachable, but the compiler cannot infer that.
		return nil
	}
}

func (o *orProjOp) Init() {
	o.input.Init()
}

// Next is part of the Operator interface.
// The idea to handle the short-circuiting logic is similar to what caseOp
// does: a logical operator has an input and two projection chains. First,
// it runs the left chain on the input batch. Then, it "subtracts" the
// tuples for which we know the result of logical operation based only on
// the left side projection (e.g. if the left side is false and we're
// doing AND operation, then the result is also false) and runs the right
// side projection only on the remaining tuples (i.e. those that were not
// "subtracted"). Next, it restores the original selection vector and
// populates the result of the logical operation.
func (o *orProjOp) Next(ctx context.Context) coldata.Batch {
	batch := o.input.Next(ctx)
	origLen := batch.Length()
	if origLen == 0 {
		return coldata.ZeroBatch
	}
	o.allocator.MaybeAddColumn(batch, coltypes.Bool, o.outputIdx)
	usesSel := false
	if sel := batch.Selection(); sel != nil {
		copy(o.origSel[:origLen], sel[:origLen])
		usesSel = true
	}

	// In order to support the short-circuiting logic, we need to be quite tricky
	// here. First, we set the input batch for the left projection to run and
	// actually run the projection.
	o.leftFeedOp.batch = batch
	batch = o.leftProjOpChain.Next(ctx)

	// Now we need to populate a selection vector on the batch in such a way that
	// those tuples that we already know the result of logical operation for do
	// not get the projection for the right side.
	// knownResult indicates the boolean value which if present on the left side
	// fully determines the result of the logical operation.
	knownResult := true
	leftCol := batch.ColVec(o.leftIdx)
	leftColVals := leftCol.Bool()
	var curIdx uint16
	if usesSel {
		sel := batch.Selection()
		origSel := o.origSel[:origLen]
		if leftCol.MaybeHasNulls() {
			leftNulls := leftCol.Nulls()
			for _, i := range origSel {
				isLeftNull := leftNulls.NullAt(i)
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		} else {
			for _, i := range origSel {
				isLeftNull := false
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		}
	} else {
		batch.SetSelection(true)
		sel := batch.Selection()
		if leftCol.MaybeHasNulls() {
			leftNulls := leftCol.Nulls()
			for i := uint16(0); i < origLen; i++ {
				isLeftNull := leftNulls.NullAt(i)
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		} else {
			for i := uint16(0); i < origLen; i++ {
				isLeftNull := false
				if isLeftNull || leftColVals[i] != knownResult {
					// We add the tuple into the selection vector if the left value is NULL or
					// it is different from knownResult.
					sel[curIdx] = i
					curIdx++
				}
			}
		}
	}

	var ranRightSide bool
	if curIdx > 0 {
		// We only run the right-side projection if there are non-zero number of
		// remaining tuples.
		batch.SetLength(curIdx)
		o.rightFeedOp.batch = batch
		batch = o.rightProjOpChain.Next(ctx)
		ranRightSide = true
	}

	// Now we need to restore the original selection vector and length.
	if usesSel {
		sel := batch.Selection()
		copy(sel[:origLen], o.origSel[:origLen])
	} else {
		batch.SetSelection(false)
	}
	batch.SetLength(origLen)

	var (
		rightCol     coldata.Vec
		rightColVals []bool
	)
	if ranRightSide {
		rightCol = batch.ColVec(o.rightIdx)
		rightColVals = rightCol.Bool()
	}
	outputCol := batch.ColVec(o.outputIdx)
	outputColVals := outputCol.Bool()
	outputNulls := outputCol.Nulls()
	// This is where we populate the output - do the actual evaluation of the
	// logical operation.
	if leftCol.MaybeHasNulls() {
		leftNulls := leftCol.Nulls()
		if rightCol != nil && rightCol.MaybeHasNulls() {
			rightNulls := rightCol.Nulls()
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := leftNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		}
	} else {
		if rightCol != nil && rightCol.MaybeHasNulls() {
			rightNulls := rightCol.Nulls()
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := rightNulls.NullAt(idx)
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:origLen] {
					idx := i
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			} else {
				if ranRightSide {
					_ = rightColVals[origLen-1]
				}
				_ = outputColVals[origLen-1]
				for i := range leftColVals[:origLen] {
					idx := uint16(i)
					isLeftNull := false
					leftVal := leftColVals[idx]
					if !isLeftNull && leftVal == knownResult {
						outputColVals[idx] = leftVal
					} else {
						isRightNull := false
						rightVal := rightColVals[idx]
						// The rules for OR'ing two booleans are:
						// 1. if at least one of the values is TRUE, then the result is also TRUE
						// 2. if both values are FALSE, then the result is also FALSE
						// 3. in all other cases (one is FALSE and the other is NULL or both are NULL),
						//    the result is NULL.
						if (leftVal && !isLeftNull) || (rightVal && !isRightNull) {
							// Rule 1: at least one boolean is TRUE.
							outputColVals[idx] = true
						} else if (!leftVal && !isLeftNull) && (!rightVal && !isRightNull) {
							// Rule 2: both booleans are FALSE.
							outputColVals[idx] = false
						} else {
							// Rule 3.
							outputNulls.SetNull(idx)
						}
					}
				}
			}
		}
	}

	return batch
}
