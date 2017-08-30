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
	"strconv"
)

type CT_RevisionMove struct {
	// Sheet Id
	SheetIdAttr uint32
	// Source
	SourceAttr string
	// Destination
	DestinationAttr string
	// Source Sheet Id
	SourceSheetIdAttr *uint32
	// Undo
	Undo []*CT_UndoInfo
	// Revision Cell Change
	Rcc []*CT_RevisionCellChange
	// Revision Format
	Rfmt    []*CT_RevisionFormatting
	RIdAttr *uint32
	UaAttr  *bool
	RaAttr  *bool
}

func NewCT_RevisionMove() *CT_RevisionMove {
	ret := &CT_RevisionMove{}
	return ret
}
func (m *CT_RevisionMove) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "source"},
		Value: fmt.Sprintf("%v", m.SourceAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "destination"},
		Value: fmt.Sprintf("%v", m.DestinationAttr)})
	if m.SourceSheetIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceSheetId"},
			Value: fmt.Sprintf("%v", *m.SourceSheetIdAttr)})
	}
	if m.RIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rId"},
			Value: fmt.Sprintf("%v", *m.RIdAttr)})
	}
	if m.UaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ua"},
			Value: fmt.Sprintf("%v", *m.UaAttr)})
	}
	if m.RaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ra"},
			Value: fmt.Sprintf("%v", *m.RaAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Undo != nil {
		seundo := xml.StartElement{Name: xml.Name{Local: "x:undo"}}
		e.EncodeElement(m.Undo, seundo)
	}
	if m.Rcc != nil {
		sercc := xml.StartElement{Name: xml.Name{Local: "x:rcc"}}
		e.EncodeElement(m.Rcc, sercc)
	}
	if m.Rfmt != nil {
		serfmt := xml.StartElement{Name: xml.Name{Local: "x:rfmt"}}
		e.EncodeElement(m.Rfmt, serfmt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RevisionMove) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "source" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SourceAttr = parsed
		}
		if attr.Name.Local == "destination" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DestinationAttr = parsed
		}
		if attr.Name.Local == "sourceSheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SourceSheetIdAttr = &pt
		}
		if attr.Name.Local == "rId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RIdAttr = &pt
		}
		if attr.Name.Local == "ua" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UaAttr = &parsed
		}
		if attr.Name.Local == "ra" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RaAttr = &parsed
		}
	}
lCT_RevisionMove:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "undo":
				tmp := NewCT_UndoInfo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Undo = append(m.Undo, tmp)
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
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionMove
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_RevisionMove) Validate() error {
	return m.ValidateWithPath("CT_RevisionMove")
}
func (m *CT_RevisionMove) ValidateWithPath(path string) error {
	for i, v := range m.Undo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Undo[%d]", path, i)); err != nil {
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
	return nil
}
