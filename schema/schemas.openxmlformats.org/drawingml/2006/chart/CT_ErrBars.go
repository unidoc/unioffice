// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_ErrBars struct {
	ErrDir     *CT_ErrDir
	ErrBarType *CT_ErrBarType
	ErrValType *CT_ErrValType
	NoEndCap   *CT_Boolean
	Plus       *CT_NumDataSource
	Minus      *CT_NumDataSource
	Val        *CT_Double
	SpPr       *drawingml.CT_ShapeProperties
	ExtLst     *CT_ExtensionList
}

func NewCT_ErrBars() *CT_ErrBars {
	ret := &CT_ErrBars{}
	ret.ErrBarType = NewCT_ErrBarType()
	ret.ErrValType = NewCT_ErrValType()
	return ret
}
func (m *CT_ErrBars) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ErrDir != nil {
		seerrDir := xml.StartElement{Name: xml.Name{Local: "errDir"}}
		e.EncodeElement(m.ErrDir, seerrDir)
	}
	seerrBarType := xml.StartElement{Name: xml.Name{Local: "errBarType"}}
	e.EncodeElement(m.ErrBarType, seerrBarType)
	seerrValType := xml.StartElement{Name: xml.Name{Local: "errValType"}}
	e.EncodeElement(m.ErrValType, seerrValType)
	if m.NoEndCap != nil {
		senoEndCap := xml.StartElement{Name: xml.Name{Local: "noEndCap"}}
		e.EncodeElement(m.NoEndCap, senoEndCap)
	}
	if m.Plus != nil {
		seplus := xml.StartElement{Name: xml.Name{Local: "plus"}}
		e.EncodeElement(m.Plus, seplus)
	}
	if m.Minus != nil {
		seminus := xml.StartElement{Name: xml.Name{Local: "minus"}}
		e.EncodeElement(m.Minus, seminus)
	}
	if m.Val != nil {
		seval := xml.StartElement{Name: xml.Name{Local: "val"}}
		e.EncodeElement(m.Val, seval)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ErrBars) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ErrBarType = NewCT_ErrBarType()
	m.ErrValType = NewCT_ErrValType()
lCT_ErrBars:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "errDir":
				m.ErrDir = NewCT_ErrDir()
				if err := d.DecodeElement(m.ErrDir, &el); err != nil {
					return err
				}
			case "errBarType":
				if err := d.DecodeElement(m.ErrBarType, &el); err != nil {
					return err
				}
			case "errValType":
				if err := d.DecodeElement(m.ErrValType, &el); err != nil {
					return err
				}
			case "noEndCap":
				m.NoEndCap = NewCT_Boolean()
				if err := d.DecodeElement(m.NoEndCap, &el); err != nil {
					return err
				}
			case "plus":
				m.Plus = NewCT_NumDataSource()
				if err := d.DecodeElement(m.Plus, &el); err != nil {
					return err
				}
			case "minus":
				m.Minus = NewCT_NumDataSource()
				if err := d.DecodeElement(m.Minus, &el); err != nil {
					return err
				}
			case "val":
				m.Val = NewCT_Double()
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
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
			break lCT_ErrBars
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ErrBars) Validate() error {
	return m.ValidateWithPath("CT_ErrBars")
}
func (m *CT_ErrBars) ValidateWithPath(path string) error {
	if m.ErrDir != nil {
		if err := m.ErrDir.ValidateWithPath(path + "/ErrDir"); err != nil {
			return err
		}
	}
	if err := m.ErrBarType.ValidateWithPath(path + "/ErrBarType"); err != nil {
		return err
	}
	if err := m.ErrValType.ValidateWithPath(path + "/ErrValType"); err != nil {
		return err
	}
	if m.NoEndCap != nil {
		if err := m.NoEndCap.ValidateWithPath(path + "/NoEndCap"); err != nil {
			return err
		}
	}
	if m.Plus != nil {
		if err := m.Plus.ValidateWithPath(path + "/Plus"); err != nil {
			return err
		}
	}
	if m.Minus != nil {
		if err := m.Minus.ValidateWithPath(path + "/Minus"); err != nil {
			return err
		}
	}
	if m.Val != nil {
		if err := m.Val.ValidateWithPath(path + "/Val"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
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
