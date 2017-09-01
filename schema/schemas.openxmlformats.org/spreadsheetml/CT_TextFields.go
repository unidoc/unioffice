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

type CT_TextFields struct {
	// Count of Fields
	CountAttr *uint32
	// Text Import Field Settings
	TextField []*CT_TextField
}

func NewCT_TextFields() *CT_TextFields {
	ret := &CT_TextFields{}
	return ret
}

func (m *CT_TextFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	setextField := xml.StartElement{Name: xml.Name{Local: "x:textField"}}
	e.EncodeElement(m.TextField, setextField)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
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
lCT_TextFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "textField":
				tmp := NewCT_TextField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TextField = append(m.TextField, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextFields
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextFields and its children
func (m *CT_TextFields) Validate() error {
	return m.ValidateWithPath("CT_TextFields")
}

// ValidateWithPath validates the CT_TextFields and its children, prefixing error messages with path
func (m *CT_TextFields) ValidateWithPath(path string) error {
	for i, v := range m.TextField {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TextField[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
