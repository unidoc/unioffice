// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_SlideMasterTextStyles struct {
	// Slide Master Title Text Style
	TitleStyle *drawingml.CT_TextListStyle
	// Slide Master Body Text Style
	BodyStyle *drawingml.CT_TextListStyle
	// Slide Master Other Text Style
	OtherStyle *drawingml.CT_TextListStyle
	ExtLst     *CT_ExtensionList
}

func NewCT_SlideMasterTextStyles() *CT_SlideMasterTextStyles {
	ret := &CT_SlideMasterTextStyles{}
	return ret
}
func (m *CT_SlideMasterTextStyles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.TitleStyle != nil {
		setitleStyle := xml.StartElement{Name: xml.Name{Local: "p:titleStyle"}}
		e.EncodeElement(m.TitleStyle, setitleStyle)
	}
	if m.BodyStyle != nil {
		sebodyStyle := xml.StartElement{Name: xml.Name{Local: "p:bodyStyle"}}
		e.EncodeElement(m.BodyStyle, sebodyStyle)
	}
	if m.OtherStyle != nil {
		seotherStyle := xml.StartElement{Name: xml.Name{Local: "p:otherStyle"}}
		e.EncodeElement(m.OtherStyle, seotherStyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SlideMasterTextStyles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideMasterTextStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "titleStyle":
				m.TitleStyle = drawingml.NewCT_TextListStyle()
				if err := d.DecodeElement(m.TitleStyle, &el); err != nil {
					return err
				}
			case "bodyStyle":
				m.BodyStyle = drawingml.NewCT_TextListStyle()
				if err := d.DecodeElement(m.BodyStyle, &el); err != nil {
					return err
				}
			case "otherStyle":
				m.OtherStyle = drawingml.NewCT_TextListStyle()
				if err := d.DecodeElement(m.OtherStyle, &el); err != nil {
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
			break lCT_SlideMasterTextStyles
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SlideMasterTextStyles) Validate() error {
	return m.ValidateWithPath("CT_SlideMasterTextStyles")
}
func (m *CT_SlideMasterTextStyles) ValidateWithPath(path string) error {
	if m.TitleStyle != nil {
		if err := m.TitleStyle.ValidateWithPath(path + "/TitleStyle"); err != nil {
			return err
		}
	}
	if m.BodyStyle != nil {
		if err := m.BodyStyle.ValidateWithPath(path + "/BodyStyle"); err != nil {
			return err
		}
	}
	if m.OtherStyle != nil {
		if err := m.OtherStyle.ValidateWithPath(path + "/OtherStyle"); err != nil {
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
