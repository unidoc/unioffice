// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_GradientStop struct {
	// Gradient Stop Position
	PositionAttr float64
	// Color
	Color *CT_Color
}

func NewCT_GradientStop() *CT_GradientStop {
	ret := &CT_GradientStop{}
	ret.Color = NewCT_Color()
	return ret
}
func (m *CT_GradientStop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "position"},
		Value: fmt.Sprintf("%v", m.PositionAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	secolor := xml.StartElement{Name: xml.Name{Local: "x:color"}}
	e.EncodeElement(m.Color, secolor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GradientStop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Color = NewCT_Color()
	for _, attr := range start.Attr {
		if attr.Name.Local == "position" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.PositionAttr = parsed
		}
	}
lCT_GradientStop:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "color":
				if err := d.DecodeElement(m.Color, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GradientStop
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GradientStop) Validate() error {
	return m.ValidateWithPath("CT_GradientStop")
}
func (m *CT_GradientStop) ValidateWithPath(path string) error {
	if err := m.Color.ValidateWithPath(path + "/Color"); err != nil {
		return err
	}
	return nil
}
