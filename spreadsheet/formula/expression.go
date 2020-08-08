// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/update"

type Expression interface {
	Eval(ctx Context, ev Evaluator) Result
	Reference(ctx Context, ev Evaluator) Reference
	String() string
	Update(updateQuery *update.UpdateQuery) Expression
}
