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
	"strconv"
)

type AG_Password struct {
	AlgorithmNameAttr *string
	HashValueAttr     *string
	SaltValueAttr     *string
	SpinCountAttr     *int32
}

func NewAG_Password() *AG_Password {
	ret := &AG_Password{}
	return ret
}
func (m *AG_Password) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:algorithmName"},
			Value: fmt.Sprintf("%v", *m.AlgorithmNameAttr)})
	}
	if m.HashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hashValue"},
			Value: fmt.Sprintf("%v", *m.HashValueAttr)})
	}
	if m.SaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:saltValue"},
			Value: fmt.Sprintf("%v", *m.SaltValueAttr)})
	}
	if m.SpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:spinCount"},
			Value: fmt.Sprintf("%v", *m.SpinCountAttr)})
	}
	start.Name.Local = "w:AG_Password"
	return nil
}
func (m *AG_Password) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "algorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgorithmNameAttr = &parsed
		}
		if attr.Name.Local == "hashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashValueAttr = &parsed
		}
		if attr.Name.Local == "saltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltValueAttr = &parsed
		}
		if attr.Name.Local == "spinCount" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.SpinCountAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Password: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *AG_Password) Validate() error {
	return m.ValidateWithPath("AG_Password")
}
func (m *AG_Password) ValidateWithPath(path string) error {
	return nil
}
