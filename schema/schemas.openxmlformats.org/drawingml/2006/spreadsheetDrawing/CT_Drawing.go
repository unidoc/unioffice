// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"log"
)

type CT_Drawing struct {
	TwoCellAnchor  *CT_TwoCellAnchor
	OneCellAnchor  *CT_OneCellAnchor
	AbsoluteAnchor *CT_AbsoluteAnchor
}

func NewCT_Drawing() *CT_Drawing {
	ret := &CT_Drawing{}
	return ret
}

func (m *CT_Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.TwoCellAnchor != nil {
		setwoCellAnchor := xml.StartElement{Name: xml.Name{Local: "xdr:twoCellAnchor"}}
		e.EncodeElement(m.TwoCellAnchor, setwoCellAnchor)
	}
	if m.OneCellAnchor != nil {
		seoneCellAnchor := xml.StartElement{Name: xml.Name{Local: "xdr:oneCellAnchor"}}
		e.EncodeElement(m.OneCellAnchor, seoneCellAnchor)
	}
	if m.AbsoluteAnchor != nil {
		seabsoluteAnchor := xml.StartElement{Name: xml.Name{Local: "xdr:absoluteAnchor"}}
		e.EncodeElement(m.AbsoluteAnchor, seabsoluteAnchor)
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
			case "twoCellAnchor":
				m.TwoCellAnchor = NewCT_TwoCellAnchor()
				if err := d.DecodeElement(m.TwoCellAnchor, &el); err != nil {
					return err
				}
			case "oneCellAnchor":
				m.OneCellAnchor = NewCT_OneCellAnchor()
				if err := d.DecodeElement(m.OneCellAnchor, &el); err != nil {
					return err
				}
			case "absoluteAnchor":
				m.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if err := d.DecodeElement(m.AbsoluteAnchor, &el); err != nil {
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
	if m.TwoCellAnchor != nil {
		if err := m.TwoCellAnchor.ValidateWithPath(path + "/TwoCellAnchor"); err != nil {
			return err
		}
	}
	if m.OneCellAnchor != nil {
		if err := m.OneCellAnchor.ValidateWithPath(path + "/OneCellAnchor"); err != nil {
			return err
		}
	}
	if m.AbsoluteAnchor != nil {
		if err := m.AbsoluteAnchor.ValidateWithPath(path + "/AbsoluteAnchor"); err != nil {
			return err
		}
	}
	return nil
}
