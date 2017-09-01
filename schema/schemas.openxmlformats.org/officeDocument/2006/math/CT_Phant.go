// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_Phant struct {
	PhantPr *CT_PhantPr
	E       *CT_OMathArg
}

func NewCT_Phant() *CT_Phant {
	ret := &CT_Phant{}
	ret.E = NewCT_OMathArg()
	return ret
}
func (m *CT_Phant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.PhantPr != nil {
		sephantPr := xml.StartElement{Name: xml.Name{Local: "m:phantPr"}}
		e.EncodeElement(m.PhantPr, sephantPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Phant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
lCT_Phant:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "phantPr":
				m.PhantPr = NewCT_PhantPr()
				if err := d.DecodeElement(m.PhantPr, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Phant
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Phant) Validate() error {
	return m.ValidateWithPath("CT_Phant")
}
func (m *CT_Phant) ValidateWithPath(path string) error {
	if m.PhantPr != nil {
		if err := m.PhantPr.ValidateWithPath(path + "/PhantPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
