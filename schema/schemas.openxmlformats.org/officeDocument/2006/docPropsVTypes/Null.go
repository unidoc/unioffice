// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
)

type Null struct {
	CT_Null
}

func NewNull() *Null {
	ret := &Null{}
	ret.CT_Null = *NewCT_Null()
	return ret
}
func (m *Null) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.CT_Null.MarshalXML(e, start)
}
func (m *Null) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Null = *NewCT_Null()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Null: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *Null) Validate() error {
	return m.ValidateWithPath("Null")
}
func (m *Null) ValidateWithPath(path string) error {
	if err := m.CT_Null.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
