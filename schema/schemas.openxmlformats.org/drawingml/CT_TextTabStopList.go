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
)

type CT_TextTabStopList struct {
	Tab []*CT_TextTabStop
}

func NewCT_TextTabStopList() *CT_TextTabStopList {
	ret := &CT_TextTabStopList{}
	return ret
}

func (m *CT_TextTabStopList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Tab != nil {
		setab := xml.StartElement{Name: xml.Name{Local: "a:tab"}}
		e.EncodeElement(m.Tab, setab)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextTabStopList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextTabStopList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tab":
				tmp := NewCT_TextTabStop()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tab = append(m.Tab, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextTabStopList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextTabStopList and its children
func (m *CT_TextTabStopList) Validate() error {
	return m.ValidateWithPath("CT_TextTabStopList")
}

// ValidateWithPath validates the CT_TextTabStopList and its children, prefixing error messages with path
func (m *CT_TextTabStopList) ValidateWithPath(path string) error {
	for i, v := range m.Tab {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tab[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
