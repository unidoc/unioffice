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

	"baliance.com/gooxml/spreadsheet"
	"baliance.com/gooxml/testhelper"
	"baliance.com/gooxml/zippkg"
)

func TestStyleSheetUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/styles.xml")
	if err != nil {
		t.Fatalf("error reading styles.xml")
	}
	dec := xml.NewDecoder(f)
	r := spreadsheet.NewStyleSheet(nil)
	if err := dec.Decode(r.X()); err != nil {
		t.Errorf("error decoding styles.xml: %s", err)
	}
	got := &bytes.Buffer{}
	fmt.Fprintf(got, zippkg.XMLHeader)
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(r.X()); err != nil {
		t.Errorf("error encoding styles.xml: %s", err)
	}

	testhelper.CompareGoldenXML(t, "styles.xml", got.Bytes())
}

func TestStyleSheetFonts(t *testing.T) {
	ss := spreadsheet.NewStyleSheet(nil)
	fc := len(ss.Fonts())
	ft := ss.AddFont()

	if len(ss.Fonts()) != fc+1 {
		t.Errorf("expected %d fonts, had %d", fc+1, len(ss.Fonts()))
	}
	if err := ss.RemoveFont(ft); err != nil {
		t.Errorf("expected no errors removing font, got %s", err)
	}
	if len(ss.Fonts()) != fc {
		t.Errorf("expected %d fonts, had %d", fc, len(ss.Fonts()))
	}
	if err := ss.RemoveFont(ft); err == nil {
		t.Errorf("expected an errors removing non-existent font, got none")
	}

}
