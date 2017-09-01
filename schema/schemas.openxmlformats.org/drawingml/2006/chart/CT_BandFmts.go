// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_BandFmts struct {
	BandFmt []*CT_BandFmt
}

func NewCT_BandFmts() *CT_BandFmts {
	ret := &CT_BandFmts{}
	return ret
}

func (m *CT_BandFmts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.BandFmt != nil {
		sebandFmt := xml.StartElement{Name: xml.Name{Local: "bandFmt"}}
		e.EncodeElement(m.BandFmt, sebandFmt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BandFmts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BandFmts:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bandFmt":
				tmp := NewCT_BandFmt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BandFmt = append(m.BandFmt, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BandFmts
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BandFmts and its children
func (m *CT_BandFmts) Validate() error {
	return m.ValidateWithPath("CT_BandFmts")
}

// ValidateWithPath validates the CT_BandFmts and its children, prefixing error messages with path
func (m *CT_BandFmts) ValidateWithPath(path string) error {
	for i, v := range m.BandFmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BandFmt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
