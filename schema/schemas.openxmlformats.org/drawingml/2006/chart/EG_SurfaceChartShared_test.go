// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart_test

import (
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

func TestEG_SurfaceChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_SurfaceChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_SurfaceChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_SurfaceChartShared should validate: %s", err)
	}
}
