// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_DocVars struct {
	// Single Document Variable
	DocVar []*CT_DocVar
}

func NewCT_DocVars() *CT_DocVars {
	ret := &CT_DocVars{}
	return ret
}

func (m *CT_DocVars) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.DocVar != nil {
		sedocVar := xml.StartElement{Name: xml.Name{Local: "w:docVar"}}
		e.EncodeElement(m.DocVar, sedocVar)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocVars) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocVars:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "docVar":
				tmp := NewCT_DocVar()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DocVar = append(m.DocVar, tmp)
			default:
				log.Printf("skipping unsupported element on CT_DocVars %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocVars
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocVars and its children
func (m *CT_DocVars) Validate() error {
	return m.ValidateWithPath("CT_DocVars")
}

// ValidateWithPath validates the CT_DocVars and its children, prefixing error messages with path
func (m *CT_DocVars) ValidateWithPath(path string) error {
	for i, v := range m.DocVar {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DocVar[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
