// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_DocParts struct {
	// Glossary Document Entry
	DocPart []*CT_DocPart
}

func NewCT_DocParts() *CT_DocParts {
	ret := &CT_DocParts{}
	return ret
}
func (m *CT_DocParts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.DocPart != nil {
		sedocPart := xml.StartElement{Name: xml.Name{Local: "w:docPart"}}
		e.EncodeElement(m.DocPart, sedocPart)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DocParts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocParts:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "docPart":
				tmp := NewCT_DocPart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DocPart = append(m.DocPart, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocParts
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DocParts) Validate() error {
	return m.ValidateWithPath("CT_DocParts")
}
func (m *CT_DocParts) ValidateWithPath(path string) error {
	for i, v := range m.DocPart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DocPart[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
