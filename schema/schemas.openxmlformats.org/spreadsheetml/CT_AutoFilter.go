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

type CT_AutoFilter struct {
	// Cell or Range Reference
	RefAttr *string
	// AutoFilter Column
	FilterColumn []*CT_FilterColumn
	// Sort State for Auto Filter
	SortState *CT_SortState
	ExtLst    *CT_ExtensionList
}

func NewCT_AutoFilter() *CT_AutoFilter {
	ret := &CT_AutoFilter{}
	return ret
}
func (m *CT_AutoFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
			Value: fmt.Sprintf("%v", *m.RefAttr)})
	}
	e.EncodeToken(start)
	if m.FilterColumn != nil {
		sefilterColumn := xml.StartElement{Name: xml.Name{Local: "x:filterColumn"}}
		e.EncodeElement(m.FilterColumn, sefilterColumn)
	}
	if m.SortState != nil {
		sesortState := xml.StartElement{Name: xml.Name{Local: "x:sortState"}}
		e.EncodeElement(m.SortState, sesortState)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AutoFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = &parsed
		}
	}
lCT_AutoFilter:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "filterColumn":
				tmp := NewCT_FilterColumn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FilterColumn = append(m.FilterColumn, tmp)
			case "sortState":
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AutoFilter
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_AutoFilter) Validate() error {
	return m.ValidateWithPath("CT_AutoFilter")
}
func (m *CT_AutoFilter) ValidateWithPath(path string) error {
	for i, v := range m.FilterColumn {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FilterColumn[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.SortState != nil {
		if err := m.SortState.ValidateWithPath(path + "/SortState"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
