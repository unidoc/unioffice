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
	"log"
	"strconv"
	"time"
)

type CT_Array struct {
	LBoundsAttr  int32
	UBoundsAttr  int32
	BaseTypeAttr ST_ArrayBaseType
	Variant      []*Variant
	I1           []int8
	I2           []int16
	I4           []int32
	Int          []int32
	Ui1          []uint8
	Ui2          []uint16
	Ui4          []uint32
	Uint         []uint32
	R4           []float32
	R8           []float64
	Decimal      []float64
	Bstr         []string
	Date         []time.Time
	Bool         []bool
	Error        []string
	Cy           []string
}

func NewCT_Array() *CT_Array {
	ret := &CT_Array{}
	ret.BaseTypeAttr = ST_ArrayBaseType(1)
	return ret
}

func (m *CT_Array) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lBounds"},
		Value: fmt.Sprintf("%v", m.LBoundsAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uBounds"},
		Value: fmt.Sprintf("%v", m.UBoundsAttr)})
	attr, err := m.BaseTypeAttr.MarshalXMLAttr(xml.Name{Local: "baseType"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.Variant != nil {
		sevariant := xml.StartElement{Name: xml.Name{Local: "vt:variant"}}
		e.EncodeElement(m.Variant, sevariant)
	}
	if m.I1 != nil {
		sei1 := xml.StartElement{Name: xml.Name{Local: "vt:i1"}}
		e.EncodeElement(m.I1, sei1)
	}
	if m.I2 != nil {
		sei2 := xml.StartElement{Name: xml.Name{Local: "vt:i2"}}
		e.EncodeElement(m.I2, sei2)
	}
	if m.I4 != nil {
		sei4 := xml.StartElement{Name: xml.Name{Local: "vt:i4"}}
		e.EncodeElement(m.I4, sei4)
	}
	if m.Int != nil {
		seint := xml.StartElement{Name: xml.Name{Local: "vt:int"}}
		e.EncodeElement(m.Int, seint)
	}
	if m.Ui1 != nil {
		seui1 := xml.StartElement{Name: xml.Name{Local: "vt:ui1"}}
		e.EncodeElement(m.Ui1, seui1)
	}
	if m.Ui2 != nil {
		seui2 := xml.StartElement{Name: xml.Name{Local: "vt:ui2"}}
		e.EncodeElement(m.Ui2, seui2)
	}
	if m.Ui4 != nil {
		seui4 := xml.StartElement{Name: xml.Name{Local: "vt:ui4"}}
		e.EncodeElement(m.Ui4, seui4)
	}
	if m.Uint != nil {
		seuint := xml.StartElement{Name: xml.Name{Local: "vt:uint"}}
		e.EncodeElement(m.Uint, seuint)
	}
	if m.R4 != nil {
		ser4 := xml.StartElement{Name: xml.Name{Local: "vt:r4"}}
		e.EncodeElement(m.R4, ser4)
	}
	if m.R8 != nil {
		ser8 := xml.StartElement{Name: xml.Name{Local: "vt:r8"}}
		e.EncodeElement(m.R8, ser8)
	}
	if m.Decimal != nil {
		sedecimal := xml.StartElement{Name: xml.Name{Local: "vt:decimal"}}
		e.EncodeElement(m.Decimal, sedecimal)
	}
	if m.Bstr != nil {
		sebstr := xml.StartElement{Name: xml.Name{Local: "vt:bstr"}}
		e.EncodeElement(m.Bstr, sebstr)
	}
	if m.Date != nil {
		sedate := xml.StartElement{Name: xml.Name{Local: "vt:date"}}
		e.EncodeElement(m.Date, sedate)
	}
	if m.Bool != nil {
		sebool := xml.StartElement{Name: xml.Name{Local: "vt:bool"}}
		e.EncodeElement(m.Bool, sebool)
	}
	if m.Error != nil {
		seerror := xml.StartElement{Name: xml.Name{Local: "vt:error"}}
		e.EncodeElement(m.Error, seerror)
	}
	if m.Cy != nil {
		secy := xml.StartElement{Name: xml.Name{Local: "vt:cy"}}
		e.EncodeElement(m.Cy, secy)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Array) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BaseTypeAttr = ST_ArrayBaseType(1)
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
lCT_Array:
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
			break lCT_Array
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Array and its children
func (m *CT_Array) Validate() error {
	return m.ValidateWithPath("CT_Array")
}

// ValidateWithPath validates the CT_Array and its children, prefixing error messages with path
func (m *CT_Array) ValidateWithPath(path string) error {
	if m.BaseTypeAttr == ST_ArrayBaseTypeUnset {
		return fmt.Errorf("%s/BaseTypeAttr is a mandatory field", path)
	}
	if err := m.BaseTypeAttr.ValidateWithPath(path + "/BaseTypeAttr"); err != nil {
		return err
	}
	for i, v := range m.Variant {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Variant[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, s := range m.Error {
		if !ST_ErrorPatternRe.MatchString(s) {
			return fmt.Errorf(`%s/m.Error[%d] must match '%s' (have %v)`, path, i, ST_ErrorPatternRe, s)
		}
	}
	for i, s := range m.Cy {
		if !ST_CyPatternRe.MatchString(s) {
			return fmt.Errorf(`%s/m.Cy[%d] must match '%s' (have %v)`, path, i, ST_CyPatternRe, s)
		}
	}
	return nil
}
