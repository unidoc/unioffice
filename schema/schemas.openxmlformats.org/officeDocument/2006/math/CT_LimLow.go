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

type CT_LimLow struct {
	LimLowPr *CT_LimLowPr
	E        *CT_OMathArg
	Lim      *CT_OMathArg
}

func NewCT_LimLow() *CT_LimLow {
	ret := &CT_LimLow{}
	ret.E = NewCT_OMathArg()
	ret.Lim = NewCT_OMathArg()
	return ret
}
func (m *CT_LimLow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.LimLowPr != nil {
		selimLowPr := xml.StartElement{Name: xml.Name{Local: "m:limLowPr"}}
		e.EncodeElement(m.LimLowPr, selimLowPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	selim := xml.StartElement{Name: xml.Name{Local: "m:lim"}}
	e.EncodeElement(m.Lim, selim)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LimLow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
	m.Lim = NewCT_OMathArg()
lCT_LimLow:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "limLowPr":
				m.LimLowPr = NewCT_LimLowPr()
				if err := d.DecodeElement(m.LimLowPr, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			case "lim":
				if err := d.DecodeElement(m.Lim, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LimLow
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_LimLow) Validate() error {
	return m.ValidateWithPath("CT_LimLow")
}
func (m *CT_LimLow) ValidateWithPath(path string) error {
	if m.LimLowPr != nil {
		if err := m.LimLowPr.ValidateWithPath(path + "/LimLowPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	if err := m.Lim.ValidateWithPath(path + "/Lim"); err != nil {
		return err
	}
	return nil
}
