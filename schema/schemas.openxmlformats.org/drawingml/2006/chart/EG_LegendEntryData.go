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

type EG_LegendEntryData struct {
	TxPr *drawingml.CT_TextBody
}

func NewEG_LegendEntryData() *EG_LegendEntryData {
	ret := &EG_LegendEntryData{}
	return ret
}

func (m *EG_LegendEntryData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	return nil
}

func (m *EG_LegendEntryData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_LegendEntryData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "txPr":
				m.TxPr = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_LegendEntryData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_LegendEntryData and its children
func (m *EG_LegendEntryData) Validate() error {
	return m.ValidateWithPath("EG_LegendEntryData")
}

// ValidateWithPath validates the EG_LegendEntryData and its children, prefixing error messages with path
func (m *EG_LegendEntryData) ValidateWithPath(path string) error {
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	return nil
}
