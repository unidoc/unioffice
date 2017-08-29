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
	"strconv"
)

type CT_TLOleChartTargetElement struct {
	// Type
	TypeAttr ST_TLChartSubelementType
	// Level
	LvlAttr *uint32
}

func NewCT_TLOleChartTargetElement() *CT_TLOleChartTargetElement {
	ret := &CT_TLOleChartTargetElement{}
	return ret
}
func (m *CT_TLOleChartTargetElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.LvlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lvl"},
			Value: fmt.Sprintf("%v", *m.LvlAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLOleChartTargetElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "lvl" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.LvlAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLOleChartTargetElement: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TLOleChartTargetElement) Validate() error {
	return m.ValidateWithPath("CT_TLOleChartTargetElement")
}
func (m *CT_TLOleChartTargetElement) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_TLChartSubelementTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
