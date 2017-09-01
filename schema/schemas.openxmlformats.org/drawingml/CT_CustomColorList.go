// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_CustomColorList struct {
	CustClr []*CT_CustomColor
}

func NewCT_CustomColorList() *CT_CustomColorList {
	ret := &CT_CustomColorList{}
	return ret
}
func (m *CT_CustomColorList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.CustClr != nil {
		secustClr := xml.StartElement{Name: xml.Name{Local: "a:custClr"}}
		e.EncodeElement(m.CustClr, secustClr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CustomColorList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomColorList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "custClr":
				tmp := NewCT_CustomColor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustClr = append(m.CustClr, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomColorList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CustomColorList) Validate() error {
	return m.ValidateWithPath("CT_CustomColorList")
}
func (m *CT_CustomColorList) ValidateWithPath(path string) error {
	for i, v := range m.CustClr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustClr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
