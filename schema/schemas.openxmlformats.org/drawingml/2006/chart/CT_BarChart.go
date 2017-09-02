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

type CT_BarChart struct {
	BarDir     *CT_BarDir
	Grouping   *CT_BarGrouping
	VaryColors *CT_Boolean
	Ser        []*CT_BarSer
	DLbls      *CT_DLbls
	GapWidth   *CT_GapAmount
	Overlap    *CT_Overlap
	SerLines   []*CT_ChartLines
	AxId       []*CT_UnsignedInt
	ExtLst     *CT_ExtensionList
}

func NewCT_BarChart() *CT_BarChart {
	ret := &CT_BarChart{}
	ret.BarDir = NewCT_BarDir()
	return ret
}

func (m *CT_BarChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sebarDir := xml.StartElement{Name: xml.Name{Local: "barDir"}}
	e.EncodeElement(m.BarDir, sebarDir)
	if m.Grouping != nil {
		segrouping := xml.StartElement{Name: xml.Name{Local: "grouping"}}
		e.EncodeElement(m.Grouping, segrouping)
	}
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
	if m.GapWidth != nil {
		segapWidth := xml.StartElement{Name: xml.Name{Local: "gapWidth"}}
		e.EncodeElement(m.GapWidth, segapWidth)
	}
	if m.Overlap != nil {
		seoverlap := xml.StartElement{Name: xml.Name{Local: "overlap"}}
		e.EncodeElement(m.Overlap, seoverlap)
	}
	if m.SerLines != nil {
		seserLines := xml.StartElement{Name: xml.Name{Local: "serLines"}}
		e.EncodeElement(m.SerLines, seserLines)
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

func (m *CT_BarChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BarDir = NewCT_BarDir()
lCT_BarChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "barDir":
				if err := d.DecodeElement(m.BarDir, &el); err != nil {
					return err
				}
			case "grouping":
				m.Grouping = NewCT_BarGrouping()
				if err := d.DecodeElement(m.Grouping, &el); err != nil {
					return err
				}
			case "varyColors":
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case "ser":
				tmp := NewCT_BarSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case "dLbls":
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case "gapWidth":
				m.GapWidth = NewCT_GapAmount()
				if err := d.DecodeElement(m.GapWidth, &el); err != nil {
					return err
				}
			case "overlap":
				m.Overlap = NewCT_Overlap()
				if err := d.DecodeElement(m.Overlap, &el); err != nil {
					return err
				}
			case "serLines":
				tmp := NewCT_ChartLines()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SerLines = append(m.SerLines, tmp)
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
				log.Printf("skipping unsupported element on CT_BarChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BarChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BarChart and its children
func (m *CT_BarChart) Validate() error {
	return m.ValidateWithPath("CT_BarChart")
}

// ValidateWithPath validates the CT_BarChart and its children, prefixing error messages with path
func (m *CT_BarChart) ValidateWithPath(path string) error {
	if err := m.BarDir.ValidateWithPath(path + "/BarDir"); err != nil {
		return err
	}
	if m.Grouping != nil {
		if err := m.Grouping.ValidateWithPath(path + "/Grouping"); err != nil {
			return err
		}
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
	if m.GapWidth != nil {
		if err := m.GapWidth.ValidateWithPath(path + "/GapWidth"); err != nil {
			return err
		}
	}
	if m.Overlap != nil {
		if err := m.Overlap.ValidateWithPath(path + "/Overlap"); err != nil {
			return err
		}
	}
	for i, v := range m.SerLines {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SerLines[%d]", path, i)); err != nil {
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
