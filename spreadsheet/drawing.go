// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
	c "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
	sd "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"
)

type Drawing struct {
	wb *Workbook
	x  *sd.WsDr
}

// X returns the inner wrapped XML type.
func (d Drawing) X() *sd.WsDr {
	return d.x
}

func (d Drawing) InitializeDefault() {
	d.x.TwoCellAnchor = sd.NewCT_TwoCellAnchor()
	d.x.TwoCellAnchor.EditAsAttr = sd.ST_EditAsOneCell

	// provide a default size so its visible, if from/to are both 0,0 then the
	// chart won't show up.
	d.x.TwoCellAnchor.From.Col = 5
	d.x.TwoCellAnchor.From.Row = 0
	d.x.TwoCellAnchor.From.ColOff.ST_CoordinateUnqualified = gooxml.Int64(0)
	d.x.TwoCellAnchor.To.Col = 10
	d.x.TwoCellAnchor.To.Row = 20
	d.x.TwoCellAnchor.To.ColOff.ST_CoordinateUnqualified = gooxml.Int64(0)
}
func (d Drawing) AddChart() Chart {
	chart := crt.NewChartSpace()
	d.wb.charts = append(d.wb.charts, chart)
	chrt := Chart{chart}

	fn := gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.ChartContentType, len(d.wb.charts))
	d.wb.ContentTypes.AddOverride(fn, gooxml.ChartContentType)

	var chartID string
	// add relationship from drawing to the chart
	for i, dr := range d.wb.drawings {
		if dr == d.x {
			fn := gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, gooxml.ChartType, len(d.wb.charts))
			rel := d.wb.drawingRels[i].AddRelationship(fn, gooxml.ChartType)
			chartID = rel.ID()
			break
		}
	}

	// maybe use a one cell anchor?
	if d.x.TwoCellAnchor == nil {
		d.InitializeDefault()
	}
	d.x.TwoCellAnchor.Choice = &sd.EG_ObjectChoicesChoice{}
	d.x.TwoCellAnchor.Choice.GraphicFrame = sd.NewCT_GraphicalObjectFrame()
	d.x.TwoCellAnchor.Choice.GraphicFrame.Graphic = dml.NewGraphic()
	d.x.TwoCellAnchor.Choice.GraphicFrame.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/chart"
	c := c.NewChart()
	c.IdAttr = chartID
	d.x.TwoCellAnchor.Choice.GraphicFrame.Graphic.GraphicData.Any = []gooxml.Any{c}

	//chart.Chart.PlotVisOnly = crt.NewCT_Boolean()
	//chart.Chart.PlotVisOnly.ValAttr = gooxml.Bool(true)

	chrt.Properties().SetSolidFill(color.White)
	chrt.SetDisplayBlanksAs(crt.ST_DispBlanksAsGap)
	return chrt
}

// TopLeft allows manipulating the top left position of the drawing.
func (d Drawing) TopLeft() CellMarker {
	if d.x.TwoCellAnchor != nil {
		if d.x.TwoCellAnchor.From == nil {
			d.x.TwoCellAnchor.From = sd.NewCT_Marker()
		}
		return CellMarker{d.x.TwoCellAnchor.From}
	}
	if d.x.OneCellAnchor != nil {
		if d.x.OneCellAnchor.From == nil {
			d.x.OneCellAnchor.From = sd.NewCT_Marker()
		}
		return CellMarker{d.x.OneCellAnchor.From}
	}

	// this will crash if used...
	return CellMarker{}
}

// BottomRight allows manipulating the bottom right position of the drawing.
func (d Drawing) BottomRight() CellMarker {
	if d.x.TwoCellAnchor != nil {
		if d.x.TwoCellAnchor.To == nil {
			d.x.TwoCellAnchor.To = sd.NewCT_Marker()
		}
		return CellMarker{d.x.TwoCellAnchor.To}
	}

	// this will crash if used...
	return CellMarker{}
}
