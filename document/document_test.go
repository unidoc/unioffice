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
