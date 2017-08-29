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

type CT_Connections struct {
	// Connection
	Connection []*CT_Connection
}

func NewCT_Connections() *CT_Connections {
	ret := &CT_Connections{}
	return ret
}
func (m *CT_Connections) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seconnection := xml.StartElement{Name: xml.Name{Local: "x:connection"}}
	e.EncodeElement(m.Connection, seconnection)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Connections) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Connections:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "connection":
				tmp := NewCT_Connection()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Connection = append(m.Connection, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Connections
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Connections) Validate() error {
	return m.ValidateWithPath("CT_Connections")
}
func (m *CT_Connections) ValidateWithPath(path string) error {
	for i, v := range m.Connection {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Connection[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
