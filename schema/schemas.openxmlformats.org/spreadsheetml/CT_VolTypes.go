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

type CT_VolTypes struct {
	// Volatile Dependency Type
	VolType []*CT_VolType
	ExtLst  *CT_ExtensionList
}

func NewCT_VolTypes() *CT_VolTypes {
	ret := &CT_VolTypes{}
	return ret
}

func (m *CT_VolTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sevolType := xml.StartElement{Name: xml.Name{Local: "x:volType"}}
	e.EncodeElement(m.VolType, sevolType)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VolTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_VolTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "volType":
				tmp := NewCT_VolType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.VolType = append(m.VolType, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VolTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VolTypes and its children
func (m *CT_VolTypes) Validate() error {
	return m.ValidateWithPath("CT_VolTypes")
}

// ValidateWithPath validates the CT_VolTypes and its children, prefixing error messages with path
func (m *CT_VolTypes) ValidateWithPath(path string) error {
	for i, v := range m.VolType {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/VolType[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
