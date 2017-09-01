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

type Blip struct {
	CT_Blip
}

func NewBlip() *Blip {
	ret := &Blip{}
	ret.CT_Blip = *NewCT_Blip()
	return ret
}

func (m *Blip) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:blip"
	return m.CT_Blip.MarshalXML(e, start)
}

func (m *Blip) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Blip = *NewCT_Blip()
	for _, attr := range start.Attr {
		if attr.Name.Local == "cstate" {
			m.CstateAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "embed" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EmbedAttr = &parsed
		}
		if attr.Name.Local == "link" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LinkAttr = &parsed
		}
	}
lBlip:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "alphaBiLevel":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaBiLevel, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "alphaCeiling":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaCeiling, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "alphaFloor":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaFloor, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "alphaInv":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaInv, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "alphaMod":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaMod, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "alphaModFix":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaModFix, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "alphaRepl":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaRepl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "biLevel":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.BiLevel, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "blur":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Blur, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "clrChange":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.ClrChange, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "clrRepl":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.ClrRepl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "duotone":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Duotone, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "fillOverlay":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.FillOverlay, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "grayscl":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Grayscl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "hsl":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Hsl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "lum":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Lum, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "tint":
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Tint, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Blip %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lBlip
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Blip and its children
func (m *Blip) Validate() error {
	return m.ValidateWithPath("Blip")
}

// ValidateWithPath validates the Blip and its children, prefixing error messages with path
func (m *Blip) ValidateWithPath(path string) error {
	if err := m.CT_Blip.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
