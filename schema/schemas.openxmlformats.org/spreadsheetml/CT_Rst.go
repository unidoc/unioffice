// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml"
)

type CT_Rst struct {
	// Text
	T *string
	// Rich Text Run
	R []*CT_RElt
	// Phonetic Run
	RPh []*CT_PhoneticRun
	// Phonetic Properties
	PhoneticPr *CT_PhoneticPr
}

func NewCT_Rst() *CT_Rst {
	ret := &CT_Rst{}
	return ret
}
func (m *CT_Rst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.T != nil {
		set := xml.StartElement{Name: xml.Name{Local: "x:t"}}
		gooxml.AddPreserveSpaceAttr(&set, *m.T)
		e.EncodeElement(m.T, set)
	}
	if m.R != nil {
		ser := xml.StartElement{Name: xml.Name{Local: "x:r"}}
		e.EncodeElement(m.R, ser)
	}
	if m.RPh != nil {
		serPh := xml.StartElement{Name: xml.Name{Local: "x:rPh"}}
		e.EncodeElement(m.RPh, serPh)
	}
	if m.PhoneticPr != nil {
		sephoneticPr := xml.StartElement{Name: xml.Name{Local: "x:phoneticPr"}}
		e.EncodeElement(m.PhoneticPr, sephoneticPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Rst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Rst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "t":
				m.T = new(string)
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			case "r":
				tmp := NewCT_RElt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.R = append(m.R, tmp)
			case "rPh":
				tmp := NewCT_PhoneticRun()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RPh = append(m.RPh, tmp)
			case "phoneticPr":
				m.PhoneticPr = NewCT_PhoneticPr()
				if err := d.DecodeElement(m.PhoneticPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Rst
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Rst) Validate() error {
	return m.ValidateWithPath("CT_Rst")
}
func (m *CT_Rst) ValidateWithPath(path string) error {
	for i, v := range m.R {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/R[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.RPh {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RPh[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.PhoneticPr != nil {
		if err := m.PhoneticPr.ValidateWithPath(path + "/PhoneticPr"); err != nil {
			return err
		}
	}
	return nil
}
