// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_WrapSquare struct {
	WrapTextAttr ST_WrapText
	DistTAttr    *uint32
	DistBAttr    *uint32
	DistLAttr    *uint32
	DistRAttr    *uint32
	EffectExtent *CT_EffectExtent
}

func NewCT_WrapSquare() *CT_WrapSquare {
	ret := &CT_WrapSquare{}
	ret.WrapTextAttr = ST_WrapText(1)
	return ret
}

func (m *CT_WrapSquare) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.WrapTextAttr.MarshalXMLAttr(xml.Name{Local: "wrapText"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"},
			Value: fmt.Sprintf("%v", *m.DistTAttr)})
	}
	if m.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"},
			Value: fmt.Sprintf("%v", *m.DistBAttr)})
	}
	if m.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"},
			Value: fmt.Sprintf("%v", *m.DistLAttr)})
	}
	if m.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"},
			Value: fmt.Sprintf("%v", *m.DistRAttr)})
	}
	e.EncodeToken(start)
	if m.EffectExtent != nil {
		seeffectExtent := xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}
		e.EncodeElement(m.EffectExtent, seeffectExtent)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WrapSquare) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.WrapTextAttr = ST_WrapText(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "wrapText" {
			m.WrapTextAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "distT" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistTAttr = &pt
		}
		if attr.Name.Local == "distB" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistBAttr = &pt
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
		}
	}
lCT_WrapSquare:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "effectExtent":
				m.EffectExtent = NewCT_EffectExtent()
				if err := d.DecodeElement(m.EffectExtent, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_WrapSquare %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WrapSquare
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WrapSquare and its children
func (m *CT_WrapSquare) Validate() error {
	return m.ValidateWithPath("CT_WrapSquare")
}

// ValidateWithPath validates the CT_WrapSquare and its children, prefixing error messages with path
func (m *CT_WrapSquare) ValidateWithPath(path string) error {
	if m.WrapTextAttr == ST_WrapTextUnset {
		return fmt.Errorf("%s/WrapTextAttr is a mandatory field", path)
	}
	if err := m.WrapTextAttr.ValidateWithPath(path + "/WrapTextAttr"); err != nil {
		return err
	}
	if m.EffectExtent != nil {
		if err := m.EffectExtent.ValidateWithPath(path + "/EffectExtent"); err != nil {
			return err
		}
	}
	return nil
}
