// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

// Package update contains definitions needed for updating references after removing rows/columns.
package update

// UpdateAction is the type for update types constants.
type UpdateAction byte

const (
	// UpdateActionRemoveColumn means updating references after removing a column.
	UpdateActionRemoveColumn UpdateAction = iota
)

// UpdateQuery contains terms of how to update references after removing row/column.
type UpdateQuery struct {
	// UpdateType is one of the update types like UpdateActionRemoveColumn.
	UpdateType UpdateAction

	// ColumnIdx is the index of the column removed.
	ColumnIdx uint32

	// SheetToUpdate contains the name of the sheet on which removing happened.
	SheetToUpdate string

	// UpdateCurrentSheet is true if references without sheet prefix should be updated as well.
	UpdateCurrentSheet bool
}
