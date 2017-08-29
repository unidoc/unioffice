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

type CT_I struct {
	// Item Type
	TAttr ST_ItemType
	// Repeated Items Count
	RAttr *uint32
	// Data Field Index
	IAttr *uint32
	// Row / Column Item Index
	X []*CT_X
}

func NewCT_I() *CT_I {
	ret := &CT_I{}
	return ret
}
func (m *CT_I) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TAttr != ST_ItemTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
			Value: fmt.Sprintf("%v", *m.RAttr)})
	}
	if m.IAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i"},
			Value: fmt.Sprintf("%v", *m.IAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "x:x"}}
		e.EncodeElement(m.X, sex)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_I) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "r" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.RAttr = &pt
		}
		if attr.Name.Local == "i" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.IAttr = &pt
		}
	}
lCT_I:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "x":
				tmp := NewCT_X()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_I
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_I) Validate() error {
	return m.ValidateWithPath("CT_I")
}
func (m *CT_I) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
