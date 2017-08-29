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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_TableStyleList struct {
	DefAttr  string
	TblStyle []*CT_TableStyle
}

func NewCT_TableStyleList() *CT_TableStyleList {
	ret := &CT_TableStyleList{}
	return ret
}
func (m *CT_TableStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "def"},
		Value: fmt.Sprintf("%v", m.DefAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	if m.TblStyle != nil {
		setblStyle := xml.StartElement{Name: xml.Name{Local: "a:tblStyle"}}
		e.EncodeElement(m.TblStyle, setblStyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TableStyleList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "def" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefAttr = parsed
		}
	}
lCT_TableStyleList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblStyle":
				tmp := NewCT_TableStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblStyle = append(m.TblStyle, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyleList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TableStyleList) Validate() error {
	return m.ValidateWithPath("CT_TableStyleList")
}
func (m *CT_TableStyleList) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.DefAttr) {
		return fmt.Errorf(`%s/m.DefAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.DefAttr)
	}
	for i, v := range m.TblStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TblStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
