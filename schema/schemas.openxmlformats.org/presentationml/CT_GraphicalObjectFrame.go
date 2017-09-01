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

type CT_GraphicalObjectFrame struct {
	BwModeAttr drawingml.ST_BlackWhiteMode
	// Non-Visual Properties for a Graphic Frame
	NvGraphicFramePr *CT_GraphicalObjectFrameNonVisual
	// 2D Transform for Graphic Frame
	Xfrm    *drawingml.CT_Transform2D
	Graphic *drawingml.Graphic
	// Extension List with Modification Flag
	ExtLst *CT_ExtensionListModify
}

func NewCT_GraphicalObjectFrame() *CT_GraphicalObjectFrame {
	ret := &CT_GraphicalObjectFrame{}
	ret.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	ret.Xfrm = drawingml.NewCT_Transform2D()
	ret.Graphic = drawingml.NewGraphic()
	return ret
}
func (m *CT_GraphicalObjectFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	e.EncodeToken(start)
	senvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "p:nvGraphicFramePr"}}
	e.EncodeElement(m.NvGraphicFramePr, senvGraphicFramePr)
	sexfrm := xml.StartElement{Name: xml.Name{Local: "p:xfrm"}}
	e.EncodeElement(m.Xfrm, sexfrm)
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GraphicalObjectFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	m.Xfrm = drawingml.NewCT_Transform2D()
	m.Graphic = drawingml.NewGraphic()
	for _, attr := range start.Attr {
		if attr.Name.Local == "bwMode" {
			m.BwModeAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_GraphicalObjectFrame:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nvGraphicFramePr":
				if err := d.DecodeElement(m.NvGraphicFramePr, &el); err != nil {
					return err
				}
			case "xfrm":
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case "graphic":
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
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
			break lCT_GraphicalObjectFrame
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GraphicalObjectFrame) Validate() error {
	return m.ValidateWithPath("CT_GraphicalObjectFrame")
}
func (m *CT_GraphicalObjectFrame) ValidateWithPath(path string) error {
	if err := m.BwModeAttr.ValidateWithPath(path + "/BwModeAttr"); err != nil {
		return err
	}
	if err := m.NvGraphicFramePr.ValidateWithPath(path + "/NvGraphicFramePr"); err != nil {
		return err
	}
	if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
		return err
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
