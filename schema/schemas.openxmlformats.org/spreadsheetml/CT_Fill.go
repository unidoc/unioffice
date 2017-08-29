// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_Fill struct {
	// Pattern
	PatternFill *CT_PatternFill
	// Gradient
	GradientFill *CT_GradientFill
}

func NewCT_Fill() *CT_Fill {
	ret := &CT_Fill{}
	return ret
}
func (m *CT_Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.PatternFill != nil {
		sepatternFill := xml.StartElement{Name: xml.Name{Local: "x:patternFill"}}
		e.EncodeElement(m.PatternFill, sepatternFill)
	}
	if m.GradientFill != nil {
		segradientFill := xml.StartElement{Name: xml.Name{Local: "x:gradientFill"}}
		e.EncodeElement(m.GradientFill, segradientFill)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Fill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Fill:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "patternFill":
				m.PatternFill = NewCT_PatternFill()
				if err := d.DecodeElement(m.PatternFill, &el); err != nil {
					return err
				}
			case "gradientFill":
				m.GradientFill = NewCT_GradientFill()
				if err := d.DecodeElement(m.GradientFill, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Fill
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Fill) Validate() error {
	return m.ValidateWithPath("CT_Fill")
}
func (m *CT_Fill) ValidateWithPath(path string) error {
	if m.PatternFill != nil {
		if err := m.PatternFill.ValidateWithPath(path + "/PatternFill"); err != nil {
			return err
		}
	}
	if m.GradientFill != nil {
		if err := m.GradientFill.ValidateWithPath(path + "/GradientFill"); err != nil {
			return err
		}
	}
	return nil
}
