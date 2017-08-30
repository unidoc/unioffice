// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml_test

import (
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/presentationml"
)

func TestCT_InOutTransitionConstructor(t *testing.T) {
	v := presentationml.NewCT_InOutTransition()
	if v == nil {
		t.Errorf("presentationml.NewCT_InOutTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed presentationml.CT_InOutTransition should validate: %s", err)
	}
}
