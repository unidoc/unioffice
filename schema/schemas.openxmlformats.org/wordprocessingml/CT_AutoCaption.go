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
)

type CT_AutoCaption struct {
	// Identifier of Object to be Automatically Captioned
	NameAttr string
	// Caption Used for Automatic Captioning
	CaptionAttr string
}

func NewCT_AutoCaption() *CT_AutoCaption {
	ret := &CT_AutoCaption{}
	return ret
}
func (m *CT_AutoCaption) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:caption"},
		Value: fmt.Sprintf("%v", m.CaptionAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AutoCaption) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AutoCaption: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_AutoCaption) Validate() error {
	return m.ValidateWithPath("CT_AutoCaption")
}
func (m *CT_AutoCaption) ValidateWithPath(path string) error {
	return nil
}
