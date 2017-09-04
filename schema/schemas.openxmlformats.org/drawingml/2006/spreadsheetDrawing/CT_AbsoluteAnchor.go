// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_AbsoluteAnchor struct {
	Pos        *drawingml.CT_Point2D
	Ext        *drawingml.CT_PositiveSize2D
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

func NewCT_AbsoluteAnchor() *CT_AbsoluteAnchor {
	ret := &CT_AbsoluteAnchor{}
	ret.Pos = drawingml.NewCT_Point2D()
	ret.Ext = drawingml.NewCT_PositiveSize2D()
	ret.ClientData = NewCT_AnchorClientData()
	return ret
}

func (m *CT_AbsoluteAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sepos := xml.StartElement{Name: xml.Name{Local: "xdr:pos"}}
	e.EncodeElement(m.Pos, sepos)
	seext := xml.StartElement{Name: xml.Name{Local: "xdr:ext"}}
	e.EncodeElement(m.Ext, seext)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	seclientData := xml.StartElement{Name: xml.Name{Local: "xdr:clientData"}}
	e.EncodeElement(m.ClientData, seclientData)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AbsoluteAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pos = drawingml.NewCT_Point2D()
	m.Ext = drawingml.NewCT_PositiveSize2D()
	m.ClientData = NewCT_AnchorClientData()
lCT_AbsoluteAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pos":
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case "ext":
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			case "sp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Sp, &el); err != nil {
					return err
				}
			case "grpSp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GrpSp, &el); err != nil {
					return err
				}
			case "graphicFrame":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GraphicFrame, &el); err != nil {
					return err
				}
			case "cxnSp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.CxnSp, &el); err != nil {
					return err
				}
			case "pic":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Pic, &el); err != nil {
					return err
				}
			case "contentPart":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.ContentPart, &el); err != nil {
					return err
				}
			case "clientData":
				if err := d.DecodeElement(m.ClientData, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_AbsoluteAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AbsoluteAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AbsoluteAnchor and its children
func (m *CT_AbsoluteAnchor) Validate() error {
	return m.ValidateWithPath("CT_AbsoluteAnchor")
}

// ValidateWithPath validates the CT_AbsoluteAnchor and its children, prefixing error messages with path
func (m *CT_AbsoluteAnchor) ValidateWithPath(path string) error {
	if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
		return err
	}
	if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if err := m.ClientData.ValidateWithPath(path + "/ClientData"); err != nil {
		return err
	}
	return nil
}
