// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package picture

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Picture struct {
	NvPicPr  *CT_PictureNonVisual
	BlipFill *drawingml.CT_BlipFillProperties
	SpPr     *drawingml.CT_ShapeProperties
}

func NewCT_Picture() *CT_Picture {
	ret := &CT_Picture{}
	ret.NvPicPr = NewCT_PictureNonVisual()
	ret.BlipFill = drawingml.NewCT_BlipFillProperties()
	ret.SpPr = drawingml.NewCT_ShapeProperties()
	return ret
}

func (m *CT_Picture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	senvPicPr := xml.StartElement{Name: xml.Name{Local: "pic:nvPicPr"}}
	e.EncodeElement(m.NvPicPr, senvPicPr)
	seblipFill := xml.StartElement{Name: xml.Name{Local: "pic:blipFill"}}
	e.EncodeElement(m.BlipFill, seblipFill)
	sespPr := xml.StartElement{Name: xml.Name{Local: "pic:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Picture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvPicPr = NewCT_PictureNonVisual()
	m.BlipFill = drawingml.NewCT_BlipFillProperties()
	m.SpPr = drawingml.NewCT_ShapeProperties()
lCT_Picture:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nvPicPr":
				if err := d.DecodeElement(m.NvPicPr, &el); err != nil {
					return err
				}
			case "blipFill":
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case "spPr":
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Picture
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Picture and its children
func (m *CT_Picture) Validate() error {
	return m.ValidateWithPath("CT_Picture")
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (m *CT_Picture) ValidateWithPath(path string) error {
	if err := m.NvPicPr.ValidateWithPath(path + "/NvPicPr"); err != nil {
		return err
	}
	if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	return nil
}
