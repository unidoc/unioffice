// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/chart"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/measurement"

	"baliance.com/gooxml/schema/soo/dml"
	c "baliance.com/gooxml/schema/soo/dml/chart"
	crt "baliance.com/gooxml/schema/soo/dml/chart"
	sd "baliance.com/gooxml/schema/soo/dml/spreadsheetDrawing"
)

// Drawing is a drawing overlay on a sheet.  Only a single drawing is allowed
// per sheet, so to display multiple charts and images on a single sheet, they
// must be added to the same drawing.
type Drawing struct {
	wb *Workbook
	x  *sd.WsDr
}

// X returns the inner wrapped XML type.
func (d Drawing) X() *sd.WsDr {
	return d.x
}

// AddChart adds an chart to a drawing, returning the chart and an anchor that
// can be used to position the chart within the sheet.
func (d Drawing) AddChart(at AnchorType) (chart.Chart, Anchor) {
	chartSpace := crt.NewChartSpace()
	d.wb.charts = append(d.wb.charts, chartSpace)

	fn := gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.ChartContentType, len(d.wb.charts))
	d.wb.ContentTypes.AddOverride(fn, gooxml.ChartContentType)

	var chartID string

	// add relationship from drawing to the chart
	for i, dr := range d.wb.drawings {
		if dr == d.x {
			fn := gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, gooxml.DrawingType, gooxml.ChartType, len(d.wb.charts))
			rel := d.wb.drawingRels[i].AddRelationship(fn, gooxml.ChartType)
			chartID = rel.ID()
			break
		}
	}

	var anc Anchor
	var gf *sd.CT_GraphicalObjectFrame
	switch at {
	case AnchorTypeAbsolute:
		aa := defaultAbsoluteAnchor()
		d.x.EG_Anchor = append(d.x.EG_Anchor, &sd.EG_Anchor{AbsoluteAnchor: aa})
		aa.Choice = &sd.EG_ObjectChoicesChoice{}
		aa.Choice.GraphicFrame = sd.NewCT_GraphicalObjectFrame()
		gf = aa.Choice.GraphicFrame
		anc = AbsoluteAnchor{aa}
	case AnchorTypeOneCell:
		oca := defaultOneCelAnchor()
		d.x.EG_Anchor = append(d.x.EG_Anchor, &sd.EG_Anchor{OneCellAnchor: oca})
		oca.Choice = &sd.EG_ObjectChoicesChoice{}
		oca.Choice.GraphicFrame = sd.NewCT_GraphicalObjectFrame()
		gf = oca.Choice.GraphicFrame
		anc = OneCellAnchor{oca}
	case AnchorTypeTwoCell:
		tca := defaultTwoCellAnchor()
		d.x.EG_Anchor = append(d.x.EG_Anchor, &sd.EG_Anchor{TwoCellAnchor: tca})
		tca.Choice = &sd.EG_ObjectChoicesChoice{}
		tca.Choice.GraphicFrame = sd.NewCT_GraphicalObjectFrame()
		gf = tca.Choice.GraphicFrame
		anc = TwoCellAnchor{tca}
	}

	// required by Mac Excel
	gf.NvGraphicFramePr = sd.NewCT_GraphicalObjectFrameNonVisual()
	gf.NvGraphicFramePr.CNvPr.IdAttr = 2
	gf.NvGraphicFramePr.CNvPr.NameAttr = "Chart"

	gf.Graphic = dml.NewGraphic()
	gf.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/chart"
	c := c.NewChart()
	c.IdAttr = chartID
	gf.Graphic.GraphicData.Any = []gooxml.Any{c}

	//chart.Chart.PlotVisOnly = crt.NewCT_Boolean()
	//chart.Chart.PlotVisOnly.ValAttr = gooxml.Bool(true)

	chrt := chart.MakeChart(chartSpace)
	chrt.Properties().SetSolidFill(color.White)
	chrt.SetDisplayBlanksAs(crt.ST_DispBlanksAsGap)
	return chrt, anc
}

// AddImage adds an image with a paricular anchor type, returning an anchor to
// allow adusting the image size/position.
func (d Drawing) AddImage(img common.ImageRef, at AnchorType) Anchor {
	imgIdx := 0
	for i, ig := range d.wb.Images {
		if ig == img {
			imgIdx = i + 1
			break
		}
	}

	var imgID string
	for i, dr := range d.wb.drawings {
		if dr == d.x {
			fn := fmt.Sprintf("../media/image%d.%s", imgIdx, img.Format())
			rel := d.wb.drawingRels[i].AddRelationship(fn, gooxml.ImageType)
			imgID = rel.ID()
			break
		}
	}

	var anc Anchor
	var pic *sd.CT_Picture
	switch at {
	case AnchorTypeAbsolute:
		aa := defaultAbsoluteAnchor()
		d.x.EG_Anchor = append(d.x.EG_Anchor, &sd.EG_Anchor{AbsoluteAnchor: aa})
		aa.Choice = &sd.EG_ObjectChoicesChoice{}
		aa.Choice.Pic = sd.NewCT_Picture()
		pic = aa.Choice.Pic
		anc = AbsoluteAnchor{aa}
	case AnchorTypeOneCell:
		oca := defaultOneCelAnchor()
		d.x.EG_Anchor = append(d.x.EG_Anchor, &sd.EG_Anchor{OneCellAnchor: oca})
		oca.Choice = &sd.EG_ObjectChoicesChoice{}
		oca.Choice.Pic = sd.NewCT_Picture()
		pic = oca.Choice.Pic
		anc = OneCellAnchor{oca}
	case AnchorTypeTwoCell:
		tca := defaultTwoCellAnchor()
		d.x.EG_Anchor = append(d.x.EG_Anchor, &sd.EG_Anchor{TwoCellAnchor: tca})
		tca.Choice = &sd.EG_ObjectChoicesChoice{}
		tca.Choice.Pic = sd.NewCT_Picture()
		pic = tca.Choice.Pic
		anc = TwoCellAnchor{tca}
	}

	pic.NvPicPr.CNvPr.IdAttr = 0
	pic.NvPicPr.CNvPr.NameAttr = "Image"
	pic.BlipFill.Blip = dml.NewCT_Blip()
	pic.BlipFill.Blip.EmbedAttr = gooxml.String(imgID)
	pic.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	pic.SpPr = dml.NewCT_ShapeProperties()
	pic.SpPr.Xfrm = dml.NewCT_Transform2D()
	pic.SpPr.Xfrm.Off = dml.NewCT_Point2D()
	pic.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	pic.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	pic.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	pic.SpPr.Xfrm.Ext.CxAttr = int64(float64(img.Size().X*measurement.Pixel72) / measurement.EMU)
	pic.SpPr.Xfrm.Ext.CyAttr = int64(float64(img.Size().Y*measurement.Pixel72) / measurement.EMU)
	pic.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	pic.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect
	pic.SpPr.Ln = dml.NewCT_LineProperties()
	pic.SpPr.Ln.NoFill = dml.NewCT_NoFillProperties()

	return anc
}

func defaultAbsoluteAnchor() *sd.CT_AbsoluteAnchor {
	aa := sd.NewCT_AbsoluteAnchor()

	return aa
}

func defaultOneCelAnchor() *sd.CT_OneCellAnchor {
	oca := sd.NewCT_OneCellAnchor()
	return oca
}

func defaultTwoCellAnchor() *sd.CT_TwoCellAnchor {
	tca := sd.NewCT_TwoCellAnchor()
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

	return tca
}
