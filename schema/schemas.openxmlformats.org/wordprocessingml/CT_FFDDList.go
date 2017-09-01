// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_FFDDList struct {
	// Drop-Down List Selection
	Result *CT_DecimalNumber
	// Default Drop-Down List Item Index
	Default *CT_DecimalNumber
	// Drop-Down List Entry
	ListEntry []*CT_String
}

func NewCT_FFDDList() *CT_FFDDList {
	ret := &CT_FFDDList{}
	return ret
}
func (m *CT_FFDDList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Result != nil {
		seresult := xml.StartElement{Name: xml.Name{Local: "w:result"}}
		e.EncodeElement(m.Result, seresult)
	}
	if m.Default != nil {
		sedefault := xml.StartElement{Name: xml.Name{Local: "w:default"}}
		e.EncodeElement(m.Default, sedefault)
	}
	if m.ListEntry != nil {
		selistEntry := xml.StartElement{Name: xml.Name{Local: "w:listEntry"}}
		e.EncodeElement(m.ListEntry, selistEntry)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FFDDList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FFDDList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "result":
				m.Result = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.Result, &el); err != nil {
					return err
				}
			case "default":
				m.Default = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.Default, &el); err != nil {
					return err
				}
			case "listEntry":
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ListEntry = append(m.ListEntry, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FFDDList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FFDDList) Validate() error {
	return m.ValidateWithPath("CT_FFDDList")
}
func (m *CT_FFDDList) ValidateWithPath(path string) error {
	if m.Result != nil {
		if err := m.Result.ValidateWithPath(path + "/Result"); err != nil {
			return err
		}
	}
	if m.Default != nil {
		if err := m.Default.ValidateWithPath(path + "/Default"); err != nil {
			return err
		}
	}
	for i, v := range m.ListEntry {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ListEntry[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
