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
	"strings"
	"testing"

	"baliance.com/gooxml"
)

func TestXSDAny(t *testing.T) {
	any := gooxml.XSDAny{}
	anyXml := `<w:settings xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml" xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" xmlns:sl="http://schemas.openxmlformats.org/schemaLibrary/2006/main" mc:Ignorable="w14 w15 w16se">
	<w:hdrShapeDefaults><o:shapedefaults v:ext="edit" spidmax="2049"><o:idmap v:ext="edit" data="1"/></o:shapedefaults></w:hdrShapeDefaults></w:settings>`

	exp := `<m:settings ma:Ignorable="w14 w15 w16se" xmlns:m="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:ma="http://schemas.openxmlformats.org/markup-compatibility" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml">&#xA;&#x9;<m:hdrShapeDefaults><o:shapedefaults v:ext="edit" spidmax="2049"><o:idmap v:ext="edit" data="1"></o:idmap></o:shapedefaults></m:hdrShapeDefaults></m:settings>`
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
