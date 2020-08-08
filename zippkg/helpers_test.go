// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package zippkg_test

import "testing"
import "github.com/unidoc/unioffice/zippkg"

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
