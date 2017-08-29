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

type CT_Tx struct {
	Choice *CT_TxChoice
}

func NewCT_Tx() *CT_Tx {
	ret := &CT_Tx{}
	ret.Choice = NewCT_TxChoice()
	return ret
}
func (m *CT_Tx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	m.Choice.MarshalXML(e, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Tx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_TxChoice()
lCT_Tx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "strRef":
				m.Choice = NewCT_TxChoice()
				if err := d.DecodeElement(&m.Choice.StrRef, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "rich":
				m.Choice = NewCT_TxChoice()
				if err := d.DecodeElement(&m.Choice.Rich, &el); err != nil {
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
			break lCT_Tx
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Tx) Validate() error {
	return m.ValidateWithPath("CT_Tx")
}
func (m *CT_Tx) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
