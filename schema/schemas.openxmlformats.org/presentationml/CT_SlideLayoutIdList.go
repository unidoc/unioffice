// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_SlideLayoutIdList struct {
	// Slide Layout Id
	SldLayoutId []*CT_SlideLayoutIdListEntry
}

func NewCT_SlideLayoutIdList() *CT_SlideLayoutIdList {
	ret := &CT_SlideLayoutIdList{}
	return ret
}
func (m *CT_SlideLayoutIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.SldLayoutId != nil {
		sesldLayoutId := xml.StartElement{Name: xml.Name{Local: "p:sldLayoutId"}}
		e.EncodeElement(m.SldLayoutId, sesldLayoutId)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SlideLayoutIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideLayoutIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sldLayoutId":
				tmp := NewCT_SlideLayoutIdListEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SldLayoutId = append(m.SldLayoutId, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideLayoutIdList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SlideLayoutIdList) Validate() error {
	return m.ValidateWithPath("CT_SlideLayoutIdList")
}
func (m *CT_SlideLayoutIdList) ValidateWithPath(path string) error {
	for i, v := range m.SldLayoutId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SldLayoutId[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
