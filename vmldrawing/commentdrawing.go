// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vmldrawing

import (
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/urn/schemas_microsoft_com/office/excel"

	st "baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	vml "baliance.com/gooxml/schema/urn/schemas_microsoft_com/vml"
)

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing() *Container {
	c := NewContainer()
	c.Layout = vml.NewOfcShapelayout()
	c.Layout.ExtAttr = vml.ST_ExtEdit
	c.Layout.Idmap = vml.NewOfcCT_IdMap()
	c.Layout.Idmap.DataAttr = gooxml.String("1")
	c.Layout.Idmap.ExtAttr = vml.ST_ExtEdit

	c.ShapeType = vml.NewShapetype()
	c.ShapeType.IdAttr = gooxml.String("_x0000_t202")
	c.ShapeType.CoordsizeAttr = gooxml.String("21600,21600")
	c.ShapeType.SptAttr = gooxml.Float32(202)
	c.ShapeType.PathAttr = gooxml.String("m0,0l0,21600,21600,21600,21600,0xe")

	se := vml.NewEG_ShapeElements()
	c.ShapeType.EG_ShapeElements = append(c.ShapeType.EG_ShapeElements, se)
	se.Path = vml.NewPath()
	se.Path.GradientshapeokAttr = st.ST_TrueFalseT
	se.Path.ConnecttypeAttr = vml.OfcST_ConnectTypeRect

	return c
}

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape(col, row int64) *vml.Shape {
	shape := vml.NewShape()
	shape.IdAttr = gooxml.String(fmt.Sprintf("cs_%d_%d", col, row))
	shape.TypeAttr = gooxml.String("#_x0000_t202")
	// visibility of the comment box is controlled by this visibility style
	shape.StyleAttr = gooxml.String("position:absolute;margin-left:80pt;margin-top:2pt;width:104pt;height:76pt;z-index:1;visibility:hidden")
	shape.FillcolorAttr = gooxml.String("#fbf6d6")
	shape.StrokecolorAttr = gooxml.String("#edeaa1")
	//shape.InsetmodeAttr = vml.OfcST_InsetModeAuto

	fill := vml.NewEG_ShapeElements()
	fill.Fill = vml.NewFill()
	fill.Fill.Color2Attr = gooxml.String("#fbfe82")
	fill.Fill.AngleAttr = gooxml.Float64(-180)
	fill.Fill.TypeAttr = vml.ST_FillTypeGradient
	fill.Fill.Fill = vml.NewOfcFill()
	fill.Fill.Fill.ExtAttr = vml.ST_ExtView
	fill.Fill.Fill.TypeAttr = vml.OfcST_FillTypeGradientUnscaled

	shape.EG_ShapeElements = append(shape.EG_ShapeElements, fill)

	shadow := vml.NewEG_ShapeElements()
	shadow.Shadow = vml.NewShadow()
	shadow.Shadow.OnAttr = st.ST_TrueFalseT
	shadow.Shadow.ObscuredAttr = st.ST_TrueFalseT
	shape.EG_ShapeElements = append(shape.EG_ShapeElements, shadow)

	fpath := vml.NewEG_ShapeElements()
	fpath.Path = vml.NewPath()
	fpath.Path.ConnecttypeAttr = vml.OfcST_ConnectTypeNone
	shape.EG_ShapeElements = append(shape.EG_ShapeElements, fpath)

	tb := vml.NewEG_ShapeElements()
	tb.Textbox = vml.NewTextbox()
	tb.Textbox.StyleAttr = gooxml.String("mso-direction-alt:auto")
	// TODO: add div?
	shape.EG_ShapeElements = append(shape.EG_ShapeElements, tb)

	cd := vml.NewEG_ShapeElements()
	cd.ClientData = excel.NewClientData()
	cd.ClientData.ObjectTypeAttr = excel.ST_ObjectTypeNote
	cd.ClientData.MoveWithCells = st.ST_TrueFalseBlankT
	cd.ClientData.SizeWithCells = st.ST_TrueFalseBlankT
	cd.ClientData.Anchor = gooxml.String("1, 15, 0, 2, 2, 54, 5, 3")
	cd.ClientData.AutoFill = st.ST_TrueFalseBlankFalse
	cd.ClientData.Row = gooxml.Int64(row)
	cd.ClientData.Column = gooxml.Int64(col)
	shape.EG_ShapeElements = append(shape.EG_ShapeElements, cd)

	return shape
}
