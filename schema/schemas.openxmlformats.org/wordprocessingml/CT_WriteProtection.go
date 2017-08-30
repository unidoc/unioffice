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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_WriteProtection struct {
	// Recommend Write Protection in User Interface
	RecommendedAttr                *sharedTypes.ST_OnOff
	CryptProviderTypeAttr          sharedTypes.ST_CryptProv
	CryptAlgorithmClassAttr        sharedTypes.ST_AlgClass
	CryptAlgorithmTypeAttr         sharedTypes.ST_AlgType
	CryptAlgorithmSidAttr          *int32
	CryptSpinCountAttr             *int32
	CryptProviderAttr              *string
	AlgIdExtAttr                   *string
	AlgIdExtSourceAttr             *string
	CryptProviderTypeExtAttr       *string
	CryptProviderTypeExtSourceAttr *string
	HashAttr                       *string
	SaltAttr                       *string
}

func NewCT_WriteProtection() *CT_WriteProtection {
	ret := &CT_WriteProtection{}
	return ret
}
func (m *CT_WriteProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RecommendedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:recommended"},
			Value: fmt.Sprintf("%v", *m.RecommendedAttr)})
	}
	if m.CryptProviderTypeAttr != sharedTypes.ST_CryptProvUnset {
		attr, err := m.CryptProviderTypeAttr.MarshalXMLAttr(xml.Name{Local: "w:cryptProviderType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmClassAttr != sharedTypes.ST_AlgClassUnset {
		attr, err := m.CryptAlgorithmClassAttr.MarshalXMLAttr(xml.Name{Local: "w:cryptAlgorithmClass"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmTypeAttr != sharedTypes.ST_AlgTypeUnset {
		attr, err := m.CryptAlgorithmTypeAttr.MarshalXMLAttr(xml.Name{Local: "w:cryptAlgorithmType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmSidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptAlgorithmSid"},
			Value: fmt.Sprintf("%v", *m.CryptAlgorithmSidAttr)})
	}
	if m.CryptSpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptSpinCount"},
			Value: fmt.Sprintf("%v", *m.CryptSpinCountAttr)})
	}
	if m.CryptProviderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptProvider"},
			Value: fmt.Sprintf("%v", *m.CryptProviderAttr)})
	}
	if m.AlgIdExtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:algIdExt"},
			Value: fmt.Sprintf("%v", *m.AlgIdExtAttr)})
	}
	if m.AlgIdExtSourceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:algIdExtSource"},
			Value: fmt.Sprintf("%v", *m.AlgIdExtSourceAttr)})
	}
	if m.CryptProviderTypeExtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptProviderTypeExt"},
			Value: fmt.Sprintf("%v", *m.CryptProviderTypeExtAttr)})
	}
	if m.CryptProviderTypeExtSourceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptProviderTypeExtSource"},
			Value: fmt.Sprintf("%v", *m.CryptProviderTypeExtSourceAttr)})
	}
	if m.HashAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hash"},
			Value: fmt.Sprintf("%v", *m.HashAttr)})
	}
	if m.SaltAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:salt"},
			Value: fmt.Sprintf("%v", *m.SaltAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WriteProtection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "recommended" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.RecommendedAttr = &parsed
		}
		if attr.Name.Local == "cryptProviderType" {
			m.CryptProviderTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cryptAlgorithmClass" {
			m.CryptAlgorithmClassAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cryptAlgorithmType" {
			m.CryptAlgorithmTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cryptAlgorithmSid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.CryptAlgorithmSidAttr = &pt
		}
		if attr.Name.Local == "cryptSpinCount" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.CryptSpinCountAttr = &pt
		}
		if attr.Name.Local == "cryptProvider" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderAttr = &parsed
		}
		if attr.Name.Local == "algIdExt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgIdExtAttr = &parsed
		}
		if attr.Name.Local == "algIdExtSource" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgIdExtSourceAttr = &parsed
		}
		if attr.Name.Local == "cryptProviderTypeExt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderTypeExtAttr = &parsed
		}
		if attr.Name.Local == "cryptProviderTypeExtSource" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderTypeExtSourceAttr = &parsed
		}
		if attr.Name.Local == "hash" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashAttr = &parsed
		}
		if attr.Name.Local == "salt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WriteProtection: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_WriteProtection) Validate() error {
	return m.ValidateWithPath("CT_WriteProtection")
}
func (m *CT_WriteProtection) ValidateWithPath(path string) error {
	if m.RecommendedAttr != nil {
		if err := m.RecommendedAttr.ValidateWithPath(path + "/RecommendedAttr"); err != nil {
			return err
		}
	}
	if err := m.CryptProviderTypeAttr.ValidateWithPath(path + "/CryptProviderTypeAttr"); err != nil {
		return err
	}
	if err := m.CryptAlgorithmClassAttr.ValidateWithPath(path + "/CryptAlgorithmClassAttr"); err != nil {
		return err
	}
	if err := m.CryptAlgorithmTypeAttr.ValidateWithPath(path + "/CryptAlgorithmTypeAttr"); err != nil {
		return err
	}
	return nil
}
