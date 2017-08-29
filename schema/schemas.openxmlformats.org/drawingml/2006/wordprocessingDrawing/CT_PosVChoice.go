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

type CT_PosVChoice struct {
	Align     ST_AlignV
	PosOffset *int32
}

func NewCT_PosVChoice() *CT_PosVChoice {
	ret := &CT_PosVChoice{}
	return ret
}
func (m *CT_PosVChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Align != ST_AlignVUnset {
		sealign := xml.StartElement{Name: xml.Name{Local: "wp:align"}}
		e.EncodeElement(m.Align, sealign)
	}
	if m.PosOffset != nil {
		seposOffset := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
		e.EncodeElement(m.PosOffset, seposOffset)
	}
	return nil
}
func (m *CT_PosVChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PosVChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "align":
				m.Align = ST_AlignVUnset
				if err := d.DecodeElement(m.Align, &el); err != nil {
					return err
				}
			case "posOffset":
				m.PosOffset = new(int32)
				if err := d.DecodeElement(m.PosOffset, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PosVChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PosVChoice) Validate() error {
	return m.ValidateWithPath("CT_PosVChoice")
}
func (m *CT_PosVChoice) ValidateWithPath(path string) error {
	if err := m.Align.ValidateWithPath(path + "/Align"); err != nil {
		return err
	}
	return nil
}
