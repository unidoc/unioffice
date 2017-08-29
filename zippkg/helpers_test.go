// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package zippkg_test

import "testing"
import "baliance.com/gooxml/zippkg"

func TestRelsPathFor(t *testing.T) {
	td := []struct {
		Inp string
		Exp string
	}{{"/", "/_rels/.rels"},
		{"/xl/workbook.xml", "/xl/_rels/workbook.xml.rels"}}
	for _, tc := range td {
		if got := zippkg.RelationsPathFor(tc.Inp); got != tc.Exp {
			t.Errorf("expected RelsPathFor(%s) = %s, got %s", tc.Inp, tc.Exp, got)
		}
	}
}
