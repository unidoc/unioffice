// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"github.com/unidoc/unioffice"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

type DataLabels struct {
	x *crt.CT_DLbls
}

func MakeDataLabels(x *crt.CT_DLbls) DataLabels {
	return DataLabels{x}
}
func (d DataLabels) ensureChoice() {
	if d.x.Choice == nil {
		d.x.Choice = crt.NewCT_DLblsChoice()
	}
}

func (d DataLabels) SetPosition(p crt.ST_DLblPos) {
	d.ensureChoice()
	d.x.Choice.DLblPos = crt.NewCT_DLblPos()
	d.x.Choice.DLblPos.ValAttr = p
}
func (d DataLabels) SetShowLegendKey(b bool) {
	d.ensureChoice()
	d.x.Choice.ShowLegendKey = crt.NewCT_Boolean()
	d.x.Choice.ShowLegendKey.ValAttr = unioffice.Bool(b)
}

func (d DataLabels) SetShowValue(b bool) {
	d.ensureChoice()
	d.x.Choice.ShowVal = crt.NewCT_Boolean()
	d.x.Choice.ShowVal.ValAttr = unioffice.Bool(b)
}

func (d DataLabels) SetShowCategoryName(b bool) {
	d.ensureChoice()
	d.x.Choice.ShowCatName = crt.NewCT_Boolean()
	d.x.Choice.ShowCatName.ValAttr = unioffice.Bool(b)
}

func (d DataLabels) SetShowSeriesName(b bool) {
	d.ensureChoice()
	d.x.Choice.ShowSerName = crt.NewCT_Boolean()
	d.x.Choice.ShowSerName.ValAttr = unioffice.Bool(b)
}

func (d DataLabels) SetShowPercent(b bool) {
	d.ensureChoice()
	d.x.Choice.ShowPercent = crt.NewCT_Boolean()
	d.x.Choice.ShowPercent.ValAttr = unioffice.Bool(b)
}

func (d DataLabels) SetShowLeaderLines(b bool) {
	d.ensureChoice()
	d.x.Choice.ShowLeaderLines = crt.NewCT_Boolean()
	d.x.Choice.ShowLeaderLines.ValAttr = unioffice.Bool(b)
}
