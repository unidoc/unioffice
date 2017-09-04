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

	"baliance.com/gooxml"
)

type CT_TLMediaNodeAudio struct {
	// Is Narration
	IsNarrationAttr *bool
	// Common Media Node Properties
	CMediaNode *CT_TLCommonMediaNodeData
}

func NewCT_TLMediaNodeAudio() *CT_TLMediaNodeAudio {
	ret := &CT_TLMediaNodeAudio{}
	ret.IsNarrationAttr = gooxml.Bool(false)
	ret.CMediaNode = NewCT_TLCommonMediaNodeData()
	return ret
}

func (m *CT_TLMediaNodeAudio) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IsNarrationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "isNarration"},
			Value: fmt.Sprintf("%d", b2i(*m.IsNarrationAttr))})
	}
	e.EncodeToken(start)
	secMediaNode := xml.StartElement{Name: xml.Name{Local: "p:cMediaNode"}}
	e.EncodeElement(m.CMediaNode, secMediaNode)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLMediaNodeAudio) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.IsNarrationAttr = gooxml.Bool(false)
	m.CMediaNode = NewCT_TLCommonMediaNodeData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "isNarration" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IsNarrationAttr = &parsed
		}
	}
lCT_TLMediaNodeAudio:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cMediaNode":
				if err := d.DecodeElement(m.CMediaNode, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLMediaNodeAudio %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLMediaNodeAudio
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLMediaNodeAudio and its children
func (m *CT_TLMediaNodeAudio) Validate() error {
	return m.ValidateWithPath("CT_TLMediaNodeAudio")
}

// ValidateWithPath validates the CT_TLMediaNodeAudio and its children, prefixing error messages with path
func (m *CT_TLMediaNodeAudio) ValidateWithPath(path string) error {
	if err := m.CMediaNode.ValidateWithPath(path + "/CMediaNode"); err != nil {
		return err
	}
	return nil
}
