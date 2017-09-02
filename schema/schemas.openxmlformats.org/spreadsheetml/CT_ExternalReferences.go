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

type CT_ExternalReferences struct {
	// External Reference
	ExternalReference []*CT_ExternalReference
}

func NewCT_ExternalReferences() *CT_ExternalReferences {
	ret := &CT_ExternalReferences{}
	return ret
}

func (m *CT_ExternalReferences) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seexternalReference := xml.StartElement{Name: xml.Name{Local: "x:externalReference"}}
	e.EncodeElement(m.ExternalReference, seexternalReference)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalReferences) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalReferences:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "externalReference":
				tmp := NewCT_ExternalReference()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ExternalReference = append(m.ExternalReference, tmp)
			default:
				log.Printf("skipping unsupported element on CT_ExternalReferences %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalReferences
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalReferences and its children
func (m *CT_ExternalReferences) Validate() error {
	return m.ValidateWithPath("CT_ExternalReferences")
}

// ValidateWithPath validates the CT_ExternalReferences and its children, prefixing error messages with path
func (m *CT_ExternalReferences) ValidateWithPath(path string) error {
	for i, v := range m.ExternalReference {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ExternalReference[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
