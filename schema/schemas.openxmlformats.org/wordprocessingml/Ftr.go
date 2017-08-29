// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

type Ftr struct {
	CT_HdrFtr
}

func NewFtr() *Ftr {
	ret := &Ftr{}
	ret.CT_HdrFtr = *NewCT_HdrFtr()
	return ret
}
func (m *Ftr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:ftr"
	return m.CT_HdrFtr.MarshalXML(e, start)
}
func (m *Ftr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_HdrFtr = *NewCT_HdrFtr()
lFtr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "altChunk":
				tmp := NewCT_AltChunk()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AltChunk = append(m.AltChunk, tmp)
			case "customXml":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.CustomXml = NewCT_CustomXmlBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case "sdt":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.Sdt = NewCT_SdtBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case "p":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_P()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.P = append(tmpcontentblockcontent.P, tmp)
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case "tbl":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_Tbl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.Tbl = append(tmpcontentblockcontent.Tbl, tmp)
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case "proofErr":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "ins":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "del":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveFrom":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveTo":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "bookmarkStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "bookmarkEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeStart":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeEnd":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "oMathPara":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case "oMath":
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lFtr
		case xml.CharData:
		}
	}
	return nil
}
func (m *Ftr) Validate() error {
	return m.ValidateWithPath("Ftr")
}
func (m *Ftr) ValidateWithPath(path string) error {
	if err := m.CT_HdrFtr.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
