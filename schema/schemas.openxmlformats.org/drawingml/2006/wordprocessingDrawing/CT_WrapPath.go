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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_WrapPath struct {
	EditedAttr *bool
	Start      *drawingml.CT_Point2D
	LineTo     []*drawingml.CT_Point2D
}

func NewCT_WrapPath() *CT_WrapPath {
	ret := &CT_WrapPath{}
	ret.Start = drawingml.NewCT_Point2D()
	return ret
}
func (m *CT_WrapPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.EditedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "edited"},
			Value: fmt.Sprintf("%v", *m.EditedAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sestart := xml.StartElement{Name: xml.Name{Local: "wp:start"}}
	e.EncodeElement(m.Start, sestart)
	selineTo := xml.StartElement{Name: xml.Name{Local: "wp:lineTo"}}
	e.EncodeElement(m.LineTo, selineTo)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WrapPath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Start = drawingml.NewCT_Point2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "edited" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EditedAttr = &parsed
		}
	}
lCT_WrapPath:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "start":
				if err := d.DecodeElement(m.Start, &el); err != nil {
					return err
				}
			case "lineTo":
				tmp := drawingml.NewCT_Point2D()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LineTo = append(m.LineTo, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WrapPath
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_WrapPath) Validate() error {
	return m.ValidateWithPath("CT_WrapPath")
}
func (m *CT_WrapPath) ValidateWithPath(path string) error {
	if err := m.Start.ValidateWithPath(path + "/Start"); err != nil {
		return err
	}
	for i, v := range m.LineTo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LineTo[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
