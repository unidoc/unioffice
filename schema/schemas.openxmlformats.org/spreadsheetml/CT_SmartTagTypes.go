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

type CT_SmartTagTypes struct {
	// Smart Tag Type
	SmartTagType []*CT_SmartTagType
}

func NewCT_SmartTagTypes() *CT_SmartTagTypes {
	ret := &CT_SmartTagTypes{}
	return ret
}
func (m *CT_SmartTagTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.SmartTagType != nil {
		sesmartTagType := xml.StartElement{Name: xml.Name{Local: "x:smartTagType"}}
		e.EncodeElement(m.SmartTagType, sesmartTagType)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SmartTagTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SmartTagTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "smartTagType":
				tmp := NewCT_SmartTagType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SmartTagType = append(m.SmartTagType, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SmartTagTypes
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SmartTagTypes) Validate() error {
	return m.ValidateWithPath("CT_SmartTagTypes")
}
func (m *CT_SmartTagTypes) ValidateWithPath(path string) error {
	for i, v := range m.SmartTagType {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SmartTagType[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
