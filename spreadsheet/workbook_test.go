// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.
package spreadsheet_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/sml"

	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/testhelper"
	"github.com/unidoc/unioffice/zippkg"
)

func TestSimpleWorkbook(t *testing.T) {
	wb := spreadsheet.New()
	defer wb.Close()
	testVersion := "00.8000"
	wb.AppProperties.X().AppVersion = &testVersion
	sheet := wb.AddSheet()
	sheet.Cell("A1").SetString("Hello World!")

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid sheet: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareGoldenZip(t, "simple-2.xlsx", got.Bytes())
}

func TestConstructor(t *testing.T) {
	wb := spreadsheet.New()
	defer wb.Close()
	if wb == nil {
		t.Errorf("expected a non-nil workbook")
	}
	if wb.X() == nil {
		t.Errorf("expected a non-nil workbook XML entity")
	}
	if wb.ContentTypes.X() == nil {
		t.Errorf("expected a non-nil content-types")
	}
	if wb.CoreProperties.X() == nil {
		t.Errorf("expected a non-nil core properties")
	}
	if wb.AppProperties.X() == nil {
		t.Errorf("expected a non-nil app properties")
	}
	if wb.Rels.X() == nil {
		t.Errorf("expected a non-nil relationship")
	}
	if wb.StyleSheet.X() == nil {
		t.Errorf("expected a non-nil stylesheet")
	}
}
func TestWorkbookUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/workbook.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	r := sml.NewWorkbook()
	if err := dec.Decode(r); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}
	got := &bytes.Buffer{}
	fmt.Fprintf(got, zippkg.XMLHeader)
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(r); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "workbook.xml", got.Bytes())
}

func TestSimpleSheet(t *testing.T) {
	wb := spreadsheet.New()
	defer wb.Close()
	testVersion := "00.8000"
	wb.AppProperties.X().AppVersion = &testVersion
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetString("testing 123")
	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareGoldenZip(t, "simple-1.xlsx", got.Bytes())
}
func TestOpen(t *testing.T) {
	wb, err := spreadsheet.Open("testdata/simple-1.xlsx")
	if err != nil {
		t.Errorf("error opening workbook: %s", err)
	}
	defer wb.Close()

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareZip(t, "simple-1.xlsx", got.Bytes(), true)
}

func TestOpenExcel2016(t *testing.T) {
	wb, err := spreadsheet.Open("../testdata/Office2016/Excel-Windows.xlsx")
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}
	defer wb.Close()

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	wb.Save(&got)
	//testhelper.CompareZip(t, "../../testdata/Office2016/Excel-Windows.xlsx", got.Bytes())
}

func TestSheetCount(t *testing.T) {
	wb := spreadsheet.New()
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	defer wb.Close()
	if wb.SheetCount() != 0 {
		t.Errorf("expected 0 sheets, got %d", wb.SheetCount())
	}
	wb.AddSheet()
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	if wb.SheetCount() != 1 {
		t.Errorf("expected 1 sheets, got %d", wb.SheetCount())
	}
	wb.AddSheet()
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	if wb.SheetCount() != 2 {
		t.Errorf("expected 2 sheets, got %d", wb.SheetCount())
	}
}

func TestPreserveSpace(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	row := sheet.AddRow()
	values := []string{"  foo  ", " bar \t", "foo\r\nbar", "\t\r\nfoo\t123\r\n"}
	for i, s := range values {
		cell := row.AddCell()
		if i%2 == 0 {
			cell.SetInlineString(s)
		} else {
			cell.SetString(s)
		}

	}

	buf := bytes.Buffer{}
	if err := ss.Save(&buf); err != nil {
		log.Fatalf("error saving spreadsheet: %s", err)
	}

	op, err := spreadsheet.Read(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		log.Fatalf("error loading saved spreadsheet: %s", err)
	}
	sheets := op.Sheets()
	if len(sheets) != 1 {
		t.Errorf("expected 1 sheet, got %d", len(sheets))
	}
	rows := sheets[0].Rows()
	if len(rows) != 1 {
		t.Errorf("expected 1 row, got %d", len(rows))
	}
	cells := rows[0].Cells()
	if len(cells) != 4 {
		t.Errorf("expected 4 cells, got %d", len(cells))
	}
	for i, c := range cells {
		got, err := c.GetRawValue()
		if err != nil {
			t.Errorf("error retrieving cell value: %s", err)
		}
		if got != values[i] {
			t.Errorf("expected cell value = %q, got %q", values[i], got)
		}
	}
}

func TestAddDefinedName(t *testing.T) {
	wb := spreadsheet.New()
	defer wb.Close()
	if len(wb.DefinedNames()) != 0 {
		t.Errorf("expeced no defined names on new wb")
	}
	name := "foo"
	ref := "bar"
	dn := wb.AddDefinedName(name, ref)
	if len(wb.DefinedNames()) != 1 {
		t.Errorf("expected 1 defined names on wb")
	}
	if dn.Name() != name {
		t.Errorf("expected name = %s, got %s", name, dn.Name())
	}
	if dn.Content() != ref {
		t.Errorf("expected content = %s, got %s", ref, dn.Content())
	}
}

func ExampleWorkbook_AddDefinedName() {
	wb := spreadsheet.New()
	defer wb.Close()
	sheet := wb.AddSheet()
	productNames := wb.AddDefinedName("ProductNames", sheet.RangeReference("A2:A6"))
	// now 'ProductNames' can be used in formulas, charts, etc.
	fmt.Printf("%s refers to %s", productNames.Name(), productNames.Content())
	// Output: ProductNames refers to 'Sheet 1'!$A$2:$A$6
}

