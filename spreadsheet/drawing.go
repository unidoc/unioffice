// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/chart"
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

func (d Drawing) InitializeDefaults() {

}

func (d Drawing) AddChart() (chart.Chart, Anchor) {
	chartSpace := crt.NewChartSpace()
	d.wb.charts = append(d.wb.charts, chartSpace)

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

	tca := sd.NewCT_TwoCellAnchor()

	d.x.EG_Anchor = append(d.x.EG_Anchor, &sd.EG_Anchor{TwoCellAnchor: tca})
	tca.EditAsAttr = sd.ST_EditAsOneCell

	// provide a default size so its visible, if from/to are both 0,0 then the
	// chart won't show up.
	tca.From.Col = 5
	tca.From.Row = 0
	// Mac Excel requires the offsets be present
	tca.From.ColOff.ST_CoordinateUnqualified = gooxml.Int64(0)
	tca.From.RowOff.ST_CoordinateUnqualified = gooxml.Int64(0)
	tca.To.Col = 10
	tca.To.Row = 20
	tca.To.ColOff.ST_CoordinateUnqualified = gooxml.Int64(0)
	tca.To.RowOff.ST_CoordinateUnqualified = gooxml.Int64(0)

	tca.Choice = &sd.EG_ObjectChoicesChoice{}
	tca.Choice.GraphicFrame = sd.NewCT_GraphicalObjectFrame()

	// required by Mac Excel
	tca.Choice.GraphicFrame.NvGraphicFramePr = sd.NewCT_GraphicalObjectFrameNonVisual()
	tca.Choice.GraphicFrame.NvGraphicFramePr.CNvPr.IdAttr = 2
	tca.Choice.GraphicFrame.NvGraphicFramePr.CNvPr.NameAttr = "Chart"

	tca.Choice.GraphicFrame.Graphic = dml.NewGraphic()
	tca.Choice.GraphicFrame.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/chart"
	c := c.NewChart()
	c.IdAttr = chartID
	tca.Choice.GraphicFrame.Graphic.GraphicData.Any = []gooxml.Any{c}

	//chart.Chart.PlotVisOnly = crt.NewCT_Boolean()
	//chart.Chart.PlotVisOnly.ValAttr = gooxml.Bool(true)

	chrt := chart.MakeChart(chartSpace)
	chrt.Properties().SetSolidFill(color.White)
	chrt.SetDisplayBlanksAs(crt.ST_DispBlanksAsGap)
	return chrt, TwoCellAnchor{tca}
}
