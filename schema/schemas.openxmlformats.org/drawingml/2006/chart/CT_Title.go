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

type CT_Title struct {
	Tx      *CT_Tx
	Layout  *CT_Layout
	Overlay *CT_Boolean
	SpPr    *drawingml.CT_ShapeProperties
	TxPr    *drawingml.CT_TextBody
	ExtLst  *CT_ExtensionList
}

func NewCT_Title() *CT_Title {
	ret := &CT_Title{}
	return ret
}

func (m *CT_Title) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Tx != nil {
		setx := xml.StartElement{Name: xml.Name{Local: "tx"}}
		e.EncodeElement(m.Tx, setx)
	}
	if m.Layout != nil {
		selayout := xml.StartElement{Name: xml.Name{Local: "layout"}}
		e.EncodeElement(m.Layout, selayout)
	}
	if m.Overlay != nil {
		seoverlay := xml.StartElement{Name: xml.Name{Local: "overlay"}}
		e.EncodeElement(m.Overlay, seoverlay)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Title) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Title:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tx":
				m.Tx = NewCT_Tx()
				if err := d.DecodeElement(m.Tx, &el); err != nil {
					return err
				}
			case "layout":
				m.Layout = NewCT_Layout()
				if err := d.DecodeElement(m.Layout, &el); err != nil {
					return err
				}
			case "overlay":
				m.Overlay = NewCT_Boolean()
				if err := d.DecodeElement(m.Overlay, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "txPr":
				m.TxPr = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Title %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Title
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Title and its children
func (m *CT_Title) Validate() error {
	return m.ValidateWithPath("CT_Title")
}

// ValidateWithPath validates the CT_Title and its children, prefixing error messages with path
func (m *CT_Title) ValidateWithPath(path string) error {
	if m.Tx != nil {
		if err := m.Tx.ValidateWithPath(path + "/Tx"); err != nil {
			return err
		}
	}
	if m.Layout != nil {
		if err := m.Layout.ValidateWithPath(path + "/Layout"); err != nil {
			return err
		}
	}
	if m.Overlay != nil {
		if err := m.Overlay.ValidateWithPath(path + "/Overlay"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
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
