// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_Dxf struct {
	// Font Properties
	Font *CT_Font
	// Number Format
	NumFmt *CT_NumFmt
	// Fill
	Fill *CT_Fill
	// Alignment
	Alignment *CT_CellAlignment
	// Border Properties
	Border *CT_Border
	// Protection Properties
	Protection *CT_CellProtection
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Dxf() *CT_Dxf {
	ret := &CT_Dxf{}
	return ret
}
func (m *CT_Dxf) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Font != nil {
		sefont := xml.StartElement{Name: xml.Name{Local: "x:font"}}
		e.EncodeElement(m.Font, sefont)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "x:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "x:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.Alignment != nil {
		sealignment := xml.StartElement{Name: xml.Name{Local: "x:alignment"}}
		e.EncodeElement(m.Alignment, sealignment)
	}
	if m.Border != nil {
		seborder := xml.StartElement{Name: xml.Name{Local: "x:border"}}
		e.EncodeElement(m.Border, seborder)
	}
	if m.Protection != nil {
		seprotection := xml.StartElement{Name: xml.Name{Local: "x:protection"}}
		e.EncodeElement(m.Protection, seprotection)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Dxf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Dxf:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "font":
				m.Font = NewCT_Font()
				if err := d.DecodeElement(m.Font, &el); err != nil {
					return err
				}
			case "numFmt":
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case "fill":
				m.Fill = NewCT_Fill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case "alignment":
				m.Alignment = NewCT_CellAlignment()
				if err := d.DecodeElement(m.Alignment, &el); err != nil {
					return err
				}
			case "border":
				m.Border = NewCT_Border()
				if err := d.DecodeElement(m.Border, &el); err != nil {
					return err
				}
			case "protection":
				m.Protection = NewCT_CellProtection()
				if err := d.DecodeElement(m.Protection, &el); err != nil {
					return err
				}
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
			break lCT_Dxf
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Dxf) Validate() error {
	return m.ValidateWithPath("CT_Dxf")
}
func (m *CT_Dxf) ValidateWithPath(path string) error {
	if m.Font != nil {
		if err := m.Font.ValidateWithPath(path + "/Font"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.Alignment != nil {
		if err := m.Alignment.ValidateWithPath(path + "/Alignment"); err != nil {
			return err
		}
	}
	if m.Border != nil {
		if err := m.Border.ValidateWithPath(path + "/Border"); err != nil {
			return err
		}
	}
	if m.Protection != nil {
		if err := m.Protection.ValidateWithPath(path + "/Protection"); err != nil {
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
