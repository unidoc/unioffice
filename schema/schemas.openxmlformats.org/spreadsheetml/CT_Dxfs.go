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
	"strconv"
)

type CT_Dxfs struct {
	// Format Count
	CountAttr *uint32
	// Formatting
	Dxf []*CT_Dxf
}

func NewCT_Dxfs() *CT_Dxfs {
	ret := &CT_Dxfs{}
	return ret
}
func (m *CT_Dxfs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Dxf != nil {
		sedxf := xml.StartElement{Name: xml.Name{Local: "x:dxf"}}
		e.EncodeElement(m.Dxf, sedxf)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Dxfs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_Dxfs:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dxf":
				tmp := NewCT_Dxf()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Dxf = append(m.Dxf, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Dxfs
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Dxfs) Validate() error {
	return m.ValidateWithPath("CT_Dxfs")
}
func (m *CT_Dxfs) ValidateWithPath(path string) error {
	for i, v := range m.Dxf {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Dxf[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
