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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_WordprocessingShapeChoice struct {
	CNvSpPr *drawingml.CT_NonVisualDrawingShapeProps
	CNvCnPr *drawingml.CT_NonVisualConnectorProperties
}

func NewCT_WordprocessingShapeChoice() *CT_WordprocessingShapeChoice {
	ret := &CT_WordprocessingShapeChoice{}
	return ret
}

func (m *CT_WordprocessingShapeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CNvSpPr != nil {
		secNvSpPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvSpPr"}}
		e.EncodeElement(m.CNvSpPr, secNvSpPr)
	}
	if m.CNvCnPr != nil {
		secNvCnPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvCnPr"}}
		e.EncodeElement(m.CNvCnPr, secNvCnPr)
	}
	return nil
}

func (m *CT_WordprocessingShapeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WordprocessingShapeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvSpPr":
				m.CNvSpPr = drawingml.NewCT_NonVisualDrawingShapeProps()
				if err := d.DecodeElement(m.CNvSpPr, &el); err != nil {
					return err
				}
			case "cNvCnPr":
				m.CNvCnPr = drawingml.NewCT_NonVisualConnectorProperties()
				if err := d.DecodeElement(m.CNvCnPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_WordprocessingShapeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WordprocessingShapeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WordprocessingShapeChoice and its children
func (m *CT_WordprocessingShapeChoice) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingShapeChoice")
}

// ValidateWithPath validates the CT_WordprocessingShapeChoice and its children, prefixing error messages with path
func (m *CT_WordprocessingShapeChoice) ValidateWithPath(path string) error {
	if m.CNvSpPr != nil {
		if err := m.CNvSpPr.ValidateWithPath(path + "/CNvSpPr"); err != nil {
			return err
		}
	}
	if m.CNvCnPr != nil {
		if err := m.CNvCnPr.ValidateWithPath(path + "/CNvCnPr"); err != nil {
			return err
		}
	}
	return nil
}
