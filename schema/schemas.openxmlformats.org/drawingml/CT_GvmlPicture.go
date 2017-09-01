// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_GvmlPicture struct {
	NvPicPr  *CT_GvmlPictureNonVisual
	BlipFill *CT_BlipFillProperties
	SpPr     *CT_ShapeProperties
	Style    *CT_ShapeStyle
	ExtLst   *CT_OfficeArtExtensionList
}

func NewCT_GvmlPicture() *CT_GvmlPicture {
	ret := &CT_GvmlPicture{}
	ret.NvPicPr = NewCT_GvmlPictureNonVisual()
	ret.BlipFill = NewCT_BlipFillProperties()
	ret.SpPr = NewCT_ShapeProperties()
	return ret
}
func (m *CT_GvmlPicture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	senvPicPr := xml.StartElement{Name: xml.Name{Local: "a:nvPicPr"}}
	e.EncodeElement(m.NvPicPr, senvPicPr)
	seblipFill := xml.StartElement{Name: xml.Name{Local: "a:blipFill"}}
	e.EncodeElement(m.BlipFill, seblipFill)
	sespPr := xml.StartElement{Name: xml.Name{Local: "a:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "a:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GvmlPicture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvPicPr = NewCT_GvmlPictureNonVisual()
	m.BlipFill = NewCT_BlipFillProperties()
	m.SpPr = NewCT_ShapeProperties()
lCT_GvmlPicture:
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
			case "style":
				m.Style = NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_GvmlPicture
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GvmlPicture) Validate() error {
	return m.ValidateWithPath("CT_GvmlPicture")
}
func (m *CT_GvmlPicture) ValidateWithPath(path string) error {
	if err := m.NvPicPr.ValidateWithPath(path + "/NvPicPr"); err != nil {
		return err
	}
	if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
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
