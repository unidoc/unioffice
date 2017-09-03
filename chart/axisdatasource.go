// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"

type AxisDataSource struct {
	x *crt.CT_AxDataSource
}

func MakeAxisDataSource(x *crt.CT_AxDataSource) AxisDataSource {
	return AxisDataSource{x}
}

func (a AxisDataSource) SetReference(s string) {
	a.x.Choice = crt.NewCT_AxDataSourceChoice()
	a.x.Choice.StrRef = crt.NewCT_StrRef()
	a.x.Choice.StrRef.F = s
}
