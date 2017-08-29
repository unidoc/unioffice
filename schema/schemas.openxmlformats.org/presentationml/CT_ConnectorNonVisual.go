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

type CT_ConnectorNonVisual struct {
	// Non-Visual Drawing Properties
	CNvPr *drawingml.CT_NonVisualDrawingProps
	// Non-Visual Connector Shape Drawing Properties
	CNvCxnSpPr *drawingml.CT_NonVisualConnectorProperties
	// Application Non-Visual Drawing Properties
	NvPr *CT_ApplicationNonVisualDrawingProps
}

func NewCT_ConnectorNonVisual() *CT_ConnectorNonVisual {
	ret := &CT_ConnectorNonVisual{}
	ret.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.CNvCxnSpPr = drawingml.NewCT_NonVisualConnectorProperties()
	ret.NvPr = NewCT_ApplicationNonVisualDrawingProps()
	return ret
}
func (m *CT_ConnectorNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	secNvPr := xml.StartElement{Name: xml.Name{Local: "p:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "p:cNvCxnSpPr"}}
	e.EncodeElement(m.CNvCxnSpPr, secNvCxnSpPr)
	senvPr := xml.StartElement{Name: xml.Name{Local: "p:nvPr"}}
	e.EncodeElement(m.NvPr, senvPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ConnectorNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	m.CNvCxnSpPr = drawingml.NewCT_NonVisualConnectorProperties()
	m.NvPr = NewCT_ApplicationNonVisualDrawingProps()
lCT_ConnectorNonVisual:
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
			case "cNvCxnSpPr":
				if err := d.DecodeElement(m.CNvCxnSpPr, &el); err != nil {
					return err
				}
			case "nvPr":
				if err := d.DecodeElement(m.NvPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConnectorNonVisual
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ConnectorNonVisual) Validate() error {
	return m.ValidateWithPath("CT_ConnectorNonVisual")
}
func (m *CT_ConnectorNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvCxnSpPr.ValidateWithPath(path + "/CNvCxnSpPr"); err != nil {
		return err
	}
	if err := m.NvPr.ValidateWithPath(path + "/NvPr"); err != nil {
		return err
	}
	return nil
}
