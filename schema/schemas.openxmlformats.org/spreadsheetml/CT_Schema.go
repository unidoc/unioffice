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

	"baliance.com/gooxml"
)

type CT_Schema struct {
	// Schema ID
	IDAttr string
	// Schema Reference
	SchemaRefAttr *string
	// Schema Root Namespace
	NamespaceAttr *string
	// Schema Language
	SchemaLanguageAttr *string
	Any                Any
}

func NewCT_Schema() *CT_Schema {
	ret := &CT_Schema{}
	return ret
}
func (m *CT_Schema) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ID"},
		Value: fmt.Sprintf("%v", m.IDAttr)})
	if m.SchemaRefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "SchemaRef"},
			Value: fmt.Sprintf("%v", *m.SchemaRefAttr)})
	}
	if m.NamespaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Namespace"},
			Value: fmt.Sprintf("%v", *m.NamespaceAttr)})
	}
	if m.SchemaLanguageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "SchemaLanguage"},
			Value: fmt.Sprintf("%v", *m.SchemaLanguageAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Any != nil {
		m.Any.MarshalXML(e, start)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IDAttr = parsed
		}
		if attr.Name.Local == "SchemaRef" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SchemaRefAttr = &parsed
		}
		if attr.Name.Local == "Namespace" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NamespaceAttr = &parsed
		}
		if attr.Name.Local == "SchemaLanguage" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SchemaLanguageAttr = &parsed
		}
	}
lCT_Schema:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			default:
				if anyEl, err := gooxml.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lCT_Schema
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Schema) Validate() error {
	return m.ValidateWithPath("CT_Schema")
}
func (m *CT_Schema) ValidateWithPath(path string) error {
	return nil
}
