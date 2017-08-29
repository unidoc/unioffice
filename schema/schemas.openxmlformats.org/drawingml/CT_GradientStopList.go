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

type CT_GradientStopList struct {
	Gs []*CT_GradientStop
}

func NewCT_GradientStopList() *CT_GradientStopList {
	ret := &CT_GradientStopList{}
	return ret
}
func (m *CT_GradientStopList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	segs := xml.StartElement{Name: xml.Name{Local: "a:gs"}}
	e.EncodeElement(m.Gs, segs)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GradientStopList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GradientStopList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "gs":
				tmp := NewCT_GradientStop()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Gs = append(m.Gs, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GradientStopList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GradientStopList) Validate() error {
	return m.ValidateWithPath("CT_GradientStopList")
}
func (m *CT_GradientStopList) ValidateWithPath(path string) error {
	for i, v := range m.Gs {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Gs[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
