// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"log"
)

type EG_WrapType struct {
	Choice *EG_WrapTypeChoice
}

func NewEG_WrapType() *EG_WrapType {
	ret := &EG_WrapType{}
	return ret
}
func (m *EG_WrapType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	return nil
}
func (m *EG_WrapType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_WrapType:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "wrapNone":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapNone, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapSquare":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapSquare, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapTight":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTight, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapThrough":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapThrough, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapTopAndBottom":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTopAndBottom, &el); err != nil {
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
			break lEG_WrapType
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_WrapType) Validate() error {
	return m.ValidateWithPath("EG_WrapType")
}
func (m *EG_WrapType) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
