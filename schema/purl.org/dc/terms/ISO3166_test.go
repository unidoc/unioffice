// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms_test

import (
	"testing"

	"baliance.com/gooxml/schema/purl.org/dc/terms"
)

func TestISO3166Constructor(t *testing.T) {
	v := terms.NewISO3166()
	if v == nil {
		t.Errorf("terms.NewISO3166 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ISO3166 should validate: %s", err)
	}
}
