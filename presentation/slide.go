// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentation

import (
	"errors"

	"baliance.com/gooxml/schema/soo/pml"
)

type Slide struct {
	sid *pml.CT_SlideIdListEntry
	x   *pml.Sld
}

// X returns the inner wrapped XML type.
func (s Slide) X() *pml.Sld {
	return s.x
}

// PlaceHolders returns all of the content place holders within a given slide.
func (s Slide) PlaceHolders() []PlaceHolder {
	ret := []PlaceHolder{}
	for _, spChc := range s.x.CSld.SpTree.Choice {
		for _, sp := range spChc.Sp {
			if sp.NvSpPr != nil && sp.NvSpPr.NvPr != nil && sp.NvSpPr.NvPr.Ph != nil {
				ret = append(ret, PlaceHolder{sp, s.x})
			}
		}
	}
	return ret
}

// GetPlaceholder returns a placeholder given its type.  If there are multiplace
// placeholders of the same type, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (s Slide) GetPlaceholder(t pml.ST_PlaceholderType) (PlaceHolder, error) {
	for _, spChc := range s.x.CSld.SpTree.Choice {
		for _, sp := range spChc.Sp {
			if sp.NvSpPr != nil && sp.NvSpPr.NvPr != nil && sp.NvSpPr.NvPr.Ph != nil {
				if sp.NvSpPr.NvPr.Ph.TypeAttr == t {
					return PlaceHolder{sp, s.x}, nil
				}
			}
		}
	}
	return PlaceHolder{}, errors.New("unable to find placeholder")
}

// GetPlaceholderByIndex returns a placeholder given its index.  If there are multiplace
// placeholders of the same index, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (s Slide) GetPlaceholderByIndex(idx uint32) (PlaceHolder, error) {
	for _, spChc := range s.x.CSld.SpTree.Choice {
		for _, sp := range spChc.Sp {
			if sp.NvSpPr != nil && sp.NvSpPr.NvPr != nil && sp.NvSpPr.NvPr.Ph != nil {
				if (idx == 0 && sp.NvSpPr.NvPr.Ph.IdxAttr == nil) ||
					(sp.NvSpPr.NvPr.Ph.IdxAttr != nil && *sp.NvSpPr.NvPr.Ph.IdxAttr == idx) {
					return PlaceHolder{sp, s.x}, nil
				}
			}
		}
	}
	return PlaceHolder{}, errors.New("unable to find placeholder")
}
