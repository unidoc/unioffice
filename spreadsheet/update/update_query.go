// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package update

// UpdateToLeft updates reference to point one cell left.
// Update types constants.
const (
	REMOVE_COLUMN byte = iota
	REMOVE_ROW
)

type UpdateQuery struct {
	UpdateType byte
	ColumnIdx uint32
	RowIdx uint32
	SheetToUpdate string
	UpdateCurrentSheet bool
}
