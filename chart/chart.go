// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"math/rand"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/drawing"

	"github.com/unidoc/unioffice/schema/soo/dml"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

// Chart is a generic chart.
type Chart struct {
	x *crt.ChartSpace
}

func MakeChart(x *crt.ChartSpace) Chart {
	return Chart{x}
}

// X returns the inner wrapped XML type.
func (c Chart) X() *crt.ChartSpace {
	return c.x
}

// AddLineChart adds a new line chart to a chart.
func (c Chart) AddLineChart() LineChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.LineChart = crt.NewCT_LineChart()
	chc.LineChart.Grouping = crt.NewCT_Grouping()
	chc.LineChart.Grouping.ValAttr = crt.ST_GroupingStandard
	return LineChart{x: chc.LineChart}
}

func setup3DChart(c *crt.CT_Chart) {
	c.View3D = crt.NewCT_View3D()
	c.View3D.RotX = crt.NewCT_RotX()
	c.View3D.RotX.ValAttr = unioffice.Int8(15)
	c.View3D.RotY = crt.NewCT_RotY()
	c.View3D.RotY.ValAttr = unioffice.Uint16(20)
	c.View3D.RAngAx = crt.NewCT_Boolean()
	c.View3D.RAngAx.ValAttr = unioffice.Bool(false)

	c.Floor = crt.NewCT_Surface()
	c.Floor.Thickness = crt.NewCT_Thickness()
	c.Floor.Thickness.ValAttr.Uint32 = unioffice.Uint32(0)

	c.SideWall = crt.NewCT_Surface()
	c.SideWall.Thickness = crt.NewCT_Thickness()
	c.SideWall.Thickness.ValAttr.Uint32 = unioffice.Uint32(0)

	c.BackWall = crt.NewCT_Surface()
	c.BackWall.Thickness = crt.NewCT_Thickness()
	c.BackWall.Thickness.ValAttr.Uint32 = unioffice.Uint32(0)

}

// AddLine3DChart adds a new 3D line chart to a chart.
func (c Chart) AddLine3DChart() Line3DChart {
	setup3DChart(c.x.Chart)
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.Line3DChart = crt.NewCT_Line3DChart()
	chc.Line3DChart.Grouping = crt.NewCT_Grouping()
	chc.Line3DChart.Grouping.ValAttr = crt.ST_GroupingStandard
	return Line3DChart{x: chc.Line3DChart}
}

// AddStockChart adds a new stock chart.
func (c Chart) AddStockChart() StockChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.StockChart = crt.NewCT_StockChart()

	b := StockChart{x: chc.StockChart}
	b.InitializeDefaults()
	return b
}

// AddBarChart adds a new bar chart to a chart.
func (c Chart) AddBarChart() BarChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.BarChart = crt.NewCT_BarChart()
	chc.BarChart.Grouping = crt.NewCT_BarGrouping()
	chc.BarChart.Grouping.ValAttr = crt.ST_BarGroupingStandard

	b := BarChart{x: chc.BarChart}
	b.InitializeDefaults()
	return b
}

// AddBar3DChart adds a new 3D bar chart to a chart.
func (c Chart) AddBar3DChart() Bar3DChart {
	setup3DChart(c.x.Chart)
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.Bar3DChart = crt.NewCT_Bar3DChart()
	chc.Bar3DChart.Grouping = crt.NewCT_BarGrouping()
	chc.Bar3DChart.Grouping.ValAttr = crt.ST_BarGroupingStandard

	b := Bar3DChart{x: chc.Bar3DChart}
	b.InitializeDefaults()
	return b
}

// AddAreaChart adds a new area chart to a chart.
func (c Chart) AddAreaChart() AreaChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.AreaChart = crt.NewCT_AreaChart()

	b := AreaChart{x: chc.AreaChart}
	b.InitializeDefaults()
	return b
}

// AddArea3DChart adds a new area chart to a chart.
func (c Chart) AddArea3DChart() Area3DChart {
	setup3DChart(c.x.Chart)
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.Area3DChart = crt.NewCT_Area3DChart()

	b := Area3DChart{x: chc.Area3DChart}
	b.InitializeDefaults()
	return b
}

// AddRadarChart adds a new radar chart to a chart.
func (c Chart) AddRadarChart() RadarChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.RadarChart = crt.NewCT_RadarChart()

	b := RadarChart{x: chc.RadarChart}
	b.InitializeDefaults()
	return b
}

// AddPieChart adds a new pie chart to a chart.
func (c Chart) AddPieChart() PieChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.PieChart = crt.NewCT_PieChart()

	b := PieChart{x: chc.PieChart}
	b.InitializeDefaults()
	return b
}

