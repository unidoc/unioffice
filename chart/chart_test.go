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
	crt "baliance.com/gooxml/schema/soo/dml/chart"
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
func TestDateAxis(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)

	c.AddDateAxis()
	c.AddCategoryAxis()
	c.AddLineChart()
}
func TestArea3DChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddArea3DChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Area3D chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}
	s.SetText("foo")
	if *s.X().Tx.Choice.V != "foo" {
		t.Errorf("expected foo")
	}

	s.CategoryAxis()
	s.Values()

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestAreaChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddAreaChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Area chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}
	s.SetText("foo")
	if *s.X().Tx.Choice.V != "foo" {
		t.Errorf("expected foo")
	}

	s.CategoryAxis()
	s.Values()

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestBar3DChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddBar3DChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Bar3D chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	s.SetText("foo")
	if *s.X().Tx.Choice.V != "foo" {
		t.Errorf("expected foo")
	}

	s.CategoryAxis()
	s.Values()

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestBarChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddBarChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Bar chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestBubbleChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddBubbleChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Bubble chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	s.SetText("foo")
	if *s.X().Tx.Choice.V != "foo" {
		t.Errorf("expected foo")
	}

	s.CategoryAxis()
	s.Values()
	s.BubbleSizes()

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestDoughnutChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddDoughnutChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Doughnut chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}
}

func TestLine3DChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddLine3DChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Line3D chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestLineChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddLineChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Line chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestPieChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddPieChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Pie chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

}
func TestPie3DChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddPie3DChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Pie3D chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}
}

func TestStockChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddStockChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Stock chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestSurfaceChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddSurfaceChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Surface chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}
func TestSurface3DChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddSurface3DChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Surface3D chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestRadarChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddRadarChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Radar chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}

func TestPieOfPieChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddPieOfPieChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil PieOfPie chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}
}

func TestScatterChart(t *testing.T) {
	spc := crt.NewChartSpace()
	c := chart.MakeChart(spc)
	crt := c.AddScatterChart()
	if crt.X() == nil {
		t.Errorf("expected non-nil Scatter chart")
	}

	s := crt.AddSeries()
	if s.X() == nil {
		t.Errorf("expected non-nil series")
	}

	va := c.AddValueAxis()
	crt.AddAxis(va)
}
