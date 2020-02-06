// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/reference"

func moveColumnLeft(column string, colIdxToRemove uint32) string {
	colIdx := reference.ColumnToIndex(column)
	if colIdx == colIdxToRemove {
		return "#REF!"
	} else if colIdx > colIdxToRemove {
		return reference.IndexToColumn(colIdx - 1)
	} else {
		return column
	}
}
