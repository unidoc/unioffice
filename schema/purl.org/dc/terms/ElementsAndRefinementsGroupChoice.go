// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/purl.org/dc/elements"
)

type ElementsAndRefinementsGroupChoice struct {
	AnyEl []*elements.AnyEl
}

func NewElementsAndRefinementsGroupChoice() *ElementsAndRefinementsGroupChoice {
	ret := &ElementsAndRefinementsGroupChoice{}
	return ret
}
func (m *ElementsAndRefinementsGroupChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AnyEl != nil {
		seany := xml.StartElement{Name: xml.Name{Local: "dc:any"}}
		e.EncodeElement(m.AnyEl, seany)
	}
	return nil
}
func (m *ElementsAndRefinementsGroupChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsAndRefinementsGroupChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "any":
				tmp := elements.NewAnyEl()
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
			break lElementsAndRefinementsGroupChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *ElementsAndRefinementsGroupChoice) Validate() error {
	return m.ValidateWithPath("ElementsAndRefinementsGroupChoice")
}
func (m *ElementsAndRefinementsGroupChoice) ValidateWithPath(path string) error {
	for i, v := range m.AnyEl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AnyEl[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
