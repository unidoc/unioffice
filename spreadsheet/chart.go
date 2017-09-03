// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"

type Chart struct {
	x *crt.ChartSpace
}

// X returns the inner wrapped XML type.
func (c Chart) X() *crt.ChartSpace {
	return c.x
}

func (c Chart) AddLineChart() LineChart {
	chc := crt.NewCT_PlotAreaChoice()
	c.x.Chart.PlotArea.Choice = append(c.x.Chart.PlotArea.Choice, chc)
	chc.LineChart = crt.NewCT_LineChart()
	chc.LineChart.Grouping = crt.NewCT_Grouping()
	chc.LineChart.Grouping.ValAttr = crt.ST_GroupingStandard
	return LineChart{chc.LineChart}
}
