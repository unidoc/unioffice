// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart_test

import (
	"testing"

	"baliance.com/gooxml/chart"
	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

func TestTitle(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)

	if c.X().Chart.Title != nil {
		t.Errorf("initial title should be nil")
	}
	if c.X().Chart.AutoTitleDeleted != nil {
		t.Errorf("initial title deleted should be nil")
	}

	title := c.AddTitle()
	if c.X().Chart.Title == nil {
		t.Errorf("initial title should not be nil")
	}
	if c.X().Chart.AutoTitleDeleted == nil {
		t.Errorf("initial AutoTitleDeleted should not be nil")
	}
	if c.X().Chart.AutoTitleDeleted.ValAttr == nil || *c.X().Chart.AutoTitleDeleted.ValAttr {
		t.Errorf("AutoTitleDeleted must be false, was %v", c.X().Chart.AutoTitleDeleted.ValAttr)
	}
	title.SetText("testing")
	if c.X().Chart.Title.Tx.Choice.Rich.P[0].EG_TextRun[0].R.T != "testing" {
		t.Errorf("expected text = testing, got %s", c.X().Chart.Title.Tx.Choice.Rich.P[0].EG_TextRun[0].R.T)
	}

	c.RemoveTitle()

	if c.X().Chart.Title != nil {
		t.Errorf("after remove, title should be nil")
	}
	if !*c.X().Chart.AutoTitleDeleted.ValAttr {
		t.Errorf("after remove, title deleted should be true")
	}
}
