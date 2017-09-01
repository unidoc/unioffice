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

type CT_Divs struct {
	// Information About Single HTML div Element
	Div []*CT_Div
}

func NewCT_Divs() *CT_Divs {
	ret := &CT_Divs{}
	return ret
}
func (m *CT_Divs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sediv := xml.StartElement{Name: xml.Name{Local: "w:div"}}
	e.EncodeElement(m.Div, sediv)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Divs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Divs:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "div":
				tmp := NewCT_Div()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Div = append(m.Div, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Divs
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Divs) Validate() error {
	return m.ValidateWithPath("CT_Divs")
}
func (m *CT_Divs) ValidateWithPath(path string) error {
	for i, v := range m.Div {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Div[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
