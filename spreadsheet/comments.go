// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/schema/soo/sml"
	"baliance.com/gooxml/vmldrawing"
)

// Comments is the container for comments for a single sheet.
type Comments struct {
	w *Workbook
	x *sml.Comments
}

// MakeComments constructs a new Comments wrapper.
func MakeComments(w *Workbook, x *sml.Comments) Comments {
	return Comments{w, x}
}

// X returns the inner wrapped XML type.
func (c Comments) X() *sml.Comments {
	return c.x
}

// Comments returns the list of comments for this sheet
func (c Comments) Comments() []Comment {
	ret := []Comment{}
	for _, cmt := range c.x.CommentList.Comment {
		ret = append(ret, Comment{c.w, cmt, c.x})
	}
	return ret
}

func (c Comments) getOrCreateAuthor(author string) uint32 {
	for i, knownAuthor := range c.x.Authors.Author {
		if knownAuthor == author {
			return uint32(i)
		}
	}

	// didn't find the author, so add a new one
	authIdx := uint32(len(c.x.Authors.Author))
	c.x.Authors.Author = append(c.x.Authors.Author, author)
	return authIdx
}

// AddComment adds a new comment and returns a RichText which will contain the
// styled comment text.
func (c Comments) AddComment(cellRef string, author string) RichText {

	cmt := sml.NewCT_Comment()
	c.x.CommentList.Comment = append(c.x.CommentList.Comment, cmt)
	cmt.RefAttr = cellRef
	cmt.AuthorIdAttr = c.getOrCreateAuthor(author)
	cmt.Text = sml.NewCT_Rst()
	return RichText{cmt.Text}
}

// AddCommentWithStyle adds a new comment styled in a default way
func (c Comments) AddCommentWithStyle(cellRef string, author string, comment string) {
	rt := c.AddComment(cellRef, author)
	run := rt.AddRun()
	run.SetBold(true)
	run.SetSize(10)
	run.SetColor(color.Black)
	run.SetFont("Calibri")
	run.SetText(author + ":")

	run = rt.AddRun()
	run.SetSize(10)
	run.SetFont("Calibri")
	run.SetColor(color.Black)
	run.SetText("\r\n" + comment + "\r\n")

	col, rowIdx, _ := ParseCellReference(cellRef)
	colIdx := ColumnToIndex(col)
	c.w.vmlDrawings[0].Shape = append(c.w.vmlDrawings[0].Shape, vmldrawing.NewCommentShape(int64(colIdx), int64(rowIdx-1)))
}
