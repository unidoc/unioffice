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

	"baliance.com/gooxml"
)

type CT_RegularTextRun struct {
	RPr *CT_TextCharacterProperties
	T   string
}

func NewCT_RegularTextRun() *CT_RegularTextRun {
	ret := &CT_RegularTextRun{}
	return ret
}
func (m *CT_RegularTextRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "a:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	set := xml.StartElement{Name: xml.Name{Local: "a:t"}}
	gooxml.AddPreserveSpaceAttr(&set, m.T)
	e.EncodeElement(m.T, set)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RegularTextRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RegularTextRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rPr":
				m.RPr = NewCT_TextCharacterProperties()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case "t":
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RegularTextRun
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_RegularTextRun) Validate() error {
	return m.ValidateWithPath("CT_RegularTextRun")
}
func (m *CT_RegularTextRun) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
