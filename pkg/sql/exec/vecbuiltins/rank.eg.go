// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package vecbuiltins

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/exec"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/execerror"
)

type rankInitFields struct {
	exec.OneInputNode
	// distinctCol is the output column of the chain of ordered distinct
	// operators in which true will indicate that a new rank needs to be assigned
	// to the corresponding tuple.
	distinctCol     []bool
	outputColIdx    int
	partitionColIdx int
}

type rankNoPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &rankNoPartitionOp{}

func (r *rankNoPartitionOp) Init() {
	r.Input().Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankNoPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(coltypes.Int64)
	} else if r.outputColIdx > batch.Width() {
		execerror.VectorizedInternalPanic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}

	if batch.Length() == 0 {
		return batch
	}

	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[sel[i]] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank
				r.rankIncrement++
			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
		}
	}
	return batch
}

type rankWithPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &rankWithPartitionOp{}

func (r *rankWithPartitionOp) Init() {
	r.Input().Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankWithPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	if r.partitionColIdx == batch.Width() {
		batch.AppendCol(coltypes.Bool)
	} else if r.partitionColIdx > batch.Width() {
		execerror.VectorizedInternalPanic("unexpected: column partitionColIdx is neither present nor the next to be appended")
	}
	partitionCol := batch.ColVec(r.partitionColIdx).Bool()

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(coltypes.Int64)
	} else if r.outputColIdx > batch.Width() {
		execerror.VectorizedInternalPanic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}

	if batch.Length() == 0 {
		return batch
	}

	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[sel[i]] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[sel[i]] = 1
				continue
			}
			if r.distinctCol[sel[i]] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank
				r.rankIncrement++
			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[i] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[i] = 1
				continue
			}
			if r.distinctCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
		}
	}
	return batch
}

type denseRankNoPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &denseRankNoPartitionOp{}

func (r *denseRankNoPartitionOp) Init() {
	r.Input().Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *denseRankNoPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(coltypes.Int64)
	} else if r.outputColIdx > batch.Width() {
		execerror.VectorizedInternalPanic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}

	if batch.Length() == 0 {
		return batch
	}

	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[sel[i]] {
				r.rank++
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank

			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
		}
	}
	return batch
}

type denseRankWithPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &denseRankWithPartitionOp{}

func (r *denseRankWithPartitionOp) Init() {
	r.Input().Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *denseRankWithPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	if r.partitionColIdx == batch.Width() {
		batch.AppendCol(coltypes.Bool)
	} else if r.partitionColIdx > batch.Width() {
		execerror.VectorizedInternalPanic("unexpected: column partitionColIdx is neither present nor the next to be appended")
	}
	partitionCol := batch.ColVec(r.partitionColIdx).Bool()

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(coltypes.Int64)
	} else if r.outputColIdx > batch.Width() {
		execerror.VectorizedInternalPanic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}

	if batch.Length() == 0 {
		return batch
	}

	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[sel[i]] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[sel[i]] = 1
				continue
			}
			if r.distinctCol[sel[i]] {
				r.rank++
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank

			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[i] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[i] = 1
				continue
			}
			if r.distinctCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
		}
	}
	return batch
}
