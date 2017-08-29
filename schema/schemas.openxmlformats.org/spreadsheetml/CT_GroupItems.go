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
)

type CT_GroupItems struct {
	// Items Created Count
	CountAttr *uint32
	// No Value
	M []*CT_Missing
	// Numeric Value
	N []*CT_Number
	// Boolean
	B []*CT_Boolean
	// Error Value
	E []*CT_Error
	// Character Value
	S []*CT_String
	// Date Time
	D []*CT_DateTime
}

func NewCT_GroupItems() *CT_GroupItems {
	ret := &CT_GroupItems{}
	return ret
}
func (m *CT_GroupItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "x:m"}}
		e.EncodeElement(m.M, sem)
	}
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "x:n"}}
		e.EncodeElement(m.N, sen)
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "x:b"}}
		e.EncodeElement(m.B, seb)
	}
	if m.E != nil {
		see := xml.StartElement{Name: xml.Name{Local: "x:e"}}
		e.EncodeElement(m.E, see)
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "x:s"}}
		e.EncodeElement(m.S, ses)
	}
	if m.D != nil {
		sed := xml.StartElement{Name: xml.Name{Local: "x:d"}}
		e.EncodeElement(m.D, sed)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GroupItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_GroupItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "m":
				tmp := NewCT_Missing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case "n":
				tmp := NewCT_Number()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			case "b":
				tmp := NewCT_Boolean()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.B = append(m.B, tmp)
			case "e":
				tmp := NewCT_Error()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			case "s":
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			case "d":
				tmp := NewCT_DateTime()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.D = append(m.D, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupItems
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GroupItems) Validate() error {
	return m.ValidateWithPath("CT_GroupItems")
}
func (m *CT_GroupItems) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.B {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/B[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.D {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/D[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
