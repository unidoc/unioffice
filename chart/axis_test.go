// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart_test

import (
	"testing"

	"github.com/unidoc/unioffice/chart"
)

func TestNullAxis(t *testing.T) {
	if chart.NullAxis.AxisID() != 0 {
		t.Errorf("expected null axis to have ID 0, go %d", chart.NullAxis.AxisID())
	}
}
