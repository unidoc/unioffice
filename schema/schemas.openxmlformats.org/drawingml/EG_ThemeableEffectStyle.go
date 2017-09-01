// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_ThemeableEffectStyle struct {
	Effect    *CT_EffectProperties
	EffectRef *CT_StyleMatrixReference
}

func NewEG_ThemeableEffectStyle() *EG_ThemeableEffectStyle {
	ret := &EG_ThemeableEffectStyle{}
	return ret
}

func (m *EG_ThemeableEffectStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Effect != nil {
		seeffect := xml.StartElement{Name: xml.Name{Local: "a:effect"}}
		e.EncodeElement(m.Effect, seeffect)
	}
	if m.EffectRef != nil {
		seeffectRef := xml.StartElement{Name: xml.Name{Local: "a:effectRef"}}
		e.EncodeElement(m.EffectRef, seeffectRef)
	}
	return nil
}

func (m *EG_ThemeableEffectStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ThemeableEffectStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "effect":
				m.Effect = NewCT_EffectProperties()
				if err := d.DecodeElement(m.Effect, &el); err != nil {
					return err
				}
			case "effectRef":
				m.EffectRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.EffectRef, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ThemeableEffectStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ThemeableEffectStyle and its children
func (m *EG_ThemeableEffectStyle) Validate() error {
	return m.ValidateWithPath("EG_ThemeableEffectStyle")
}

// ValidateWithPath validates the EG_ThemeableEffectStyle and its children, prefixing error messages with path
func (m *EG_ThemeableEffectStyle) ValidateWithPath(path string) error {
	if m.Effect != nil {
		if err := m.Effect.ValidateWithPath(path + "/Effect"); err != nil {
			return err
		}
	}
	if m.EffectRef != nil {
		if err := m.EffectRef.ValidateWithPath(path + "/EffectRef"); err != nil {
			return err
		}
	}
	return nil
}
