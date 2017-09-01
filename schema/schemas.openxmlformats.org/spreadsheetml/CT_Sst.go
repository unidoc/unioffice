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

type CT_Sst struct {
	// String Count
	CountAttr *uint32
	// Unique String Count
	UniqueCountAttr *uint32
	// String Item
	Si     []*CT_Rst
	ExtLst *CT_ExtensionList
}

func NewCT_Sst() *CT_Sst {
	ret := &CT_Sst{}
	return ret
}

func (m *CT_Sst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.UniqueCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueCount"},
			Value: fmt.Sprintf("%v", *m.UniqueCountAttr)})
	}
	e.EncodeToken(start)
	if m.Si != nil {
		sesi := xml.StartElement{Name: xml.Name{Local: "x:si"}}
		e.EncodeElement(m.Si, sesi)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Sst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		if attr.Name.Local == "uniqueCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.UniqueCountAttr = &pt
		}
	}
lCT_Sst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "si":
				tmp := NewCT_Rst()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Si = append(m.Si, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Sst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Sst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Sst and its children
func (m *CT_Sst) Validate() error {
	return m.ValidateWithPath("CT_Sst")
}

// ValidateWithPath validates the CT_Sst and its children, prefixing error messages with path
func (m *CT_Sst) ValidateWithPath(path string) error {
	for i, v := range m.Si {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Si[%d]", path, i)); err != nil {
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
