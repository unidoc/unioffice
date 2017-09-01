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

type CT_TableStyles struct {
	// Table Style Count
	CountAttr *uint32
	// Default Table Style
	DefaultTableStyleAttr *string
	// Default Pivot Style
	DefaultPivotStyleAttr *string
	// Table Style
	TableStyle []*CT_TableStyle
}

func NewCT_TableStyles() *CT_TableStyles {
	ret := &CT_TableStyles{}
	return ret
}
func (m *CT_TableStyles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.DefaultTableStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultTableStyle"},
			Value: fmt.Sprintf("%v", *m.DefaultTableStyleAttr)})
	}
	if m.DefaultPivotStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultPivotStyle"},
			Value: fmt.Sprintf("%v", *m.DefaultPivotStyleAttr)})
	}
	e.EncodeToken(start)
	if m.TableStyle != nil {
		setableStyle := xml.StartElement{Name: xml.Name{Local: "x:tableStyle"}}
		e.EncodeElement(m.TableStyle, setableStyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TableStyles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
		if attr.Name.Local == "defaultTableStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefaultTableStyleAttr = &parsed
		}
		if attr.Name.Local == "defaultPivotStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefaultPivotStyleAttr = &parsed
		}
	}
lCT_TableStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tableStyle":
				tmp := NewCT_TableStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TableStyle = append(m.TableStyle, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyles
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TableStyles) Validate() error {
	return m.ValidateWithPath("CT_TableStyles")
}
func (m *CT_TableStyles) ValidateWithPath(path string) error {
	for i, v := range m.TableStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TableStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
