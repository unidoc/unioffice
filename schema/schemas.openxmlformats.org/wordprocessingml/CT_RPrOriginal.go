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
	"log"
)

type CT_RPrOriginal struct {
	EG_RPrBase []*EG_RPrBase
}

func NewCT_RPrOriginal() *CT_RPrOriginal {
	ret := &CT_RPrOriginal{}
	return ret
}
func (m *CT_RPrOriginal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.EG_RPrBase != nil {
		for _, c := range m.EG_RPrBase {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RPrOriginal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RPrOriginal:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rStyle":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.RStyle = NewCT_String()
				if err := d.DecodeElement(tmprprbase.RStyle, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "rFonts":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.RFonts = NewCT_Fonts()
				if err := d.DecodeElement(tmprprbase.RFonts, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "b":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.B = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.B, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "bCs":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.BCs = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.BCs, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "i":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.I = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.I, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "iCs":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.ICs = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.ICs, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "caps":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Caps = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Caps, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "smallCaps":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.SmallCaps = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.SmallCaps, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "strike":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Strike = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Strike, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "dstrike":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Dstrike = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Dstrike, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "outline":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Outline = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Outline, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "shadow":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Shadow = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Shadow, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "emboss":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Emboss = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Emboss, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "imprint":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Imprint = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Imprint, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "noProof":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.NoProof = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.NoProof, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "snapToGrid":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.SnapToGrid = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.SnapToGrid, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "vanish":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Vanish = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Vanish, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "webHidden":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.WebHidden = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.WebHidden, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "color":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Color = NewCT_Color()
				if err := d.DecodeElement(tmprprbase.Color, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "spacing":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Spacing = NewCT_SignedTwipsMeasure()
				if err := d.DecodeElement(tmprprbase.Spacing, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "w":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.W = NewCT_TextScale()
				if err := d.DecodeElement(tmprprbase.W, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "kern":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Kern = NewCT_HpsMeasure()
				if err := d.DecodeElement(tmprprbase.Kern, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "position":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Position = NewCT_SignedHpsMeasure()
				if err := d.DecodeElement(tmprprbase.Position, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "sz":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Sz = NewCT_HpsMeasure()
				if err := d.DecodeElement(tmprprbase.Sz, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "szCs":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.SzCs = NewCT_HpsMeasure()
				if err := d.DecodeElement(tmprprbase.SzCs, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "highlight":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Highlight = NewCT_Highlight()
				if err := d.DecodeElement(tmprprbase.Highlight, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "u":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.U = NewCT_Underline()
				if err := d.DecodeElement(tmprprbase.U, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "effect":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Effect = NewCT_TextEffect()
				if err := d.DecodeElement(tmprprbase.Effect, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "bdr":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Bdr = NewCT_Border()
				if err := d.DecodeElement(tmprprbase.Bdr, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "shd":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Shd = NewCT_Shd()
				if err := d.DecodeElement(tmprprbase.Shd, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "fitText":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.FitText = NewCT_FitText()
				if err := d.DecodeElement(tmprprbase.FitText, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "vertAlign":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.VertAlign = NewCT_VerticalAlignRun()
				if err := d.DecodeElement(tmprprbase.VertAlign, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "rtl":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Rtl = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Rtl, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "cs":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Cs = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.Cs, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "em":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Em = NewCT_Em()
				if err := d.DecodeElement(tmprprbase.Em, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "lang":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.Lang = NewCT_Language()
				if err := d.DecodeElement(tmprprbase.Lang, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "eastAsianLayout":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.EastAsianLayout = NewCT_EastAsianLayout()
				if err := d.DecodeElement(tmprprbase.EastAsianLayout, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "specVanish":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.SpecVanish = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.SpecVanish, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			case "oMath":
				tmprprbase := NewEG_RPrBase()
				tmprprbase.OMath = NewCT_OnOff()
				if err := d.DecodeElement(tmprprbase.OMath, &el); err != nil {
					return err
				}
				m.EG_RPrBase = append(m.EG_RPrBase, tmprprbase)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RPrOriginal
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_RPrOriginal) Validate() error {
	return m.ValidateWithPath("CT_RPrOriginal")
}
func (m *CT_RPrOriginal) ValidateWithPath(path string) error {
	for i, v := range m.EG_RPrBase {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_RPrBase[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
