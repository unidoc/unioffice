// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"log"
	"strconv"
	"time"
)

type Array struct {
	CT_Array
}

func NewArray() *Array {
	ret := &Array{}
	ret.CT_Array = *NewCT_Array()
	return ret
}
func (m *Array) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.CT_Array.MarshalXML(e, start)
}
func (m *Array) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Array = *NewCT_Array()
	for _, attr := range start.Attr {
		if attr.Name.Local == "lBounds" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.LBoundsAttr = int32(parsed)
		}
		if attr.Name.Local == "uBounds" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.UBoundsAttr = int32(parsed)
		}
		if attr.Name.Local == "baseType" {
			m.BaseTypeAttr.UnmarshalXMLAttr(attr)
		}
	}
lArray:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "variant":
				tmp := NewVariant()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Variant = append(m.Variant, tmp)
			case "i1":
				var tmp int8
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I1 = append(m.I1, tmp)
			case "i2":
				var tmp int16
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I2 = append(m.I2, tmp)
			case "i4":
				var tmp int32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I4 = append(m.I4, tmp)
			case "int":
				var tmp int32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Int = append(m.Int, tmp)
			case "ui1":
				var tmp uint8
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui1 = append(m.Ui1, tmp)
			case "ui2":
				var tmp uint16
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui2 = append(m.Ui2, tmp)
			case "ui4":
				var tmp uint32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui4 = append(m.Ui4, tmp)
			case "uint":
				var tmp uint32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Uint = append(m.Uint, tmp)
			case "r4":
				var tmp float32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.R4 = append(m.R4, tmp)
			case "r8":
				var tmp float64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.R8 = append(m.R8, tmp)
			case "decimal":
				var tmp float64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Decimal = append(m.Decimal, tmp)
			case "bstr":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Bstr = append(m.Bstr, tmp)
			case "date":
				var tmp time.Time
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Date = append(m.Date, tmp)
			case "bool":
				var tmp bool
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Bool = append(m.Bool, tmp)
			case "error":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Error = append(m.Error, tmp)
			case "cy":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Cy = append(m.Cy, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lArray
		case xml.CharData:
		}
	}
	return nil
}
func (m *Array) Validate() error {
	return m.ValidateWithPath("Array")
}
func (m *Array) ValidateWithPath(path string) error {
	if err := m.CT_Array.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
