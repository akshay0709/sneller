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

package plan

import (
	"strings"

	"github.com/SnellerInc/sneller/expr"
	"github.com/SnellerInc/sneller/ion"
	"github.com/SnellerInc/sneller/vm"
)

// Unnest joins a row on a list-like field
// within that row and computes a projection
// plus an optional conditional clause
type Unnest struct {
	Nonterminal // source op
	Expr        expr.Node
	Result      string
}

func (u *Unnest) rewrite(rw expr.Rewriter) {
	u.From.rewrite(rw)
	u.Expr = expr.Rewrite(rw, u.Expr)
}

func (u *Unnest) encode(dst *ion.Buffer, st *ion.Symtab) error {
	dst.BeginStruct(-1)
	settype("unnest", dst, st)
	dst.BeginField(st.Intern("expr"))
	u.Expr.Encode(dst, st)
	dst.BeginField(st.Intern("result"))
	dst.WriteString(u.Result)
	dst.EndStruct()
	return nil
}

func decodeSel(dst *vm.Selection, st *ion.Symtab, src []byte) error {
	bind, err := expr.DecodeBindings(st, src)
	if err != nil {
		return err
	}
	*dst = bind
	return nil
}

func (u *Unnest) setfield(d Decoder, name string, st *ion.Symtab, body []byte) error {
	switch name {
	case "result":
		s, _, err := ion.ReadString(body)
		if err != nil {
			return err
		}
		u.Result = s
	case "expr":
		e, _, err := expr.Decode(st, body)
		if err != nil {
			return err
		}
		u.Expr = e
	default:
		return errUnexpectedField
	}
	return nil
}

func (u *Unnest) String() string {
	var out strings.Builder
	out.WriteString("UNNEST ")
	out.WriteString(expr.ToString(u.Expr))
	out.WriteString(" AS ")
	out.WriteString(u.Result)
	return out.String()
}

func (u *Unnest) wrap(dst vm.QuerySink, ep *ExecParams) func(TableHandle) error {
	op, err := vm.NewUnnest(dst, u.Expr, u.Result)
	if err != nil {
		return delay(err)
	}
	return u.From.wrap(op, ep)
}
