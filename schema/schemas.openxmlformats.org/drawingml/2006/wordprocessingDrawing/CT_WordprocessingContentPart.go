// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_WordprocessingContentPart struct {
	BwModeAttr      drawingml.ST_BlackWhiteMode
	IdAttr          string
	NvContentPartPr *CT_WordprocessingContentPartNonVisual
	Xfrm            *drawingml.CT_Transform2D
	ExtLst          *drawingml.CT_OfficeArtExtensionList
}

func NewCT_WordprocessingContentPart() *CT_WordprocessingContentPart {
	ret := &CT_WordprocessingContentPart{}
	return ret
}
func (m *CT_WordprocessingContentPart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BwModeAttr != drawingml.ST_BlackWhiteModeUnset {
		attr, err := m.BwModeAttr.MarshalXMLAttr(xml.Name{Local: "bwMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.NvContentPartPr != nil {
		senvContentPartPr := xml.StartElement{Name: xml.Name{Local: "wp:nvContentPartPr"}}
		e.EncodeElement(m.NvContentPartPr, senvContentPartPr)
	}
	if m.Xfrm != nil {
		sexfrm := xml.StartElement{Name: xml.Name{Local: "wp:xfrm"}}
		e.EncodeElement(m.Xfrm, sexfrm)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WordprocessingContentPart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bwMode" {
			m.BwModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
lCT_WordprocessingContentPart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nvContentPartPr":
				m.NvContentPartPr = NewCT_WordprocessingContentPartNonVisual()
				if err := d.DecodeElement(m.NvContentPartPr, &el); err != nil {
					return err
				}
			case "xfrm":
				m.Xfrm = drawingml.NewCT_Transform2D()
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
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
			break lCT_WordprocessingContentPart
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_WordprocessingContentPart) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingContentPart")
}
func (m *CT_WordprocessingContentPart) ValidateWithPath(path string) error {
	if err := m.BwModeAttr.ValidateWithPath(path + "/BwModeAttr"); err != nil {
		return err
	}
	if m.NvContentPartPr != nil {
		if err := m.NvContentPartPr.ValidateWithPath(path + "/NvContentPartPr"); err != nil {
			return err
		}
	}
	if m.Xfrm != nil {
		if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
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
