// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// NumberingLevel is the definition for numbering for a particular level within
// a NumberingDefinition.
type NumberingLevel struct {
	x *wml.CT_Lvl
}

// X returns the inner wrapped XML type.
func (n NumberingLevel) X() *wml.CT_Lvl {
	return n.x
}

// SetFormat sets the numbering format.
func (n NumberingLevel) SetFormat(f wml.ST_NumberFormat) {
	if n.x.NumFmt == nil {
		n.x.NumFmt = wml.NewCT_NumFmt()
	}
	n.x.NumFmt.ValAttr = f
}

// SetText sets the text to be used in bullet mode.
func (n NumberingLevel) SetText(t string) {
	if t == "" {
		n.x.LvlText = nil
	} else {
		n.x.LvlText = wml.NewCT_LevelText()
		n.x.LvlText.ValAttr = unioffice.String(t)
	}
}

// Properties returns the numbering level paragraph properties.
func (n NumberingLevel) Properties() ParagraphStyleProperties {
	if n.x.PPr == nil {
		n.x.PPr = wml.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{n.x.PPr}
}

// SetAlignment sets the paragraph alignment
func (n NumberingLevel) SetAlignment(j wml.ST_Jc) {
	if j == wml.ST_JcUnset {
		n.x.LvlJc = nil
	} else {
		n.x.LvlJc = wml.NewCT_Jc()
		n.x.LvlJc.ValAttr = j
	}
}

// RunProperties returns the RunProperties controlling numbering level font, etc.
func (n NumberingLevel) RunProperties() RunProperties {
	if n.x.RPr == nil {
		n.x.RPr = wml.NewCT_RPr()
	}
	return RunProperties{n.x.RPr}
}
