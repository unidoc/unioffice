// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_OMathPara struct {
	OMathParaPr *CT_OMathParaPr
	OMath       []*CT_OMath
}

func NewCT_OMathPara() *CT_OMathPara {
	ret := &CT_OMathPara{}
	return ret
}
func (m *CT_OMathPara) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.OMathParaPr != nil {
		seoMathParaPr := xml.StartElement{Name: xml.Name{Local: "m:oMathParaPr"}}
		e.EncodeElement(m.OMathParaPr, seoMathParaPr)
	}
	seoMath := xml.StartElement{Name: xml.Name{Local: "m:oMath"}}
	e.EncodeElement(m.OMath, seoMath)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OMathPara) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OMathPara:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "oMathParaPr":
				m.OMathParaPr = NewCT_OMathParaPr()
				if err := d.DecodeElement(m.OMathParaPr, &el); err != nil {
					return err
				}
			case "oMath":
				tmp := NewCT_OMath()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.OMath = append(m.OMath, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OMathPara
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_OMathPara) Validate() error {
	return m.ValidateWithPath("CT_OMathPara")
}
func (m *CT_OMathPara) ValidateWithPath(path string) error {
	if m.OMathParaPr != nil {
		if err := m.OMathParaPr.ValidateWithPath(path + "/OMathParaPr"); err != nil {
			return err
		}
	}
	for i, v := range m.OMath {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/OMath[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
