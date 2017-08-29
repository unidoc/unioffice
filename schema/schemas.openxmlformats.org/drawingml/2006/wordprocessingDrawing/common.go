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

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

func ParseUnionST_Coordinate(s string) (drawingml.ST_Coordinate, error) {
	return drawingml.ParseUnionST_Coordinate(s)
}

type ST_WrapText byte

const (
	ST_WrapTextUnset     ST_WrapText = 0
	ST_WrapTextBothSides ST_WrapText = 1
	ST_WrapTextLeft      ST_WrapText = 2
	ST_WrapTextRight     ST_WrapText = 3
	ST_WrapTextLargest   ST_WrapText = 4
)

func (e ST_WrapText) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WrapTextUnset:
		attr.Value = ""
	case ST_WrapTextBothSides:
		attr.Value = "bothSides"
	case ST_WrapTextLeft:
		attr.Value = "left"
	case ST_WrapTextRight:
		attr.Value = "right"
	case ST_WrapTextLargest:
		attr.Value = "largest"
	}
	return attr, nil
}
func (e *ST_WrapText) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bothSides":
		*e = 1
	case "left":
		*e = 2
	case "right":
		*e = 3
	case "largest":
		*e = 4
	}
	return nil
}
func (m ST_WrapText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_WrapText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "bothSides":
			*m = 1
		case "left":
			*m = 2
		case "right":
			*m = 3
		case "largest":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_WrapText) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bothSides"
	case 2:
		return "left"
	case 3:
		return "right"
	case 4:
		return "largest"
	}
	return ""
}
func (m ST_WrapText) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_WrapText) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AlignH byte

const (
	ST_AlignHUnset   ST_AlignH = 0
	ST_AlignHLeft    ST_AlignH = 1
	ST_AlignHRight   ST_AlignH = 2
	ST_AlignHCenter  ST_AlignH = 3
	ST_AlignHInside  ST_AlignH = 4
	ST_AlignHOutside ST_AlignH = 5
)

func (e ST_AlignH) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AlignHUnset:
		attr.Value = ""
	case ST_AlignHLeft:
		attr.Value = "left"
	case ST_AlignHRight:
		attr.Value = "right"
	case ST_AlignHCenter:
		attr.Value = "center"
	case ST_AlignHInside:
		attr.Value = "inside"
	case ST_AlignHOutside:
		attr.Value = "outside"
	}
	return attr, nil
}
func (e *ST_AlignH) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "left":
		*e = 1
	case "right":
		*e = 2
	case "center":
		*e = 3
	case "inside":
		*e = 4
	case "outside":
		*e = 5
	}
	return nil
}
func (m ST_AlignH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_AlignH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "left":
			*m = 1
		case "right":
			*m = 2
		case "center":
			*m = 3
		case "inside":
			*m = 4
		case "outside":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_AlignH) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "left"
	case 2:
		return "right"
	case 3:
		return "center"
	case 4:
		return "inside"
	case 5:
		return "outside"
	}
	return ""
}
func (m ST_AlignH) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_AlignH) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RelFromH byte

const (
	ST_RelFromHUnset         ST_RelFromH = 0
	ST_RelFromHMargin        ST_RelFromH = 1
	ST_RelFromHPage          ST_RelFromH = 2
	ST_RelFromHColumn        ST_RelFromH = 3
	ST_RelFromHCharacter     ST_RelFromH = 4
	ST_RelFromHLeftMargin    ST_RelFromH = 5
	ST_RelFromHRightMargin   ST_RelFromH = 6
	ST_RelFromHInsideMargin  ST_RelFromH = 7
	ST_RelFromHOutsideMargin ST_RelFromH = 8
)

func (e ST_RelFromH) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RelFromHUnset:
		attr.Value = ""
	case ST_RelFromHMargin:
		attr.Value = "margin"
	case ST_RelFromHPage:
		attr.Value = "page"
	case ST_RelFromHColumn:
		attr.Value = "column"
	case ST_RelFromHCharacter:
		attr.Value = "character"
	case ST_RelFromHLeftMargin:
		attr.Value = "leftMargin"
	case ST_RelFromHRightMargin:
		attr.Value = "rightMargin"
	case ST_RelFromHInsideMargin:
		attr.Value = "insideMargin"
	case ST_RelFromHOutsideMargin:
		attr.Value = "outsideMargin"
	}
	return attr, nil
}
func (e *ST_RelFromH) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "margin":
		*e = 1
	case "page":
		*e = 2
	case "column":
		*e = 3
	case "character":
		*e = 4
	case "leftMargin":
		*e = 5
	case "rightMargin":
		*e = 6
	case "insideMargin":
		*e = 7
	case "outsideMargin":
		*e = 8
	}
	return nil
}
func (m ST_RelFromH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_RelFromH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "margin":
			*m = 1
		case "page":
			*m = 2
		case "column":
			*m = 3
		case "character":
			*m = 4
		case "leftMargin":
			*m = 5
		case "rightMargin":
			*m = 6
		case "insideMargin":
			*m = 7
		case "outsideMargin":
			*m = 8
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_RelFromH) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "margin"
	case 2:
		return "page"
	case 3:
		return "column"
	case 4:
		return "character"
	case 5:
		return "leftMargin"
	case 6:
		return "rightMargin"
	case 7:
		return "insideMargin"
	case 8:
		return "outsideMargin"
	}
	return ""
}
func (m ST_RelFromH) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_RelFromH) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AlignV byte

