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

type CT_DdeItem struct {
	// DDE Name
	NameAttr *string
	// Object Linking TechnologyE
	OleAttr *bool
	// Advise
	AdviseAttr *bool
	// Data is an Image
	PreferPicAttr *bool
	// DDE Name Values
	Values *CT_DdeValues
}

func NewCT_DdeItem() *CT_DdeItem {
	ret := &CT_DdeItem{}
	return ret
}
func (m *CT_DdeItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.OleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ole"},
			Value: fmt.Sprintf("%v", *m.OleAttr)})
	}
	if m.AdviseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advise"},
			Value: fmt.Sprintf("%v", *m.AdviseAttr)})
	}
	if m.PreferPicAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preferPic"},
			Value: fmt.Sprintf("%v", *m.PreferPicAttr)})
	}
	e.EncodeToken(start)
	if m.Values != nil {
		sevalues := xml.StartElement{Name: xml.Name{Local: "x:values"}}
		e.EncodeElement(m.Values, sevalues)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DdeItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "ole" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OleAttr = &parsed
		}
		if attr.Name.Local == "advise" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdviseAttr = &parsed
		}
		if attr.Name.Local == "preferPic" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreferPicAttr = &parsed
		}
	}
lCT_DdeItem:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "values":
				m.Values = NewCT_DdeValues()
				if err := d.DecodeElement(m.Values, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DdeItem
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DdeItem) Validate() error {
	return m.ValidateWithPath("CT_DdeItem")
}
func (m *CT_DdeItem) ValidateWithPath(path string) error {
	if m.Values != nil {
		if err := m.Values.ValidateWithPath(path + "/Values"); err != nil {
			return err
		}
	}
	return nil
}
