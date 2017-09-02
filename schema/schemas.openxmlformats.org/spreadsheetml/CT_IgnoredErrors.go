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

type CT_IgnoredErrors struct {
	// Ignored Error
	IgnoredError []*CT_IgnoredError
	ExtLst       *CT_ExtensionList
}

func NewCT_IgnoredErrors() *CT_IgnoredErrors {
	ret := &CT_IgnoredErrors{}
	return ret
}

func (m *CT_IgnoredErrors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seignoredError := xml.StartElement{Name: xml.Name{Local: "x:ignoredError"}}
	e.EncodeElement(m.IgnoredError, seignoredError)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_IgnoredErrors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_IgnoredErrors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ignoredError":
				tmp := NewCT_IgnoredError()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.IgnoredError = append(m.IgnoredError, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_IgnoredErrors %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_IgnoredErrors
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_IgnoredErrors and its children
func (m *CT_IgnoredErrors) Validate() error {
	return m.ValidateWithPath("CT_IgnoredErrors")
}

// ValidateWithPath validates the CT_IgnoredErrors and its children, prefixing error messages with path
func (m *CT_IgnoredErrors) ValidateWithPath(path string) error {
	for i, v := range m.IgnoredError {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/IgnoredError[%d]", path, i)); err != nil {
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
