// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Shape struct {
	// Use Background Fill
	UseBgFillAttr *bool
	// Non-Visual Properties for a Shape
	NvSpPr *CT_ShapeNonVisual
	SpPr   *drawingml.CT_ShapeProperties
	// Shape Style
	Style *drawingml.CT_ShapeStyle
	// Shape Text Body
	TxBody *drawingml.CT_TextBody
	ExtLst *CT_ExtensionListModify
}

func NewCT_Shape() *CT_Shape {
	ret := &CT_Shape{}
	ret.NvSpPr = NewCT_ShapeNonVisual()
	ret.SpPr = drawingml.NewCT_ShapeProperties()
	return ret
}
func (m *CT_Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.UseBgFillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useBgFill"},
			Value: fmt.Sprintf("%v", *m.UseBgFillAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	senvSpPr := xml.StartElement{Name: xml.Name{Local: "p:nvSpPr"}}
	e.EncodeElement(m.NvSpPr, senvSpPr)
	sespPr := xml.StartElement{Name: xml.Name{Local: "p:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "p:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.TxBody != nil {
		setxBody := xml.StartElement{Name: xml.Name{Local: "p:txBody"}}
		e.EncodeElement(m.TxBody, setxBody)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvSpPr = NewCT_ShapeNonVisual()
	m.SpPr = drawingml.NewCT_ShapeProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "useBgFill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseBgFillAttr = &parsed
		}
	}
lCT_Shape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nvSpPr":
				if err := d.DecodeElement(m.NvSpPr, &el); err != nil {
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
			case "txBody":
				m.TxBody = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxBody, &el); err != nil {
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
			break lCT_Shape
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Shape) Validate() error {
	return m.ValidateWithPath("CT_Shape")
}
func (m *CT_Shape) ValidateWithPath(path string) error {
	if err := m.NvSpPr.ValidateWithPath(path + "/NvSpPr"); err != nil {
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
	if m.TxBody != nil {
		if err := m.TxBody.ValidateWithPath(path + "/TxBody"); err != nil {
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
