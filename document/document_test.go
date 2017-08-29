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
	testhelper.CompareZip(t, "simple-1.docx", got.Bytes())
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
