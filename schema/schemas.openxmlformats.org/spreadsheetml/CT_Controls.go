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

type CT_Controls struct {
	// Embedded Control
	Control []*CT_Control
}

func NewCT_Controls() *CT_Controls {
	ret := &CT_Controls{}
	return ret
}
func (m *CT_Controls) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secontrol := xml.StartElement{Name: xml.Name{Local: "x:control"}}
	e.EncodeElement(m.Control, secontrol)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Controls) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Controls:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "control":
				tmp := NewCT_Control()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Control = append(m.Control, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Controls
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Controls) Validate() error {
	return m.ValidateWithPath("CT_Controls")
}
func (m *CT_Controls) ValidateWithPath(path string) error {
	for i, v := range m.Control {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Control[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
