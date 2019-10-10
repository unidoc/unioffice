//name: zhexiao
//date: 2019-10-10
//create object shape for document xml
//================================start
package wml

import (
	"encoding/xml"
	"fmt"
	"github.com/unidoc/unioffice"
)

type CT_Shape struct {
	IdAttr    *string
	TypeAttr  *string
	StyleAttr *string
	OleAttr   *string

	Imagedata *CT_Imagedata
}

func NewCT_Shape() *CT_Shape {
	ret := &CT_Shape{}
	return ret
}

func (m *CT_Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
			Value: fmt.Sprintf("%v", *m.TypeAttr)})
	}

	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}

	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}

	if m.OleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:ole"},
			Value: fmt.Sprintf("%v", *m.OleAttr)})
	}

	_ = e.EncodeToken(start)
	_ = e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}

		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = &parsed
			continue
		}

		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}

		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "ole" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OleAttr = &parsed
			continue
		}
	}
lCT_Shape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "imagedata"}:
				m.Imagedata = NewCT_Imagedata()
				if err := d.DecodeElement(&m.Imagedata, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Shape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Shape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObject and its children
func (m *CT_Shape) Validate() error {
	return m.ValidateWithPath("CT_Shape")
}

// ValidateWithPath validates the CT_OleObject and its children, prefixing error messages with path
func (m *CT_Shape) ValidateWithPath(path string) error {
	return nil
}

//================================end
