// Copyright 2018 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// ParagraphSpacing controls the spacing for a paragraph and its lines.
type ParagraphSpacing struct {
	x *wml.CT_Spacing
}

// SetBefore sets the spacing that comes before the paragraph.
func (p ParagraphSpacing) SetBefore(before measurement.Distance) {
	p.x.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.BeforeAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(before / measurement.Twips))
}

// SetAfter sets the spacing that comes after the paragraph.
func (p ParagraphSpacing) SetAfter(after measurement.Distance) {
	p.x.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.AfterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(after / measurement.Twips))
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (p ParagraphSpacing) SetLineSpacing(d measurement.Distance, rule wml.ST_LineSpacingRule) {
	if rule == wml.ST_LineSpacingRuleUnset {
		p.x.LineRuleAttr = wml.ST_LineSpacingRuleUnset
		p.x.LineAttr = nil
	} else {
		p.x.LineRuleAttr = rule
		p.x.LineAttr = &wml.ST_SignedTwipsMeasure{}
		p.x.LineAttr.Int64 = unioffice.Int64(int64(d / measurement.Twips))
	}
}

// SetBeforeAuto controls if spacing before a paragraph is automatically determined.
func (p ParagraphSpacing) SetBeforeAuto(b bool) {
	if b {
		p.x.BeforeAutospacingAttr = &sharedTypes.ST_OnOff{}
		p.x.BeforeAutospacingAttr.Bool = unioffice.Bool(true)
	} else {
		p.x.BeforeAutospacingAttr = nil
	}
}

// SetAfterAuto controls if spacing after a paragraph is automatically determined.
func (p ParagraphSpacing) SetAfterAuto(b bool) {
	if b {
		p.x.AfterAutospacingAttr = &sharedTypes.ST_OnOff{}
		p.x.AfterAutospacingAttr.Bool = unioffice.Bool(true)
	} else {
		p.x.AfterAutospacingAttr = nil
	}
}
