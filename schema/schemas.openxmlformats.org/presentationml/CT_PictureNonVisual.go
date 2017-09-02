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

type CT_PictureNonVisual struct {
	CNvPr *drawingml.CT_NonVisualDrawingProps
	// Non-Visual Picture Drawing Properties
	CNvPicPr *drawingml.CT_NonVisualPictureProperties
	NvPr     *CT_ApplicationNonVisualDrawingProps
}

func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	ret := &CT_PictureNonVisual{}
	ret.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.CNvPicPr = drawingml.NewCT_NonVisualPictureProperties()
	ret.NvPr = NewCT_ApplicationNonVisualDrawingProps()
	return ret
}

func (m *CT_PictureNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "p:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvPicPr := xml.StartElement{Name: xml.Name{Local: "p:cNvPicPr"}}
	e.EncodeElement(m.CNvPicPr, secNvPicPr)
	senvPr := xml.StartElement{Name: xml.Name{Local: "p:nvPr"}}
	e.EncodeElement(m.NvPr, senvPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PictureNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	m.CNvPicPr = drawingml.NewCT_NonVisualPictureProperties()
	m.NvPr = NewCT_ApplicationNonVisualDrawingProps()
lCT_PictureNonVisual:
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
			case "cNvPicPr":
				if err := d.DecodeElement(m.CNvPicPr, &el); err != nil {
					return err
				}
			case "nvPr":
				if err := d.DecodeElement(m.NvPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PictureNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PictureNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PictureNonVisual and its children
func (m *CT_PictureNonVisual) Validate() error {
	return m.ValidateWithPath("CT_PictureNonVisual")
}

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (m *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvPicPr.ValidateWithPath(path + "/CNvPicPr"); err != nil {
		return err
	}
	if err := m.NvPr.ValidateWithPath(path + "/NvPr"); err != nil {
		return err
	}
	return nil
}
