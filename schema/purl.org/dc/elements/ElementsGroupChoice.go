// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements

import (
	"encoding/xml"
	"fmt"
	"log"
)

type ElementsGroupChoice struct {
	AnyEl []*AnyEl
}

func NewElementsGroupChoice() *ElementsGroupChoice {
	ret := &ElementsGroupChoice{}
	return ret
}
func (m *ElementsGroupChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AnyEl != nil {
		seany := xml.StartElement{Name: xml.Name{Local: "dc:any"}}
		e.EncodeElement(m.AnyEl, seany)
	}
	return nil
}
func (m *ElementsGroupChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsGroupChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "any":
				tmp := NewAnyEl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AnyEl = append(m.AnyEl, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementsGroupChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *ElementsGroupChoice) Validate() error {
	return m.ValidateWithPath("ElementsGroupChoice")
}
func (m *ElementsGroupChoice) ValidateWithPath(path string) error {
	for i, v := range m.AnyEl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AnyEl[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