const (
	ST_AlignVUnset   ST_AlignV = 0
	ST_AlignVTop     ST_AlignV = 1
	ST_AlignVBottom  ST_AlignV = 2
	ST_AlignVCenter  ST_AlignV = 3
	ST_AlignVInside  ST_AlignV = 4
	ST_AlignVOutside ST_AlignV = 5
)

func (e ST_AlignV) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AlignVUnset:
		attr.Value = ""
	case ST_AlignVTop:
		attr.Value = "top"
	case ST_AlignVBottom:
		attr.Value = "bottom"
	case ST_AlignVCenter:
		attr.Value = "center"
	case ST_AlignVInside:
		attr.Value = "inside"
	case ST_AlignVOutside:
		attr.Value = "outside"
	}
	return attr, nil
}
func (e *ST_AlignV) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "bottom":
		*e = 2
	case "center":
		*e = 3
	case "inside":
		*e = 4
	case "outside":
		*e = 5
	}
	return nil
}
func (m ST_AlignV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_AlignV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "top":
			*m = 1
		case "bottom":
			*m = 2
		case "center":
			*m = 3
		case "inside":
			*m = 4
		case "outside":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_AlignV) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "top"
	case 2:
		return "bottom"
	case 3:
		return "center"
	case 4:
		return "inside"
	case 5:
		return "outside"
	}
	return ""
}
func (m ST_AlignV) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_AlignV) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RelFromV byte

const (
	ST_RelFromVUnset         ST_RelFromV = 0
	ST_RelFromVMargin        ST_RelFromV = 1
	ST_RelFromVPage          ST_RelFromV = 2
	ST_RelFromVParagraph     ST_RelFromV = 3
	ST_RelFromVLine          ST_RelFromV = 4
	ST_RelFromVTopMargin     ST_RelFromV = 5
	ST_RelFromVBottomMargin  ST_RelFromV = 6
	ST_RelFromVInsideMargin  ST_RelFromV = 7
	ST_RelFromVOutsideMargin ST_RelFromV = 8
)

func (e ST_RelFromV) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RelFromVUnset:
		attr.Value = ""
	case ST_RelFromVMargin:
		attr.Value = "margin"
	case ST_RelFromVPage:
		attr.Value = "page"
	case ST_RelFromVParagraph:
		attr.Value = "paragraph"
	case ST_RelFromVLine:
		attr.Value = "line"
	case ST_RelFromVTopMargin:
		attr.Value = "topMargin"
	case ST_RelFromVBottomMargin:
		attr.Value = "bottomMargin"
	case ST_RelFromVInsideMargin:
		attr.Value = "insideMargin"
	case ST_RelFromVOutsideMargin:
		attr.Value = "outsideMargin"
	}
	return attr, nil
}
func (e *ST_RelFromV) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "margin":
		*e = 1
	case "page":
		*e = 2
	case "paragraph":
		*e = 3
	case "line":
		*e = 4
	case "topMargin":
		*e = 5
	case "bottomMargin":
		*e = 6
	case "insideMargin":
		*e = 7
	case "outsideMargin":
		*e = 8
	}
	return nil
}
func (m ST_RelFromV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_RelFromV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "margin":
			*m = 1
		case "page":
			*m = 2
		case "paragraph":
			*m = 3
		case "line":
			*m = 4
		case "topMargin":
			*m = 5
		case "bottomMargin":
			*m = 6
		case "insideMargin":
			*m = 7
		case "outsideMargin":
			*m = 8
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_RelFromV) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "margin"
	case 2:
		return "page"
	case 3:
		return "paragraph"
	case 4:
		return "line"
	case 5:
		return "topMargin"
	case 6:
		return "bottomMargin"
	case 7:
		return "insideMargin"
	case 8:
		return "outsideMargin"
	}
	return ""
}
func (m ST_RelFromV) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_RelFromV) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_EffectExtent", NewCT_EffectExtent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_Inline", NewCT_Inline)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WrapPath", NewCT_WrapPath)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WrapNone", NewCT_WrapNone)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WrapSquare", NewCT_WrapSquare)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WrapTight", NewCT_WrapTight)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WrapThrough", NewCT_WrapThrough)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WrapTopBottom", NewCT_WrapTopBottom)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_PosH", NewCT_PosH)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_PosV", NewCT_PosV)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_Anchor", NewCT_Anchor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_TxbxContent", NewCT_TxbxContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_TextboxInfo", NewCT_TextboxInfo)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_LinkedTextboxInformation", NewCT_LinkedTextboxInformation)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WordprocessingShape", NewCT_WordprocessingShape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_GraphicFrame", NewCT_GraphicFrame)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WordprocessingContentPartNonVisual", NewCT_WordprocessingContentPartNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WordprocessingContentPart", NewCT_WordprocessingContentPart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WordprocessingGroup", NewCT_WordprocessingGroup)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "CT_WordprocessingCanvas", NewCT_WordprocessingCanvas)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "wpc", NewWpc)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "wgp", NewWgp)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "wsp", NewWsp)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "inline", NewInline)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "anchor", NewAnchor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", "EG_WrapType", NewEG_WrapType)
}
