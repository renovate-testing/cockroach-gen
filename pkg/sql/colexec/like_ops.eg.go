// Code generated by execgen; DO NOT EDIT.

package colexec

import (
	"bytes"
	"context"
	"regexp"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
)

type selPrefixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p *selPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projPrefixBytesBytesConstOp struct {
	projConstOpBase
	constArg []byte
}

func (p projPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the projection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	if projVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		projVec.Nulls().UnsetNulls()
	}
	projCol := projVec.Bool()
	if vec.Nulls().MaybeHasNulls() {
		colNulls := vec.Nulls()
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = bytes.HasPrefix(arg, p.constArg)
				}
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = bytes.HasPrefix(arg, p.constArg)
				}
			}
		}
		colNullsCopy := colNulls.Copy()
		projVec.SetNulls(&colNullsCopy)
	} else {
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				arg := col.Get(i)
				projCol[i] = bytes.HasPrefix(arg, p.constArg)
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				arg := col.Get(i)
				projCol[i] = bytes.HasPrefix(arg, p.constArg)
			}
		}
	}
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

func (p projPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selSuffixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p *selSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projSuffixBytesBytesConstOp struct {
	projConstOpBase
	constArg []byte
}

func (p projSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the projection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	if projVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		projVec.Nulls().UnsetNulls()
	}
	projCol := projVec.Bool()
	if vec.Nulls().MaybeHasNulls() {
		colNulls := vec.Nulls()
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = bytes.HasSuffix(arg, p.constArg)
				}
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = bytes.HasSuffix(arg, p.constArg)
				}
			}
		}
		colNullsCopy := colNulls.Copy()
		projVec.SetNulls(&colNullsCopy)
	} else {
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				arg := col.Get(i)
				projCol[i] = bytes.HasSuffix(arg, p.constArg)
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				arg := col.Get(i)
				projCol[i] = bytes.HasSuffix(arg, p.constArg)
			}
		}
	}
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

func (p projSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selRegexpBytesBytesConstOp struct {
	selConstOpBase
	constArg *regexp.Regexp
}

func (p *selRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p *selRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}

type projRegexpBytesBytesConstOp struct {
	projConstOpBase
	constArg *regexp.Regexp
}

func (p projRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the projection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	if projVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		projVec.Nulls().UnsetNulls()
	}
	projCol := projVec.Bool()
	if vec.Nulls().MaybeHasNulls() {
		colNulls := vec.Nulls()
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = p.constArg.Match(arg)
				}
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = p.constArg.Match(arg)
				}
			}
		}
		colNullsCopy := colNulls.Copy()
		projVec.SetNulls(&colNullsCopy)
	} else {
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				arg := col.Get(i)
				projCol[i] = p.constArg.Match(arg)
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				arg := col.Get(i)
				projCol[i] = p.constArg.Match(arg)
			}
		}
	}
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

func (p projRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}

type selNotPrefixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selNotPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p *selNotPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projNotPrefixBytesBytesConstOp struct {
	projConstOpBase
	constArg []byte
}

func (p projNotPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the projection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	if projVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		projVec.Nulls().UnsetNulls()
	}
	projCol := projVec.Bool()
	if vec.Nulls().MaybeHasNulls() {
		colNulls := vec.Nulls()
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = !bytes.HasPrefix(arg, p.constArg)
				}
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = !bytes.HasPrefix(arg, p.constArg)
				}
			}
		}
		colNullsCopy := colNulls.Copy()
		projVec.SetNulls(&colNullsCopy)
	} else {
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				arg := col.Get(i)
				projCol[i] = !bytes.HasPrefix(arg, p.constArg)
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				arg := col.Get(i)
				projCol[i] = !bytes.HasPrefix(arg, p.constArg)
			}
		}
	}
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

func (p projNotPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selNotSuffixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selNotSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p *selNotSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projNotSuffixBytesBytesConstOp struct {
	projConstOpBase
	constArg []byte
}

func (p projNotSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the projection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	if projVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		projVec.Nulls().UnsetNulls()
	}
	projCol := projVec.Bool()
	if vec.Nulls().MaybeHasNulls() {
		colNulls := vec.Nulls()
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = !bytes.HasSuffix(arg, p.constArg)
				}
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = !bytes.HasSuffix(arg, p.constArg)
				}
			}
		}
		colNullsCopy := colNulls.Copy()
		projVec.SetNulls(&colNullsCopy)
	} else {
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				arg := col.Get(i)
				projCol[i] = !bytes.HasSuffix(arg, p.constArg)
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				arg := col.Get(i)
				projCol[i] = !bytes.HasSuffix(arg, p.constArg)
			}
		}
	}
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

func (p projNotSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selNotRegexpBytesBytesConstOp struct {
	selConstOpBase
	constArg *regexp.Regexp
}

func (p *selNotRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					isNull := nulls.NullAt(i)
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					isNull := false
					if cmp && !isNull {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p *selNotRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}

type projNotRegexpBytesBytesConstOp struct {
	projConstOpBase
	constArg *regexp.Regexp
}

func (p projNotRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `decimalScratch` local variable of type `decimalOverloadScratch`.
	decimalScratch := p.decimalScratch
	// However, the scratch is not used in all of the projection operators, so
	// we add this to go around "unused" error.
	_ = decimalScratch
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	if projVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		projVec.Nulls().UnsetNulls()
	}
	projCol := projVec.Bool()
	if vec.Nulls().MaybeHasNulls() {
		colNulls := vec.Nulls()
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = !p.constArg.Match(arg)
				}
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				if !colNulls.NullAt(i) {
					// We only want to perform the projection operation if the value is not null.
					arg := col.Get(i)
					projCol[i] = !p.constArg.Match(arg)
				}
			}
		}
		colNullsCopy := colNulls.Copy()
		projVec.SetNulls(&colNullsCopy)
	} else {
		if sel := batch.Selection(); sel != nil {
			sel = sel[:n]
			for _, i := range sel {
				arg := col.Get(i)
				projCol[i] = !p.constArg.Match(arg)
			}
		} else {
			col = col
			_ = 0
			_ = n
			_ = projCol[n-1]
			for i := 0; i < n; i++ {
				arg := col.Get(i)
				projCol[i] = !p.constArg.Match(arg)
			}
		}
	}
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

func (p projNotRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}
