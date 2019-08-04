// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type LineProperties struct {
	x *dml.CT_LineProperties
}

// X returns the inner wrapped XML type.
func (l LineProperties) X() *dml.CT_LineProperties {
	return l.x
}

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (l LineProperties) SetWidth(w measurement.Distance) {
	l.x.WAttr = unioffice.Int32(int32(w / measurement.EMU))
}

func (l LineProperties) clearFill() {
	l.x.NoFill = nil
	l.x.GradFill = nil
	l.x.SolidFill = nil
	l.x.PattFill = nil
}

func (l LineProperties) SetNoFill() {
	l.clearFill()
	l.x.NoFill = dml.NewCT_NoFillProperties()
}

func (l LineProperties) SetSolidFill(c color.Color) {
	l.clearFill()
	l.x.SolidFill = dml.NewCT_SolidColorFillProperties()
	l.x.SolidFill.SrgbClr = dml.NewCT_SRgbColor()
	l.x.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// LineJoin is the type of line join
type LineJoin byte

// LineJoin types
const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
	LineJoinMiter
)

// SetJoin sets the line join style.
func (l LineProperties) SetJoin(e LineJoin) {
	l.x.Round = nil
	l.x.Miter = nil
	l.x.Bevel = nil
	switch e {
	case LineJoinRound:
		l.x.Round = dml.NewCT_LineJoinRound()
	case LineJoinBevel:
		l.x.Bevel = dml.NewCT_LineJoinBevel()
	case LineJoinMiter:
		l.x.Miter = dml.NewCT_LineJoinMiterProperties()
	}
}