// AddPie3DChart adds a new pie chart to a chart.
func (c Chart) AddPie3DChart() Pie3DChart {
	setup3DChart(c.x.Chart)
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.Pie3DChart = crt.NewCT_Pie3DChart()

	b := Pie3DChart{x: chc.Pie3DChart}
	b.InitializeDefaults()
	return b
}

// AddSurfaceChart adds a new surface chart to a chart.
func (c Chart) AddSurfaceChart() SurfaceChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.SurfaceChart = crt.NewCT_SurfaceChart()

	// The surface chart is in 3d, only rotated to an initial straight down view and
	// perspective = 0
	setup3DChart(c.x.Chart)
	c.x.Chart.View3D.RotX.ValAttr = unioffice.Int8(90)
	c.x.Chart.View3D.RotY.ValAttr = unioffice.Uint16(0)
	c.x.Chart.View3D.Perspective = crt.NewCT_Perspective()
	c.x.Chart.View3D.Perspective.ValAttr = unioffice.Uint8(0)

	b := SurfaceChart{x: chc.SurfaceChart}
	b.InitializeDefaults()
	return b
}

// AddSurface3DChart adds a new 3D surface chart to a chart.
func (c Chart) AddSurface3DChart() Surface3DChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.Surface3DChart = crt.NewCT_Surface3DChart()

	setup3DChart(c.x.Chart)
	b := Surface3DChart{x: chc.Surface3DChart}
	b.InitializeDefaults()
	return b
}

// AddPieOfPieChart adds a new pie chart to a chart.
func (c Chart) AddPieOfPieChart() PieOfPieChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.OfPieChart = crt.NewCT_OfPieChart()

	b := PieOfPieChart{x: chc.OfPieChart}
	b.InitializeDefaults()
	return b
}

// AddDoughnutChart adds a new doughnut (pie with a hole in the center) chart to a chart.
func (c Chart) AddDoughnutChart() DoughnutChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.DoughnutChart = crt.NewCT_DoughnutChart()
	b := DoughnutChart{x: chc.DoughnutChart}
	b.InitializeDefaults()
	return b
}

// AddScatterChart adds a scatter (X/Y) chart.
func (c Chart) AddScatterChart() ScatterChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.ScatterChart = crt.NewCT_ScatterChart()

	b := ScatterChart{x: chc.ScatterChart}
	b.InitializeDefaults()
	return b
}

// AddBubbleChart adds a new bubble chart.
func (c Chart) AddBubbleChart() BubbleChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.BubbleChart = crt.NewCT_BubbleChart()

	b := BubbleChart{x: chc.BubbleChart}
	b.InitializeDefaults()
	return b
}

// Properties returns the chart's shape properties.
func (c Chart) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}

// SetDisplayBlanksAs controls how missing values are displayed.
func (c Chart) SetDisplayBlanksAs(v crt.ST_DispBlanksAs) {
	c.x.Chart.DispBlanksAs = crt.NewCT_DispBlanksAs()
	c.x.Chart.DispBlanksAs.ValAttr = v
}

// AddValueAxis adds a value axis to the chart.
func (c Chart) AddValueAxis() ValueAxis {
	va := crt.NewCT_ValAx()
	if c.x.Chart.PlotArea.CChoice == nil {
		c.x.Chart.PlotArea.CChoice = crt.NewCT_PlotAreaChoice1()
	}
	va.AxId = crt.NewCT_UnsignedInt()
	va.AxId.ValAttr = 0x7FFFFFFF & rand.Uint32()
	c.x.Chart.PlotArea.CChoice.ValAx = append(c.x.Chart.PlotArea.CChoice.ValAx, va)

	va.Delete = crt.NewCT_Boolean()
	va.Delete.ValAttr = unioffice.Bool(false)

	va.Scaling = crt.NewCT_Scaling()
	va.Scaling.Orientation = crt.NewCT_Orientation()
	va.Scaling.Orientation.ValAttr = crt.ST_OrientationMinMax

	va.Choice = &crt.EG_AxSharedChoice{}
	va.Choice.Crosses = crt.NewCT_Crosses()
	va.Choice.Crosses.ValAttr = crt.ST_CrossesAutoZero

	va.CrossBetween = crt.NewCT_CrossBetween()
	va.CrossBetween.ValAttr = crt.ST_CrossBetweenBetween

	vax := MakeValueAxis(va)
	vax.MajorGridLines().Properties().LineProperties().SetSolidFill(color.LightGray)
	vax.SetMajorTickMark(crt.ST_TickMarkOut)
	vax.SetMinorTickMark(crt.ST_TickMarkIn)
	vax.SetTickLabelPosition(crt.ST_TickLblPosNextTo)
	vax.Properties().LineProperties().SetSolidFill(color.Black)

	vax.SetPosition(crt.ST_AxPosL)
	return vax
}

