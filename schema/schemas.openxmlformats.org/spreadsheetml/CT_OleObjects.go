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

type CT_OleObjects struct {
	// Embedded Object
	OleObject []*CT_OleObject
}

func NewCT_OleObjects() *CT_OleObjects {
	ret := &CT_OleObjects{}
	return ret
}
func (m *CT_OleObjects) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seoleObject := xml.StartElement{Name: xml.Name{Local: "x:oleObject"}}
	e.EncodeElement(m.OleObject, seoleObject)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OleObjects) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OleObjects:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "oleObject":
				tmp := NewCT_OleObject()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.OleObject = append(m.OleObject, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObjects
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_OleObjects) Validate() error {
	return m.ValidateWithPath("CT_OleObjects")
}
func (m *CT_OleObjects) ValidateWithPath(path string) error {
	for i, v := range m.OleObject {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/OleObject[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
