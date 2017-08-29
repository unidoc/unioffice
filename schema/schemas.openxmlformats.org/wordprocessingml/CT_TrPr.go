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

type CT_TrPr struct {
	// Table Row Conditional Formatting
	CnfStyle []*CT_Cnf
	// Associated HTML div ID
	DivId []*CT_DecimalNumber
	// Grid Columns Before First Cell
	GridBefore []*CT_DecimalNumber
	// Grid Columns After Last Cell
	GridAfter []*CT_DecimalNumber
	// Preferred Width Before Table Row
	WBefore []*CT_TblWidth
	// Preferred Width After Table Row
	WAfter []*CT_TblWidth
	// Table Row Cannot Break Across Pages
	CantSplit []*CT_OnOff
	// Table Row Height
	TrHeight []*CT_Height
	// Repeat Table Row on Every New Page
	TblHeader []*CT_OnOff
	// Table Row Cell Spacing
	TblCellSpacing []*CT_TblWidth
	// Table Row Alignment
	Jc []*CT_JcTable
	// Hidden Table Row Marker
	Hidden     []*CT_OnOff
	Ins        *CT_TrackChange
	Del        *CT_TrackChange
	TrPrChange *CT_TrPrChange
}

func NewCT_TrPr() *CT_TrPr {
	ret := &CT_TrPr{}
	return ret
}
func (m *CT_TrPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.CnfStyle != nil {
		secnfStyle := xml.StartElement{Name: xml.Name{Local: "w:cnfStyle"}}
		e.EncodeElement(m.CnfStyle, secnfStyle)
	}
	if m.DivId != nil {
		sedivId := xml.StartElement{Name: xml.Name{Local: "w:divId"}}
		e.EncodeElement(m.DivId, sedivId)
	}
	if m.GridBefore != nil {
		segridBefore := xml.StartElement{Name: xml.Name{Local: "w:gridBefore"}}
		e.EncodeElement(m.GridBefore, segridBefore)
	}
	if m.GridAfter != nil {
		segridAfter := xml.StartElement{Name: xml.Name{Local: "w:gridAfter"}}
		e.EncodeElement(m.GridAfter, segridAfter)
	}
	if m.WBefore != nil {
		sewBefore := xml.StartElement{Name: xml.Name{Local: "w:wBefore"}}
		e.EncodeElement(m.WBefore, sewBefore)
	}
	if m.WAfter != nil {
		sewAfter := xml.StartElement{Name: xml.Name{Local: "w:wAfter"}}
		e.EncodeElement(m.WAfter, sewAfter)
	}
	if m.CantSplit != nil {
		secantSplit := xml.StartElement{Name: xml.Name{Local: "w:cantSplit"}}
		e.EncodeElement(m.CantSplit, secantSplit)
	}
	if m.TrHeight != nil {
		setrHeight := xml.StartElement{Name: xml.Name{Local: "w:trHeight"}}
		e.EncodeElement(m.TrHeight, setrHeight)
	}
	if m.TblHeader != nil {
		setblHeader := xml.StartElement{Name: xml.Name{Local: "w:tblHeader"}}
		e.EncodeElement(m.TblHeader, setblHeader)
	}
	if m.TblCellSpacing != nil {
		setblCellSpacing := xml.StartElement{Name: xml.Name{Local: "w:tblCellSpacing"}}
		e.EncodeElement(m.TblCellSpacing, setblCellSpacing)
	}
	if m.Jc != nil {
		sejc := xml.StartElement{Name: xml.Name{Local: "w:jc"}}
		e.EncodeElement(m.Jc, sejc)
	}
	if m.Hidden != nil {
		sehidden := xml.StartElement{Name: xml.Name{Local: "w:hidden"}}
		e.EncodeElement(m.Hidden, sehidden)
	}
	if m.Ins != nil {
		seins := xml.StartElement{Name: xml.Name{Local: "w:ins"}}
		e.EncodeElement(m.Ins, seins)
	}
	if m.Del != nil {
		sedel := xml.StartElement{Name: xml.Name{Local: "w:del"}}
		e.EncodeElement(m.Del, sedel)
	}
	if m.TrPrChange != nil {
		setrPrChange := xml.StartElement{Name: xml.Name{Local: "w:trPrChange"}}
		e.EncodeElement(m.TrPrChange, setrPrChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TrPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TrPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cnfStyle":
				tmp := NewCT_Cnf()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CnfStyle = append(m.CnfStyle, tmp)
			case "divId":
				tmp := NewCT_DecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DivId = append(m.DivId, tmp)
			case "gridBefore":
				tmp := NewCT_DecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GridBefore = append(m.GridBefore, tmp)
			case "gridAfter":
				tmp := NewCT_DecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GridAfter = append(m.GridAfter, tmp)
			case "wBefore":
				tmp := NewCT_TblWidth()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.WBefore = append(m.WBefore, tmp)
			case "wAfter":
				tmp := NewCT_TblWidth()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.WAfter = append(m.WAfter, tmp)
			case "cantSplit":
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CantSplit = append(m.CantSplit, tmp)
			case "trHeight":
				tmp := NewCT_Height()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TrHeight = append(m.TrHeight, tmp)
			case "tblHeader":
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblHeader = append(m.TblHeader, tmp)
			case "tblCellSpacing":
				tmp := NewCT_TblWidth()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblCellSpacing = append(m.TblCellSpacing, tmp)
			case "jc":
				tmp := NewCT_JcTable()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Jc = append(m.Jc, tmp)
			case "hidden":
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Hidden = append(m.Hidden, tmp)
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
			case "trPrChange":
				m.TrPrChange = NewCT_TrPrChange()
				if err := d.DecodeElement(m.TrPrChange, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TrPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TrPr) Validate() error {
	return m.ValidateWithPath("CT_TrPr")
}
func (m *CT_TrPr) ValidateWithPath(path string) error {
	for i, v := range m.CnfStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CnfStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.DivId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DivId[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GridBefore {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GridBefore[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GridAfter {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GridAfter[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.WBefore {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/WBefore[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.WAfter {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/WAfter[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.CantSplit {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CantSplit[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.TrHeight {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TrHeight[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.TblHeader {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TblHeader[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.TblCellSpacing {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TblCellSpacing[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Jc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Jc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Hidden {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Hidden[%d]", path, i)); err != nil {
			return err
		}
	}
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
	if m.TrPrChange != nil {
		if err := m.TrPrChange.ValidateWithPath(path + "/TrPrChange"); err != nil {
			return err
		}
	}
	return nil
}
