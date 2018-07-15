// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"bytes"
	"testing"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/testhelper"
)

func TestSimpleDoc(t *testing.T) {
	doc := document.New()
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("foo")
	got := bytes.Buffer{}
	if err := doc.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	doc.Save(&got)
	testhelper.CompareGoldenZip(t, "simple-1.docx", got.Bytes())
}

func TestOpen(t *testing.T) {
	wb, err := document.Open("testdata/simple-1.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareZip(t, "simple-1.docx", got.Bytes(), true)
}

func TestOpenHeaderFooter(t *testing.T) {
	wb, err := document.Open("testdata/header-footer-multiple.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareGoldenZip(t, "header-footer-multiple.docx", got.Bytes())
}

func TestAddParagraph(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	doc.AddParagraph()
	if len(doc.Paragraphs()) != 1 {
		t.Errorf("expected 1 paragraphs, got %d", len(doc.Paragraphs()))
	}
	doc.AddParagraph()
	if len(doc.Paragraphs()) != 2 {
		t.Errorf("expected 2 paragraphs, got %d", len(doc.Paragraphs()))
	}
}

func TestOpenWord2016(t *testing.T) {
	doc, err := document.Open("../testdata/Office2016/Word-Windows.docx")
	if err != nil {
		t.Errorf("error opening Windows Word 2016 document: %s", err)
	}
	got := bytes.Buffer{}
	if err := doc.Save(&got); err != nil {
		t.Errorf("error saving W216 file: %s", err)
	}
	testhelper.CompareGoldenZipFilesOnly(t, "../../testdata/Office2016/Word-Windows.docx", got.Bytes())
}

func TestInsertParagraph(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	p := doc.AddParagraph()
	before := doc.InsertParagraphBefore(p)
	after := doc.InsertParagraphAfter(p)
	if len(doc.Paragraphs()) != 3 {
		t.Errorf("expected 3 paragraphs, got %d", len(doc.Paragraphs()))
	}
	if doc.Paragraphs()[0].X() != before.X() {
		t.Error("InsertParagraphBefore failed")
	}
	if doc.Paragraphs()[2].X() != after.X() {
		t.Error("InsertParagraphAfter failed")
	}
}

func TestInsertTable(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	p1 := doc.AddParagraph()
	p2 := doc.AddParagraph()
	beforeP1 := doc.InsertTableBefore(p1)
	afterP1 := doc.InsertTableAfter(p1)
	beforeP2 := doc.InsertTableBefore(p2)
	afterP2 := doc.InsertTableAfter(p2)
	if len(doc.Tables()) != 4 {
		t.Errorf("expected 4 tables, got %d", len(doc.Tables()))
	}
	if doc.Tables()[0].X() != beforeP1.X() {
		t.Error("InsertTableBefore 1st paragraph failed")
	}
	if doc.Tables()[1].X() != afterP1.X() {
		t.Error("InsertTableAfter 1st paragraph failed")
	}
	if doc.Tables()[2].X() != beforeP2.X() {
		t.Error("InsertTableBefore 2nd paragraph failed")
	}
	if doc.Tables()[3].X() != afterP2.X() {
		t.Error("InsertTableAfter 2nd paragraph failed")
	}
}

func TestInsertRun(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	p := doc.AddParagraph()
	middle := p.AddRun()
	before := p.InsertRunBefore(middle)
	after := p.InsertRunAfter(middle)
	middle.AddText("middle")
	before.AddText("before")
	after.AddText("after")
	if len(p.Runs()) != 3 {
		t.Errorf("expected 3 runs, got %d", len(p.Runs()))
	}
	if p.Runs()[0].X() != before.X() {
		t.Error("InsertParagraphBefore failed")
	}
	if p.Runs()[2].X() != after.X() {
		t.Error("InsertParagraphAfter failed")
	}

	p.RemoveRun(after)

	if len(p.Runs()) != 2 {
		t.Errorf("expected 2 runs, got %d", len(p.Runs()))
	}
	if p.Runs()[0].X() != before.X() {
		t.Error("InsertParagraphBefore failed")
	}
	p.RemoveRun(before)

	if len(p.Runs()) != 1 {
		t.Errorf("expected 1 runs, got %d", len(p.Runs()))
	}

	if p.Runs()[0].X() != middle.X() {
		t.Errorf("remove failed")
	}
}

func TestInsertBookmarks(t *testing.T) {
	doc := document.New()
	if len(doc.Bookmarks()) != 0 {
		t.Errorf("expected 0 bookmarks, got %d", len(doc.Bookmarks()))
	}

	p := doc.AddParagraph()
	p.AddBookmark("bookmark1")
	p.AddBookmark("bookmark2")

	if len(doc.Bookmarks()) != 2 {
		t.Errorf("expected 2 bookmarks, got %d", len(doc.Bookmarks()))
	}
}

func TestDuplicateBookmarks(t *testing.T) {
	doc := document.New()
	if len(doc.Bookmarks()) != 0 {
		t.Errorf("expected 0 bookmarks, got %d", len(doc.Bookmarks()))
	}

	p := doc.AddParagraph()
	p.AddBookmark("bookmark1")
	p.AddBookmark("bookmark1")

	if len(doc.Bookmarks()) != 2 {
		t.Errorf("expected 2 bookmarks, got %d", len(doc.Bookmarks()))
	}

	if err := doc.Validate(); err == nil {
		t.Errorf("expected error due to duplicate bookmark names")
	}
}

func TestHeaderAndFooterImages(t *testing.T) {
	doc := document.New()
	img1, err := common.ImageFromFile("testdata/gopher.png")
	if err != nil {
		t.Fatalf("unable to create image: %s", err)
	}
	img2, err := common.ImageFromFile("testdata/gophercolor.png")
	if err != nil {
		t.Fatalf("unable to create image: %s", err)
	}
	dir1, err := doc.AddImage(img1)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}
	dir2, err := doc.AddImage(img2)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}

	if dir1.RelID() != "rId4" {
		t.Errorf("expected rId4 != %s", dir1.RelID())
	}
	if dir2.RelID() != "rId5" {
		t.Errorf("expected rId5 != %s", dir2.RelID())
	}

	hdr := doc.AddHeader()
	ftr := doc.AddFooter()
	hir1, err := hdr.AddImage(img1)
	fir1, err := ftr.AddImage(img1)
	hir2, err := hdr.AddImage(img2)
	fir2, err := ftr.AddImage(img2)
	if hir1.RelID() != "rId1" {
		t.Errorf("expected rId1 != %s", hir1.RelID())
	}
	if hir2.RelID() != "rId2" {
		t.Errorf("expected rId2 != %s", hir2.RelID())
	}
	if fir1.RelID() != "rId1" {
		t.Errorf("expected rId1 != %s", hir1.RelID())
	}
	if fir2.RelID() != "rId2" {
		t.Errorf("expected rId2 != %s", hir2.RelID())
	}
}
