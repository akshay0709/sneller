// Copyright (C) 2022 Sneller, Inc.
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"fmt"
	"unsafe"

	"github.com/SnellerInc/sneller/ion"
	"github.com/SnellerInc/sneller/ion/zion/zll"

	"golang.org/x/exp/slices"
)

type zionState struct {
	shape     zll.Shape
	buckets   zll.Buckets
	blocksize int64
}

type zionConsumer interface {
	symbolize(st *symtab, aux *auxbindings) error
	zionOk(fields []string) bool
	writeZion(state *zionState) error
}

// zionFlattener is a wrapper for rowConsumers
// that do not implement zionConsumer
type zionFlattener struct {
	rowConsumer // inherit writeRows, Close(), next(), etc.
	infields    []string

	// cached structures:
	myaux   auxbindings
	strided []vmref
	params  rowParams
	tape    []ion.Symbol
	empty   []vmref
}

// we only flatten when the number of fields is small;
// otherwise we have to allocate a bunch of space to
// write out all the vmrefs (columns * rows, 8 bytes each)
// which might actually be *larger* than the data we have to copy...
const maxFlatten = 8

func (z *zionFlattener) zionOk(fields []string) bool {
	if len(fields) > 0 && len(fields) < maxFlatten {
		z.infields = append(z.infields[:0], fields...)
		return true
	}
	return false
}

func (z *zionFlattener) symbolize(st *symtab, aux *auxbindings) error {
	if len(aux.bound) != 0 {
		panic("zionFlattener not the top element in the rowConsumer chain?")
	}
	z.tape = z.tape[:0]
	for _, name := range z.infields {
		sym, ok := st.Symbolize(name)
		if !ok {
			continue
		}
		z.tape = append(z.tape, sym)
	}
	slices.Sort(z.tape)

	// we're going to bind auxbound in symbol order
	z.myaux.reset()
	for i := range z.tape {
		z.myaux.push(st.Get(z.tape[i]))
	}
	return z.rowConsumer.symbolize(st, &z.myaux)
}

// zionflatten unpacks the contents of buckets that match 'tape'
// into the corresponding vmref slices
//
// prerequisites:
//   - len(fields) == len(tape)*zionStride
//   - len(shape) > 0
//   - len(tape) > 0
//
//go:noescape
func zionflatten(shape []byte, buckets *zll.Buckets, fields []vmref, tape []ion.Symbol) (in, out int)

const (
	//lint:ignore U1000 used in assembly
	zllBucketPos          = unsafe.Offsetof(zll.Buckets{}.Pos)
	zllBucketDecompressed = unsafe.Offsetof(zll.Buckets{}.Decompressed)

	// we try to process zionStride rows at a time from the shape
	zionStride = 256
)

func empty(src []vmref, n int) []vmref {
	if cap(src) < n {
		return make([]vmref, n)
	}
	for i := range src {
		src[i] = vmref{}
	}
	return src[:n]
}

// convert a writeZion into a writeRows
// by projecting into auxparams
func (z *zionFlattener) writeZion(state *zionState) error {
	if len(z.tape) == 0 {
		// unusual edge-case: none of the matched symbols
		// are part of the symbol table; just count
		// the number of rows and emit empty rows
		n, err := state.shape.Count()
		if err != nil {
			return err
		}
		z.params.auxbound = z.params.auxbound[:0]
		z.empty = empty(z.empty, n)
		return z.writeRows(z.empty, &z.params)
	}

	// force decompression of the buckets we want
	err := state.buckets.SelectSymbols(z.tape)
	if err != nil {
		return err
	}

	// allocate space for up to zionStride rows * columns;
	// each "column" starts at z.strided[column * zionStride:]
	// which simplifies the assembly a bit
	z.strided = sanitizeAux(z.strided, len(z.tape)*zionStride)
	posn := state.buckets.Pos
	// set slice sizes for all the fields
	z.params.auxbound = shrink(z.params.auxbound, len(z.tape))
	pos := state.shape.Start
	for pos < len(state.shape.Bits) {
		in, out := zionflatten(state.shape.Bits[pos:], &state.buckets, z.strided, z.tape)
		pos += in
		if pos > len(state.shape.Bits) {
			panic("read out-of-bounds")
		}
		if out > zionStride {
			panic("write out-of-bounds")
		}
		if out <= 0 || in <= 0 {
			err = fmt.Errorf("couldn't copy out zion data (data corruption?)")
			break
		}
		for i := range z.params.auxbound {
			z.params.auxbound[i] = sanitizeAux(z.strided[i*zionStride:], out)
		}
		// the callee is allowed to clobber this,
		// so it has to be re-zeroed on each iteration
		z.empty = empty(z.empty, out)
		err = z.writeRows(z.empty, &z.params)
		if err != nil {
			break
		}
	}
	state.buckets.Pos = posn // restore bucket positions
	return err
}
