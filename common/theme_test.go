// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package common_test

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/testhelper"
	"github.com/unidoc/unioffice/zippkg"
)

func TestThemeUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/theme.xml")
	if err != nil {
		t.Fatalf("error reading theme file")
	}
	dec := xml.NewDecoder(f)
	ct := common.NewTheme()
	if err := dec.Decode(ct.X()); err != nil {
		t.Errorf("error decoding theme: %s", err)
	}

	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(ct.X()); err != nil {
		t.Errorf("error encoding theme: %s", err)
	}

	testhelper.CompareGoldenXML(t, "theme.xml", got.Bytes())
}
