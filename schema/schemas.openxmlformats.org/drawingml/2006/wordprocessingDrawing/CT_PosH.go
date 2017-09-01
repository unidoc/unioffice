// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_PosH struct {
	RelativeFromAttr ST_RelFromH
	Choice           *CT_PosHChoice
}

func NewCT_PosH() *CT_PosH {
	ret := &CT_PosH{}
	ret.RelativeFromAttr = ST_RelFromH(1)
	ret.Choice = NewCT_PosHChoice()
	return ret
}
func (m *CT_PosH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.RelativeFromAttr.MarshalXMLAttr(xml.Name{Local: "relativeFrom"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PosH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RelativeFromAttr = ST_RelFromH(1)
	m.Choice = NewCT_PosHChoice()
	for _, attr := range start.Attr {
		if attr.Name.Local == "relativeFrom" {
			m.RelativeFromAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PosH:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "align":
				m.Choice = NewCT_PosHChoice()
				if err := d.DecodeElement(&m.Choice.Align, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "posOffset":
				m.Choice = NewCT_PosHChoice()
				if err := d.DecodeElement(&m.Choice.PosOffset, &el); err != nil {
					return err
				}
				_ = m.Choice
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PosH
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PosH) Validate() error {
	return m.ValidateWithPath("CT_PosH")
}
func (m *CT_PosH) ValidateWithPath(path string) error {
	if m.RelativeFromAttr == ST_RelFromHUnset {
		return fmt.Errorf("%s/RelativeFromAttr is a mandatory field", path)
	}
	if err := m.RelativeFromAttr.ValidateWithPath(path + "/RelativeFromAttr"); err != nil {
		return err
	}
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
