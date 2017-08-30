// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml_test

import (
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

func TestCT_SphereCoordsConstructor(t *testing.T) {
	v := drawingml.NewCT_SphereCoords()
	if v == nil {
		t.Errorf("drawingml.NewCT_SphereCoords must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed drawingml.CT_SphereCoords should validate: %s", err)
	}
}
