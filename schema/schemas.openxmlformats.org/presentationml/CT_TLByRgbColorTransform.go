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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_TLByRgbColorTransform struct {
	// Red
	RAttr drawingml.ST_FixedPercentage
	// Green
	GAttr drawingml.ST_FixedPercentage
	// Blue
	BAttr drawingml.ST_FixedPercentage
}

func NewCT_TLByRgbColorTransform() *CT_TLByRgbColorTransform {
	ret := &CT_TLByRgbColorTransform{}
	return ret
}
func (m *CT_TLByRgbColorTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "g"},
		Value: fmt.Sprintf("%v", m.GAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
		Value: fmt.Sprintf("%v", m.BAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLByRgbColorTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.RAttr = parsed
		}
		if attr.Name.Local == "g" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.GAttr = parsed
		}
		if attr.Name.Local == "b" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLByRgbColorTransform: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TLByRgbColorTransform) Validate() error {
	return m.ValidateWithPath("CT_TLByRgbColorTransform")
}
func (m *CT_TLByRgbColorTransform) ValidateWithPath(path string) error {
	if err := m.RAttr.ValidateWithPath(path + "/RAttr"); err != nil {
		return err
	}
	if err := m.GAttr.ValidateWithPath(path + "/GAttr"); err != nil {
		return err
	}
	if err := m.BAttr.ValidateWithPath(path + "/BAttr"); err != nil {
		return err
	}
	return nil
}
