// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package presentation

import (
	"github.com/unidoc/unioffice/schema/soo/dml"
	"github.com/unidoc/unioffice/schema/soo/pml"
)

// ViewProperties contains presentation specific properties.
type ViewProperties struct {
	x *pml.ViewPr
}

// NewViewProperties constructs a new ViewProperties.
func NewViewProperties() ViewProperties {
	return ViewProperties{x: pml.NewViewPr()}
}

// X returns the inner wrapped XML type.
func (p ViewProperties) X() *pml.ViewPr {
	return p.x
}

// LastViewAttr returns the LastViewAttr property.
func (p ViewProperties) LastViewAttr() pml.ST_ViewType {
	return p.x.LastViewAttr
}

// ShowCommentsAttr returns the WebPr property.
func (p ViewProperties) ShowCommentsAttr() *bool {
	return p.x.ShowCommentsAttr
}

// NormalViewPr returns the NormalViewPr property.
func (p ViewProperties) NormalViewPr() *pml.CT_NormalViewProperties {
	return p.x.NormalViewPr
}

// SlideViewPr returns the SlideViewPr property.
func (p ViewProperties) SlideViewPr() *pml.CT_SlideViewProperties {
	return p.x.SlideViewPr
}

// OutlineViewPr returns the OutlineViewPr property.
func (p ViewProperties) OutlineViewPr() *pml.CT_OutlineViewProperties {
	return p.x.OutlineViewPr
}

// NotesTextViewPr returns the NotesTextViewPr property.
func (p ViewProperties) NotesTextViewPr() *pml.CT_NotesTextViewProperties {
	return p.x.NotesTextViewPr
}

// NotesViewPr returns the NotesViewPr property.
func (p ViewProperties) NotesViewPr() *pml.CT_NotesViewProperties {
	return p.x.NotesViewPr
}

// SorterViewPr returns the SorterViewPr property.
func (p ViewProperties) SorterViewPr() *pml.CT_SlideSorterViewProperties {
	return p.x.SorterViewPr
}

// GridSpacing returns the GridSpacing property.
func (p ViewProperties) GridSpacing() *dml.CT_PositiveSize2D {
	return p.x.GridSpacing
}

// ExtLst returns the ExtLst property.
func (p ViewProperties) ExtLst() *pml.CT_ExtensionList {
	return p.x.ExtLst
}
