// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_BlipChoice struct {
	AlphaBiLevel []*CT_AlphaBiLevelEffect
	AlphaCeiling []*CT_AlphaCeilingEffect
	AlphaFloor   []*CT_AlphaFloorEffect
	AlphaInv     []*CT_AlphaInverseEffect
	AlphaMod     []*CT_AlphaModulateEffect
	AlphaModFix  []*CT_AlphaModulateFixedEffect
	AlphaRepl    []*CT_AlphaReplaceEffect
	BiLevel      []*CT_BiLevelEffect
	Blur         []*CT_BlurEffect
	ClrChange    []*CT_ColorChangeEffect
	ClrRepl      []*CT_ColorReplaceEffect
	Duotone      []*CT_DuotoneEffect
	FillOverlay  []*CT_FillOverlayEffect
	Grayscl      []*CT_GrayscaleEffect
	Hsl          []*CT_HSLEffect
	Lum          []*CT_LuminanceEffect
	Tint         []*CT_TintEffect
}

func NewCT_BlipChoice() *CT_BlipChoice {
	ret := &CT_BlipChoice{}
	return ret
}

func (m *CT_BlipChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AlphaBiLevel != nil {
		sealphaBiLevel := xml.StartElement{Name: xml.Name{Local: "a:alphaBiLevel"}}
		e.EncodeElement(m.AlphaBiLevel, sealphaBiLevel)
	}
	if m.AlphaCeiling != nil {
		sealphaCeiling := xml.StartElement{Name: xml.Name{Local: "a:alphaCeiling"}}
		e.EncodeElement(m.AlphaCeiling, sealphaCeiling)
	}
	if m.AlphaFloor != nil {
		sealphaFloor := xml.StartElement{Name: xml.Name{Local: "a:alphaFloor"}}
		e.EncodeElement(m.AlphaFloor, sealphaFloor)
	}
	if m.AlphaInv != nil {
		sealphaInv := xml.StartElement{Name: xml.Name{Local: "a:alphaInv"}}
		e.EncodeElement(m.AlphaInv, sealphaInv)
	}
	if m.AlphaMod != nil {
		sealphaMod := xml.StartElement{Name: xml.Name{Local: "a:alphaMod"}}
		e.EncodeElement(m.AlphaMod, sealphaMod)
	}
	if m.AlphaModFix != nil {
		sealphaModFix := xml.StartElement{Name: xml.Name{Local: "a:alphaModFix"}}
		e.EncodeElement(m.AlphaModFix, sealphaModFix)
	}
	if m.AlphaRepl != nil {
		sealphaRepl := xml.StartElement{Name: xml.Name{Local: "a:alphaRepl"}}
		e.EncodeElement(m.AlphaRepl, sealphaRepl)
	}
	if m.BiLevel != nil {
		sebiLevel := xml.StartElement{Name: xml.Name{Local: "a:biLevel"}}
		e.EncodeElement(m.BiLevel, sebiLevel)
	}
	if m.Blur != nil {
		seblur := xml.StartElement{Name: xml.Name{Local: "a:blur"}}
		e.EncodeElement(m.Blur, seblur)
	}
	if m.ClrChange != nil {
		seclrChange := xml.StartElement{Name: xml.Name{Local: "a:clrChange"}}
		e.EncodeElement(m.ClrChange, seclrChange)
	}
	if m.ClrRepl != nil {
		seclrRepl := xml.StartElement{Name: xml.Name{Local: "a:clrRepl"}}
		e.EncodeElement(m.ClrRepl, seclrRepl)
	}
	if m.Duotone != nil {
		seduotone := xml.StartElement{Name: xml.Name{Local: "a:duotone"}}
		e.EncodeElement(m.Duotone, seduotone)
	}
	if m.FillOverlay != nil {
		sefillOverlay := xml.StartElement{Name: xml.Name{Local: "a:fillOverlay"}}
		e.EncodeElement(m.FillOverlay, sefillOverlay)
	}
	if m.Grayscl != nil {
		segrayscl := xml.StartElement{Name: xml.Name{Local: "a:grayscl"}}
		e.EncodeElement(m.Grayscl, segrayscl)
	}
	if m.Hsl != nil {
		sehsl := xml.StartElement{Name: xml.Name{Local: "a:hsl"}}
		e.EncodeElement(m.Hsl, sehsl)
	}
	if m.Lum != nil {
		selum := xml.StartElement{Name: xml.Name{Local: "a:lum"}}
		e.EncodeElement(m.Lum, selum)
	}
	if m.Tint != nil {
		setint := xml.StartElement{Name: xml.Name{Local: "a:tint"}}
		e.EncodeElement(m.Tint, setint)
	}
	return nil
}

func (m *CT_BlipChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BlipChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "alphaBiLevel":
				tmp := NewCT_AlphaBiLevelEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaBiLevel = append(m.AlphaBiLevel, tmp)
			case "alphaCeiling":
				tmp := NewCT_AlphaCeilingEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaCeiling = append(m.AlphaCeiling, tmp)
			case "alphaFloor":
				tmp := NewCT_AlphaFloorEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaFloor = append(m.AlphaFloor, tmp)
			case "alphaInv":
				tmp := NewCT_AlphaInverseEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaInv = append(m.AlphaInv, tmp)
			case "alphaMod":
				tmp := NewCT_AlphaModulateEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaMod = append(m.AlphaMod, tmp)
			case "alphaModFix":
				tmp := NewCT_AlphaModulateFixedEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaModFix = append(m.AlphaModFix, tmp)
			case "alphaRepl":
				tmp := NewCT_AlphaReplaceEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaRepl = append(m.AlphaRepl, tmp)
			case "biLevel":
				tmp := NewCT_BiLevelEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BiLevel = append(m.BiLevel, tmp)
			case "blur":
				tmp := NewCT_BlurEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Blur = append(m.Blur, tmp)
			case "clrChange":
				tmp := NewCT_ColorChangeEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ClrChange = append(m.ClrChange, tmp)
			case "clrRepl":
				tmp := NewCT_ColorReplaceEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ClrRepl = append(m.ClrRepl, tmp)
			case "duotone":
				tmp := NewCT_DuotoneEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Duotone = append(m.Duotone, tmp)
			case "fillOverlay":
				tmp := NewCT_FillOverlayEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FillOverlay = append(m.FillOverlay, tmp)
			case "grayscl":
				tmp := NewCT_GrayscaleEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Grayscl = append(m.Grayscl, tmp)
			case "hsl":
				tmp := NewCT_HSLEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Hsl = append(m.Hsl, tmp)
			case "lum":
				tmp := NewCT_LuminanceEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Lum = append(m.Lum, tmp)
			case "tint":
				tmp := NewCT_TintEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tint = append(m.Tint, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BlipChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BlipChoice and its children
func (m *CT_BlipChoice) Validate() error {
	return m.ValidateWithPath("CT_BlipChoice")
}

// ValidateWithPath validates the CT_BlipChoice and its children, prefixing error messages with path
func (m *CT_BlipChoice) ValidateWithPath(path string) error {
	for i, v := range m.AlphaBiLevel {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaBiLevel[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaCeiling {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaCeiling[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaFloor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaFloor[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaInv {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaInv[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaMod {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaMod[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaModFix {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaModFix[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaRepl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaRepl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BiLevel {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BiLevel[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Blur {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Blur[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ClrChange {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ClrChange[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ClrRepl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ClrRepl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Duotone {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Duotone[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.FillOverlay {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FillOverlay[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Grayscl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Grayscl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Hsl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Hsl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Lum {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Lum[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Tint {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tint[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
