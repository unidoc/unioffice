// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_PlotAreaChoice struct {
	AreaChart      []*CT_AreaChart
	Area3DChart    []*CT_Area3DChart
	LineChart      []*CT_LineChart
	Line3DChart    []*CT_Line3DChart
	StockChart     []*CT_StockChart
	RadarChart     []*CT_RadarChart
	ScatterChart   []*CT_ScatterChart
	PieChart       []*CT_PieChart
	Pie3DChart     []*CT_Pie3DChart
	DoughnutChart  []*CT_DoughnutChart
	BarChart       []*CT_BarChart
	Bar3DChart     []*CT_Bar3DChart
	OfPieChart     []*CT_OfPieChart
	SurfaceChart   []*CT_SurfaceChart
	Surface3DChart []*CT_Surface3DChart
	BubbleChart    []*CT_BubbleChart
}

func NewCT_PlotAreaChoice() *CT_PlotAreaChoice {
	ret := &CT_PlotAreaChoice{}
	return ret
}
func (m *CT_PlotAreaChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AreaChart != nil {
		seareaChart := xml.StartElement{Name: xml.Name{Local: "areaChart"}}
		e.EncodeElement(m.AreaChart, seareaChart)
	}
	if m.Area3DChart != nil {
		searea3DChart := xml.StartElement{Name: xml.Name{Local: "area3DChart"}}
		e.EncodeElement(m.Area3DChart, searea3DChart)
	}
	if m.LineChart != nil {
		selineChart := xml.StartElement{Name: xml.Name{Local: "lineChart"}}
		e.EncodeElement(m.LineChart, selineChart)
	}
	if m.Line3DChart != nil {
		seline3DChart := xml.StartElement{Name: xml.Name{Local: "line3DChart"}}
		e.EncodeElement(m.Line3DChart, seline3DChart)
	}
	if m.StockChart != nil {
		sestockChart := xml.StartElement{Name: xml.Name{Local: "stockChart"}}
		e.EncodeElement(m.StockChart, sestockChart)
	}
	if m.RadarChart != nil {
		seradarChart := xml.StartElement{Name: xml.Name{Local: "radarChart"}}
		e.EncodeElement(m.RadarChart, seradarChart)
	}
	if m.ScatterChart != nil {
		sescatterChart := xml.StartElement{Name: xml.Name{Local: "scatterChart"}}
		e.EncodeElement(m.ScatterChart, sescatterChart)
	}
	if m.PieChart != nil {
		sepieChart := xml.StartElement{Name: xml.Name{Local: "pieChart"}}
		e.EncodeElement(m.PieChart, sepieChart)
	}
	if m.Pie3DChart != nil {
		sepie3DChart := xml.StartElement{Name: xml.Name{Local: "pie3DChart"}}
		e.EncodeElement(m.Pie3DChart, sepie3DChart)
	}
	if m.DoughnutChart != nil {
		sedoughnutChart := xml.StartElement{Name: xml.Name{Local: "doughnutChart"}}
		e.EncodeElement(m.DoughnutChart, sedoughnutChart)
	}
	if m.BarChart != nil {
		sebarChart := xml.StartElement{Name: xml.Name{Local: "barChart"}}
		e.EncodeElement(m.BarChart, sebarChart)
	}
	if m.Bar3DChart != nil {
		sebar3DChart := xml.StartElement{Name: xml.Name{Local: "bar3DChart"}}
		e.EncodeElement(m.Bar3DChart, sebar3DChart)
	}
	if m.OfPieChart != nil {
		seofPieChart := xml.StartElement{Name: xml.Name{Local: "ofPieChart"}}
		e.EncodeElement(m.OfPieChart, seofPieChart)
	}
	if m.SurfaceChart != nil {
		sesurfaceChart := xml.StartElement{Name: xml.Name{Local: "surfaceChart"}}
		e.EncodeElement(m.SurfaceChart, sesurfaceChart)
	}
	if m.Surface3DChart != nil {
		sesurface3DChart := xml.StartElement{Name: xml.Name{Local: "surface3DChart"}}
		e.EncodeElement(m.Surface3DChart, sesurface3DChart)
	}
	if m.BubbleChart != nil {
		sebubbleChart := xml.StartElement{Name: xml.Name{Local: "bubbleChart"}}
		e.EncodeElement(m.BubbleChart, sebubbleChart)
	}
	return nil
}
func (m *CT_PlotAreaChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PlotAreaChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "areaChart":
				tmp := NewCT_AreaChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AreaChart = append(m.AreaChart, tmp)
			case "area3DChart":
				tmp := NewCT_Area3DChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Area3DChart = append(m.Area3DChart, tmp)
			case "lineChart":
				tmp := NewCT_LineChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LineChart = append(m.LineChart, tmp)
			case "line3DChart":
				tmp := NewCT_Line3DChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Line3DChart = append(m.Line3DChart, tmp)
			case "stockChart":
				tmp := NewCT_StockChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.StockChart = append(m.StockChart, tmp)
			case "radarChart":
				tmp := NewCT_RadarChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RadarChart = append(m.RadarChart, tmp)
			case "scatterChart":
				tmp := NewCT_ScatterChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ScatterChart = append(m.ScatterChart, tmp)
			case "pieChart":
				tmp := NewCT_PieChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PieChart = append(m.PieChart, tmp)
			case "pie3DChart":
				tmp := NewCT_Pie3DChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pie3DChart = append(m.Pie3DChart, tmp)
			case "doughnutChart":
				tmp := NewCT_DoughnutChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DoughnutChart = append(m.DoughnutChart, tmp)
			case "barChart":
				tmp := NewCT_BarChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BarChart = append(m.BarChart, tmp)
			case "bar3DChart":
				tmp := NewCT_Bar3DChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Bar3DChart = append(m.Bar3DChart, tmp)
			case "ofPieChart":
				tmp := NewCT_OfPieChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.OfPieChart = append(m.OfPieChart, tmp)
			case "surfaceChart":
				tmp := NewCT_SurfaceChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SurfaceChart = append(m.SurfaceChart, tmp)
			case "surface3DChart":
				tmp := NewCT_Surface3DChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Surface3DChart = append(m.Surface3DChart, tmp)
			case "bubbleChart":
				tmp := NewCT_BubbleChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BubbleChart = append(m.BubbleChart, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PlotAreaChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PlotAreaChoice) Validate() error {
	return m.ValidateWithPath("CT_PlotAreaChoice")
}
func (m *CT_PlotAreaChoice) ValidateWithPath(path string) error {
	for i, v := range m.AreaChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AreaChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Area3DChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Area3DChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.LineChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LineChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Line3DChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Line3DChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.StockChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/StockChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.RadarChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RadarChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ScatterChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ScatterChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.PieChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PieChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Pie3DChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pie3DChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.DoughnutChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DoughnutChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BarChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BarChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Bar3DChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Bar3DChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.OfPieChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/OfPieChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.SurfaceChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SurfaceChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Surface3DChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Surface3DChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BubbleChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BubbleChart[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