// AddCategoryAxis adds a category axis.
func (c Chart) AddCategoryAxis() CategoryAxis {
	ca := crt.NewCT_CatAx()
	if c.x.Chart.PlotArea.CChoice == nil {
		c.x.Chart.PlotArea.CChoice = crt.NewCT_PlotAreaChoice1()
	}

	ca.AxId = crt.NewCT_UnsignedInt()
	ca.AxId.ValAttr = 0x7FFFFFFF & rand.Uint32()
	c.x.Chart.PlotArea.CChoice.CatAx = append(c.x.Chart.PlotArea.CChoice.CatAx, ca)

	ca.Auto = crt.NewCT_Boolean()
	ca.Auto.ValAttr = unioffice.Bool(true)

	ca.Delete = crt.NewCT_Boolean()
	ca.Delete.ValAttr = unioffice.Bool(false)

	cax := MakeCategoryAxis(ca)
	cax.InitializeDefaults()

	return cax
}

// AddDateAxis adds a value axis to the chart.
func (c Chart) AddDateAxis() DateAxis {
	va := crt.NewCT_DateAx()
	if c.x.Chart.PlotArea.CChoice == nil {
		c.x.Chart.PlotArea.CChoice = crt.NewCT_PlotAreaChoice1()
	}
	va.AxId = crt.NewCT_UnsignedInt()
	va.AxId.ValAttr = 0x7FFFFFFF & rand.Uint32()
	c.x.Chart.PlotArea.CChoice.DateAx = append(c.x.Chart.PlotArea.CChoice.DateAx, va)

	va.Delete = crt.NewCT_Boolean()
	va.Delete.ValAttr = unioffice.Bool(false)

	va.Scaling = crt.NewCT_Scaling()
	va.Scaling.Orientation = crt.NewCT_Orientation()
	va.Scaling.Orientation.ValAttr = crt.ST_OrientationMinMax

	va.Choice = &crt.EG_AxSharedChoice{}
	va.Choice.Crosses = crt.NewCT_Crosses()
	va.Choice.Crosses.ValAttr = crt.ST_CrossesAutoZero

	vax := DateAxis{va}
	vax.MajorGridLines().Properties().LineProperties().SetSolidFill(color.LightGray)
	vax.SetMajorTickMark(crt.ST_TickMarkOut)
	vax.SetMinorTickMark(crt.ST_TickMarkIn)
	vax.SetTickLabelPosition(crt.ST_TickLblPosNextTo)
	vax.Properties().LineProperties().SetSolidFill(color.Black)

	vax.SetPosition(crt.ST_AxPosL)
	return vax
}
func (c Chart) AddSeriesAxis() SeriesAxis {
	sa := crt.NewCT_SerAx()
	if c.x.Chart.PlotArea.CChoice == nil {
		c.x.Chart.PlotArea.CChoice = crt.NewCT_PlotAreaChoice1()
	}

	sa.AxId = crt.NewCT_UnsignedInt()
	sa.AxId.ValAttr = 0x7FFFFFFF & rand.Uint32()
	c.x.Chart.PlotArea.CChoice.SerAx = append(c.x.Chart.PlotArea.CChoice.SerAx, sa)

	sa.Delete = crt.NewCT_Boolean()
	sa.Delete.ValAttr = unioffice.Bool(false)

	sax := MakeSeriesAxis(sa)
	sax.InitializeDefaults()

	return sax
}

// RemoveLegend removes the legend if the chart has one.
func (c Chart) RemoveLegend() {
	c.x.Chart.Legend = nil
}

// AddLegend adds a legend to a chart, replacing any existing legend.
func (c Chart) AddLegend() Legend {
	c.x.Chart.Legend = crt.NewCT_Legend()
	leg := MakeLegend(c.x.Chart.Legend)
	leg.InitializeDefaults()
	return leg
}

// RemoveTitle removes any existing title from the chart.
func (c Chart) RemoveTitle() {
	c.x.Chart.Title = nil
	c.x.Chart.AutoTitleDeleted = crt.NewCT_Boolean()
	c.x.Chart.AutoTitleDeleted.ValAttr = unioffice.Bool(true)
}

// AddTitle sets a new title on the chart.
func (c Chart) AddTitle() Title {
	c.x.Chart.Title = crt.NewCT_Title()
	c.x.Chart.Title.Overlay = crt.NewCT_Boolean()
	c.x.Chart.Title.Overlay.ValAttr = unioffice.Bool(false)

	c.x.Chart.AutoTitleDeleted = crt.NewCT_Boolean()
	c.x.Chart.AutoTitleDeleted.ValAttr = unioffice.Bool(false)

	title := MakeTitle(c.x.Chart.Title)
	title.InitializeDefaults()
	return title
}
