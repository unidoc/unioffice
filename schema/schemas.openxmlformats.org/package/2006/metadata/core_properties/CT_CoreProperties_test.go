// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package core_properties_test

import (
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/metadata/core_properties"
)

func TestCT_CorePropertiesConstructor(t *testing.T) {
	v := core_properties.NewCT_CoreProperties()
	if v == nil {
		t.Errorf("core_properties.NewCT_CoreProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed core_properties.CT_CoreProperties should validate: %s", err)
	}
}
