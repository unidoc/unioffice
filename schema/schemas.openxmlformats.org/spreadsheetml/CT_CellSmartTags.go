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

type CT_CellSmartTags struct {
	// Reference
	RAttr string
	// Cell Smart Tag
	CellSmartTag []*CT_CellSmartTag
}

func NewCT_CellSmartTags() *CT_CellSmartTags {
	ret := &CT_CellSmartTags{}
	return ret
}
func (m *CT_CellSmartTags) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	e.EncodeToken(start)
	secellSmartTag := xml.StartElement{Name: xml.Name{Local: "x:cellSmartTag"}}
	e.EncodeElement(m.CellSmartTag, secellSmartTag)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CellSmartTags) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = parsed
		}
	}
lCT_CellSmartTags:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cellSmartTag":
				tmp := NewCT_CellSmartTag()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CellSmartTag = append(m.CellSmartTag, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellSmartTags
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CellSmartTags) Validate() error {
	return m.ValidateWithPath("CT_CellSmartTags")
}
func (m *CT_CellSmartTags) ValidateWithPath(path string) error {
	for i, v := range m.CellSmartTag {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CellSmartTag[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
