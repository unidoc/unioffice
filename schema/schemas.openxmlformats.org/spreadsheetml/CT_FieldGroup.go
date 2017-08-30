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

type CT_FieldGroup struct {
	// Parent
	ParAttr *uint32
	// Field Base
	BaseAttr *uint32
	// Range Grouping Properties
	RangePr *CT_RangePr
	// Discrete Grouping Properties
	DiscretePr *CT_DiscretePr
	// OLAP Group Items
	GroupItems *CT_GroupItems
}

func NewCT_FieldGroup() *CT_FieldGroup {
	ret := &CT_FieldGroup{}
	return ret
}
func (m *CT_FieldGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ParAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "par"},
			Value: fmt.Sprintf("%v", *m.ParAttr)})
	}
	if m.BaseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "base"},
			Value: fmt.Sprintf("%v", *m.BaseAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.RangePr != nil {
		serangePr := xml.StartElement{Name: xml.Name{Local: "x:rangePr"}}
		e.EncodeElement(m.RangePr, serangePr)
	}
	if m.DiscretePr != nil {
		sediscretePr := xml.StartElement{Name: xml.Name{Local: "x:discretePr"}}
		e.EncodeElement(m.DiscretePr, sediscretePr)
	}
	if m.GroupItems != nil {
		segroupItems := xml.StartElement{Name: xml.Name{Local: "x:groupItems"}}
		e.EncodeElement(m.GroupItems, segroupItems)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FieldGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "par" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ParAttr = &pt
		}
		if attr.Name.Local == "base" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BaseAttr = &pt
		}
	}
lCT_FieldGroup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rangePr":
				m.RangePr = NewCT_RangePr()
				if err := d.DecodeElement(m.RangePr, &el); err != nil {
					return err
				}
			case "discretePr":
				m.DiscretePr = NewCT_DiscretePr()
				if err := d.DecodeElement(m.DiscretePr, &el); err != nil {
					return err
				}
			case "groupItems":
				m.GroupItems = NewCT_GroupItems()
				if err := d.DecodeElement(m.GroupItems, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FieldGroup
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FieldGroup) Validate() error {
	return m.ValidateWithPath("CT_FieldGroup")
}
func (m *CT_FieldGroup) ValidateWithPath(path string) error {
	if m.RangePr != nil {
		if err := m.RangePr.ValidateWithPath(path + "/RangePr"); err != nil {
			return err
		}
	}
	if m.DiscretePr != nil {
		if err := m.DiscretePr.ValidateWithPath(path + "/DiscretePr"); err != nil {
			return err
		}
	}
	if m.GroupItems != nil {
		if err := m.GroupItems.ValidateWithPath(path + "/GroupItems"); err != nil {
			return err
		}
	}
	return nil
}
