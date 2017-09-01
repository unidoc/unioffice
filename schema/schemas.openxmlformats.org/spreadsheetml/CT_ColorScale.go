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
)

type CT_ColorScale struct {
	// Conditional Format Value Object
	Cfvo []*CT_Cfvo
	// Color Gradiant Interpolation
	Color []*CT_Color
}

func NewCT_ColorScale() *CT_ColorScale {
	ret := &CT_ColorScale{}
	return ret
}
func (m *CT_ColorScale) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secfvo := xml.StartElement{Name: xml.Name{Local: "x:cfvo"}}
	e.EncodeElement(m.Cfvo, secfvo)
	secolor := xml.StartElement{Name: xml.Name{Local: "x:color"}}
	e.EncodeElement(m.Color, secolor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ColorScale) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorScale:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cfvo":
				tmp := NewCT_Cfvo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cfvo = append(m.Cfvo, tmp)
			case "color":
				tmp := NewCT_Color()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Color = append(m.Color, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorScale
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ColorScale) Validate() error {
	return m.ValidateWithPath("CT_ColorScale")
}
func (m *CT_ColorScale) ValidateWithPath(path string) error {
	for i, v := range m.Cfvo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cfvo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Color {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Color[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
