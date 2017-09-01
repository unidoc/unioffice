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

type CT_ParaRPr struct {
	// Inserted Paragraph
	Ins *CT_TrackChange
	// Deleted Paragraph
	Del *CT_TrackChange
	// Move Source Paragraph
	MoveFrom *CT_TrackChange
	// Move Destination Paragraph
	MoveTo     *CT_TrackChange
	EG_RPrBase []*EG_RPrBase
	// Revision Information for Run Properties on the Paragraph Mark
	RPrChange *CT_ParaRPrChange
}

func NewCT_ParaRPr() *CT_ParaRPr {
	ret := &CT_ParaRPr{}
	return ret
}
func (m *CT_ParaRPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Ins != nil {
		seins := xml.StartElement{Name: xml.Name{Local: "w:ins"}}
		e.EncodeElement(m.Ins, seins)
	}
	if m.Del != nil {
		sedel := xml.StartElement{Name: xml.Name{Local: "w:del"}}
		e.EncodeElement(m.Del, sedel)
	}
	if m.MoveFrom != nil {
		semoveFrom := xml.StartElement{Name: xml.Name{Local: "w:moveFrom"}}
		e.EncodeElement(m.MoveFrom, semoveFrom)
	}
	if m.MoveTo != nil {
		semoveTo := xml.StartElement{Name: xml.Name{Local: "w:moveTo"}}
		e.EncodeElement(m.MoveTo, semoveTo)
	}
	if m.EG_RPrBase != nil {
		for _, c := range m.EG_RPrBase {
			c.MarshalXML(e, start)
		}
	}
	if m.RPrChange != nil {
		serPrChange := xml.StartElement{Name: xml.Name{Local: "w:rPrChange"}}
		e.EncodeElement(m.RPrChange, serPrChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ParaRPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ParaRPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ins":
				m.Ins = NewCT_TrackChange()
				if err := d.DecodeElement(m.Ins, &el); err != nil {
					return err
				}
			case "del":
				m.Del = NewCT_TrackChange()
				if err := d.DecodeElement(m.Del, &el); err != nil {
					return err
				}
			case "moveFrom":
				m.MoveFrom = NewCT_TrackChange()
				if err := d.DecodeElement(m.MoveFrom, &el); err != nil {
					return err
				}
			case "moveTo":
				m.MoveTo = NewCT_TrackChange()
				if err := d.DecodeElement(m.MoveTo, &el); err != nil {
					return err
				}
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
			case "rPrChange":
				m.RPrChange = NewCT_ParaRPrChange()
				if err := d.DecodeElement(m.RPrChange, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ParaRPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ParaRPr) Validate() error {
	return m.ValidateWithPath("CT_ParaRPr")
}
func (m *CT_ParaRPr) ValidateWithPath(path string) error {
	if m.Ins != nil {
		if err := m.Ins.ValidateWithPath(path + "/Ins"); err != nil {
			return err
		}
	}
	if m.Del != nil {
		if err := m.Del.ValidateWithPath(path + "/Del"); err != nil {
			return err
		}
	}
	if m.MoveFrom != nil {
		if err := m.MoveFrom.ValidateWithPath(path + "/MoveFrom"); err != nil {
			return err
		}
	}
	if m.MoveTo != nil {
		if err := m.MoveTo.ValidateWithPath(path + "/MoveTo"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_RPrBase {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_RPrBase[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.RPrChange != nil {
		if err := m.RPrChange.ValidateWithPath(path + "/RPrChange"); err != nil {
			return err
		}
	}
	return nil
}
