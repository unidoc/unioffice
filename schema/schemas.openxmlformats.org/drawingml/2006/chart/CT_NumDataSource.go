// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_NumDataSource struct {
	Choice *CT_NumDataSourceChoice
}

func NewCT_NumDataSource() *CT_NumDataSource {
	ret := &CT_NumDataSource{}
	ret.Choice = NewCT_NumDataSourceChoice()
	return ret
}
func (m *CT_NumDataSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	m.Choice.MarshalXML(e, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NumDataSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_NumDataSourceChoice()
lCT_NumDataSource:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "numRef":
				m.Choice = NewCT_NumDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumRef, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "numLit":
				m.Choice = NewCT_NumDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumLit, &el); err != nil {
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
			break lCT_NumDataSource
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NumDataSource) Validate() error {
	return m.ValidateWithPath("CT_NumDataSource")
}
func (m *CT_NumDataSource) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
