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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type EG_SerShared struct {
	Idx   *CT_UnsignedInt
	Order *CT_UnsignedInt
	Tx    *CT_SerTx
	SpPr  *drawingml.CT_ShapeProperties
}

func NewEG_SerShared() *EG_SerShared {
	ret := &EG_SerShared{}
	ret.Idx = NewCT_UnsignedInt()
	ret.Order = NewCT_UnsignedInt()
	return ret
}

func (m *EG_SerShared) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	seidx := xml.StartElement{Name: xml.Name{Local: "idx"}}
	e.EncodeElement(m.Idx, seidx)
	seorder := xml.StartElement{Name: xml.Name{Local: "order"}}
	e.EncodeElement(m.Order, seorder)
	if m.Tx != nil {
		setx := xml.StartElement{Name: xml.Name{Local: "tx"}}
		e.EncodeElement(m.Tx, setx)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	return nil
}

func (m *EG_SerShared) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
	m.Order = NewCT_UnsignedInt()
lEG_SerShared:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "idx":
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case "order":
				if err := d.DecodeElement(m.Order, &el); err != nil {
					return err
				}
			case "tx":
				m.Tx = NewCT_SerTx()
				if err := d.DecodeElement(m.Tx, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_SerShared %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_SerShared
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_SerShared and its children
func (m *EG_SerShared) Validate() error {
	return m.ValidateWithPath("EG_SerShared")
}

// ValidateWithPath validates the EG_SerShared and its children, prefixing error messages with path
func (m *EG_SerShared) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if err := m.Order.ValidateWithPath(path + "/Order"); err != nil {
		return err
	}
	if m.Tx != nil {
		if err := m.Tx.ValidateWithPath(path + "/Tx"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	return nil
}
