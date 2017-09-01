// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_CTCategories struct {
	Cat []*CT_CTCategory
}

func NewCT_CTCategories() *CT_CTCategories {
	ret := &CT_CTCategories{}
	return ret
}

func (m *CT_CTCategories) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Cat != nil {
		secat := xml.StartElement{Name: xml.Name{Local: "cat"}}
		e.EncodeElement(m.Cat, secat)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CTCategories) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CTCategories:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cat":
				tmp := NewCT_CTCategory()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cat = append(m.Cat, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CTCategories
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CTCategories and its children
func (m *CT_CTCategories) Validate() error {
	return m.ValidateWithPath("CT_CTCategories")
}

// ValidateWithPath validates the CT_CTCategories and its children, prefixing error messages with path
func (m *CT_CTCategories) ValidateWithPath(path string) error {
	for i, v := range m.Cat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
