// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package unioffice_test

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/unidoc/unioffice"
)

func TestXSDAny(t *testing.T) {
	any := unioffice.XSDAny{}
	anyXml := `<w:settings xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml" xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" xmlns:sl="http://schemas.openxmlformats.org/schemaLibrary/2006/main" mc:Ignorable="w14 w15 w16se"><w:hdrShapeDefaults><o:shapedefaults v:ext="edit" spidmax="2049"><o:idmap v:ext="edit" data="1"/>foobar</o:shapedefaults></w:hdrShapeDefaults></w:settings>`
	exp := `<w:settings mc:Ignorable="w14 w15 w16se" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml"><w:hdrShapeDefaults><o:shapedefaults v:ext="edit" spidmax="2049"><o:idmap v:ext="edit" data="1"></o:idmap>foobar</o:shapedefaults></w:hdrShapeDefaults></w:settings>`
	dec := xml.NewDecoder(strings.NewReader(anyXml))
	if err := dec.Decode(&any); err != nil {
		t.Errorf("error decoding XSDAny: %s", err)
	}
	buf := bytes.Buffer{}
	enc := xml.NewEncoder(&buf)
	if err := enc.Encode(&any); err != nil {
		t.Errorf("error encoding XSDAny: %s", err)
	}
	if buf.String() != exp {
		t.Errorf("expected %s, got %s", exp, buf.String())
	}
}
