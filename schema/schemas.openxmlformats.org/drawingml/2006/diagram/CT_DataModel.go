// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_DataModel struct {
	PtLst  *CT_PtList
	CxnLst *CT_CxnList
	Bg     *drawingml.CT_BackgroundFormatting
	Whole  *drawingml.CT_WholeE2oFormatting
	ExtLst *drawingml.CT_OfficeArtExtensionList
}

func NewCT_DataModel() *CT_DataModel {
	ret := &CT_DataModel{}
	ret.PtLst = NewCT_PtList()
	return ret
}

func (m *CT_DataModel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	septLst := xml.StartElement{Name: xml.Name{Local: "ptLst"}}
	e.EncodeElement(m.PtLst, septLst)
	if m.CxnLst != nil {
		secxnLst := xml.StartElement{Name: xml.Name{Local: "cxnLst"}}
		e.EncodeElement(m.CxnLst, secxnLst)
	}
	if m.Bg != nil {
		sebg := xml.StartElement{Name: xml.Name{Local: "bg"}}
		e.EncodeElement(m.Bg, sebg)
	}
	if m.Whole != nil {
		sewhole := xml.StartElement{Name: xml.Name{Local: "whole"}}
		e.EncodeElement(m.Whole, sewhole)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataModel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PtLst = NewCT_PtList()
lCT_DataModel:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ptLst":
				if err := d.DecodeElement(m.PtLst, &el); err != nil {
					return err
				}
			case "cxnLst":
				m.CxnLst = NewCT_CxnList()
				if err := d.DecodeElement(m.CxnLst, &el); err != nil {
					return err
				}
			case "bg":
				m.Bg = drawingml.NewCT_BackgroundFormatting()
				if err := d.DecodeElement(m.Bg, &el); err != nil {
					return err
				}
			case "whole":
				m.Whole = drawingml.NewCT_WholeE2oFormatting()
				if err := d.DecodeElement(m.Whole, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_DataModel %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataModel
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DataModel and its children
func (m *CT_DataModel) Validate() error {
	return m.ValidateWithPath("CT_DataModel")
}

// ValidateWithPath validates the CT_DataModel and its children, prefixing error messages with path
func (m *CT_DataModel) ValidateWithPath(path string) error {
	if err := m.PtLst.ValidateWithPath(path + "/PtLst"); err != nil {
		return err
	}
	if m.CxnLst != nil {
		if err := m.CxnLst.ValidateWithPath(path + "/CxnLst"); err != nil {
			return err
		}
	}
	if m.Bg != nil {
		if err := m.Bg.ValidateWithPath(path + "/Bg"); err != nil {
			return err
		}
	}
	if m.Whole != nil {
		if err := m.Whole.ValidateWithPath(path + "/Whole"); err != nil {
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
