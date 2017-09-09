// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"testing"

	"baliance.com/gooxml/spreadsheet"
)

func TestComments(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	c := sheet.Comments()

	expRef := "A1"
	expAuth := "John Doe"
	c.AddCommentWithStyle(expRef, expAuth, "This is my comment")
	if c.X().Authors == nil {
		t.Fatalf("author should be non-nil")
	}
	if len(c.X().Authors.Author) != 1 {
		t.Errorf("expected one author, got %v", c.X().Authors.Author)
	}

	auth := c.X().Authors.Author[0]
	if auth != expAuth {
		t.Errorf("expected author = %s, got %s", expAuth, auth)
	}
	if c.X().CommentList == nil {
		t.Fatalf("commentlist should be non-nil")
	}
	if len(c.X().CommentList.Comment) != 1 {
		t.Errorf("expected one comment, got %v", c.X().CommentList.Comment)
	}
	cmt := c.X().CommentList.Comment[0]
	if cmt.AuthorIdAttr != 0 {
		t.Errorf("expected author ID = 0, got %d", cmt.AuthorIdAttr)
	}
	if cmt.RefAttr != expRef {
		t.Errorf("expected ref = %s, got %s", expRef, cmt.RefAttr)
	}
}

func TestCommentsReusesAuthorIDs(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	c := sheet.Comments()

	c.AddCommentWithStyle("A1", "foo", "This is my comment")
	if c.X().Authors == nil {
		t.Fatalf("author should be non-nil")
	}
	if len(c.X().Authors.Author) != 1 {
		t.Errorf("expected one author, got %v", c.X().Authors.Author)
	}

	c.AddCommentWithStyle("B1", "foo", "This is another comment")
	if c.X().Authors == nil {
		t.Fatalf("author should be non-nil")
	}
	if len(c.X().Authors.Author) != 1 {
		t.Errorf("expected one author, got %v", c.X().Authors.Author)
	}

	c.AddCommentWithStyle("C1", "bar", "This is the last comment")
	if c.X().Authors == nil {
		t.Fatalf("author should be non-nil")
	}
	if len(c.X().Authors.Author) != 2 {
		t.Errorf("expected two authors, got %v", c.X().Authors.Author)
	}

}
