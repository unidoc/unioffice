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

type CT_StyleDefinitionHeaderLst struct {
	StyleDefHdr []*CT_StyleDefinitionHeader
}

func NewCT_StyleDefinitionHeaderLst() *CT_StyleDefinitionHeaderLst {
	ret := &CT_StyleDefinitionHeaderLst{}
	return ret
}
func (m *CT_StyleDefinitionHeaderLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.StyleDefHdr != nil {
		sestyleDefHdr := xml.StartElement{Name: xml.Name{Local: "styleDefHdr"}}
		e.EncodeElement(m.StyleDefHdr, sestyleDefHdr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_StyleDefinitionHeaderLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_StyleDefinitionHeaderLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "styleDefHdr":
				tmp := NewCT_StyleDefinitionHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.StyleDefHdr = append(m.StyleDefHdr, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StyleDefinitionHeaderLst
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_StyleDefinitionHeaderLst) Validate() error {
	return m.ValidateWithPath("CT_StyleDefinitionHeaderLst")
}
func (m *CT_StyleDefinitionHeaderLst) ValidateWithPath(path string) error {
	for i, v := range m.StyleDefHdr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/StyleDefHdr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
