// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"baliance.com/gooxml/schema/soo/sml"
	"baliance.com/gooxml/testhelper"
	"baliance.com/gooxml/zippkg"
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
