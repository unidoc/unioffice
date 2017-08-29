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

type CT_VolMain struct {
	// First String
	FirstAttr string
	// Topic
	Tp []*CT_VolTopic
}

func NewCT_VolMain() *CT_VolMain {
	ret := &CT_VolMain{}
	return ret
}
func (m *CT_VolMain) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "first"},
		Value: fmt.Sprintf("%v", m.FirstAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	setp := xml.StartElement{Name: xml.Name{Local: "x:tp"}}
	e.EncodeElement(m.Tp, setp)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_VolMain) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "first" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FirstAttr = parsed
		}
	}
lCT_VolMain:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tp":
				tmp := NewCT_VolTopic()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tp = append(m.Tp, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VolMain
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_VolMain) Validate() error {
	return m.ValidateWithPath("CT_VolMain")
}
func (m *CT_VolMain) ValidateWithPath(path string) error {
	for i, v := range m.Tp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tp[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
