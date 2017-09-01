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

type CT_RPR struct {
	Lit    *CT_OnOff
	Choice *CT_RPRChoice
	Brk    *CT_ManualBreak
	Aln    *CT_OnOff
}

func NewCT_RPR() *CT_RPR {
	ret := &CT_RPR{}
	return ret
}
func (m *CT_RPR) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Lit != nil {
		selit := xml.StartElement{Name: xml.Name{Local: "m:lit"}}
		e.EncodeElement(m.Lit, selit)
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.Brk != nil {
		sebrk := xml.StartElement{Name: xml.Name{Local: "m:brk"}}
		e.EncodeElement(m.Brk, sebrk)
	}
	if m.Aln != nil {
		sealn := xml.StartElement{Name: xml.Name{Local: "m:aln"}}
		e.EncodeElement(m.Aln, sealn)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RPR) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RPR:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "lit":
				m.Lit = NewCT_OnOff()
				if err := d.DecodeElement(m.Lit, &el); err != nil {
					return err
				}
			case "nor":
				m.Choice = NewCT_RPRChoice()
				if err := d.DecodeElement(&m.Choice.Nor, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "brk":
				m.Brk = NewCT_ManualBreak()
				if err := d.DecodeElement(m.Brk, &el); err != nil {
					return err
				}
			case "aln":
				m.Aln = NewCT_OnOff()
				if err := d.DecodeElement(m.Aln, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RPR
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_RPR) Validate() error {
	return m.ValidateWithPath("CT_RPR")
}
func (m *CT_RPR) ValidateWithPath(path string) error {
	if m.Lit != nil {
		if err := m.Lit.ValidateWithPath(path + "/Lit"); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.Brk != nil {
		if err := m.Brk.ValidateWithPath(path + "/Brk"); err != nil {
			return err
		}
	}
	if m.Aln != nil {
		if err := m.Aln.ValidateWithPath(path + "/Aln"); err != nil {
			return err
		}
	}
	return nil
}
