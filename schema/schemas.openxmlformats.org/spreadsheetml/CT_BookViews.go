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

type CT_BookViews struct {
	// Workbook View
	WorkbookView []*CT_BookView
}

func NewCT_BookViews() *CT_BookViews {
	ret := &CT_BookViews{}
	return ret
}
func (m *CT_BookViews) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seworkbookView := xml.StartElement{Name: xml.Name{Local: "x:workbookView"}}
	e.EncodeElement(m.WorkbookView, seworkbookView)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_BookViews) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BookViews:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "workbookView":
				tmp := NewCT_BookView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.WorkbookView = append(m.WorkbookView, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BookViews
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_BookViews) Validate() error {
	return m.ValidateWithPath("CT_BookViews")
}
func (m *CT_BookViews) ValidateWithPath(path string) error {
	for i, v := range m.WorkbookView {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/WorkbookView[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
