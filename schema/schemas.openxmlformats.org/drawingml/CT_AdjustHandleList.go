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

type CT_AdjustHandleList struct {
	AhXY    []*CT_XYAdjustHandle
	AhPolar []*CT_PolarAdjustHandle
}

func NewCT_AdjustHandleList() *CT_AdjustHandleList {
	ret := &CT_AdjustHandleList{}
	return ret
}
func (m *CT_AdjustHandleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.AhXY != nil {
		seahXY := xml.StartElement{Name: xml.Name{Local: "a:ahXY"}}
		e.EncodeElement(m.AhXY, seahXY)
	}
	if m.AhPolar != nil {
		seahPolar := xml.StartElement{Name: xml.Name{Local: "a:ahPolar"}}
		e.EncodeElement(m.AhPolar, seahPolar)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AdjustHandleList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AdjustHandleList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ahXY":
				tmp := NewCT_XYAdjustHandle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AhXY = append(m.AhXY, tmp)
			case "ahPolar":
				tmp := NewCT_PolarAdjustHandle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AhPolar = append(m.AhPolar, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AdjustHandleList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_AdjustHandleList) Validate() error {
	return m.ValidateWithPath("CT_AdjustHandleList")
}
func (m *CT_AdjustHandleList) ValidateWithPath(path string) error {
	for i, v := range m.AhXY {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AhXY[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AhPolar {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AhPolar[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
