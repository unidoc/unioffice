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

	"baliance.com/gooxml"
)

type CT_PivotSource struct {
	Name   string
	FmtId  *CT_UnsignedInt
	ExtLst []*CT_ExtensionList
}

func NewCT_PivotSource() *CT_PivotSource {
	ret := &CT_PivotSource{}
	ret.FmtId = NewCT_UnsignedInt()
	return ret
}
func (m *CT_PivotSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	sename := xml.StartElement{Name: xml.Name{Local: "name"}}
	gooxml.AddPreserveSpaceAttr(&sename, m.Name)
	e.EncodeElement(m.Name, sename)
	sefmtId := xml.StartElement{Name: xml.Name{Local: "fmtId"}}
	e.EncodeElement(m.FmtId, sefmtId)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FmtId = NewCT_UnsignedInt()
lCT_PivotSource:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "name":
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case "fmtId":
				if err := d.DecodeElement(m.FmtId, &el); err != nil {
					return err
				}
			case "extLst":
				tmp := NewCT_ExtensionList()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ExtLst = append(m.ExtLst, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotSource
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotSource) Validate() error {
	return m.ValidateWithPath("CT_PivotSource")
}
func (m *CT_PivotSource) ValidateWithPath(path string) error {
	if err := m.FmtId.ValidateWithPath(path + "/FmtId"); err != nil {
		return err
	}
	for i, v := range m.ExtLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ExtLst[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
