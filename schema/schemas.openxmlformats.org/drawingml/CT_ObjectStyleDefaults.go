// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_ObjectStyleDefaults struct {
	SpDef  *CT_DefaultShapeDefinition
	LnDef  *CT_DefaultShapeDefinition
	TxDef  *CT_DefaultShapeDefinition
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_ObjectStyleDefaults() *CT_ObjectStyleDefaults {
	ret := &CT_ObjectStyleDefaults{}
	return ret
}
func (m *CT_ObjectStyleDefaults) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.SpDef != nil {
		sespDef := xml.StartElement{Name: xml.Name{Local: "a:spDef"}}
		e.EncodeElement(m.SpDef, sespDef)
	}
	if m.LnDef != nil {
		selnDef := xml.StartElement{Name: xml.Name{Local: "a:lnDef"}}
		e.EncodeElement(m.LnDef, selnDef)
	}
	if m.TxDef != nil {
		setxDef := xml.StartElement{Name: xml.Name{Local: "a:txDef"}}
		e.EncodeElement(m.TxDef, setxDef)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ObjectStyleDefaults) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ObjectStyleDefaults:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "spDef":
				m.SpDef = NewCT_DefaultShapeDefinition()
				if err := d.DecodeElement(m.SpDef, &el); err != nil {
					return err
				}
			case "lnDef":
				m.LnDef = NewCT_DefaultShapeDefinition()
				if err := d.DecodeElement(m.LnDef, &el); err != nil {
					return err
				}
			case "txDef":
				m.TxDef = NewCT_DefaultShapeDefinition()
				if err := d.DecodeElement(m.TxDef, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_ObjectStyleDefaults
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ObjectStyleDefaults) Validate() error {
	return m.ValidateWithPath("CT_ObjectStyleDefaults")
}
func (m *CT_ObjectStyleDefaults) ValidateWithPath(path string) error {
	if m.SpDef != nil {
		if err := m.SpDef.ValidateWithPath(path + "/SpDef"); err != nil {
			return err
		}
	}
	if m.LnDef != nil {
		if err := m.LnDef.ValidateWithPath(path + "/LnDef"); err != nil {
			return err
		}
	}
	if m.TxDef != nil {
		if err := m.TxDef.ValidateWithPath(path + "/TxDef"); err != nil {
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
