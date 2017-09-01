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

type CT_Footnotes struct {
	// Footnote Content
	Footnote []*CT_FtnEdn
}

func NewCT_Footnotes() *CT_Footnotes {
	ret := &CT_Footnotes{}
	return ret
}

func (m *CT_Footnotes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Footnote != nil {
		sefootnote := xml.StartElement{Name: xml.Name{Local: "w:footnote"}}
		e.EncodeElement(m.Footnote, sefootnote)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Footnotes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Footnotes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "footnote":
				tmp := NewCT_FtnEdn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Footnote = append(m.Footnote, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Footnotes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Footnotes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Footnotes and its children
func (m *CT_Footnotes) Validate() error {
	return m.ValidateWithPath("CT_Footnotes")
}

// ValidateWithPath validates the CT_Footnotes and its children, prefixing error messages with path
func (m *CT_Footnotes) ValidateWithPath(path string) error {
	for i, v := range m.Footnote {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Footnote[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
