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
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

type CT_Body struct {
	EG_BlockLevelElts []*EG_BlockLevelElts
	// Document Final Section Properties
	SectPr *CT_SectPr
}

func NewCT_Body() *CT_Body {
	ret := &CT_Body{}
	return ret
}

func (m *CT_Body) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.EG_BlockLevelElts != nil {
		for _, c := range m.EG_BlockLevelElts {
			c.MarshalXML(e, start)
		}
	}
	if m.SectPr != nil {
		sesectPr := xml.StartElement{Name: xml.Name{Local: "w:sectPr"}}
		e.EncodeElement(m.SectPr, sesectPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Body:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "altChunk":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmp := NewCT_AltChunk()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpblocklevelelts.AltChunk = append(tmpblocklevelelts.AltChunk, tmp)
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
			case "customXml":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.CustomXml = NewCT_CustomXmlBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case "sdt":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.Sdt = NewCT_SdtBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case "p":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_P()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.P = append(tmpcontentblockcontent.P, tmp)
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case "tbl":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_Tbl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.Tbl = append(tmpcontentblockcontent.Tbl, tmp)
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case "proofErr":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "ins":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "del":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveFrom":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveTo":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case "bookmarkStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "bookmarkEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeStart":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeEnd":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "oMathPara":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case "oMath":
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case "sectPr":
				m.SectPr = NewCT_SectPr()
				if err := d.DecodeElement(m.SectPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Body %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Body
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Body and its children
func (m *CT_Body) Validate() error {
	return m.ValidateWithPath("CT_Body")
}

// ValidateWithPath validates the CT_Body and its children, prefixing error messages with path
func (m *CT_Body) ValidateWithPath(path string) error {
	for i, v := range m.EG_BlockLevelElts {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_BlockLevelElts[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.SectPr != nil {
		if err := m.SectPr.ValidateWithPath(path + "/SectPr"); err != nil {
			return err
		}
	}
	return nil
}
