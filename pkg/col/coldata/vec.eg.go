// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package coldata

import (
	"fmt"
	"time"

	"github.com/cockroachdb/apd/v2"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
)

func (m *memColumn) Append(args SliceArgs) {
	switch m.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bool()
			toCol := m.Bool()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.BytesFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bytes()
			toCol := m.Bytes()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol.AppendSlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				// We need to truncate toCol before appending to it, so in case of Bytes,
				// we append an empty slice.
				toCol.AppendSlice(toCol, args.DestIdx, 0, 0)
				// We will be getting all values below to be appended, regardless of
				// whether the value is NULL. It is possible that Bytes' invariant of
				// non-decreasing offsets on the source is currently not maintained, so
				// we explicitly enforce it.
				maxIdx := 0
				for _, selIdx := range sel {
					if selIdx > maxIdx {
						maxIdx = selIdx
					}
				}
				fromCol.UpdateOffsetsToBeNonDecreasing(maxIdx + 1)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol.AppendVal(val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.DecimalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Decimal()
			toCol := m.Decimal()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				{
					__desiredCap := args.DestIdx + args.SrcEndIdx - args.SrcStartIdx
					if cap(toCol) >= __desiredCap {
						toCol = toCol[:__desiredCap]
					} else {
						__prevCap := cap(toCol)
						__capToAllocate := __desiredCap
						if __capToAllocate < 2*__prevCap {
							__capToAllocate = 2 * __prevCap
						}
						__new_slice := make([]apd.Decimal, __desiredCap, __capToAllocate)
						copy(__new_slice, toCol[:args.DestIdx])
						toCol = __new_slice
					}
					__src_slice := fromCol[args.SrcStartIdx:args.SrcEndIdx]
					__dst_slice := toCol[args.DestIdx:]
					for __i := range __src_slice {
						__dst_slice[__i].Set(&__src_slice[__i])
					}
				}
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, apd.Decimal{})
					toCol[len(toCol)-1].Set(&val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.IntFamily:
		switch m.t.Width() {
		case 16:
			fromCol := args.Src.Int16()
			toCol := m.Int16()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		case 32:
			fromCol := args.Src.Int32()
			toCol := m.Int32()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		case -1:
		default:
			fromCol := args.Src.Int64()
			toCol := m.Int64()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.FloatFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Float64()
			toCol := m.Float64()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.TimestampTZFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Timestamp()
			toCol := m.Timestamp()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.IntervalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Interval()
			toCol := m.Interval()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol[0:args.DestIdx]
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx) //gcassert:inline
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Datum()
			toCol := m.Datum()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol.AppendSlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Slice(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol.AppendVal(val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", m.t))
	}
}

func (m *memColumn) Copy(args CopySliceArgs) {
	if !args.SelOnDest {
		// We're about to overwrite this entire range, so unset all the nulls.
		m.Nulls().UnsetNullRange(args.DestIdx, args.DestIdx+(args.SrcEndIdx-args.SrcStartIdx))
	}
	// } else {
	// SelOnDest indicates that we're applying the input selection vector as a lens
	// into the output vector as well. We'll set the non-nulls by hand below.
	// }

	switch m.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bool()
			toCol := m.Bool()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx] = v
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx] = v
					}
				}
				return
			}
			// No Sel.
			copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
			m.nulls.set(args.SliceArgs)
		}
	case types.BytesFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bytes()
			toCol := m.Bytes()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol.Set(selIdx, v)
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol.Set(selIdx, v)
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol.Set(i+args.DestIdx, v)
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol.Set(i+args.DestIdx, v)
					}
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args.SliceArgs)
		}
	case types.DecimalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Decimal()
			toCol := m.Decimal()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx].Set(&v)
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx].Set(&v)
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx].Set(&v)
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx].Set(&v)
					}
				}
				return
			}
			// No Sel.
			{
				__tgt_slice := toCol[args.DestIdx:]
				__src_slice := fromCol[args.SrcStartIdx:args.SrcEndIdx]
				for __i := range __src_slice {
					__tgt_slice[__i].Set(&__src_slice[__i])
				}
			}
			m.nulls.set(args.SliceArgs)
		}
	case types.IntFamily:
		switch m.t.Width() {
		case 16:
			fromCol := args.Src.Int16()
			toCol := m.Int16()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx] = v
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx] = v
					}
				}
				return
			}
			// No Sel.
			copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
			m.nulls.set(args.SliceArgs)
		case 32:
			fromCol := args.Src.Int32()
			toCol := m.Int32()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx] = v
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx] = v
					}
				}
				return
			}
			// No Sel.
			copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
			m.nulls.set(args.SliceArgs)
		case -1:
		default:
			fromCol := args.Src.Int64()
			toCol := m.Int64()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx] = v
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx] = v
					}
				}
				return
			}
			// No Sel.
			copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
			m.nulls.set(args.SliceArgs)
		}
	case types.FloatFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Float64()
			toCol := m.Float64()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx] = v
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx] = v
					}
				}
				return
			}
			// No Sel.
			copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
			m.nulls.set(args.SliceArgs)
		}
	case types.TimestampTZFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Timestamp()
			toCol := m.Timestamp()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx] = v
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx] = v
					}
				}
				return
			}
			// No Sel.
			copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
			m.nulls.set(args.SliceArgs)
		}
	case types.IntervalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Interval()
			toCol := m.Interval()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								m.nulls.UnsetNull(selIdx)
								toCol[selIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[selIdx] = v
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx) //gcassert:inline
								toCol[i+args.DestIdx] = v
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx) //gcassert:inline
						toCol[i+args.DestIdx] = v
					}
				}
				return
			}
			// No Sel.
			copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
			m.nulls.set(args.SliceArgs)
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Datum()
			toCol := m.Datum()
			if args.Sel != nil {
				sel := args.Sel
				if args.SelOnDest {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								// Remove an unused warning in some cases.
								_ = i
								m.nulls.SetNull(selIdx)
							} else {
								v := fromCol.Get(selIdx)
								m.nulls.UnsetNull(selIdx)
								toCol.Set(selIdx, v)
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx)
						toCol.Set(selIdx, v)
					}
				} else {
					if args.Src.MaybeHasNulls() {
						nulls := args.Src.Nulls()
						for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
							if nulls.NullAt(selIdx) {
								m.nulls.SetNull(i + args.DestIdx)
							} else {
								v := fromCol.Get(selIdx)
								toCol.Set(i+args.DestIdx, v)
							}
						}
						return
					}
					// No Nulls.
					for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						selIdx := sel[args.SrcStartIdx+i]
						v := fromCol.Get(selIdx)
						toCol.Set(i+args.DestIdx, v)
					}
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args.SliceArgs)
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", m.t))
	}
}

