// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_DocPartBehaviors struct {
	// Entry Insertion Behavior
	Behavior []*CT_DocPartBehavior
}

func NewCT_DocPartBehaviors() *CT_DocPartBehaviors {
	ret := &CT_DocPartBehaviors{}
	return ret
}

func (m *CT_DocPartBehaviors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Behavior != nil {
		sebehavior := xml.StartElement{Name: xml.Name{Local: "w:behavior"}}
		e.EncodeElement(m.Behavior, sebehavior)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPartBehaviors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocPartBehaviors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "behavior":
				tmp := NewCT_DocPartBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Behavior = append(m.Behavior, tmp)
			default:
				log.Printf("skipping unsupported element on CT_DocPartBehaviors %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocPartBehaviors
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocPartBehaviors and its children
func (m *CT_DocPartBehaviors) Validate() error {
	return m.ValidateWithPath("CT_DocPartBehaviors")
}

// ValidateWithPath validates the CT_DocPartBehaviors and its children, prefixing error messages with path
func (m *CT_DocPartBehaviors) ValidateWithPath(path string) error {
	for i, v := range m.Behavior {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Behavior[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
