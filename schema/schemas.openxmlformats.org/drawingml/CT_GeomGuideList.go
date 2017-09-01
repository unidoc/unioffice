// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_GeomGuideList struct {
	Gd []*CT_GeomGuide
}

func NewCT_GeomGuideList() *CT_GeomGuideList {
	ret := &CT_GeomGuideList{}
	return ret
}

func (m *CT_GeomGuideList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Gd != nil {
		segd := xml.StartElement{Name: xml.Name{Local: "a:gd"}}
		e.EncodeElement(m.Gd, segd)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GeomGuideList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GeomGuideList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "gd":
				tmp := NewCT_GeomGuide()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Gd = append(m.Gd, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GeomGuideList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GeomGuideList and its children
func (m *CT_GeomGuideList) Validate() error {
	return m.ValidateWithPath("CT_GeomGuideList")
}

// ValidateWithPath validates the CT_GeomGuideList and its children, prefixing error messages with path
func (m *CT_GeomGuideList) ValidateWithPath(path string) error {
	for i, v := range m.Gd {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Gd[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
