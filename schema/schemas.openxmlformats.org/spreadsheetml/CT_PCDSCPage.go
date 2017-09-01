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

type CT_PCDSCPage struct {
	// Page Item String Count
	CountAttr *uint32
	// Page Item
	PageItem []*CT_PageItem
}

func NewCT_PCDSCPage() *CT_PCDSCPage {
	ret := &CT_PCDSCPage{}
	return ret
}

func (m *CT_PCDSCPage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.PageItem != nil {
		sepageItem := xml.StartElement{Name: xml.Name{Local: "x:pageItem"}}
		e.EncodeElement(m.PageItem, sepageItem)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PCDSCPage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_PCDSCPage:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pageItem":
				tmp := NewCT_PageItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PageItem = append(m.PageItem, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PCDSCPage
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PCDSCPage and its children
func (m *CT_PCDSCPage) Validate() error {
	return m.ValidateWithPath("CT_PCDSCPage")
}

// ValidateWithPath validates the CT_PCDSCPage and its children, prefixing error messages with path
func (m *CT_PCDSCPage) ValidateWithPath(path string) error {
	for i, v := range m.PageItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PageItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
