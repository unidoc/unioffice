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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Vector struct {
	BaseTypeAttr ST_VectorBaseType
	SizeAttr     uint32
	Variant      []*Variant
	I1           []int8
	I2           []int16
	I4           []int32
	I8           []int64
	Ui1          []uint8
	Ui2          []uint16
	Ui4          []uint32
	Ui8          []uint64
	R4           []float32
	R8           []float64
	Lpstr        []string
	Lpwstr       []string
	Bstr         []string
	Date         []time.Time
	Filetime     []time.Time
	Bool         []bool
	Cy           []string
	Error        []string
	Clsid        []string
}

func NewCT_Vector() *CT_Vector {
	ret := &CT_Vector{}
	ret.BaseTypeAttr = ST_VectorBaseType(1)
	return ret
}

func (m *CT_Vector) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.BaseTypeAttr.MarshalXMLAttr(xml.Name{Local: "baseType"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "size"},
		Value: fmt.Sprintf("%v", m.SizeAttr)})
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
	if m.I8 != nil {
		sei8 := xml.StartElement{Name: xml.Name{Local: "vt:i8"}}
		e.EncodeElement(m.I8, sei8)
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
	if m.Ui8 != nil {
		seui8 := xml.StartElement{Name: xml.Name{Local: "vt:ui8"}}
		e.EncodeElement(m.Ui8, seui8)
	}
	if m.R4 != nil {
		ser4 := xml.StartElement{Name: xml.Name{Local: "vt:r4"}}
		e.EncodeElement(m.R4, ser4)
	}
	if m.R8 != nil {
		ser8 := xml.StartElement{Name: xml.Name{Local: "vt:r8"}}
		e.EncodeElement(m.R8, ser8)
	}
	if m.Lpstr != nil {
		selpstr := xml.StartElement{Name: xml.Name{Local: "vt:lpstr"}}
		e.EncodeElement(m.Lpstr, selpstr)
	}
	if m.Lpwstr != nil {
		selpwstr := xml.StartElement{Name: xml.Name{Local: "vt:lpwstr"}}
		e.EncodeElement(m.Lpwstr, selpwstr)
	}
	if m.Bstr != nil {
		sebstr := xml.StartElement{Name: xml.Name{Local: "vt:bstr"}}
		e.EncodeElement(m.Bstr, sebstr)
	}
	if m.Date != nil {
		sedate := xml.StartElement{Name: xml.Name{Local: "vt:date"}}
		e.EncodeElement(m.Date, sedate)
	}
	if m.Filetime != nil {
		sefiletime := xml.StartElement{Name: xml.Name{Local: "vt:filetime"}}
		e.EncodeElement(m.Filetime, sefiletime)
	}
	if m.Bool != nil {
		sebool := xml.StartElement{Name: xml.Name{Local: "vt:bool"}}
		e.EncodeElement(m.Bool, sebool)
	}
	if m.Cy != nil {
		secy := xml.StartElement{Name: xml.Name{Local: "vt:cy"}}
		e.EncodeElement(m.Cy, secy)
	}
	if m.Error != nil {
		seerror := xml.StartElement{Name: xml.Name{Local: "vt:error"}}
		e.EncodeElement(m.Error, seerror)
	}
	if m.Clsid != nil {
		seclsid := xml.StartElement{Name: xml.Name{Local: "vt:clsid"}}
		e.EncodeElement(m.Clsid, seclsid)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Vector) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BaseTypeAttr = ST_VectorBaseType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "baseType" {
			m.BaseTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "size" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SizeAttr = uint32(parsed)
		}
	}
lCT_Vector:
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
			case "i8":
				var tmp int64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I8 = append(m.I8, tmp)
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
			case "ui8":
				var tmp uint64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui8 = append(m.Ui8, tmp)
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
			case "lpstr":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Lpstr = append(m.Lpstr, tmp)
			case "lpwstr":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Lpwstr = append(m.Lpwstr, tmp)
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
			case "filetime":
				var tmp time.Time
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Filetime = append(m.Filetime, tmp)
			case "bool":
				var tmp bool
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Bool = append(m.Bool, tmp)
			case "cy":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Cy = append(m.Cy, tmp)
			case "error":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Error = append(m.Error, tmp)
			case "clsid":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Clsid = append(m.Clsid, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Vector
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Vector and its children
func (m *CT_Vector) Validate() error {
	return m.ValidateWithPath("CT_Vector")
}

// ValidateWithPath validates the CT_Vector and its children, prefixing error messages with path
func (m *CT_Vector) ValidateWithPath(path string) error {
	if m.BaseTypeAttr == ST_VectorBaseTypeUnset {
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
	for i, s := range m.Cy {
		if !ST_CyPatternRe.MatchString(s) {
			return fmt.Errorf(`%s/m.Cy[%d] must match '%s' (have %v)`, path, i, ST_CyPatternRe, s)
		}
	}
	for i, s := range m.Error {
		if !ST_ErrorPatternRe.MatchString(s) {
			return fmt.Errorf(`%s/m.Error[%d] must match '%s' (have %v)`, path, i, ST_ErrorPatternRe, s)
		}
	}
	for i, s := range m.Clsid {
		if !sharedTypes.ST_GuidPatternRe.MatchString(s) {
			return fmt.Errorf(`%s/m.Clsid[%d] must match '%s' (have %v)`, path, i, sharedTypes.ST_GuidPatternRe, s)
		}
	}
	return nil
}
