// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_ExternalData struct {
	IdAttr     string
	AutoUpdate *CT_Boolean
}

func NewCT_ExternalData() *CT_ExternalData {
	ret := &CT_ExternalData{}
	return ret
}
func (m *CT_ExternalData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.AutoUpdate != nil {
		seautoUpdate := xml.StartElement{Name: xml.Name{Local: "autoUpdate"}}
		e.EncodeElement(m.AutoUpdate, seautoUpdate)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ExternalData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
lCT_ExternalData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "autoUpdate":
				m.AutoUpdate = NewCT_Boolean()
				if err := d.DecodeElement(m.AutoUpdate, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalData
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ExternalData) Validate() error {
	return m.ValidateWithPath("CT_ExternalData")
}
func (m *CT_ExternalData) ValidateWithPath(path string) error {
	if m.AutoUpdate != nil {
		if err := m.AutoUpdate.ValidateWithPath(path + "/AutoUpdate"); err != nil {
			return err
		}
	}
	return nil
}
