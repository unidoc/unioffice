// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package presentation

import (
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/schema/soo/dml"
	"github.com/unidoc/unioffice/schema/soo/pml"
)

// Image is an image within a slide.
type Image struct {
	x *pml.CT_Picture
}

// Properties returns the properties of the TextBox.
func (i Image) Properties() drawing.ShapeProperties {
	if i.x.SpPr == nil {
		i.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(i.x.SpPr)
}
