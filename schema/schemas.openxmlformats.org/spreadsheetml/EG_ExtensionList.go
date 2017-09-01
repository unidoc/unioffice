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

type EG_ExtensionList struct {
	// Extension
	Ext []*CT_Extension
}

func NewEG_ExtensionList() *EG_ExtensionList {
	ret := &EG_ExtensionList{}
	return ret
}

func (m *EG_ExtensionList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "x:ext"}}
		e.EncodeElement(m.Ext, seext)
	}
	return nil
}

func (m *EG_ExtensionList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ExtensionList:
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
			break lEG_ExtensionList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ExtensionList and its children
func (m *EG_ExtensionList) Validate() error {
	return m.ValidateWithPath("EG_ExtensionList")
}

// ValidateWithPath validates the EG_ExtensionList and its children, prefixing error messages with path
func (m *EG_ExtensionList) ValidateWithPath(path string) error {
	for i, v := range m.Ext {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ext[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
