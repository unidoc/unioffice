// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Hyperlinks struct {
	// Hyperlink
	Hyperlink []*CT_Hyperlink
}

func NewCT_Hyperlinks() *CT_Hyperlinks {
	ret := &CT_Hyperlinks{}
	return ret
}
func (m *CT_Hyperlinks) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sehyperlink := xml.StartElement{Name: xml.Name{Local: "x:hyperlink"}}
	e.EncodeElement(m.Hyperlink, sehyperlink)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Hyperlinks) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Hyperlinks:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "hyperlink":
				tmp := NewCT_Hyperlink()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Hyperlink = append(m.Hyperlink, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Hyperlinks
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Hyperlinks) Validate() error {
	return m.ValidateWithPath("CT_Hyperlinks")
}
func (m *CT_Hyperlinks) ValidateWithPath(path string) error {
	for i, v := range m.Hyperlink {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Hyperlink[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
