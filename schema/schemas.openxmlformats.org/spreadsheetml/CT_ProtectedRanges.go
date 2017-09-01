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

type CT_ProtectedRanges struct {
	// Protected Range
	ProtectedRange []*CT_ProtectedRange
}

func NewCT_ProtectedRanges() *CT_ProtectedRanges {
	ret := &CT_ProtectedRanges{}
	return ret
}

func (m *CT_ProtectedRanges) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	seprotectedRange := xml.StartElement{Name: xml.Name{Local: "x:protectedRange"}}
	e.EncodeElement(m.ProtectedRange, seprotectedRange)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ProtectedRanges) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ProtectedRanges:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "protectedRange":
				tmp := NewCT_ProtectedRange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ProtectedRange = append(m.ProtectedRange, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ProtectedRanges
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ProtectedRanges and its children
func (m *CT_ProtectedRanges) Validate() error {
	return m.ValidateWithPath("CT_ProtectedRanges")
}

// ValidateWithPath validates the CT_ProtectedRanges and its children, prefixing error messages with path
func (m *CT_ProtectedRanges) ValidateWithPath(path string) error {
	for i, v := range m.ProtectedRange {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ProtectedRange[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
