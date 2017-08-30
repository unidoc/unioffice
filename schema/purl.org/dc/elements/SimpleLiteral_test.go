// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements_test

import (
	"testing"

	"baliance.com/gooxml/schema/purl.org/dc/elements"
)

func TestSimpleLiteralConstructor(t *testing.T) {
	v := elements.NewSimpleLiteral()
	if v == nil {
		t.Errorf("elements.NewSimpleLiteral must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.SimpleLiteral should validate: %s", err)
	}
}
