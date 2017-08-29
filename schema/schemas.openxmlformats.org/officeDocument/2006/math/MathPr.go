// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type MathPr struct {
	CT_MathPr
}

func NewMathPr() *MathPr {
	ret := &MathPr{}
	ret.CT_MathPr = *NewCT_MathPr()
	return ret
}
func (m *MathPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "m:mathPr"
	return m.CT_MathPr.MarshalXML(e, start)
}
func (m *MathPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_MathPr = *NewCT_MathPr()
lMathPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "mathFont":
				m.MathFont = NewCT_String()
				if err := d.DecodeElement(m.MathFont, &el); err != nil {
					return err
				}
			case "brkBin":
				m.BrkBin = NewCT_BreakBin()
				if err := d.DecodeElement(m.BrkBin, &el); err != nil {
					return err
				}
			case "brkBinSub":
				m.BrkBinSub = NewCT_BreakBinSub()
				if err := d.DecodeElement(m.BrkBinSub, &el); err != nil {
					return err
				}
			case "smallFrac":
				m.SmallFrac = NewCT_OnOff()
				if err := d.DecodeElement(m.SmallFrac, &el); err != nil {
					return err
				}
			case "dispDef":
				m.DispDef = NewCT_OnOff()
				if err := d.DecodeElement(m.DispDef, &el); err != nil {
					return err
				}
			case "lMargin":
				m.LMargin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.LMargin, &el); err != nil {
					return err
				}
			case "rMargin":
				m.RMargin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.RMargin, &el); err != nil {
					return err
				}
			case "defJc":
				m.DefJc = NewCT_OMathJc()
				if err := d.DecodeElement(m.DefJc, &el); err != nil {
					return err
				}
			case "preSp":
				m.PreSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.PreSp, &el); err != nil {
					return err
				}
			case "postSp":
				m.PostSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.PostSp, &el); err != nil {
					return err
				}
			case "interSp":
				m.InterSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.InterSp, &el); err != nil {
					return err
				}
			case "intraSp":
				m.IntraSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.IntraSp, &el); err != nil {
					return err
				}
			case "wrapIndent":
				m.Choice = NewCT_MathPrChoice()
				if err := d.DecodeElement(&m.Choice.WrapIndent, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapRight":
				m.Choice = NewCT_MathPrChoice()
				if err := d.DecodeElement(&m.Choice.WrapRight, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "intLim":
				m.IntLim = NewCT_LimLoc()
				if err := d.DecodeElement(m.IntLim, &el); err != nil {
					return err
				}
			case "naryLim":
				m.NaryLim = NewCT_LimLoc()
				if err := d.DecodeElement(m.NaryLim, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lMathPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *MathPr) Validate() error {
	return m.ValidateWithPath("MathPr")
}
func (m *MathPr) ValidateWithPath(path string) error {
	if err := m.CT_MathPr.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
