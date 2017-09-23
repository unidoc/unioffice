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

	"baliance.com/gooxml"

	"baliance.com/gooxml/schema/soo/wml"
	"baliance.com/gooxml/zippkg"
)

func TestCreatorUnknownType(t *testing.T) {
	el, err := gooxml.CreateElement(xml.StartElement{Name: xml.Name{Local: "foo", Space: "bar"}})
	if el == nil || err != nil {
		t.Errorf("CreateElement should never return nil: %s", err)
	}
	if _, ok := el.(*gooxml.XSDAny); !ok {
		t.Errorf("CreateElement should return XSDAny for unknown types")
	}
}

func TestCreatorKnownType(t *testing.T) {
	el, err := gooxml.CreateElement(xml.StartElement{Name: xml.Name{Local: "CT_Settings", Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"}})
	if el == nil || err != nil {
		t.Errorf("CreateElement should never return nil: %s", err)
	}
	if _, ok := el.(*wml.CT_Settings); !ok {
		t.Errorf("CreateElement should return the element requested, got %T", el)
	}
}
func TestRawEncode(t *testing.T) {
	f, err := os.Open("testdata/settings.xml")
	if err != nil {
		t.Fatalf("error reading settings file")
	}
	dec := xml.NewDecoder(f)
	var got *bytes.Buffer

	// should round trip multiple times with no changes after
	// the first encoding
	for i := 0; i < 5; i++ {
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
	beg := strings.LastIndex(xmlStr, "<w:hdrShapeDefaults>")
	end := strings.LastIndex(xmlStr, "</w:hdrShapeDefaults>")

	gotRaw := xmlStr[beg+20 : end]
	exp := "<o:shapedefaults xmlns=\"urn:schemas-microsoft-com:office:office\" xmlns:o=\"urn:schemas-microsoft-com:office:office\" xmlns:r=\"http://schemas.openxmlformats.org/officeDocument/2006/relationships\" xmlns:s=\"http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes\" xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:xml=\"http://www.w3.org/XML/1998/namespace\" spidmax=\"2049\" ext=\"edit\"/>"
	if gotRaw != exp {
		t.Errorf("expected\n%q\ngot\n%q\n", exp, gotRaw)
	}

}
