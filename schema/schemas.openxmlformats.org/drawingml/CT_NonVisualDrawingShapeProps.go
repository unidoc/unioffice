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
	"strconv"
)

type CT_NonVisualDrawingShapeProps struct {
	TxBoxAttr *bool
	SpLocks   *CT_ShapeLocking
	ExtLst    *CT_OfficeArtExtensionList
}

func NewCT_NonVisualDrawingShapeProps() *CT_NonVisualDrawingShapeProps {
	ret := &CT_NonVisualDrawingShapeProps{}
	return ret
}
func (m *CT_NonVisualDrawingShapeProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TxBoxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "txBox"},
			Value: fmt.Sprintf("%v", *m.TxBoxAttr)})
	}
	e.EncodeToken(start)
	if m.SpLocks != nil {
		sespLocks := xml.StartElement{Name: xml.Name{Local: "a:spLocks"}}
		e.EncodeElement(m.SpLocks, sespLocks)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NonVisualDrawingShapeProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "txBox" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TxBoxAttr = &parsed
		}
	}
lCT_NonVisualDrawingShapeProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "spLocks":
				m.SpLocks = NewCT_ShapeLocking()
				if err := d.DecodeElement(m.SpLocks, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_NonVisualDrawingShapeProps
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NonVisualDrawingShapeProps) Validate() error {
	return m.ValidateWithPath("CT_NonVisualDrawingShapeProps")
}
func (m *CT_NonVisualDrawingShapeProps) ValidateWithPath(path string) error {
	if m.SpLocks != nil {
		if err := m.SpLocks.ValidateWithPath(path + "/SpLocks"); err != nil {
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
