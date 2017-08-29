// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_TextRun struct {
	R   *CT_RegularTextRun
	Br  *CT_TextLineBreak
	Fld *CT_TextField
}

func NewEG_TextRun() *EG_TextRun {
	ret := &EG_TextRun{}
	return ret
}
func (m *EG_TextRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.R != nil {
		ser := xml.StartElement{Name: xml.Name{Local: "a:r"}}
		e.EncodeElement(m.R, ser)
	}
	if m.Br != nil {
		sebr := xml.StartElement{Name: xml.Name{Local: "a:br"}}
		e.EncodeElement(m.Br, sebr)
	}
	if m.Fld != nil {
		sefld := xml.StartElement{Name: xml.Name{Local: "a:fld"}}
		e.EncodeElement(m.Fld, sefld)
	}
	return nil
}
func (m *EG_TextRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "r":
				m.R = NewCT_RegularTextRun()
				if err := d.DecodeElement(m.R, &el); err != nil {
					return err
				}
			case "br":
				m.Br = NewCT_TextLineBreak()
				if err := d.DecodeElement(m.Br, &el); err != nil {
					return err
				}
			case "fld":
				m.Fld = NewCT_TextField()
				if err := d.DecodeElement(m.Fld, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextRun
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_TextRun) Validate() error {
	return m.ValidateWithPath("EG_TextRun")
}
func (m *EG_TextRun) ValidateWithPath(path string) error {
	if m.R != nil {
		if err := m.R.ValidateWithPath(path + "/R"); err != nil {
			return err
		}
	}
	if m.Br != nil {
		if err := m.Br.ValidateWithPath(path + "/Br"); err != nil {
			return err
		}
	}
	if m.Fld != nil {
		if err := m.Fld.ValidateWithPath(path + "/Fld"); err != nil {
			return err
		}
	}
	return nil
}
