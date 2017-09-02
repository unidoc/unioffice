// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/presentationml"
)

func TestCT_TLAnimVariantStringValConstructor(t *testing.T) {
	v := presentationml.NewCT_TLAnimVariantStringVal()
	if v == nil {
		t.Errorf("presentationml.NewCT_TLAnimVariantStringVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed presentationml.CT_TLAnimVariantStringVal should validate: %s", err)
	}
}

func TestCT_TLAnimVariantStringValMarshalUnmarshal(t *testing.T) {
	v := presentationml.NewCT_TLAnimVariantStringVal()
	buf, _ := xml.Marshal(v)
	v2 := presentationml.NewCT_TLAnimVariantStringVal()
	xml.Unmarshal(buf, v2)
}
