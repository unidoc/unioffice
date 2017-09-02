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

type EG_WrapTypeChoice struct {
	WrapNone         *CT_WrapNone
	WrapSquare       *CT_WrapSquare
	WrapTight        *CT_WrapTight
	WrapThrough      *CT_WrapThrough
	WrapTopAndBottom *CT_WrapTopBottom
}

func NewEG_WrapTypeChoice() *EG_WrapTypeChoice {
	ret := &EG_WrapTypeChoice{}
	return ret
}

func (m *EG_WrapTypeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WrapNone != nil {
		sewrapNone := xml.StartElement{Name: xml.Name{Local: "wp:wrapNone"}}
		e.EncodeElement(m.WrapNone, sewrapNone)
	}
	if m.WrapSquare != nil {
		sewrapSquare := xml.StartElement{Name: xml.Name{Local: "wp:wrapSquare"}}
		e.EncodeElement(m.WrapSquare, sewrapSquare)
	}
	if m.WrapTight != nil {
		sewrapTight := xml.StartElement{Name: xml.Name{Local: "wp:wrapTight"}}
		e.EncodeElement(m.WrapTight, sewrapTight)
	}
	if m.WrapThrough != nil {
		sewrapThrough := xml.StartElement{Name: xml.Name{Local: "wp:wrapThrough"}}
		e.EncodeElement(m.WrapThrough, sewrapThrough)
	}
	if m.WrapTopAndBottom != nil {
		sewrapTopAndBottom := xml.StartElement{Name: xml.Name{Local: "wp:wrapTopAndBottom"}}
		e.EncodeElement(m.WrapTopAndBottom, sewrapTopAndBottom)
	}
	return nil
}

func (m *EG_WrapTypeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_WrapTypeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "wrapNone":
				m.WrapNone = NewCT_WrapNone()
				if err := d.DecodeElement(m.WrapNone, &el); err != nil {
					return err
				}
			case "wrapSquare":
				m.WrapSquare = NewCT_WrapSquare()
				if err := d.DecodeElement(m.WrapSquare, &el); err != nil {
					return err
				}
			case "wrapTight":
				m.WrapTight = NewCT_WrapTight()
				if err := d.DecodeElement(m.WrapTight, &el); err != nil {
					return err
				}
			case "wrapThrough":
				m.WrapThrough = NewCT_WrapThrough()
				if err := d.DecodeElement(m.WrapThrough, &el); err != nil {
					return err
				}
			case "wrapTopAndBottom":
				m.WrapTopAndBottom = NewCT_WrapTopBottom()
				if err := d.DecodeElement(m.WrapTopAndBottom, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_WrapTypeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_WrapTypeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_WrapTypeChoice and its children
func (m *EG_WrapTypeChoice) Validate() error {
	return m.ValidateWithPath("EG_WrapTypeChoice")
}

// ValidateWithPath validates the EG_WrapTypeChoice and its children, prefixing error messages with path
func (m *EG_WrapTypeChoice) ValidateWithPath(path string) error {
	if m.WrapNone != nil {
		if err := m.WrapNone.ValidateWithPath(path + "/WrapNone"); err != nil {
			return err
		}
	}
	if m.WrapSquare != nil {
		if err := m.WrapSquare.ValidateWithPath(path + "/WrapSquare"); err != nil {
			return err
		}
	}
	if m.WrapTight != nil {
		if err := m.WrapTight.ValidateWithPath(path + "/WrapTight"); err != nil {
			return err
		}
	}
	if m.WrapThrough != nil {
		if err := m.WrapThrough.ValidateWithPath(path + "/WrapThrough"); err != nil {
			return err
		}
	}
	if m.WrapTopAndBottom != nil {
		if err := m.WrapTopAndBottom.ValidateWithPath(path + "/WrapTopAndBottom"); err != nil {
			return err
		}
	}
	return nil
}
