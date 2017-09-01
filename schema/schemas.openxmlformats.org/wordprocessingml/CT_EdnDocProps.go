// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_EdnDocProps struct {
	// Endnote Placement
	Pos *CT_EdnPos
	// Endnote Numbering Format
	NumFmt *CT_NumFmt
	// Footnote and Endnote Numbering Starting Value
	NumStart *CT_DecimalNumber
	// Footnote and Endnote Numbering Restart Location
	NumRestart *CT_NumRestart
	Endnote    []*CT_FtnEdnSepRef
}

func NewCT_EdnDocProps() *CT_EdnDocProps {
	ret := &CT_EdnDocProps{}
	return ret
}
func (m *CT_EdnDocProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Pos != nil {
		sepos := xml.StartElement{Name: xml.Name{Local: "w:pos"}}
		e.EncodeElement(m.Pos, sepos)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "w:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.NumStart != nil {
		senumStart := xml.StartElement{Name: xml.Name{Local: "w:numStart"}}
		e.EncodeElement(m.NumStart, senumStart)
	}
	if m.NumRestart != nil {
		senumRestart := xml.StartElement{Name: xml.Name{Local: "w:numRestart"}}
		e.EncodeElement(m.NumRestart, senumRestart)
	}
	if m.Endnote != nil {
		seendnote := xml.StartElement{Name: xml.Name{Local: "w:endnote"}}
		e.EncodeElement(m.Endnote, seendnote)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_EdnDocProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EdnDocProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pos":
				m.Pos = NewCT_EdnPos()
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case "numFmt":
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case "numStart":
				m.NumStart = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumStart, &el); err != nil {
					return err
				}
			case "numRestart":
				m.NumRestart = NewCT_NumRestart()
				if err := d.DecodeElement(m.NumRestart, &el); err != nil {
					return err
				}
			case "endnote":
				tmp := NewCT_FtnEdnSepRef()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Endnote = append(m.Endnote, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EdnDocProps
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_EdnDocProps) Validate() error {
	return m.ValidateWithPath("CT_EdnDocProps")
}
func (m *CT_EdnDocProps) ValidateWithPath(path string) error {
	if m.Pos != nil {
		if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.NumStart != nil {
		if err := m.NumStart.ValidateWithPath(path + "/NumStart"); err != nil {
			return err
		}
	}
	if m.NumRestart != nil {
		if err := m.NumRestart.ValidateWithPath(path + "/NumRestart"); err != nil {
			return err
		}
	}
	for i, v := range m.Endnote {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Endnote[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
