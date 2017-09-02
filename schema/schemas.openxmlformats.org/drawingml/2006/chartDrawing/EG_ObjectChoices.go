// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"log"
)

type EG_ObjectChoices struct {
	Choice *EG_ObjectChoicesChoice
}

func NewEG_ObjectChoices() *EG_ObjectChoices {
	ret := &EG_ObjectChoices{}
	return ret
}

func (m *EG_ObjectChoices) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	return nil
}

func (m *EG_ObjectChoices) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ObjectChoices:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Sp, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "grpSp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GrpSp, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "graphicFrame":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GraphicFrame, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "cxnSp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.CxnSp, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "pic":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Pic, &el); err != nil {
					return err
				}
				_ = m.Choice
			default:
				log.Printf("skipping unsupported element on EG_ObjectChoices %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ObjectChoices
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ObjectChoices and its children
func (m *EG_ObjectChoices) Validate() error {
	return m.ValidateWithPath("EG_ObjectChoices")
}

// ValidateWithPath validates the EG_ObjectChoices and its children, prefixing error messages with path
func (m *EG_ObjectChoices) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
