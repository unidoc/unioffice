// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package update

// Update types constants.
const (
	REMOVE_COLUMN byte = iota
	REMOVE_ROW
)

// UpdateQuery contains terms of how to update references after removing row/column.
// UpdateType is one of the update types like REMOVE_COLUMN or REMOVE_ROW.
// ColumnIdx is the index of the column removed.
// RowIdx is the index of the row removed.
// SheetToUpdate contains the name of the sheet on which removing happened.
// UpdateCurrentSheet is true if references without sheet prefix should be updated as well.
type UpdateQuery struct {
	UpdateType byte
	ColumnIdx uint32
	RowIdx uint32
	SheetToUpdate string
	UpdateCurrentSheet bool
}
