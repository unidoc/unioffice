// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements

import (
	"encoding/xml"
	"fmt"
)

type AnyEl struct {
	SimpleLiteral
}

func NewAnyEl() *AnyEl {
	ret := &AnyEl{}
	ret.SimpleLiteral = *NewSimpleLiteral()
	return ret
}
func (m *AnyEl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.SimpleLiteral.MarshalXML(e, start)
}
func (m *AnyEl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SimpleLiteral = *NewSimpleLiteral()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AnyEl: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *AnyEl) Validate() error {
	return m.ValidateWithPath("AnyEl")
}
func (m *AnyEl) ValidateWithPath(path string) error {
	if err := m.SimpleLiteral.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
