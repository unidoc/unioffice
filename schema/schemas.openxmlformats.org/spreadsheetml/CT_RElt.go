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

	"baliance.com/gooxml"
)

type CT_RElt struct {
	// Run Properties
	RPr *CT_RPrElt
	// Text
	T string
}

func NewCT_RElt() *CT_RElt {
	ret := &CT_RElt{}
	return ret
}

func (m *CT_RElt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "x:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	set := xml.StartElement{Name: xml.Name{Local: "x:t"}}
	gooxml.AddPreserveSpaceAttr(&set, m.T)
	e.EncodeElement(m.T, set)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RElt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RElt:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rPr":
				m.RPr = NewCT_RPrElt()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case "t":
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RElt
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RElt and its children
func (m *CT_RElt) Validate() error {
	return m.ValidateWithPath("CT_RElt")
}

// ValidateWithPath validates the CT_RElt and its children, prefixing error messages with path
func (m *CT_RElt) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
