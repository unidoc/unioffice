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

type CT_Cols struct {
	// Column Width & Formatting
	Col []*CT_Col
}

func NewCT_Cols() *CT_Cols {
	ret := &CT_Cols{}
	return ret
}
func (m *CT_Cols) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	secol := xml.StartElement{Name: xml.Name{Local: "x:col"}}
	e.EncodeElement(m.Col, secol)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Cols) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Cols:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "col":
				tmp := NewCT_Col()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Col = append(m.Col, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Cols
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Cols) Validate() error {
	return m.ValidateWithPath("CT_Cols")
}
func (m *CT_Cols) ValidateWithPath(path string) error {
	for i, v := range m.Col {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Col[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
