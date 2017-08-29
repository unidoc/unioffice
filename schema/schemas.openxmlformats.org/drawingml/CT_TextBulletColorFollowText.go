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
)

type CT_TextBulletColorFollowText struct {
}

func NewCT_TextBulletColorFollowText() *CT_TextBulletColorFollowText {
	ret := &CT_TextBulletColorFollowText{}
	return ret
}
func (m *CT_TextBulletColorFollowText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TextBulletColorFollowText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextBulletColorFollowText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TextBulletColorFollowText) Validate() error {
	return m.ValidateWithPath("CT_TextBulletColorFollowText")
}
func (m *CT_TextBulletColorFollowText) ValidateWithPath(path string) error {
	return nil
}
