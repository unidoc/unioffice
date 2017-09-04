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

type CT_DdeValues struct {
	// Rows
	RowsAttr *uint32
	// Columns
	ColsAttr *uint32
	// Value
	Value []*CT_DdeValue
}

func NewCT_DdeValues() *CT_DdeValues {
	ret := &CT_DdeValues{}
	ret.RowsAttr = gooxml.Uint32(1)
	ret.ColsAttr = gooxml.Uint32(1)
	return ret
}

func (m *CT_DdeValues) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rows"},
			Value: fmt.Sprintf("%v", *m.RowsAttr)})
	}
	if m.ColsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cols"},
			Value: fmt.Sprintf("%v", *m.ColsAttr)})
	}
	e.EncodeToken(start)
	sevalue := xml.StartElement{Name: xml.Name{Local: "x:value"}}
	e.EncodeElement(m.Value, sevalue)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DdeValues) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RowsAttr = gooxml.Uint32(1)
	m.ColsAttr = gooxml.Uint32(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "rows" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RowsAttr = &pt
		}
		if attr.Name.Local == "cols" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ColsAttr = &pt
		}
	}
lCT_DdeValues:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "value":
				tmp := NewCT_DdeValue()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Value = append(m.Value, tmp)
			default:
				log.Printf("skipping unsupported element on CT_DdeValues %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DdeValues
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DdeValues and its children
func (m *CT_DdeValues) Validate() error {
	return m.ValidateWithPath("CT_DdeValues")
}

// ValidateWithPath validates the CT_DdeValues and its children, prefixing error messages with path
func (m *CT_DdeValues) ValidateWithPath(path string) error {
	for i, v := range m.Value {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Value[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
