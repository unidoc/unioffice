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

type CT_IndexedColors struct {
	// RGB Color
	RgbColor []*CT_RgbColor
}

func NewCT_IndexedColors() *CT_IndexedColors {
	ret := &CT_IndexedColors{}
	return ret
}
func (m *CT_IndexedColors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	sergbColor := xml.StartElement{Name: xml.Name{Local: "x:rgbColor"}}
	e.EncodeElement(m.RgbColor, sergbColor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_IndexedColors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_IndexedColors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rgbColor":
				tmp := NewCT_RgbColor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RgbColor = append(m.RgbColor, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_IndexedColors
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_IndexedColors) Validate() error {
	return m.ValidateWithPath("CT_IndexedColors")
}
func (m *CT_IndexedColors) ValidateWithPath(path string) error {
	for i, v := range m.RgbColor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RgbColor[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
