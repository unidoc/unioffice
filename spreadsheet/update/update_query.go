// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
