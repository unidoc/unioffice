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

type CT_ColorMappingOverride struct {
	Choice *CT_ColorMappingOverrideChoice
}

func NewCT_ColorMappingOverride() *CT_ColorMappingOverride {
	ret := &CT_ColorMappingOverride{}
	ret.Choice = NewCT_ColorMappingOverrideChoice()
	return ret
}

func (m *CT_ColorMappingOverride) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "a:CT_ColorMappingOverride"
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorMappingOverride) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_ColorMappingOverrideChoice()
lCT_ColorMappingOverride:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "masterClrMapping":
				m.Choice = NewCT_ColorMappingOverrideChoice()
				if err := d.DecodeElement(&m.Choice.MasterClrMapping, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "overrideClrMapping":
				m.Choice = NewCT_ColorMappingOverrideChoice()
				if err := d.DecodeElement(&m.Choice.OverrideClrMapping, &el); err != nil {
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
			break lCT_ColorMappingOverride
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorMappingOverride and its children
func (m *CT_ColorMappingOverride) Validate() error {
	return m.ValidateWithPath("CT_ColorMappingOverride")
}

// ValidateWithPath validates the CT_ColorMappingOverride and its children, prefixing error messages with path
func (m *CT_ColorMappingOverride) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
