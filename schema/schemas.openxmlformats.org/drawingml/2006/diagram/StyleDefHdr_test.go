// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/diagram"
)

func TestStyleDefHdrConstructor(t *testing.T) {
	v := diagram.NewStyleDefHdr()
	if v == nil {
		t.Errorf("diagram.NewStyleDefHdr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.StyleDefHdr should validate: %s", err)
	}
}

func TestStyleDefHdrMarshalUnmarshal(t *testing.T) {
	v := diagram.NewStyleDefHdr()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewStyleDefHdr()
	xml.Unmarshal(buf, v2)
}
