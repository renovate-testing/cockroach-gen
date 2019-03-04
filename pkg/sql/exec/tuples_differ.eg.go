// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package exec

import (
	"bytes"

	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/pkg/errors"
)

// tuplesDiffer takes in two ColVecs as well as tuple indices to check whether
// the tuples differ.
func tuplesDiffer(
	t types.T, aColVec ColVec, aTupleIdx int, bColVec ColVec, bTupleIdx int, differ *bool,
) error {
	switch t {
	case types.Bool:
		aCol := aColVec.Bool()
		bCol := bColVec.Bool()
		var unique bool
		unique = aCol[aTupleIdx] != bCol[bTupleIdx]
		*differ = *differ || unique
		return nil
	case types.Bytes:
		aCol := aColVec.Bytes()
		bCol := bColVec.Bytes()
		var unique bool
		unique = !bytes.Equal(aCol[aTupleIdx], bCol[bTupleIdx])
		*differ = *differ || unique
		return nil
	case types.Decimal:
		aCol := aColVec.Decimal()
		bCol := bColVec.Decimal()
		var unique bool
		unique = tree.CompareDecimals(&aCol[aTupleIdx], &bCol[bTupleIdx]) != 0
		*differ = *differ || unique
		return nil
	case types.Int8:
		aCol := aColVec.Int8()
		bCol := bColVec.Int8()
		var unique bool
		unique = aCol[aTupleIdx] != bCol[bTupleIdx]
		*differ = *differ || unique
		return nil
	case types.Int16:
		aCol := aColVec.Int16()
		bCol := bColVec.Int16()
		var unique bool
		unique = aCol[aTupleIdx] != bCol[bTupleIdx]
		*differ = *differ || unique
		return nil
	case types.Int32:
		aCol := aColVec.Int32()
		bCol := bColVec.Int32()
		var unique bool
		unique = aCol[aTupleIdx] != bCol[bTupleIdx]
		*differ = *differ || unique
		return nil
	case types.Int64:
		aCol := aColVec.Int64()
		bCol := bColVec.Int64()
		var unique bool
		unique = aCol[aTupleIdx] != bCol[bTupleIdx]
		*differ = *differ || unique
		return nil
	case types.Float32:
		aCol := aColVec.Float32()
		bCol := bColVec.Float32()
		var unique bool
		unique = aCol[aTupleIdx] != bCol[bTupleIdx]
		*differ = *differ || unique
		return nil
	case types.Float64:
		aCol := aColVec.Float64()
		bCol := bColVec.Float64()
		var unique bool
		unique = aCol[aTupleIdx] != bCol[bTupleIdx]
		*differ = *differ || unique
		return nil
	default:
		return errors.Errorf("unsupported tuplesDiffer type %s", t)
	}
}
