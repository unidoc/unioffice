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

type CT_CalcChain struct {
	// Cell
	C      []*CT_CalcCell
	ExtLst *CT_ExtensionList
}

func NewCT_CalcChain() *CT_CalcChain {
	ret := &CT_CalcChain{}
	return ret
}

func (m *CT_CalcChain) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sec := xml.StartElement{Name: xml.Name{Local: "x:c"}}
	e.EncodeElement(m.C, sec)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CalcChain) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CalcChain:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "c":
				tmp := NewCT_CalcCell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.C = append(m.C, tmp)
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
			break lCT_CalcChain
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CalcChain and its children
func (m *CT_CalcChain) Validate() error {
	return m.ValidateWithPath("CT_CalcChain")
}

// ValidateWithPath validates the CT_CalcChain and its children, prefixing error messages with path
func (m *CT_CalcChain) ValidateWithPath(path string) error {
	for i, v := range m.C {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/C[%d]", path, i)); err != nil {
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
