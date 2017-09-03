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

type CT_Drawing struct {
	RelSizeAnchor *CT_RelSizeAnchor
	AbsSizeAnchor *CT_AbsSizeAnchor
}

func NewCT_Drawing() *CT_Drawing {
	ret := &CT_Drawing{}
	return ret
}

func (m *CT_Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "CT_Drawing"
	e.EncodeToken(start)
	if m.RelSizeAnchor != nil {
		serelSizeAnchor := xml.StartElement{Name: xml.Name{Local: "relSizeAnchor"}}
		e.EncodeElement(m.RelSizeAnchor, serelSizeAnchor)
	}
	if m.AbsSizeAnchor != nil {
		seabsSizeAnchor := xml.StartElement{Name: xml.Name{Local: "absSizeAnchor"}}
		e.EncodeElement(m.AbsSizeAnchor, seabsSizeAnchor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Drawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Drawing:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "relSizeAnchor":
				m.RelSizeAnchor = NewCT_RelSizeAnchor()
				if err := d.DecodeElement(m.RelSizeAnchor, &el); err != nil {
					return err
				}
			case "absSizeAnchor":
				m.AbsSizeAnchor = NewCT_AbsSizeAnchor()
				if err := d.DecodeElement(m.AbsSizeAnchor, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Drawing %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Drawing
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Drawing and its children
func (m *CT_Drawing) Validate() error {
	return m.ValidateWithPath("CT_Drawing")
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (m *CT_Drawing) ValidateWithPath(path string) error {
	if m.RelSizeAnchor != nil {
		if err := m.RelSizeAnchor.ValidateWithPath(path + "/RelSizeAnchor"); err != nil {
			return err
		}
	}
	if m.AbsSizeAnchor != nil {
		if err := m.AbsSizeAnchor.ValidateWithPath(path + "/AbsSizeAnchor"); err != nil {
			return err
		}
	}
	return nil
}
