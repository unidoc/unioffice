// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

func TestCT_StretchInfoPropertiesConstructor(t *testing.T) {
	v := drawingml.NewCT_StretchInfoProperties()
	if v == nil {
		t.Errorf("drawingml.NewCT_StretchInfoProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed drawingml.CT_StretchInfoProperties should validate: %s", err)
	}
}

func TestCT_StretchInfoPropertiesMarshalUnmarshal(t *testing.T) {
	v := drawingml.NewCT_StretchInfoProperties()
	buf, _ := xml.Marshal(v)
	v2 := drawingml.NewCT_StretchInfoProperties()
	xml.Unmarshal(buf, v2)
}
