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

type CT_DataRefs struct {
	// Data Consolidation Reference Count
	CountAttr *uint32
	// Data Consolidation Reference
	DataRef []*CT_DataRef
}

func NewCT_DataRefs() *CT_DataRefs {
	ret := &CT_DataRefs{}
	return ret
}

func (m *CT_DataRefs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.DataRef != nil {
		sedataRef := xml.StartElement{Name: xml.Name{Local: "x:dataRef"}}
		e.EncodeElement(m.DataRef, sedataRef)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataRefs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_DataRefs:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dataRef":
				tmp := NewCT_DataRef()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DataRef = append(m.DataRef, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataRefs
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DataRefs and its children
func (m *CT_DataRefs) Validate() error {
	return m.ValidateWithPath("CT_DataRefs")
}

// ValidateWithPath validates the CT_DataRefs and its children, prefixing error messages with path
func (m *CT_DataRefs) ValidateWithPath(path string) error {
	for i, v := range m.DataRef {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DataRef[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
