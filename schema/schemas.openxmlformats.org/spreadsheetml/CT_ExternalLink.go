// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_ExternalLink struct {
	Choice *CT_ExternalLinkChoice
	ExtLst *CT_ExtensionList
}

func NewCT_ExternalLink() *CT_ExternalLink {
	ret := &CT_ExternalLink{}
	return ret
}
func (m *CT_ExternalLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ExternalLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalLink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "externalBook":
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.ExternalBook, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "ddeLink":
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.DdeLink, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "oleLink":
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.OleLink, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalLink
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ExternalLink) Validate() error {
	return m.ValidateWithPath("CT_ExternalLink")
}
func (m *CT_ExternalLink) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
