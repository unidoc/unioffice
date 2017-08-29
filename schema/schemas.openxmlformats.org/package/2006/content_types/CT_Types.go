// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package content_types

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Types struct {
	Default  []*Default
	Override []*Override
}

func NewCT_Types() *CT_Types {
	ret := &CT_Types{}
	return ret
}
func (m *CT_Types) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Default != nil {
		seDefault := xml.StartElement{Name: xml.Name{Local: "Default"}}
		e.EncodeElement(m.Default, seDefault)
	}
	if m.Override != nil {
		seOverride := xml.StartElement{Name: xml.Name{Local: "Override"}}
		e.EncodeElement(m.Override, seOverride)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Types) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Types:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "Default":
				tmp := NewDefault()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Default = append(m.Default, tmp)
			case "Override":
				tmp := NewOverride()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Override = append(m.Override, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Types
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Types) Validate() error {
	return m.ValidateWithPath("CT_Types")
}
func (m *CT_Types) ValidateWithPath(path string) error {
	for i, v := range m.Default {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Default[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Override {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Override[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
