// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_DLbls struct {
	DLbl   []*CT_DLbl
	Choice *CT_DLblsChoice
	ExtLst *CT_ExtensionList
}

func NewCT_DLbls() *CT_DLbls {
	ret := &CT_DLbls{}
	return ret
}
func (m *CT_DLbls) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.DLbl != nil {
		sedLbl := xml.StartElement{Name: xml.Name{Local: "dLbl"}}
		e.EncodeElement(m.DLbl, sedLbl)
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DLbls) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DLbls:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dLbl":
				tmp := NewCT_DLbl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DLbl = append(m.DLbl, tmp)
			case "delete":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.Delete, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "numFmt":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.NumFmt, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "spPr":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.SpPr, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "txPr":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.TxPr, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "dLblPos":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.DLblPos, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showLegendKey":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.ShowLegendKey, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showVal":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.ShowVal, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showCatName":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.ShowCatName, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showSerName":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.ShowSerName, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showPercent":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.ShowPercent, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showBubbleSize":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.ShowBubbleSize, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "separator":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.Separator, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showLeaderLines":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.ShowLeaderLines, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "leaderLines":
				m.Choice = NewCT_DLblsChoice()
				if err := d.DecodeElement(&m.Choice.LeaderLines, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DLbls
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DLbls) Validate() error {
	return m.ValidateWithPath("CT_DLbls")
}
func (m *CT_DLbls) ValidateWithPath(path string) error {
	for i, v := range m.DLbl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DLbl[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
