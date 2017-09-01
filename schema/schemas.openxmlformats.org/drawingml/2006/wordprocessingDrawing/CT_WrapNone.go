// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
)

type CT_WrapNone struct {
}

func NewCT_WrapNone() *CT_WrapNone {
	ret := &CT_WrapNone{}
	return ret
}

func (m *CT_WrapNone) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WrapNone) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WrapNone: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_WrapNone and its children
func (m *CT_WrapNone) Validate() error {
	return m.ValidateWithPath("CT_WrapNone")
}

// ValidateWithPath validates the CT_WrapNone and its children, prefixing error messages with path
func (m *CT_WrapNone) ValidateWithPath(path string) error {
	return nil
}
