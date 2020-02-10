// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/update"

type Expression interface {
	Eval(ctx Context, ev Evaluator) Result
	Reference(ctx Context, ev Evaluator) Reference
	String() string
	Update(updateQuery *update.UpdateQuery) Expression
}
