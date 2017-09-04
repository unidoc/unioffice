// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"

// BarChart is a 2D bar chart.
type BarChart struct {
	x *crt.CT_BarChart
}

// X returns the inner wrapped XML type.
func (c BarChart) X() *crt.CT_BarChart {
	return c.x
}
