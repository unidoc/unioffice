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

type CT_RevisionSheetRename struct {
	// Sheet Id
	SheetIdAttr uint32
	// Old Sheet Name
	OldNameAttr string
	// New Sheet Name
	NewNameAttr string
	ExtLst      *CT_ExtensionList
	RIdAttr     *uint32
	UaAttr      *bool
	RaAttr      *bool
}

func NewCT_RevisionSheetRename() *CT_RevisionSheetRename {
	ret := &CT_RevisionSheetRename{}
	return ret
}
func (m *CT_RevisionSheetRename) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldName"},
		Value: fmt.Sprintf("%v", m.OldNameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "newName"},
		Value: fmt.Sprintf("%v", m.NewNameAttr)})
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
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RevisionSheetRename) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "oldName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OldNameAttr = parsed
		}
		if attr.Name.Local == "newName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NewNameAttr = parsed
		}
		if attr.Name.Local == "rId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
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
lCT_RevisionSheetRename:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
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
			break lCT_RevisionSheetRename
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_RevisionSheetRename) Validate() error {
	return m.ValidateWithPath("CT_RevisionSheetRename")
}
func (m *CT_RevisionSheetRename) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
