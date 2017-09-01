// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_ColorTransformHeaderLst struct {
	ColorsDefHdr []*CT_ColorTransformHeader
}

func NewCT_ColorTransformHeaderLst() *CT_ColorTransformHeaderLst {
	ret := &CT_ColorTransformHeaderLst{}
	return ret
}

func (m *CT_ColorTransformHeaderLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.ColorsDefHdr != nil {
		secolorsDefHdr := xml.StartElement{Name: xml.Name{Local: "colorsDefHdr"}}
		e.EncodeElement(m.ColorsDefHdr, secolorsDefHdr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorTransformHeaderLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorTransformHeaderLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "colorsDefHdr":
				tmp := NewCT_ColorTransformHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ColorsDefHdr = append(m.ColorsDefHdr, tmp)
			default:
				log.Printf("skipping unsupported element on CT_ColorTransformHeaderLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorTransformHeaderLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorTransformHeaderLst and its children
func (m *CT_ColorTransformHeaderLst) Validate() error {
	return m.ValidateWithPath("CT_ColorTransformHeaderLst")
}

// ValidateWithPath validates the CT_ColorTransformHeaderLst and its children, prefixing error messages with path
func (m *CT_ColorTransformHeaderLst) ValidateWithPath(path string) error {
	for i, v := range m.ColorsDefHdr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ColorsDefHdr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
