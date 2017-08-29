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

type CT_LineChart struct {
	Grouping   *CT_Grouping
	VaryColors *CT_Boolean
	Ser        []*CT_LineSer
	DLbls      *CT_DLbls
	DropLines  *CT_ChartLines
	HiLowLines *CT_ChartLines
	UpDownBars *CT_UpDownBars
	Marker     *CT_Boolean
	Smooth     *CT_Boolean
	AxId       []*CT_UnsignedInt
	ExtLst     *CT_ExtensionList
}

func NewCT_LineChart() *CT_LineChart {
	ret := &CT_LineChart{}
	ret.Grouping = NewCT_Grouping()
	return ret
}
func (m *CT_LineChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	segrouping := xml.StartElement{Name: xml.Name{Local: "grouping"}}
	e.EncodeElement(m.Grouping, segrouping)
	if m.VaryColors != nil {
		sevaryColors := xml.StartElement{Name: xml.Name{Local: "varyColors"}}
		e.EncodeElement(m.VaryColors, sevaryColors)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "ser"}}
		e.EncodeElement(m.Ser, seser)
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.DropLines != nil {
		sedropLines := xml.StartElement{Name: xml.Name{Local: "dropLines"}}
		e.EncodeElement(m.DropLines, sedropLines)
	}
	if m.HiLowLines != nil {
		sehiLowLines := xml.StartElement{Name: xml.Name{Local: "hiLowLines"}}
		e.EncodeElement(m.HiLowLines, sehiLowLines)
	}
	if m.UpDownBars != nil {
		seupDownBars := xml.StartElement{Name: xml.Name{Local: "upDownBars"}}
		e.EncodeElement(m.UpDownBars, seupDownBars)
	}
	if m.Marker != nil {
		semarker := xml.StartElement{Name: xml.Name{Local: "marker"}}
		e.EncodeElement(m.Marker, semarker)
	}
	if m.Smooth != nil {
		sesmooth := xml.StartElement{Name: xml.Name{Local: "smooth"}}
		e.EncodeElement(m.Smooth, sesmooth)
	}
	seaxId := xml.StartElement{Name: xml.Name{Local: "axId"}}
	e.EncodeElement(m.AxId, seaxId)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LineChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Grouping = NewCT_Grouping()
lCT_LineChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "grouping":
				if err := d.DecodeElement(m.Grouping, &el); err != nil {
					return err
				}
			case "varyColors":
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case "ser":
				tmp := NewCT_LineSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case "dLbls":
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case "dropLines":
				m.DropLines = NewCT_ChartLines()
				if err := d.DecodeElement(m.DropLines, &el); err != nil {
					return err
				}
			case "hiLowLines":
				m.HiLowLines = NewCT_ChartLines()
				if err := d.DecodeElement(m.HiLowLines, &el); err != nil {
					return err
				}
			case "upDownBars":
				m.UpDownBars = NewCT_UpDownBars()
				if err := d.DecodeElement(m.UpDownBars, &el); err != nil {
					return err
				}
			case "marker":
				m.Marker = NewCT_Boolean()
				if err := d.DecodeElement(m.Marker, &el); err != nil {
					return err
				}
			case "smooth":
				m.Smooth = NewCT_Boolean()
				if err := d.DecodeElement(m.Smooth, &el); err != nil {
					return err
				}
			case "axId":
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AxId = append(m.AxId, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LineChart
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_LineChart) Validate() error {
	return m.ValidateWithPath("CT_LineChart")
}
func (m *CT_LineChart) ValidateWithPath(path string) error {
	if err := m.Grouping.ValidateWithPath(path + "/Grouping"); err != nil {
		return err
	}
	if m.VaryColors != nil {
		if err := m.VaryColors.ValidateWithPath(path + "/VaryColors"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	if m.DropLines != nil {
		if err := m.DropLines.ValidateWithPath(path + "/DropLines"); err != nil {
			return err
		}
	}
	if m.HiLowLines != nil {
		if err := m.HiLowLines.ValidateWithPath(path + "/HiLowLines"); err != nil {
			return err
		}
	}
	if m.UpDownBars != nil {
		if err := m.UpDownBars.ValidateWithPath(path + "/UpDownBars"); err != nil {
			return err
		}
	}
	if m.Marker != nil {
		if err := m.Marker.ValidateWithPath(path + "/Marker"); err != nil {
			return err
		}
	}
	if m.Smooth != nil {
		if err := m.Smooth.ValidateWithPath(path + "/Smooth"); err != nil {
			return err
		}
	}
	for i, v := range m.AxId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AxId[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
