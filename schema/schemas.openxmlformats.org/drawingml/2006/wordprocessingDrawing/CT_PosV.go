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

type CT_PosV struct {
	RelativeFromAttr ST_RelFromV
	Choice           *CT_PosVChoice
}

func NewCT_PosV() *CT_PosV {
	ret := &CT_PosV{}
	ret.Choice = NewCT_PosVChoice()
	return ret
}
func (m *CT_PosV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.RelativeFromAttr.MarshalXMLAttr(xml.Name{Local: "relativeFrom"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	start.Attr = nil
	m.Choice.MarshalXML(e, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PosV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_PosVChoice()
	for _, attr := range start.Attr {
		if attr.Name.Local == "relativeFrom" {
			m.RelativeFromAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PosV:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "align":
				m.Choice = NewCT_PosVChoice()
				if err := d.DecodeElement(&m.Choice.Align, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "posOffset":
				m.Choice = NewCT_PosVChoice()
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
			break lCT_PosV
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PosV) Validate() error {
	return m.ValidateWithPath("CT_PosV")
}
func (m *CT_PosV) ValidateWithPath(path string) error {
	if m.RelativeFromAttr == ST_RelFromVUnset {
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
