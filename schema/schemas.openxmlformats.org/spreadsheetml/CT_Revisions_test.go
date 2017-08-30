// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml_test

import (
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

func TestCT_RevisionsConstructor(t *testing.T) {
	v := spreadsheetml.NewCT_Revisions()
	if v == nil {
		t.Errorf("spreadsheetml.NewCT_Revisions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetml.CT_Revisions should validate: %s", err)
	}
}
