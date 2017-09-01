// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_GuideList struct {
	// A Guide
	Guide []*CT_Guide
}

func NewCT_GuideList() *CT_GuideList {
	ret := &CT_GuideList{}
	return ret
}

func (m *CT_GuideList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Guide != nil {
		seguide := xml.StartElement{Name: xml.Name{Local: "p:guide"}}
		e.EncodeElement(m.Guide, seguide)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GuideList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GuideList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "guide":
				tmp := NewCT_Guide()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Guide = append(m.Guide, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GuideList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GuideList and its children
func (m *CT_GuideList) Validate() error {
	return m.ValidateWithPath("CT_GuideList")
}

// ValidateWithPath validates the CT_GuideList and its children, prefixing error messages with path
func (m *CT_GuideList) ValidateWithPath(path string) error {
	for i, v := range m.Guide {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Guide[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
