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

	"baliance.com/gooxml"
)

type CT_MetadataTypes struct {
	// Metadata Type Count
	CountAttr *uint32
	// Metadata Type Information
	MetadataType []*CT_MetadataType
}

func NewCT_MetadataTypes() *CT_MetadataTypes {
	ret := &CT_MetadataTypes{}
	ret.CountAttr = gooxml.Uint32(0)
	return ret
}

func (m *CT_MetadataTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	semetadataType := xml.StartElement{Name: xml.Name{Local: "x:metadataType"}}
	e.EncodeElement(m.MetadataType, semetadataType)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MetadataTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CountAttr = gooxml.Uint32(0)
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_MetadataTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "metadataType":
				tmp := NewCT_MetadataType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.MetadataType = append(m.MetadataType, tmp)
			default:
				log.Printf("skipping unsupported element on CT_MetadataTypes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MetadataTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MetadataTypes and its children
func (m *CT_MetadataTypes) Validate() error {
	return m.ValidateWithPath("CT_MetadataTypes")
}

// ValidateWithPath validates the CT_MetadataTypes and its children, prefixing error messages with path
func (m *CT_MetadataTypes) ValidateWithPath(path string) error {
	for i, v := range m.MetadataType {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/MetadataType[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
