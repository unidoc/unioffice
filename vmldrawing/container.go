// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vmldrawing

import (
	"encoding/xml"

	"github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml"
)

type Container struct {
	Layout    *vml.OfcShapelayout
	ShapeType *vml.Shapetype
	Shape     []*vml.Shape
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "urn:schemas-microsoft-com:office:excel"})
	start.Name.Local = "xml"
	e.EncodeToken(start)
	if c.Layout != nil {
		se := xml.StartElement{Name: xml.Name{Local: "o:shapelayout"}}
		e.EncodeElement(c.Layout, se)
	}
	if c.ShapeType != nil {
		se := xml.StartElement{Name: xml.Name{Local: "v:shapetype"}}
		e.EncodeElement(c.ShapeType, se)
	}
	for _, s := range c.Shape {
		se := xml.StartElement{Name: xml.Name{Local: "v:shape"}}
		e.EncodeElement(s, se)
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (c *Container) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	c.Shape = nil
outer:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "shapelayout":
				c.Layout = vml.NewOfcShapelayout()
				if err := d.DecodeElement(c.Layout, &el); err != nil {
					return err
				}
			case "shapetype":
				c.ShapeType = vml.NewShapetype()
				if err := d.DecodeElement(c.ShapeType, &el); err != nil {
					return err
				}
			case "shape":
				shp := vml.NewShape()
				if err := d.DecodeElement(shp, &el); err != nil {
					return err
				}
				c.Shape = append(c.Shape, shp)
			}
		case xml.EndElement:
			break outer
		}
	}
	return nil
}
