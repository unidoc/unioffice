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

func TestAG_SectPrAttributesConstructor(t *testing.T) {
	v := wordprocessingml.NewAG_SectPrAttributes()
	if v == nil {
		t.Errorf("wordprocessingml.NewAG_SectPrAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.AG_SectPrAttributes should validate: %s", err)
	}
}

func TestAG_SectPrAttributesMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewAG_SectPrAttributes()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewAG_SectPrAttributes()
	xml.Unmarshal(buf, v2)
}
