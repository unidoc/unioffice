// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_ExtensionList struct {
	// Extension
	Ext []*CT_Extension
}

func NewCT_ExtensionList() *CT_ExtensionList {
	ret := &CT_ExtensionList{}
	return ret
}

func (m *CT_ExtensionList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "p:ext"}}
		e.EncodeElement(m.Ext, seext)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExtensionList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExtensionList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ext":
				tmp := NewCT_Extension()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ext = append(m.Ext, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExtensionList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExtensionList and its children
func (m *CT_ExtensionList) Validate() error {
	return m.ValidateWithPath("CT_ExtensionList")
}

// ValidateWithPath validates the CT_ExtensionList and its children, prefixing error messages with path
func (m *CT_ExtensionList) ValidateWithPath(path string) error {
	for i, v := range m.Ext {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ext[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
