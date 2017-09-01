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

type CT_Placeholder struct {
	// Document Part Reference
	DocPart *CT_String
}

func NewCT_Placeholder() *CT_Placeholder {
	ret := &CT_Placeholder{}
	ret.DocPart = NewCT_String()
	return ret
}
func (m *CT_Placeholder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sedocPart := xml.StartElement{Name: xml.Name{Local: "w:docPart"}}
	e.EncodeElement(m.DocPart, sedocPart)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Placeholder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.DocPart = NewCT_String()
lCT_Placeholder:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "docPart":
				if err := d.DecodeElement(m.DocPart, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Placeholder
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Placeholder) Validate() error {
	return m.ValidateWithPath("CT_Placeholder")
}
func (m *CT_Placeholder) ValidateWithPath(path string) error {
	if err := m.DocPart.ValidateWithPath(path + "/DocPart"); err != nil {
		return err
	}
	return nil
}
