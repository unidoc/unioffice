// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

func TestCT_RunTrackChangeConstructor(t *testing.T) {
	v := wordprocessingml.NewCT_RunTrackChange()
	if v == nil {
		t.Errorf("wordprocessingml.NewCT_RunTrackChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.CT_RunTrackChange should validate: %s", err)
	}
}

func TestCT_RunTrackChangeMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewCT_RunTrackChange()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewCT_RunTrackChange()
	xml.Unmarshal(buf, v2)
}