func (m *memColumn) Window(start int, end int) Vec {
	switch m.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Bool()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.BytesFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Bytes()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.DecimalFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Decimal()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.IntFamily:
		switch m.t.Width() {
		case 16:
			col := m.Int16()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		case 32:
			col := m.Int32()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		case -1:
		default:
			col := m.Int64()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.FloatFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Float64()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.TimestampTZFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Timestamp()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.IntervalFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Interval()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col[start:end],
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Datum()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Slice(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	}
	panic(fmt.Sprintf("unhandled type %s", m.t))
}

// SetValueAt is an inefficient helper to set the value in a Vec when the type
// is unknown.
func SetValueAt(v Vec, elem interface{}, rowIdx int) {
	switch t := v.Type(); v.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bool()
			newVal := elem.(bool)
			target[rowIdx] = newVal
		}
	case types.BytesFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bytes()
			newVal := elem.([]byte)
			target.Set(rowIdx, newVal)
		}
	case types.DecimalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Decimal()
			newVal := elem.(apd.Decimal)
			target[rowIdx].Set(&newVal)
		}
	case types.IntFamily:
		switch t.Width() {
		case 16:
			target := v.Int16()
			newVal := elem.(int16)
			target[rowIdx] = newVal
		case 32:
			target := v.Int32()
			newVal := elem.(int32)
			target[rowIdx] = newVal
		case -1:
		default:
			target := v.Int64()
			newVal := elem.(int64)
			target[rowIdx] = newVal
		}
	case types.FloatFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Float64()
			newVal := elem.(float64)
			target[rowIdx] = newVal
		}
	case types.TimestampTZFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Timestamp()
			newVal := elem.(time.Time)
			target[rowIdx] = newVal
		}
	case types.IntervalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Interval()
			newVal := elem.(duration.Duration)
			target[rowIdx] = newVal
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Datum()
			newVal := elem.(interface{})
			target.Set(rowIdx, newVal)
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", t))
	}
}

// GetValueAt is an inefficient helper to get the value in a Vec when the type
// is unknown.
func GetValueAt(v Vec, rowIdx int) interface{} {
	t := v.Type()
	switch v.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bool()
			return target.Get(rowIdx) //gcassert:inline
		}
	case types.BytesFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bytes()
			return target.Get(rowIdx) //gcassert:inline
		}
	case types.DecimalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Decimal()
			return target.Get(rowIdx) //gcassert:inline
		}
	case types.IntFamily:
		switch t.Width() {
		case 16:
			target := v.Int16()
			return target.Get(rowIdx) //gcassert:inline
		case 32:
			target := v.Int32()
			return target.Get(rowIdx) //gcassert:inline
		case -1:
		default:
			target := v.Int64()
			return target.Get(rowIdx) //gcassert:inline
		}
	case types.FloatFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Float64()
			return target.Get(rowIdx) //gcassert:inline
		}
	case types.TimestampTZFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Timestamp()
			return target.Get(rowIdx) //gcassert:inline
		}
	case types.IntervalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Interval()
			return target.Get(rowIdx) //gcassert:inline
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Datum()
			return target.Get(rowIdx)
		}
	}
	panic(fmt.Sprintf("unhandled type %s", t))
}
