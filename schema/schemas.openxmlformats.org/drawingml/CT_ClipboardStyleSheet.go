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
)

type CT_ClipboardStyleSheet struct {
	ThemeElements *CT_BaseStyles
	ClrMap        *CT_ColorMapping
}

func NewCT_ClipboardStyleSheet() *CT_ClipboardStyleSheet {
	ret := &CT_ClipboardStyleSheet{}
	ret.ThemeElements = NewCT_BaseStyles()
	ret.ClrMap = NewCT_ColorMapping()
	return ret
}
func (m *CT_ClipboardStyleSheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "a:CT_ClipboardStyleSheet"
	e.EncodeToken(start)
	start.Attr = nil
	sethemeElements := xml.StartElement{Name: xml.Name{Local: "a:themeElements"}}
	e.EncodeElement(m.ThemeElements, sethemeElements)
	seclrMap := xml.StartElement{Name: xml.Name{Local: "a:clrMap"}}
	e.EncodeElement(m.ClrMap, seclrMap)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ClipboardStyleSheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ThemeElements = NewCT_BaseStyles()
	m.ClrMap = NewCT_ColorMapping()
lCT_ClipboardStyleSheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "themeElements":
				if err := d.DecodeElement(m.ThemeElements, &el); err != nil {
					return err
				}
			case "clrMap":
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ClipboardStyleSheet
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ClipboardStyleSheet) Validate() error {
	return m.ValidateWithPath("CT_ClipboardStyleSheet")
}
func (m *CT_ClipboardStyleSheet) ValidateWithPath(path string) error {
	if err := m.ThemeElements.ValidateWithPath(path + "/ThemeElements"); err != nil {
		return err
	}
	if err := m.ClrMap.ValidateWithPath(path + "/ClrMap"); err != nil {
		return err
	}
	return nil
}
