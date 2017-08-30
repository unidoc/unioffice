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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Spacing struct {
	// Spacing Above Paragraph
	BeforeAttr *sharedTypes.ST_TwipsMeasure
	// Spacing Above Paragraph IN Line Units
	BeforeLinesAttr *int32
	// Automatically Determine Spacing Above Paragraph
	BeforeAutospacingAttr *sharedTypes.ST_OnOff
	// Spacing Below Paragraph
	AfterAttr *sharedTypes.ST_TwipsMeasure
	// Spacing Below Paragraph in Line Units
	AfterLinesAttr *int32
	// Automatically Determine Spacing Below Paragraph
	AfterAutospacingAttr *sharedTypes.ST_OnOff
	// Spacing Between Lines in Paragraph
	LineAttr *ST_SignedTwipsMeasure
	// Spacing Between Lines
	LineRuleAttr ST_LineSpacingRule
}

func NewCT_Spacing() *CT_Spacing {
	ret := &CT_Spacing{}
	return ret
}
func (m *CT_Spacing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BeforeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:before"},
			Value: fmt.Sprintf("%v", *m.BeforeAttr)})
	}
	if m.BeforeLinesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:beforeLines"},
			Value: fmt.Sprintf("%v", *m.BeforeLinesAttr)})
	}
	if m.BeforeAutospacingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:beforeAutospacing"},
			Value: fmt.Sprintf("%v", *m.BeforeAutospacingAttr)})
	}
	if m.AfterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:after"},
			Value: fmt.Sprintf("%v", *m.AfterAttr)})
	}
	if m.AfterLinesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:afterLines"},
			Value: fmt.Sprintf("%v", *m.AfterLinesAttr)})
	}
	if m.AfterAutospacingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:afterAutospacing"},
			Value: fmt.Sprintf("%v", *m.AfterAutospacingAttr)})
	}
	if m.LineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:line"},
			Value: fmt.Sprintf("%v", *m.LineAttr)})
	}
	if m.LineRuleAttr != ST_LineSpacingRuleUnset {
		attr, err := m.LineRuleAttr.MarshalXMLAttr(xml.Name{Local: "w:lineRule"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Spacing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "before" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.BeforeAttr = &parsed
		}
		if attr.Name.Local == "beforeLines" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.BeforeLinesAttr = &pt
		}
		if attr.Name.Local == "beforeAutospacing" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.BeforeAutospacingAttr = &parsed
		}
		if attr.Name.Local == "after" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.AfterAttr = &parsed
		}
		if attr.Name.Local == "afterLines" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.AfterLinesAttr = &pt
		}
		if attr.Name.Local == "afterAutospacing" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.AfterAutospacingAttr = &parsed
		}
		if attr.Name.Local == "line" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.LineAttr = &parsed
		}
		if attr.Name.Local == "lineRule" {
			m.LineRuleAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Spacing: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Spacing) Validate() error {
	return m.ValidateWithPath("CT_Spacing")
}
func (m *CT_Spacing) ValidateWithPath(path string) error {
	if m.BeforeAttr != nil {
		if err := m.BeforeAttr.ValidateWithPath(path + "/BeforeAttr"); err != nil {
			return err
		}
	}
	if m.BeforeAutospacingAttr != nil {
		if err := m.BeforeAutospacingAttr.ValidateWithPath(path + "/BeforeAutospacingAttr"); err != nil {
			return err
		}
	}
	if m.AfterAttr != nil {
		if err := m.AfterAttr.ValidateWithPath(path + "/AfterAttr"); err != nil {
			return err
		}
	}
	if m.AfterAutospacingAttr != nil {
		if err := m.AfterAutospacingAttr.ValidateWithPath(path + "/AfterAutospacingAttr"); err != nil {
			return err
		}
	}
	if m.LineAttr != nil {
		if err := m.LineAttr.ValidateWithPath(path + "/LineAttr"); err != nil {
			return err
		}
	}
	if err := m.LineRuleAttr.ValidateWithPath(path + "/LineRuleAttr"); err != nil {
		return err
	}
	return nil
}
