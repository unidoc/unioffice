// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_Endnotes struct {
	// Endnote Content
	Endnote *CT_FtnEdn
}

func NewCT_Endnotes() *CT_Endnotes {
	ret := &CT_Endnotes{}
	return ret
}
func (m *CT_Endnotes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Endnote != nil {
		seendnote := xml.StartElement{Name: xml.Name{Local: "w:endnote"}}
		e.EncodeElement(m.Endnote, seendnote)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Endnotes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Endnotes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "endnote":
				m.Endnote = NewCT_FtnEdn()
				if err := d.DecodeElement(m.Endnote, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Endnotes
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Endnotes) Validate() error {
	return m.ValidateWithPath("CT_Endnotes")
}
func (m *CT_Endnotes) ValidateWithPath(path string) error {
	if m.Endnote != nil {
		if err := m.Endnote.ValidateWithPath(path + "/Endnote"); err != nil {
			return err
		}
	}
	return nil
}
