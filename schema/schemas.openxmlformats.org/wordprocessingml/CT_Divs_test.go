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

func TestCT_DivsConstructor(t *testing.T) {
	v := wordprocessingml.NewCT_Divs()
	if v == nil {
		t.Errorf("wordprocessingml.NewCT_Divs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.CT_Divs should validate: %s", err)
	}
}

func TestCT_DivsMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewCT_Divs()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewCT_Divs()
	xml.Unmarshal(buf, v2)
}
