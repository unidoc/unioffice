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
	"strconv"
)

type CT_MdxMetadata struct {
	// MDX Metadata Record Count
	CountAttr *uint32
	// MDX Metadata Record
	Mdx []*CT_Mdx
}

func NewCT_MdxMetadata() *CT_MdxMetadata {
	ret := &CT_MdxMetadata{}
	return ret
}
func (m *CT_MdxMetadata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	semdx := xml.StartElement{Name: xml.Name{Local: "x:mdx"}}
	e.EncodeElement(m.Mdx, semdx)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MdxMetadata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_MdxMetadata:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "mdx":
				tmp := NewCT_Mdx()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Mdx = append(m.Mdx, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MdxMetadata
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MdxMetadata) Validate() error {
	return m.ValidateWithPath("CT_MdxMetadata")
}
func (m *CT_MdxMetadata) ValidateWithPath(path string) error {
	for i, v := range m.Mdx {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Mdx[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
