// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"testing"

	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
	"baliance.com/gooxml/zippkg"
)

func TestRawEncode(t *testing.T) {
	f, err := os.Open("testdata/settings.xml")
	if err != nil {
		t.Fatalf("error reading settings file")
	}
	dec := xml.NewDecoder(f)
	var got *bytes.Buffer
	for i := 0; i < 2; i++ {
		stng := wml.NewSettings()
		if err := dec.Decode(stng); err != nil {
			t.Errorf("error decoding settings: %s", err)
		}
		got = &bytes.Buffer{}
		fmt.Fprintf(got, zippkg.XMLHeader)
		enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
		if err := enc.Encode(stng); err != nil {
			t.Errorf("error encoding settings: %s", err)
		}

		dec = xml.NewDecoder(bytes.NewReader(got.Bytes()))
	}
	xmlStr := got.String()
	beg := strings.Index(xmlStr, "<w:hdrShapeDefaults>")
	end := strings.Index(xmlStr, "</w:hdrShapeDefaults>")

	gotRaw := xmlStr[beg+20 : end]
	exp := `<shapedefaults xmlns="urn:schemas-microsoft-com:office:office" spidmax="2049" xmlns:_="urn:schemas-microsoft-com:vml" _:ext="edit"/>`
	if gotRaw != exp {
		t.Errorf("expected\n%q\ngot\n%q\n", exp, gotRaw)
	}
}
