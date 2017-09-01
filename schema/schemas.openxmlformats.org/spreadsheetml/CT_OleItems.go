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

type CT_OleItems struct {
	// Object Link Item
	OleItem []*CT_OleItem
}

func NewCT_OleItems() *CT_OleItems {
	ret := &CT_OleItems{}
	return ret
}

func (m *CT_OleItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.OleItem != nil {
		seoleItem := xml.StartElement{Name: xml.Name{Local: "x:oleItem"}}
		e.EncodeElement(m.OleItem, seoleItem)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OleItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "oleItem":
				tmp := NewCT_OleItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.OleItem = append(m.OleItem, tmp)
			default:
				log.Printf("skipping unsupported element on CT_OleItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleItems and its children
func (m *CT_OleItems) Validate() error {
	return m.ValidateWithPath("CT_OleItems")
}

// ValidateWithPath validates the CT_OleItems and its children, prefixing error messages with path
func (m *CT_OleItems) ValidateWithPath(path string) error {
	for i, v := range m.OleItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/OleItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
