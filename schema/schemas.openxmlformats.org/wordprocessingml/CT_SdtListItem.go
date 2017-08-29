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

type CT_SdtListItem struct {
	// List Entry Display Text
	DisplayTextAttr *string
	// List Entry Value
	ValueAttr *string
}

func NewCT_SdtListItem() *CT_SdtListItem {
	ret := &CT_SdtListItem{}
	return ret
}
func (m *CT_SdtListItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DisplayTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:displayText"},
			Value: fmt.Sprintf("%v", *m.DisplayTextAttr)})
	}
	if m.ValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:value"},
			Value: fmt.Sprintf("%v", *m.ValueAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SdtListItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "displayText" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DisplayTextAttr = &parsed
		}
		if attr.Name.Local == "value" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValueAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SdtListItem: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_SdtListItem) Validate() error {
	return m.ValidateWithPath("CT_SdtListItem")
}
func (m *CT_SdtListItem) ValidateWithPath(path string) error {
	return nil
}
