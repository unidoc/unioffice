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

type CT_TLMediaNodeVideo struct {
	// Full Screen
	FullScrnAttr *bool
	// Common Media Node Properties
	CMediaNode *CT_TLCommonMediaNodeData
}

func NewCT_TLMediaNodeVideo() *CT_TLMediaNodeVideo {
	ret := &CT_TLMediaNodeVideo{}
	ret.CMediaNode = NewCT_TLCommonMediaNodeData()
	return ret
}
func (m *CT_TLMediaNodeVideo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.FullScrnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fullScrn"},
			Value: fmt.Sprintf("%v", *m.FullScrnAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	secMediaNode := xml.StartElement{Name: xml.Name{Local: "p:cMediaNode"}}
	e.EncodeElement(m.CMediaNode, secMediaNode)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLMediaNodeVideo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CMediaNode = NewCT_TLCommonMediaNodeData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "fullScrn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FullScrnAttr = &parsed
		}
	}
lCT_TLMediaNodeVideo:
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
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLMediaNodeVideo
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLMediaNodeVideo) Validate() error {
	return m.ValidateWithPath("CT_TLMediaNodeVideo")
}
func (m *CT_TLMediaNodeVideo) ValidateWithPath(path string) error {
	if err := m.CMediaNode.ValidateWithPath(path + "/CMediaNode"); err != nil {
		return err
	}
	return nil
}
