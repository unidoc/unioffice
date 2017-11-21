// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/measurement"

	"baliance.com/gooxml/schema/soo/dml"
)

type ShapeProperties struct {
	x *dml.CT_ShapeProperties
}

func MakeShapeProperties(x *dml.CT_ShapeProperties) ShapeProperties {
	return ShapeProperties{x}
}

// X returns the inner wrapped XML type.
func (s ShapeProperties) X() *dml.CT_ShapeProperties {
	return s.x
}

func (s ShapeProperties) clearFill() {
	s.x.NoFill = nil
	s.x.BlipFill = nil
	s.x.GradFill = nil
	s.x.GrpFill = nil
	s.x.SolidFill = nil
	s.x.PattFill = nil
}

func (s ShapeProperties) SetNoFill() {
	s.clearFill()
	s.x.NoFill = dml.NewCT_NoFillProperties()
}

func (s ShapeProperties) SetSolidFill(c color.Color) {
	s.clearFill()
	s.x.SolidFill = dml.NewCT_SolidColorFillProperties()
	s.x.SolidFill.SrgbClr = dml.NewCT_SRgbColor()
	s.x.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

func (s ShapeProperties) LineProperties() LineProperties {
	if s.x.Ln == nil {
		s.x.Ln = dml.NewCT_LineProperties()
	}
	return LineProperties{s.x.Ln}
}

func (s ShapeProperties) ensureXfrm() {
	if s.x.Xfrm == nil {
		s.x.Xfrm = dml.NewCT_Transform2D()
	}
}

// SetWidth sets the width of the shape.
func (s ShapeProperties) SetWidth(w measurement.Distance) {
	s.ensureXfrm()
	if s.x.Xfrm.Ext == nil {
		s.x.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	}
	s.x.Xfrm.Ext.CxAttr = int64(w / measurement.EMU)
}

// SetHeight sets the height of the shape.
func (s ShapeProperties) SetHeight(h measurement.Distance) {
	s.ensureXfrm()
	if s.x.Xfrm.Ext == nil {
		s.x.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	}
	s.x.Xfrm.Ext.CyAttr = int64(h / measurement.EMU)
}

// SetSize sets the width and height of the shape.
func (s ShapeProperties) SetSize(w, h measurement.Distance) {
	s.SetWidth(w)
	s.SetHeight(h)
}

// SetPosition sets the position of the shape.
func (s ShapeProperties) SetPosition(x, y measurement.Distance) {
	s.ensureXfrm()
	if s.x.Xfrm.Off == nil {
		s.x.Xfrm.Off = dml.NewCT_Point2D()
	}
	s.x.Xfrm.Off.XAttr.ST_CoordinateUnqualified = gooxml.Int64(int64(x / measurement.EMU))
	s.x.Xfrm.Off.YAttr.ST_CoordinateUnqualified = gooxml.Int64(int64(y / measurement.EMU))
}

// SetGeometry sets the shape type of the shape
func (s ShapeProperties) SetGeometry(g dml.ST_ShapeType) {
	if s.x.PrstGeom == nil {
		s.x.PrstGeom = dml.NewCT_PresetGeometry2D()
	}
	s.x.PrstGeom.PrstAttr = g
}

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (s ShapeProperties) SetFlipHorizontal(b bool) {
	s.ensureXfrm()
	if !b {
		s.x.Xfrm.FlipHAttr = nil
	} else {
		s.x.Xfrm.FlipHAttr = gooxml.Bool(true)
	}
}

// SetFlipVertical controls if the shape is flipped vertically.
func (s ShapeProperties) SetFlipVertical(b bool) {
	s.ensureXfrm()
	if !b {
		s.x.Xfrm.FlipVAttr = nil
	} else {
		s.x.Xfrm.FlipVAttr = gooxml.Bool(true)
	}
}
