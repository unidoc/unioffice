// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/reference"

// updateColumnToLeft gets a column reference string representation like JJ, parses it and makes a string representation of a new reference with respect to the update type in the case of a column to the left of this reference was removed (e.g. JI).
func updateColumnToLeft(column string, colIdxToRemove uint32) string {
	colIdx := reference.ColumnToIndex(column)
	if colIdx == colIdxToRemove {
		return "#REF!"
	} else if colIdx > colIdxToRemove {
		return reference.IndexToColumn(colIdx - 1)
	} else {
		return column
	}
}
