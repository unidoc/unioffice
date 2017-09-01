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

type EG_RunLevelElts struct {
	// Proofing Error Anchor
	ProofErr *CT_ProofErr
	// Range Permission Start
	PermStart *CT_PermStart
	// Range Permission End
	PermEnd *CT_Perm
	// Inserted Run Content
	Ins *CT_RunTrackChange
	// Deleted Run Content
	Del *CT_RunTrackChange
	// Move Source Run Content
	MoveFrom *CT_RunTrackChange
	// Move Destination Run Content
	MoveTo                 *CT_RunTrackChange
	EG_RangeMarkupElements []*EG_RangeMarkupElements
	EG_MathContent         []*EG_MathContent
}

func NewEG_RunLevelElts() *EG_RunLevelElts {
	ret := &EG_RunLevelElts{}
	return ret
}

func (m *EG_RunLevelElts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ProofErr != nil {
		seproofErr := xml.StartElement{Name: xml.Name{Local: "w:proofErr"}}
		e.EncodeElement(m.ProofErr, seproofErr)
	}
	if m.PermStart != nil {
		sepermStart := xml.StartElement{Name: xml.Name{Local: "w:permStart"}}
		e.EncodeElement(m.PermStart, sepermStart)
	}
	if m.PermEnd != nil {
		sepermEnd := xml.StartElement{Name: xml.Name{Local: "w:permEnd"}}
		e.EncodeElement(m.PermEnd, sepermEnd)
	}
	if m.Ins != nil {
		seins := xml.StartElement{Name: xml.Name{Local: "w:ins"}}
		e.EncodeElement(m.Ins, seins)
	}
	if m.Del != nil {
		sedel := xml.StartElement{Name: xml.Name{Local: "w:del"}}
		e.EncodeElement(m.Del, sedel)
	}
	if m.MoveFrom != nil {
		semoveFrom := xml.StartElement{Name: xml.Name{Local: "w:moveFrom"}}
		e.EncodeElement(m.MoveFrom, semoveFrom)
	}
	if m.MoveTo != nil {
		semoveTo := xml.StartElement{Name: xml.Name{Local: "w:moveTo"}}
		e.EncodeElement(m.MoveTo, semoveTo)
	}
	if m.EG_RangeMarkupElements != nil {
		for _, c := range m.EG_RangeMarkupElements {
			c.MarshalXML(e, start)
		}
	}
	if m.EG_MathContent != nil {
		for _, c := range m.EG_MathContent {
			c.MarshalXML(e, start)
		}
	}
	return nil
}

func (m *EG_RunLevelElts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_RunLevelElts:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "proofErr":
				m.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(m.ProofErr, &el); err != nil {
					return err
				}
			case "permStart":
				m.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(m.PermStart, &el); err != nil {
					return err
				}
			case "permEnd":
				m.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(m.PermEnd, &el); err != nil {
					return err
				}
			case "ins":
				m.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(m.Ins, &el); err != nil {
					return err
				}
			case "del":
				m.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(m.Del, &el); err != nil {
					return err
				}
			case "moveFrom":
				m.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(m.MoveFrom, &el); err != nil {
					return err
				}
			case "moveTo":
				m.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(m.MoveTo, &el); err != nil {
					return err
				}
			case "bookmarkStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "bookmarkEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeStart":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeEnd":
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RangeMarkupElements = append(m.EG_RangeMarkupElements, tmprangemarkupelements)
			case "oMathPara":
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_MathContent = append(m.EG_MathContent, tmpmathcontent)
			case "oMath":
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_MathContent = append(m.EG_MathContent, tmpmathcontent)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_RunLevelElts
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_RunLevelElts and its children
func (m *EG_RunLevelElts) Validate() error {
	return m.ValidateWithPath("EG_RunLevelElts")
}

// ValidateWithPath validates the EG_RunLevelElts and its children, prefixing error messages with path
func (m *EG_RunLevelElts) ValidateWithPath(path string) error {
	if m.ProofErr != nil {
		if err := m.ProofErr.ValidateWithPath(path + "/ProofErr"); err != nil {
			return err
		}
	}
	if m.PermStart != nil {
		if err := m.PermStart.ValidateWithPath(path + "/PermStart"); err != nil {
			return err
		}
	}
	if m.PermEnd != nil {
		if err := m.PermEnd.ValidateWithPath(path + "/PermEnd"); err != nil {
			return err
		}
	}
	if m.Ins != nil {
		if err := m.Ins.ValidateWithPath(path + "/Ins"); err != nil {
			return err
		}
	}
	if m.Del != nil {
		if err := m.Del.ValidateWithPath(path + "/Del"); err != nil {
			return err
		}
	}
	if m.MoveFrom != nil {
		if err := m.MoveFrom.ValidateWithPath(path + "/MoveFrom"); err != nil {
			return err
		}
	}
	if m.MoveTo != nil {
		if err := m.MoveTo.ValidateWithPath(path + "/MoveTo"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_RangeMarkupElements {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_RangeMarkupElements[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.EG_MathContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_MathContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
