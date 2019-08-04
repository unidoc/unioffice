// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vmldrawing

import (
	"fmt"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/office/excel"

	st "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	vml "github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml"
)

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing() *Container {
	c := NewContainer()
	c.Layout = vml.NewOfcShapelayout()
	c.Layout.ExtAttr = vml.ST_ExtEdit
	c.Layout.Idmap = vml.NewOfcCT_IdMap()
	c.Layout.Idmap.DataAttr = unioffice.String("1")
	c.Layout.Idmap.ExtAttr = vml.ST_ExtEdit

	c.ShapeType = vml.NewShapetype()
	c.ShapeType.IdAttr = unioffice.String("_x0000_t202")
	c.ShapeType.CoordsizeAttr = unioffice.String("21600,21600")
	c.ShapeType.SptAttr = unioffice.Float32(202)
	c.ShapeType.PathAttr = unioffice.String("m0,0l0,21600,21600,21600,21600,0xe")

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
	shape.IdAttr = unioffice.String(fmt.Sprintf("cs_%d_%d", col, row))
	shape.TypeAttr = unioffice.String("#_x0000_t202")
	// visibility of the comment box is controlled by this visibility style
	shape.StyleAttr = unioffice.String("position:absolute;margin-left:80pt;margin-top:2pt;width:104pt;height:76pt;z-index:1;visibility:hidden")
	shape.FillcolorAttr = unioffice.String("#fbf6d6")
	shape.StrokecolorAttr = unioffice.String("#edeaa1")
	//shape.InsetmodeAttr = vml.OfcST_InsetModeAuto

	fill := vml.NewEG_ShapeElements()
	fill.Fill = vml.NewFill()
	fill.Fill.Color2Attr = unioffice.String("#fbfe82")
	fill.Fill.AngleAttr = unioffice.Float64(-180)
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
	tb.Textbox.StyleAttr = unioffice.String("mso-direction-alt:auto")
	// TODO: add div?
	shape.EG_ShapeElements = append(shape.EG_ShapeElements, tb)

	cd := vml.NewEG_ShapeElements()
	cd.ClientData = excel.NewClientData()
	cd.ClientData.ObjectTypeAttr = excel.ST_ObjectTypeNote
	cd.ClientData.MoveWithCells = st.ST_TrueFalseBlankT
	cd.ClientData.SizeWithCells = st.ST_TrueFalseBlankT
	cd.ClientData.Anchor = unioffice.String("1, 15, 0, 2, 2, 54, 5, 3")
	cd.ClientData.AutoFill = st.ST_TrueFalseBlankFalse
	cd.ClientData.Row = unioffice.Int64(row)
	cd.ClientData.Column = unioffice.Int64(col)
	shape.EG_ShapeElements = append(shape.EG_ShapeElements, cd)

	return shape
}
