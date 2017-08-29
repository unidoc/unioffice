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
)

type CT_Revisions struct {
	// Revision Row Column Insert Delete
	Rrc []*CT_RevisionRowColumn
	// Revision Cell Move
	Rm []*CT_RevisionMove
	// Revision Custom View
	Rcv []*CT_RevisionCustomView
	// Revision Sheet Name
	Rsnm []*CT_RevisionSheetRename
	// Revision Insert Sheet
	Ris []*CT_RevisionInsertSheet
	// Revision Cell Change
	Rcc []*CT_RevisionCellChange
	// Revision Format
	Rfmt []*CT_RevisionFormatting
	// Revision AutoFormat
	Raf []*CT_RevisionAutoFormatting
	// Revision Defined Name
	Rdn []*CT_RevisionDefinedName
	// Revision Cell Comment
	Rcmt []*CT_RevisionComment
	// Revision Query Table
	Rqt []*CT_RevisionQueryTableField
	// Revision Merge Conflict
	Rcft []*CT_RevisionConflict
}

func NewCT_Revisions() *CT_Revisions {
	ret := &CT_Revisions{}
	return ret
}
func (m *CT_Revisions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Rrc != nil {
		serrc := xml.StartElement{Name: xml.Name{Local: "x:rrc"}}
		e.EncodeElement(m.Rrc, serrc)
	}
	if m.Rm != nil {
		serm := xml.StartElement{Name: xml.Name{Local: "x:rm"}}
		e.EncodeElement(m.Rm, serm)
	}
	if m.Rcv != nil {
		sercv := xml.StartElement{Name: xml.Name{Local: "x:rcv"}}
		e.EncodeElement(m.Rcv, sercv)
	}
	if m.Rsnm != nil {
		sersnm := xml.StartElement{Name: xml.Name{Local: "x:rsnm"}}
		e.EncodeElement(m.Rsnm, sersnm)
	}
	if m.Ris != nil {
		seris := xml.StartElement{Name: xml.Name{Local: "x:ris"}}
		e.EncodeElement(m.Ris, seris)
	}
	if m.Rcc != nil {
		sercc := xml.StartElement{Name: xml.Name{Local: "x:rcc"}}
		e.EncodeElement(m.Rcc, sercc)
	}
	if m.Rfmt != nil {
		serfmt := xml.StartElement{Name: xml.Name{Local: "x:rfmt"}}
		e.EncodeElement(m.Rfmt, serfmt)
	}
	if m.Raf != nil {
		seraf := xml.StartElement{Name: xml.Name{Local: "x:raf"}}
		e.EncodeElement(m.Raf, seraf)
	}
	if m.Rdn != nil {
		serdn := xml.StartElement{Name: xml.Name{Local: "x:rdn"}}
		e.EncodeElement(m.Rdn, serdn)
	}
	if m.Rcmt != nil {
		sercmt := xml.StartElement{Name: xml.Name{Local: "x:rcmt"}}
		e.EncodeElement(m.Rcmt, sercmt)
	}
	if m.Rqt != nil {
		serqt := xml.StartElement{Name: xml.Name{Local: "x:rqt"}}
		e.EncodeElement(m.Rqt, serqt)
	}
	if m.Rcft != nil {
		sercft := xml.StartElement{Name: xml.Name{Local: "x:rcft"}}
		e.EncodeElement(m.Rcft, sercft)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Revisions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Revisions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rrc":
				tmp := NewCT_RevisionRowColumn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rrc = append(m.Rrc, tmp)
			case "rm":
				tmp := NewCT_RevisionMove()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rm = append(m.Rm, tmp)
			case "rcv":
				tmp := NewCT_RevisionCustomView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcv = append(m.Rcv, tmp)
			case "rsnm":
				tmp := NewCT_RevisionSheetRename()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rsnm = append(m.Rsnm, tmp)
			case "ris":
				tmp := NewCT_RevisionInsertSheet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ris = append(m.Ris, tmp)
			case "rcc":
				tmp := NewCT_RevisionCellChange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcc = append(m.Rcc, tmp)
			case "rfmt":
				tmp := NewCT_RevisionFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rfmt = append(m.Rfmt, tmp)
			case "raf":
				tmp := NewCT_RevisionAutoFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Raf = append(m.Raf, tmp)
			case "rdn":
				tmp := NewCT_RevisionDefinedName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rdn = append(m.Rdn, tmp)
			case "rcmt":
				tmp := NewCT_RevisionComment()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcmt = append(m.Rcmt, tmp)
			case "rqt":
				tmp := NewCT_RevisionQueryTableField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rqt = append(m.Rqt, tmp)
			case "rcft":
				tmp := NewCT_RevisionConflict()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcft = append(m.Rcft, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Revisions
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Revisions) Validate() error {
	return m.ValidateWithPath("CT_Revisions")
}
func (m *CT_Revisions) ValidateWithPath(path string) error {
	for i, v := range m.Rrc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rrc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rm {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rm[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcv {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcv[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rsnm {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rsnm[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Ris {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ris[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rfmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rfmt[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Raf {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Raf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rdn {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rdn[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcmt[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rqt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rqt[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcft {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcft[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
