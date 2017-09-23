// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml/schema/soo/wml"
)

// Style is a style within the styles.xml file.
type Style struct {
	x *wml.CT_Style
}

// X returns the inner wrapped XML type.
func (s Style) X() *wml.CT_Style {
	return s.x
}

// Type returns the type of the style.
func (s Style) Type() wml.ST_StyleType {
	return s.x.TypeAttr
}

// StyleID returns the style ID.
func (s Style) StyleID() string {
	if s.x.StyleIdAttr == nil {
		return ""
	}
	return *s.x.StyleIdAttr
}

// Name returns the name of the style if set.
func (s Style) Name() string {
	if s.x.Name == nil {
		return ""
	}
	return s.x.Name.ValAttr
}

// SetName sets the name of the style.
func (s Style) SetName(name string) {
	s.x.Name = wml.NewCT_String()
	s.x.Name.ValAttr = name
}

// SetPrimaryStyle marks the style as a primary style.
func (s Style) SetPrimaryStyle(b bool) {
	if b {
		s.x.QFormat = wml.NewCT_OnOff()
	} else {
		s.x.QFormat = nil
	}
}

// SetUISortOrder controls the order the style is displayed in the UI.
func (s Style) SetUISortOrder(order int) {
	s.x.UiPriority = wml.NewCT_DecimalNumber()
	s.x.UiPriority.ValAttr = int64(order)
}

// SetSemiHidden controls if the style is hidden in the UI.
func (s Style) SetSemiHidden(b bool) {
	if b {
		s.x.SemiHidden = wml.NewCT_OnOff()
	} else {
		s.x.SemiHidden = nil
	}
}

// SetUnhideWhenUsed controls if a semi hidden style becomes visible when used.
func (s Style) SetUnhideWhenUsed(b bool) {
	if b {
		s.x.UnhideWhenUsed = wml.NewCT_OnOff()
	} else {
		s.x.UnhideWhenUsed = nil
	}
}

// SetBasedOn sets the style that this style is based on.
func (s Style) SetBasedOn(name string) {
	if name == "" {
		s.x.BasedOn = nil
	} else {
		s.x.BasedOn = wml.NewCT_String()
		s.x.BasedOn.ValAttr = name
	}
}

// SetLinkedStyle sets the style that this style is linked to.
func (s Style) SetLinkedStyle(name string) {
	if name == "" {
		s.x.Link = nil
	} else {
		s.x.Link = wml.NewCT_String()
		s.x.Link.ValAttr = name
	}
}

// SetNextStyle sets the style that the next paragraph will use.
func (s Style) SetNextStyle(name string) {
	if name == "" {
		s.x.Next = nil
	} else {
		s.x.Next = wml.NewCT_String()
		s.x.Next.ValAttr = name
	}
}

// ParagraphProperties returns the paragraph style properties.
func (s Style) ParagraphProperties() ParagraphStyleProperties {
	if s.x.PPr == nil {
		s.x.PPr = wml.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{s.x.PPr}
}

// RunProperties returns the run style.
func (s Style) RunProperties() RunStyleProperties {
	if s.x.RPr == nil {
		s.x.RPr = wml.NewCT_RPr()
	}
	return RunStyleProperties{s.x.RPr}
}
