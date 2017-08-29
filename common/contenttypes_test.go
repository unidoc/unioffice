// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common_test

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/testhelper"
	"baliance.com/gooxml/zippkg"
)

func TestContentTypesUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/contenttypes.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	ct := common.NewContentTypes()
	if err := dec.Decode(ct.X()); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}

	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(ct.X()); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "contenttypes.xml", got.Bytes())
}