func TestOpenComments(t *testing.T) {
	wb, err := spreadsheet.Open("./testdata/comments.xlsx")
	defer wb.Close()
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}

	sheet := wb.Sheets()[0]
	if len(sheet.Comments().Comments()) != 1 {
		t.Fatalf("sheet should have returned 1 existing comments")
	}
	cmt := sheet.Comments().Comments()[0]
	if cmt.Author() != "John Doe" {
		t.Errorf("error reading comment author")
	}
}
func TestWorkbookProtection(t *testing.T) {
	wb := spreadsheet.New()
	defer wb.Close()
	if wb.X().WorkbookProtection != nil {
		t.Errorf("expected no protection for new workbook")
	}
	wb.Protection()
	if wb.X().WorkbookProtection == nil {
		t.Errorf("expected protection")
	}
	wb.ClearProtection()
	if wb.X().WorkbookProtection != nil {
		t.Errorf("expected no protection after clear")
	}
}

func TestSheetGetName(t *testing.T) {
	wb := spreadsheet.New()
	defer wb.Close()
	s := wb.AddSheet()
	if _, err := wb.GetSheet("foo"); err == nil {
		t.Errorf("expected an error")
	}
	if _, err := wb.GetSheet(s.Name()); err != nil {
		t.Errorf("expected no error")
	}
}

// TestOpenOrderedSheets test for issue #154 where sheet title didn't match sheet content.
func TestOpenOrderedSheets(t *testing.T) {
	wb, err := spreadsheet.Open("./testdata/ordered-sheets.xlsx")
	defer wb.Close()
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}

	for i, sheet := range wb.Sheets() {
		expContent := fmt.Sprintf("%d", i+1)
		if sheet.Name() != expContent {
			t.Errorf("expected sheet %d name = %s, got %s", i, expContent, sheet.Name())
		}
		got := sheet.Cell("A1").GetFormattedValue()
		if got != expContent {
			t.Errorf("expected sheet %d cell A1 = %s, got %s", i, expContent, got)
		}
	}

}

func TestRemoveSheet(t *testing.T) {
	wb, err := spreadsheet.Open("./testdata/sheets.xlsx")
	defer wb.Close()
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}

	wasCount := wb.SheetCount()

	if err := wb.RemoveSheet(15); err == nil {
		t.Fatalf("invalid sheet index, expected error %v, got nil", spreadsheet.ErrorNotFound)
	}

	if err := wb.RemoveSheet(2); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := wb.Validate(); err != nil {
		t.Fatalf("produced invalid workbook: %v", err)
	}

	if wb.SheetCount() != (wasCount - 1) {
		t.Fatalf("expected sheets count %d, got %d", wasCount-1, wb.SheetCount())
	}
}

func TestRemoveSheetByName(t *testing.T) {
	wb, err := spreadsheet.Open("./testdata/sheets.xlsx")
	defer wb.Close()
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}

	wasCount := wb.SheetCount()

	if err := wb.RemoveSheetByName("Sheet156"); err == nil {
		t.Fatalf("invalid sheet name, expected error %v, got nil", spreadsheet.ErrorNotFound)
	}

	if err := wb.RemoveSheetByName("Sheet2"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := wb.Validate(); err != nil {
		t.Fatalf("produced invalid workbook: %v", err)
	}

	if wb.SheetCount() != (wasCount - 1) {
		t.Fatalf("expected sheets count %d, got %d", wasCount-1, wb.SheetCount())
	}
}

func TestCopySheet(t *testing.T) {
	wb, err := spreadsheet.Open("./testdata/sheets.xlsx")
	defer wb.Close()
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}

	wasCount := wb.SheetCount()

	if _, err := wb.CopySheet(15, "Copied Sheet"); err == nil {
		t.Fatalf("invalid sheet index, expected error %v, got nil", spreadsheet.ErrorNotFound)
	}

	copiedSheet, err := wb.CopySheet(1, "Copied Sheet")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := wb.Validate(); err != nil {
		t.Fatalf("produced invalid workbook: %v", err)
	}

	if copiedSheet.Name() != "Copied Sheet" {
		t.Fatalf("invalid name in the copied sheet, expected \"Copied Sheet\", got \"%s\"", copiedSheet.Name())
	}

	if wb.SheetCount() != (wasCount + 1) {
		t.Fatalf("expected sheets count %d, got %d", wasCount+1, wb.SheetCount())
	}
}

func TestCopySheetByName(t *testing.T) {
	wb, err := spreadsheet.Open("./testdata/sheets.xlsx")
	defer wb.Close()
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}

	wasCount := wb.SheetCount()

	if _, err := wb.CopySheetByName("Sheet156", "Copied Sheet"); err == nil {
		t.Fatalf("invalid sheet name, expected error %v, got nil", spreadsheet.ErrorNotFound)
	}

	copiedSheet, err := wb.CopySheetByName("Sheet2", "Copied Sheet")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := wb.Validate(); err != nil {
		t.Fatalf("produced invalid workbook: %v", err)
	}

	if copiedSheet.Name() != "Copied Sheet" {
		t.Fatalf("invalid name in the copied sheet, expected \"Copied Sheet\", got \"%s\"", copiedSheet.Name())
	}

	if wb.SheetCount() != (wasCount + 1) {
		t.Fatalf("expected sheets count %d, got %d", wasCount+1, wb.SheetCount())
	}
}
