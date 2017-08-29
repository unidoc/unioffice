// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_DocumentBase struct {
	// Document Background
	Background *CT_Background
}

func NewCT_DocumentBase() *CT_DocumentBase {
	ret := &CT_DocumentBase{}
	return ret
}
func (m *CT_DocumentBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Background != nil {
		sebackground := xml.StartElement{Name: xml.Name{Local: "w:background"}}
		e.EncodeElement(m.Background, sebackground)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DocumentBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocumentBase:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "background":
				m.Background = NewCT_Background()
				if err := d.DecodeElement(m.Background, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocumentBase
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DocumentBase) Validate() error {
	return m.ValidateWithPath("CT_DocumentBase")
}
func (m *CT_DocumentBase) ValidateWithPath(path string) error {
	if m.Background != nil {
		if err := m.Background.ValidateWithPath(path + "/Background"); err != nil {
			return err
		}
	}
	return nil
}
