// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CT_TLTimeNodeExclusive struct {
	// Common TimeNode Properties
	CTn *CT_TLCommonTimeNodeData
}

func NewCT_TLTimeNodeExclusive() *CT_TLTimeNodeExclusive {
	ret := &CT_TLTimeNodeExclusive{}
	ret.CTn = NewCT_TLCommonTimeNodeData()
	return ret
}
func (m *CT_TLTimeNodeExclusive) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secTn := xml.StartElement{Name: xml.Name{Local: "p:cTn"}}
	e.EncodeElement(m.CTn, secTn)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLTimeNodeExclusive) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CTn = NewCT_TLCommonTimeNodeData()
lCT_TLTimeNodeExclusive:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cTn":
				if err := d.DecodeElement(m.CTn, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeNodeExclusive
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLTimeNodeExclusive) Validate() error {
	return m.ValidateWithPath("CT_TLTimeNodeExclusive")
}
func (m *CT_TLTimeNodeExclusive) ValidateWithPath(path string) error {
	if err := m.CTn.ValidateWithPath(path + "/CTn"); err != nil {
		return err
	}
	return nil
}
