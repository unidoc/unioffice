//name: zhexiao(肖哲)
//date: 2019-10-10
//新增CT_Object下面的OleObject解析
//================================start
package wml

import (
	"encoding/xml"
	"fmt"
	"github.com/unidoc/unioffice"
)

type CT_OleObject struct {
	TypeAttr *string

	// Embedded Object ProgId
	ProgIdAttr *string

	// Shape Id
	ShapeIdAttr *string

	//r:id
	IdAttr *string
}

func NewCT_OleObject() *CT_OleObject {
	ret := &CT_OleObject{}
	return ret
}

func (m *CT_OleObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Type"},
			Value: fmt.Sprintf("%v", *m.TypeAttr)})
	}

	if m.ProgIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ProgID"},
			Value: fmt.Sprintf("%v", *m.ProgIdAttr)})
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ShapeID"},
		Value: fmt.Sprintf("%v", m.ShapeIdAttr)})

	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}

	_ = e.EncodeToken(start)
	_ = e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleObject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "Type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = &parsed
			continue
		}

		if attr.Name.Local == "ProgID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIdAttr = &parsed
			continue
		}

		if attr.Name.Local == "ShapeID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ShapeIdAttr = &parsed
			continue
		}

		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
	}
lCT_OleObject:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				unioffice.Log("skipping unsupported element on CT_OleObject %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObject
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObject and its children
func (m *CT_OleObject) Validate() error {
	return m.ValidateWithPath("CT_OleObject")
}

// ValidateWithPath validates the CT_OleObject and its children, prefixing error messages with path
func (m *CT_OleObject) ValidateWithPath(path string) error {
	return nil
}

//================================end
