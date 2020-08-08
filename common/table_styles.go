// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package common

import (
	"github.com/unidoc/unioffice/schema/soo/dml"
)

// TableStyles contains document specific properties.
type TableStyles struct {
	x *dml.TblStyleLst
}

// NewTableStyles constructs a new TableStyles.
func NewTableStyles() TableStyles {
	return TableStyles{x: dml.NewTblStyleLst()}
}

// X returns the inner wrapped XML type.
func (t TableStyles) X() *dml.TblStyleLst {
	return t.x
}

// DefAttr returns the DefAttr property.
func (t TableStyles) DefAttr() string {
	return t.x.DefAttr
}

// TblStyle returns the TblStyle property.
func (t TableStyles) TblStyle() []*dml.CT_TableStyle {
	return t.x.TblStyle
}
