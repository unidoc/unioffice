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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_LevelText struct {
	// Level Text
	ValAttr *string
	// Level Text Is Null Character
	NullAttr *sharedTypes.ST_OnOff
}

func NewCT_LevelText() *CT_LevelText {
	ret := &CT_LevelText{}
	return ret
}
func (m *CT_LevelText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.NullAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:null"},
			Value: fmt.Sprintf("%v", *m.NullAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LevelText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
		if attr.Name.Local == "null" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.NullAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LevelText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_LevelText) Validate() error {
	return m.ValidateWithPath("CT_LevelText")
}
func (m *CT_LevelText) ValidateWithPath(path string) error {
	if m.NullAttr != nil {
		if err := m.NullAttr.ValidateWithPath(path + "/NullAttr"); err != nil {
			return err
		}
	}
	return nil
}
