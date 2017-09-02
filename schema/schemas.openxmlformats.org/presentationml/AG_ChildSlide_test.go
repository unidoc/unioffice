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

func TestAG_ChildSlideConstructor(t *testing.T) {
	v := presentationml.NewAG_ChildSlide()
	if v == nil {
		t.Errorf("presentationml.NewAG_ChildSlide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed presentationml.AG_ChildSlide should validate: %s", err)
	}
}

func TestAG_ChildSlideMarshalUnmarshal(t *testing.T) {
	v := presentationml.NewAG_ChildSlide()
	buf, _ := xml.Marshal(v)
	v2 := presentationml.NewAG_ChildSlide()
	xml.Unmarshal(buf, v2)
}
