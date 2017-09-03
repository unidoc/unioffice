// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type Wsp struct {
	CT_WordprocessingShape
}

func NewWsp() *Wsp {
	ret := &Wsp{}
	ret.CT_WordprocessingShape = *NewCT_WordprocessingShape()
	return ret
}

func (m *Wsp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_WordprocessingShape.MarshalXML(e, start)
}

func (m *Wsp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_WordprocessingShape = *NewCT_WordprocessingShape()
	for _, attr := range start.Attr {
		if attr.Name.Local == "normalEastAsianFlow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NormalEastAsianFlowAttr = &parsed
		}
	}
lWsp:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvSpPr":
				m.Choice = NewCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvSpPr, &el); err != nil {
					return err
				}
			case "cNvCnPr":
				m.Choice = NewCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvCnPr, &el); err != nil {
					return err
				}
			case "spPr":
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "style":
				m.Style = drawingml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			case "txbx":
				m.CChoice = NewCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.CChoice.Txbx, &el); err != nil {
					return err
				}
			case "linkedTxbx":
				m.CChoice = NewCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.CChoice.LinkedTxbx, &el); err != nil {
					return err
				}
			case "bodyPr":
				if err := d.DecodeElement(m.BodyPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Wsp %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWsp
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Wsp and its children
func (m *Wsp) Validate() error {
	return m.ValidateWithPath("Wsp")
}

// ValidateWithPath validates the Wsp and its children, prefixing error messages with path
func (m *Wsp) ValidateWithPath(path string) error {
	if err := m.CT_WordprocessingShape.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
