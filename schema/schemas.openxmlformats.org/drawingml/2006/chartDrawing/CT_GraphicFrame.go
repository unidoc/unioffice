// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_GraphicFrame struct {
	MacroAttr        *string
	FPublishedAttr   *bool
	NvGraphicFramePr *CT_GraphicFrameNonVisual
	Xfrm             *drawingml.CT_Transform2D
	Graphic          *drawingml.Graphic
}

func NewCT_GraphicFrame() *CT_GraphicFrame {
	ret := &CT_GraphicFrame{}
	ret.NvGraphicFramePr = NewCT_GraphicFrameNonVisual()
	ret.Xfrm = drawingml.NewCT_Transform2D()
	ret.Graphic = drawingml.NewGraphic()
	return ret
}
func (m *CT_GraphicFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.MacroAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "macro"},
			Value: fmt.Sprintf("%v", *m.MacroAttr)})
	}
	if m.FPublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fPublished"},
			Value: fmt.Sprintf("%v", *m.FPublishedAttr)})
	}
	e.EncodeToken(start)
	senvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "nvGraphicFramePr"}}
	e.EncodeElement(m.NvGraphicFramePr, senvGraphicFramePr)
	sexfrm := xml.StartElement{Name: xml.Name{Local: "xfrm"}}
	e.EncodeElement(m.Xfrm, sexfrm)
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GraphicFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvGraphicFramePr = NewCT_GraphicFrameNonVisual()
	m.Xfrm = drawingml.NewCT_Transform2D()
	m.Graphic = drawingml.NewGraphic()
	for _, attr := range start.Attr {
		if attr.Name.Local == "macro" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MacroAttr = &parsed
		}
		if attr.Name.Local == "fPublished" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FPublishedAttr = &parsed
		}
	}
lCT_GraphicFrame:
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
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicFrame
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GraphicFrame) Validate() error {
	return m.ValidateWithPath("CT_GraphicFrame")
}
func (m *CT_GraphicFrame) ValidateWithPath(path string) error {
	if err := m.NvGraphicFramePr.ValidateWithPath(path + "/NvGraphicFramePr"); err != nil {
		return err
	}
	if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
		return err
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	return nil
}
