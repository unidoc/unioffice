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

type CT_CacheSource struct {
	// Cache Type
	TypeAttr ST_SourceType
	// Connection Index
	ConnectionIdAttr *uint32
	// Worksheet PivotCache Source
	WorksheetSource *CT_WorksheetSource
	// Consolidation Source
	Consolidation *CT_Consolidation
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_CacheSource() *CT_CacheSource {
	ret := &CT_CacheSource{}
	ret.TypeAttr = ST_SourceType(1)
	return ret
}
func (m *CT_CacheSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.ConnectionIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectionId"},
			Value: fmt.Sprintf("%v", *m.ConnectionIdAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.WorksheetSource != nil {
		seworksheetSource := xml.StartElement{Name: xml.Name{Local: "x:worksheetSource"}}
		e.EncodeElement(m.WorksheetSource, seworksheetSource)
	}
	if m.Consolidation != nil {
		seconsolidation := xml.StartElement{Name: xml.Name{Local: "x:consolidation"}}
		e.EncodeElement(m.Consolidation, seconsolidation)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CacheSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_SourceType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ConnectionIdAttr = &pt
		}
	}
lCT_CacheSource:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "worksheetSource":
				m.WorksheetSource = NewCT_WorksheetSource()
				if err := d.DecodeElement(m.WorksheetSource, &el); err != nil {
					return err
				}
			case "consolidation":
				m.Consolidation = NewCT_Consolidation()
				if err := d.DecodeElement(m.Consolidation, &el); err != nil {
					return err
				}
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
			break lCT_CacheSource
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CacheSource) Validate() error {
	return m.ValidateWithPath("CT_CacheSource")
}
func (m *CT_CacheSource) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_SourceTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.WorksheetSource != nil {
		if err := m.WorksheetSource.ValidateWithPath(path + "/WorksheetSource"); err != nil {
			return err
		}
	}
	if m.Consolidation != nil {
		if err := m.Consolidation.ValidateWithPath(path + "/Consolidation"); err != nil {
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
