// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import "baliance.com/gooxml/schema/soo/sml"

// Comment is a single comment within a sheet.
type Comment struct {
	w    *Workbook
	x    *sml.CT_Comment
	cmts *sml.Comments
}

// X returns the inner wrapped XML type.
func (c Comment) X() *sml.CT_Comment {
	return c.x
}

// CellReference returns the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (c Comment) CellReference() string {
	return c.x.RefAttr
}

// SetCellReference sets the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (c Comment) SetCellReference(cellRef string) {
	c.x.RefAttr = cellRef
}

// Author returns the author of the comment
func (c Comment) Author() string {
	if c.x.AuthorIdAttr < uint32(len(c.cmts.Authors.Author)) {
		return c.cmts.Authors.Author[c.x.AuthorIdAttr]
	}
	return ""
}

// SetAuthor sets the author of the comment. If the comment body contains the
// author's name (as is the case with Excel and Comments.AddCommentWithStyle, it
// will not be changed).  This method only changes the metadata author of the
// comment.
func (c Comment) SetAuthor(author string) {
	c.x.AuthorIdAttr = Comments{c.w, c.cmts}.getOrCreateAuthor(author)
}
