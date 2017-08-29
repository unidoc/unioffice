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

type EG_TextBullet struct {
	BuNone    *CT_TextNoBullet
	BuAutoNum *CT_TextAutonumberBullet
	BuChar    *CT_TextCharBullet
	BuBlip    *CT_TextBlipBullet
}

func NewEG_TextBullet() *EG_TextBullet {
	ret := &EG_TextBullet{}
	return ret
}
func (m *EG_TextBullet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BuNone != nil {
		sebuNone := xml.StartElement{Name: xml.Name{Local: "a:buNone"}}
		e.EncodeElement(m.BuNone, sebuNone)
	}
	if m.BuAutoNum != nil {
		sebuAutoNum := xml.StartElement{Name: xml.Name{Local: "a:buAutoNum"}}
		e.EncodeElement(m.BuAutoNum, sebuAutoNum)
	}
	if m.BuChar != nil {
		sebuChar := xml.StartElement{Name: xml.Name{Local: "a:buChar"}}
		e.EncodeElement(m.BuChar, sebuChar)
	}
	if m.BuBlip != nil {
		sebuBlip := xml.StartElement{Name: xml.Name{Local: "a:buBlip"}}
		e.EncodeElement(m.BuBlip, sebuBlip)
	}
	return nil
}
func (m *EG_TextBullet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextBullet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "buNone":
				m.BuNone = NewCT_TextNoBullet()
				if err := d.DecodeElement(m.BuNone, &el); err != nil {
					return err
				}
			case "buAutoNum":
				m.BuAutoNum = NewCT_TextAutonumberBullet()
				if err := d.DecodeElement(m.BuAutoNum, &el); err != nil {
					return err
				}
			case "buChar":
				m.BuChar = NewCT_TextCharBullet()
				if err := d.DecodeElement(m.BuChar, &el); err != nil {
					return err
				}
			case "buBlip":
				m.BuBlip = NewCT_TextBlipBullet()
				if err := d.DecodeElement(m.BuBlip, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextBullet
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_TextBullet) Validate() error {
	return m.ValidateWithPath("EG_TextBullet")
}
func (m *EG_TextBullet) ValidateWithPath(path string) error {
	if m.BuNone != nil {
		if err := m.BuNone.ValidateWithPath(path + "/BuNone"); err != nil {
			return err
		}
	}
	if m.BuAutoNum != nil {
		if err := m.BuAutoNum.ValidateWithPath(path + "/BuAutoNum"); err != nil {
			return err
		}
	}
	if m.BuChar != nil {
		if err := m.BuChar.ValidateWithPath(path + "/BuChar"); err != nil {
			return err
		}
	}
	if m.BuBlip != nil {
		if err := m.BuBlip.ValidateWithPath(path + "/BuBlip"); err != nil {
			return err
		}
	}
	return nil
}
