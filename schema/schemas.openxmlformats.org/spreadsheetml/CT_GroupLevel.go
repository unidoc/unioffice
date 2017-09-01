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

type CT_GroupLevel struct {
	// Unique Name
	UniqueNameAttr string
	// Grouping Level Display Name
	CaptionAttr string
	// User-Defined Group Level
	UserAttr *bool
	// Custom Roll Up
	CustomRollUpAttr *bool
	// OLAP Level Groups
	Groups *CT_Groups
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_GroupLevel() *CT_GroupLevel {
	ret := &CT_GroupLevel{}
	return ret
}
func (m *CT_GroupLevel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
		Value: fmt.Sprintf("%v", m.UniqueNameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
		Value: fmt.Sprintf("%v", m.CaptionAttr)})
	if m.UserAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "user"},
			Value: fmt.Sprintf("%v", *m.UserAttr)})
	}
	if m.CustomRollUpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customRollUp"},
			Value: fmt.Sprintf("%v", *m.CustomRollUpAttr)})
	}
	e.EncodeToken(start)
	if m.Groups != nil {
		segroups := xml.StartElement{Name: xml.Name{Local: "x:groups"}}
		e.EncodeElement(m.Groups, segroups)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GroupLevel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = parsed
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = parsed
		}
		if attr.Name.Local == "user" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UserAttr = &parsed
		}
		if attr.Name.Local == "customRollUp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomRollUpAttr = &parsed
		}
	}
lCT_GroupLevel:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "groups":
				m.Groups = NewCT_Groups()
				if err := d.DecodeElement(m.Groups, &el); err != nil {
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
			break lCT_GroupLevel
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GroupLevel) Validate() error {
	return m.ValidateWithPath("CT_GroupLevel")
}
func (m *CT_GroupLevel) ValidateWithPath(path string) error {
	if m.Groups != nil {
		if err := m.Groups.ValidateWithPath(path + "/Groups"); err != nil {
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
