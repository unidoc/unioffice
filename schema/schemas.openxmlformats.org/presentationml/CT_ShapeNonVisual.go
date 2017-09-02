// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_ShapeNonVisual struct {
	// Non-Visual Drawing Properties
	CNvPr *drawingml.CT_NonVisualDrawingProps
	// Non-Visual Drawing Properties for a Shape
	CNvSpPr *drawingml.CT_NonVisualDrawingShapeProps
	// Application Non-Visual Drawing Properties
	NvPr *CT_ApplicationNonVisualDrawingProps
}

func NewCT_ShapeNonVisual() *CT_ShapeNonVisual {
	ret := &CT_ShapeNonVisual{}
	ret.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.CNvSpPr = drawingml.NewCT_NonVisualDrawingShapeProps()
	ret.NvPr = NewCT_ApplicationNonVisualDrawingProps()
	return ret
}

func (m *CT_ShapeNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "p:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvSpPr := xml.StartElement{Name: xml.Name{Local: "p:cNvSpPr"}}
	e.EncodeElement(m.CNvSpPr, secNvSpPr)
	senvPr := xml.StartElement{Name: xml.Name{Local: "p:nvPr"}}
	e.EncodeElement(m.NvPr, senvPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShapeNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	m.CNvSpPr = drawingml.NewCT_NonVisualDrawingShapeProps()
	m.NvPr = NewCT_ApplicationNonVisualDrawingProps()
lCT_ShapeNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvSpPr":
				if err := d.DecodeElement(m.CNvSpPr, &el); err != nil {
					return err
				}
			case "nvPr":
				if err := d.DecodeElement(m.NvPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ShapeNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ShapeNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ShapeNonVisual and its children
func (m *CT_ShapeNonVisual) Validate() error {
	return m.ValidateWithPath("CT_ShapeNonVisual")
}

// ValidateWithPath validates the CT_ShapeNonVisual and its children, prefixing error messages with path
func (m *CT_ShapeNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvSpPr.ValidateWithPath(path + "/CNvSpPr"); err != nil {
		return err
	}
	if err := m.NvPr.ValidateWithPath(path + "/NvPr"); err != nil {
		return err
	}
	return nil
}
