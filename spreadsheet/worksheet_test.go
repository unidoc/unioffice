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
	"os"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/sml"
	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/testhelper"
	"github.com/unidoc/unioffice/zippkg"
)

func TestWorksheetUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/worksheet.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	r := sml.NewWorksheet()
	if err := dec.Decode(r); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}
	got := &bytes.Buffer{}
	fmt.Fprintf(got, zippkg.XMLHeader)
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(r); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "worksheet.xml", got.Bytes())
}

// Issue #212
func TestInsertMergedCells(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	sheet.AddMergedCells("A1", "C1")
	sheet.AddMergedCells("A2", "C2")
	sheet.AddMergedCells("A3", "C3")
	sheet.AddMergedCells("D1", "E3")
	sheet.InsertRow(2)

	// should go down a line
	for i, exp := range []string{
		"A1:C1", // before inserted row, no change
		"A3:C3", // after inserted row, moved down
		"A4:C4", // after inserted row, moved down
		"D1:E4", // covers inserted row, expanded
	} {
		got := sheet.MergedCells()[i].Reference()

		if got != exp {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}
}
