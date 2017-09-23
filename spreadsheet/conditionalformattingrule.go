// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/sml"
)

type ConditionalFormattingRule struct {
	x *sml.CT_CfRule
}

func (c ConditionalFormattingRule) InitializeDefaults() {
	c.SetType(sml.ST_CfTypeCellIs)
	c.SetOperator(sml.ST_ConditionalFormattingOperatorGreaterThan)
	c.SetPriority(1)

}

// X returns the inner wrapped XML type.
func (c ConditionalFormattingRule) X() *sml.CT_CfRule {
	return c.x
}

// SetConditionValue sets the condition value to be used for style applicaton.
func (c ConditionalFormattingRule) SetConditionValue(v string) {
	c.x.Formula = []string{v}
}

// Priority returns the rule priority
func (c ConditionalFormattingRule) Priority() int32 {
	return c.x.PriorityAttr
}

// SetPriority sets the rule priority
func (c ConditionalFormattingRule) SetPriority(p int32) {
	c.x.PriorityAttr = p
}

// Type returns the type of the rule
func (c ConditionalFormattingRule) Type() sml.ST_CfType {
	return c.x.TypeAttr
}

// SetType sets the type of the rule.
func (c ConditionalFormattingRule) SetType(t sml.ST_CfType) {
	c.x.TypeAttr = t
}

// Operator returns the operator for the rule
func (c ConditionalFormattingRule) Operator() sml.ST_ConditionalFormattingOperator {
	return c.x.OperatorAttr
}

// SetOperator sets the operator for the rule.
func (c ConditionalFormattingRule) SetOperator(t sml.ST_ConditionalFormattingOperator) {
	c.x.OperatorAttr = t
}

// SetStyle sets the style to be used for conditional rules
func (c ConditionalFormattingRule) SetStyle(d DifferentialStyle) {
	c.x.DxfIdAttr = gooxml.Uint32(d.Index())
}

func (c ConditionalFormattingRule) clear() {
	c.x.OperatorAttr = sml.ST_ConditionalFormattingOperatorUnset
	c.x.ColorScale = nil
	c.x.IconSet = nil
	c.x.Formula = nil
}

// SetColorScale configures the rule as a color scale, removing existing
// configuration.
func (c ConditionalFormattingRule) SetColorScale() ColorScale {
	c.clear()
	c.SetType(sml.ST_CfTypeColorScale)
	c.x.ColorScale = sml.NewCT_ColorScale()
	return ColorScale{c.x.ColorScale}
}

// SetIcons configures the rule as an icon scale, removing existing
// configuration.
func (c ConditionalFormattingRule) SetIcons() IconScale {
	c.clear()
	c.SetType(sml.ST_CfTypeIconSet)
	c.x.IconSet = sml.NewCT_IconSet()
	ic := IconScale{c.x.IconSet}
	ic.SetIcons(sml.ST_IconSetType3TrafficLights1)
	return ic
}

// SetDataBar configures the rule as a data bar, removing existing
// configuration.
func (c ConditionalFormattingRule) SetDataBar() DataBarScale {
	c.clear()
	c.SetType(sml.ST_CfTypeDataBar)
	c.x.DataBar = sml.NewCT_DataBar()
	db := DataBarScale{c.x.DataBar}
	db.SetShowValue(true)
	db.SetMinLength(10)
	db.SetMaxLength(90)
	return db
}
