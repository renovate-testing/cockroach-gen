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

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	// HACK: crlfmt removes the "*/}}" comment if it's the last line in the import
	// block. This was picked because it sorts after "pkg/sql/exec/execgen" and
	// has no deps.
	_ "github.com/cockroachdb/cockroach/pkg/util/bufalloc"
)

func (m *memColumn) Append(args SliceArgs) {
	switch args.ColType {
	case coltypes.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol.AppendSlice(fromCol, int(args.DestIdx), int(args.SrcStartIdx), int(args.SrcEndIdx))
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// We need to truncate toCol before appending to it, so in case of Bytes,
			// we append an empty slice.
			toCol.AppendSlice(toCol, int(args.DestIdx), 0, 0)
			// We will be getting all values below to be appended, regardless of
			// whether the value is NULL. It is possible that Bytes' invariant of
			// non-decreasing offsets on the source is currently not maintained, so
			// we explicitly enforce it.
			maxIdx := uint16(0)
			for _, selIdx := range sel {
				if selIdx > maxIdx {
					maxIdx = selIdx
				}
			}
			fromCol.UpdateOffsetsToBeNonDecreasing(uint64(maxIdx + 1))
			for _, selIdx := range sel {
				val := fromCol.Get(int(selIdx))
				toCol.AppendVal(val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			{
				__desiredCap := int(args.DestIdx) + int(args.SrcEndIdx) - int(args.SrcStartIdx)
				if cap(toCol) >= __desiredCap {
					toCol = toCol[:__desiredCap]
				} else {
					__prevCap := cap(toCol)
					__capToAllocate := __desiredCap
					if __capToAllocate < 2*__prevCap {
						__capToAllocate = 2 * __prevCap
					}
					__new_slice := make([]apd.Decimal, __desiredCap, __capToAllocate)
					copy(__new_slice, toCol[:int(args.DestIdx)])
					toCol = __new_slice
				}
				__src_slice := fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]
				__dst_slice := toCol[int(args.DestIdx):]
				for __i := range __src_slice {
					__dst_slice[__i].Set(&__src_slice[__i])
				}
			}
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, apd.Decimal{})
				toCol[len(toCol)-1].Set(&val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Timestamp:
		fromCol := args.Src.Timestamp()
		toCol := m.Timestamp()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
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

	switch args.ColType {
	case coltypes.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		m.nulls.set(args.SliceArgs)
	case coltypes.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol.Get(int(selIdx))
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol.Set(int(selIdx), v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol.Get(int(selIdx))
					toCol.Set(int(selIdx), v)
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol.Get(int(selIdx))
							toCol.Set(i+int(args.DestIdx), v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol.Get(int(selIdx))
					toCol.Set(i+int(args.DestIdx), v)
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol.Get(int(selIdx))
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol.Set(int(selIdx), v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol.Get(int(selIdx))
					toCol.Set(int(selIdx), v)
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol.Get(int(selIdx))
							toCol.Set(i+int(args.DestIdx), v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol.Get(int(selIdx))
					toCol.Set(i+int(args.DestIdx), v)
				}
			}
			return
		}
		// No Sel or Sel64.
		toCol.CopySlice(fromCol, int(args.DestIdx), int(args.SrcStartIdx), int(args.SrcEndIdx))
		m.nulls.set(args.SliceArgs)
	case coltypes.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)].Set(&v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)].Set(&v)
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)].Set(&v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)].Set(&v)
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)].Set(&v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)].Set(&v)
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)].Set(&v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)].Set(&v)
				}
			}
			return
		}
		// No Sel or Sel64.
		{
			__tgt_slice := toCol[int(args.DestIdx):]
			__src_slice := fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]
			for __i := range __src_slice {
				__tgt_slice[__i].Set(&__src_slice[__i])
			}
		}
		m.nulls.set(args.SliceArgs)
	case coltypes.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		m.nulls.set(args.SliceArgs)
	case coltypes.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		m.nulls.set(args.SliceArgs)
	case coltypes.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		m.nulls.set(args.SliceArgs)
	case coltypes.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		m.nulls.set(args.SliceArgs)
	case coltypes.Timestamp:
		fromCol := args.Src.Timestamp()
		toCol := m.Timestamp()
		if args.Sel64 != nil {
			sel := args.Sel64
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull64(uint64(selIdx))
						} else {
							v := fromCol[int(selIdx)]
							m.nulls.UnsetNull64(uint64(selIdx))
							toCol[int(selIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[int(selIdx)] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt64(uint64(selIdx)) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toCol[i+int(args.DestIdx)] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[int(args.SrcStartIdx)+i]
					v := fromCol[int(selIdx)]
					toCol[i+int(args.DestIdx)] = v
				}
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		m.nulls.set(args.SliceArgs)
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
	}
}

func (m *memColumn) Window(colType coltypes.T, start uint64, end uint64) Vec {
	switch colType {
	case coltypes.Bool:
		col := m.Bool()
		return &memColumn{
			t:     colType,
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Bytes:
		col := m.Bytes()
		return &memColumn{
			t:     colType,
			col:   col.Window(int(start), int(end)),
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Decimal:
		col := m.Decimal()
		return &memColumn{
			t:     colType,
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int16:
		col := m.Int16()
		return &memColumn{
			t:     colType,
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int32:
		col := m.Int32()
		return &memColumn{
			t:     colType,
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int64:
		col := m.Int64()
		return &memColumn{
			t:     colType,
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Float64:
		col := m.Float64()
		return &memColumn{
			t:     colType,
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Timestamp:
		col := m.Timestamp()
		return &memColumn{
			t:     colType,
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) PrettyValueAt(colIdx uint16, colType coltypes.T) string {
	if m.nulls.NullAt(colIdx) {
		return "NULL"
	}
	switch colType {
	case coltypes.Bool:
		col := m.Bool()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Bytes:
		col := m.Bytes()
		v := col.Get(int(colIdx))
		return fmt.Sprintf("%v", v)
	case coltypes.Decimal:
		col := m.Decimal()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Int16:
		col := m.Int16()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Int32:
		col := m.Int32()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Int64:
		col := m.Int64()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Float64:
		col := m.Float64()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Timestamp:
		col := m.Timestamp()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

// Helper to set the value in a Vec when the type is unknown.
func SetValueAt(v Vec, elem interface{}, rowIdx uint16, colType coltypes.T) {
	switch colType {
	case coltypes.Bool:
		target := v.Bool()
		newVal := elem.(bool)
		target[int(rowIdx)] = newVal
	case coltypes.Bytes:
		target := v.Bytes()
		newVal := elem.([]byte)
		target.Set(int(rowIdx), newVal)
	case coltypes.Decimal:
		target := v.Decimal()
		newVal := elem.(apd.Decimal)
		target[int(rowIdx)].Set(&newVal)
	case coltypes.Int16:
		target := v.Int16()
		newVal := elem.(int16)
		target[int(rowIdx)] = newVal
	case coltypes.Int32:
		target := v.Int32()
		newVal := elem.(int32)
		target[int(rowIdx)] = newVal
	case coltypes.Int64:
		target := v.Int64()
		newVal := elem.(int64)
		target[int(rowIdx)] = newVal
	case coltypes.Float64:
		target := v.Float64()
		newVal := elem.(float64)
		target[int(rowIdx)] = newVal
	case coltypes.Timestamp:
		target := v.Timestamp()
		newVal := elem.(time.Time)
		target[int(rowIdx)] = newVal
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}
