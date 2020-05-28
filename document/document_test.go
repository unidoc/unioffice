// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"github.com/unidoc/unioffice/testhelper"
)

func TestSimpleDoc(t *testing.T) {
	doc := document.New()
	testVersion := "00.8000"
	doc.AppProperties.X().AppVersion = &testVersion
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

func TestOpenStrict(t *testing.T) {
	strict, err := document.Open("testdata/strict.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	gotStrict := bytes.Buffer{}
	if err := strict.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	strict.Save(&gotStrict)
	ioutil.WriteFile("testdata/non-strict.docx", gotStrict.Bytes(), 0644)

	// run test assuming that the doc is a valid non-strict doc
	nonStrict, err := document.Open("testdata/non-strict.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	gotNonStrict := bytes.Buffer{}
	if err := nonStrict.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	nonStrict.Save(&gotNonStrict)
	testhelper.CompareZip(t, "non-strict.docx", gotNonStrict.Bytes(), true)

	os.Remove("testdata/non-strict.docx")
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
	png3x3 := []byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
		0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x03,
		0x08, 0x02, 0x00, 0x00, 0x00, 0xd9, 0x4a, 0x22,
		0xe8, 0x00, 0x00, 0x00, 0x1e, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xd7, 0x63, 0xf8, 0xc5, 0x1e, 0xf8,
		0x9d, 0xfd, 0xd7, 0x34, 0xf6, 0x5f, 0x0c, 0x10,
		0x8a, 0x9d, 0xf7, 0x17, 0x03, 0x84, 0x62, 0xf7,
		0xf9, 0x05, 0x00, 0xd2, 0x6f, 0x0d, 0x71, 0x26,
		0x33, 0x2f, 0xe1, 0x00, 0x00, 0x00, 0x00, 0x49,
		0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	}
	img3, err := common.ImageFromBytes(png3x3)
	if err != nil {
		t.Fatalf("unable to create image: %s", err)
	}

	dir1, err := doc.AddImage(img1)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}
	if dir1.RelID() != "rId4" {
		t.Errorf("expected rId4 != %s", dir1.RelID())
	}

	dir2, err := doc.AddImage(img2)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}
	if dir2.RelID() != "rId5" {
		t.Errorf("expected rId5 != %s", dir2.RelID())
	}

	dir3, err := doc.AddImage(img3)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}
	if dir3.RelID() != "rId6" {
		t.Errorf("expected rId6 != %s", dir3.RelID())
	}

	hdr := doc.AddHeader()
	ftr := doc.AddFooter()

	hir1, err := hdr.AddImage(img1)
	if err != nil {
		t.Fatalf("unable to add image to header: %s", err)
	}
	if hir1.RelID() != "rId1" {
		t.Errorf("expected rId1 != %s", hir1.RelID())
	}

	hir2, err := hdr.AddImage(img2)
	if err != nil {
		t.Fatalf("unable to add image to header: %s", err)
	}
	if hir2.RelID() != "rId2" {
		t.Errorf("expected rId2 != %s", hir2.RelID())
	}

	hir3, err := hdr.AddImage(img3)
	if err != nil {
		t.Fatalf("unable to add image to header: %s", err)
	}
	if hir3.RelID() != "rId3" {
		t.Errorf("expected rId3 != %s", hir3.RelID())
	}

	fir1, err := ftr.AddImage(img1)
	if err != nil {
		t.Fatalf("unable to add image to footer: %s", err)
	}
	if fir1.RelID() != "rId1" {
		t.Errorf("expected rId1 != %s", fir1.RelID())
	}

	fir2, err := ftr.AddImage(img2)
	if err != nil {
		t.Fatalf("unable to add image to footer: %s", err)
	}
	if fir2.RelID() != "rId2" {
		t.Errorf("expected rId2 != %s", fir2.RelID())
	}

	fir3, err := ftr.AddImage(img3)
	if err != nil {
		t.Fatalf("unable to add image to footer: %s", err)
	}
	if fir3.RelID() != "rId3" {
		t.Errorf("expected rId3 != %s", fir3.RelID())
	}
}

func TestIssue198(t *testing.T) {
	// this tests the image fixes performed as part of issue 198
	// where we were breaking jpg images
	fn := "issue198.docx"
	doc, err := document.Open("testdata/" + fn)
	if err != nil {
		t.Errorf("error reading %s: %s", fn, err)
		return
	}
	got := bytes.Buffer{}
	doc.Save(&got)
	testhelper.CompareGoldenZip(t, fn+".golden", got.Bytes())
}

func TestGetTables(t *testing.T) {
	doc := document.New()
	table := doc.AddTable()
	tables := doc.Tables()

	if len(tables) != 1 {
		t.Errorf("expected 1 table, got %d", len(tables))
		return
	}

	if table != tables[0] {
		t.Error("retrieved table != added table")
		return
	}

	tbl := document.New().AddTable().X()

	tc := table.AddRow().AddCell().X()
	elts := wml.NewEG_BlockLevelElts()
	tc.EG_BlockLevelElts = append(tc.EG_BlockLevelElts, elts)
	c := wml.NewEG_ContentBlockContent()
	elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, c)
	c.Tbl = append(c.Tbl, tbl)

	tables = doc.Tables()
	if len(tables) < 2 {
		t.Errorf("nested table not enumerated. found %d, expected 2", len(tables))
	}
}

