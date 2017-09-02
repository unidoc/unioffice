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
	"strconv"
)

type CT_BackgroundProperties struct {
	// Shade to Title
	ShadeToTitleAttr *bool
	ExtLst           *CT_ExtensionList
}

func NewCT_BackgroundProperties() *CT_BackgroundProperties {
	ret := &CT_BackgroundProperties{}
	return ret
}

func (m *CT_BackgroundProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ShadeToTitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shadeToTitle"},
			Value: fmt.Sprintf("%v", *m.ShadeToTitleAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BackgroundProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "shadeToTitle" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShadeToTitleAttr = &parsed
		}
	}
lCT_BackgroundProperties:
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
				log.Printf("skipping unsupported element on CT_BackgroundProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BackgroundProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BackgroundProperties and its children
func (m *CT_BackgroundProperties) Validate() error {
	return m.ValidateWithPath("CT_BackgroundProperties")
}

// ValidateWithPath validates the CT_BackgroundProperties and its children, prefixing error messages with path
func (m *CT_BackgroundProperties) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
