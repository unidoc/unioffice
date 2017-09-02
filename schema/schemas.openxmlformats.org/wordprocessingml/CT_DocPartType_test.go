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

func TestCT_DocPartTypeConstructor(t *testing.T) {
	v := wordprocessingml.NewCT_DocPartType()
	if v == nil {
		t.Errorf("wordprocessingml.NewCT_DocPartType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.CT_DocPartType should validate: %s", err)
	}
}

func TestCT_DocPartTypeMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewCT_DocPartType()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewCT_DocPartType()
	xml.Unmarshal(buf, v2)
}