func TestInsertParagraphInTable(t *testing.T) {
	doc := document.New()

	paraBeforeTable := doc.AddParagraph()
	paraBeforeTable.AddRun().AddText("before table")

	table := doc.InsertTableAfter(paraBeforeTable)
	tablePara1 := table.AddRow().AddCell().AddParagraph()
	tablePara1.AddRun().AddText("table paragraph 1")

	paraAfterTable := doc.AddParagraph()
	paraAfterTable.AddRun().AddText("after table")

	tablePara2 := doc.InsertParagraphBefore(tablePara1)
	tablePara2.AddRun().AddText("table paragraph before table paragraph 1")

	tablePara3 := doc.InsertParagraphAfter(tablePara1)
	tablePara3.AddRun().AddText("table paragraph after table paragraph 1")

	bles := doc.X().Body.EG_BlockLevelElts

	if len(bles) != 3 {
		t.Errorf("expected 3 block level elements, got %d", len(bles))
	}
	if len(bles[0].EG_ContentBlockContent[0].P) != 1 {
		t.Errorf("expected 1 paragraph in the first block level element, got %d", len(bles[0].EG_ContentBlockContent[0].P))
	}
	if len(bles[1].EG_ContentBlockContent[0].P) != 0 {
		t.Errorf("expected no paragraphs in the second block level element, got %d", len(bles[1].EG_ContentBlockContent[0].P))
	}
	if len(bles[1].EG_ContentBlockContent[0].Tbl) != 1 {
		t.Errorf("expected 1 table in the second block level element, got %d", len(bles[1].EG_ContentBlockContent[0].Tbl))
	}
	if len(bles[2].EG_ContentBlockContent[0].P) != 1 {
		t.Errorf("expected 1 paragraph in the third block level element, got %d", len(bles[2].EG_ContentBlockContent[0].P))
	}

	cbc := bles[1].EG_ContentBlockContent[0].Tbl[0].EG_ContentRowContent[0].Tr[0].EG_ContentCellContent[0].Tc[0].EG_BlockLevelElts[0].EG_ContentBlockContent[0]
	if len(cbc.P) != 3 {
		t.Errorf("expected 3 paragraphs in the table, got %d", len(cbc.P))
	}
	expected := []string{"table paragraph before table paragraph 1", "table paragraph 1", "table paragraph after table paragraph 1"}
	for i, p := range cbc.P {
		got := p.EG_PContent[0].EG_ContentRunContent[0].R.EG_RunInnerContent[0].T.Content
		if got != expected[i] {
			t.Errorf("expected %s in %d table paragraph, got %s", expected[i], i, got)
		}
	}
}

func TestInsertTableInTable(t *testing.T) {
	doc := document.New()

	paraBeforeTable := doc.AddParagraph()
	paraBeforeTable.AddRun().AddText("before table")

	table := doc.InsertTableAfter(paraBeforeTable)
	row := table.AddRow()
	tablePara1 := row.AddCell().AddParagraph()
	tablePara1.AddRun().AddText("table paragraph 1")

	paraAfterTable := doc.AddParagraph()
	paraAfterTable.AddRun().AddText("after table")

	tableInTable1 := doc.InsertTableBefore(tablePara1)
	ttPara1 := tableInTable1.AddRow().AddCell().AddParagraph()
	ttPara1.AddRun().AddText("table inside table before paragraph 1")

	tableInTable2 := doc.InsertTableAfter(tablePara1)
	ttPara2 := tableInTable2.AddRow().AddCell().AddParagraph()
	ttPara2.AddRun().AddText("table inside table after paragraph 1")

	bles := doc.X().Body.EG_BlockLevelElts

	if len(bles) != 3 {
		t.Errorf("expected 3 block level elements, got %d", len(bles))
	}
	if len(bles[0].EG_ContentBlockContent[0].P) != 1 {
		t.Errorf("expected 1 paragraph in the first block level element, got %d", len(bles[0].EG_ContentBlockContent[0].P))
	}
	if len(bles[1].EG_ContentBlockContent[0].P) != 0 {
		t.Errorf("expected no paragraphs in the second block level element, got %d", len(bles[1].EG_ContentBlockContent[0].P))
	}
	if len(bles[1].EG_ContentBlockContent[0].Tbl) != 1 {
		t.Errorf("expected 1 table in the second block level element, got %d", len(bles[1].EG_ContentBlockContent[0].Tbl))
	}
	if len(bles[2].EG_ContentBlockContent[0].P) != 1 {
		t.Errorf("expected 1 paragraph in the third block level element, got %d", len(bles[2].EG_ContentBlockContent[0].P))
	}

	elts := bles[1].EG_ContentBlockContent[0].Tbl[0].EG_ContentRowContent[0].Tr[0].EG_ContentCellContent[0].Tc[0].EG_BlockLevelElts
	cbc0 := elts[0].EG_ContentBlockContent[0]
	if len(cbc0.Tbl) != 1 {
		t.Errorf("expected 1 table as the first element of table, got %d", len(cbc0.Tbl))
	}
	cbc1 := elts[1].EG_ContentBlockContent[0]
	if len(cbc1.P) != 1 {
		t.Errorf("expected 1 paragraph as the second element of table, got %d", len(cbc1.P))
	}
	cbc2 := elts[2].EG_ContentBlockContent[0]
	if len(cbc2.Tbl) != 1 {
		t.Errorf("expected 1 table as the third element of table, got %d", len(cbc2.Tbl))
	}
	expected := "table inside table after paragraph 1"
	got := cbc2.Tbl[0].EG_ContentRowContent[0].Tr[0].EG_ContentCellContent[0].Tc[0].EG_BlockLevelElts[0].EG_ContentBlockContent[0].P[0].EG_PContent[0].EG_ContentRunContent[0].R.EG_RunInnerContent[0].T.Content
	if expected != got {
		t.Errorf("expected %s in the second inner table paragraph, got %s", expected, got)
	}
}
