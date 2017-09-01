// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

// FormFieldType is the type of the form field.
//go:generate stringer -type=FormFieldType
type FormFieldType byte

// Form Field Type constants
const (
	FormFieldTypeUnknown FormFieldType = iota
	FormFieldTypeText
	FormFieldTypeCheckBox
	FormFieldTypeDropDown
)

// FormField is a form within a document.  It references the document, so
// changes to the form field wil be reflected in the document if it is saved.
type FormField struct {
	x      *wml.CT_FFData
	textIC *wml.EG_RunInnerContent
}

// Type returns the type of the field.
func (f FormField) Type() FormFieldType {
	if f.x.TextInput != nil {
		return FormFieldTypeText
	} else if f.x.CheckBox != nil {
		return FormFieldTypeCheckBox
	} else if f.x.DdList != nil {
		return FormFieldTypeDropDown
	}
	return FormFieldTypeUnknown
}

// Name returns the name of the field.
func (f FormField) Name() string {
	return *f.x.Name[0].ValAttr
}

// PossibleValues returns the possible values for a FormFieldTypeDropDown.
func (f FormField) PossibleValues() []string {
	if f.x.DdList == nil {
		return nil
	}
	ret := []string{}
	for _, s := range f.x.DdList.ListEntry {
		if s == nil {
			continue
		}
		ret = append(ret, s.ValAttr)
	}
	return ret
}

// SetValue sets the value of a FormFieldTypeText or FormFieldTypeDropDown.  For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (f FormField) SetValue(v string) {
	if f.x.DdList != nil {
		for i, s := range f.PossibleValues() {
			if s == v {
				f.x.DdList.Result = wml.NewCT_DecimalNumber()
				f.x.DdList.Result.ValAttr = int64(i)
				break
			}
		}
	} else if f.x.TextInput != nil {
		f.textIC.T = wml.NewCT_Text()
		f.textIC.T.Content = v
	}
}

// Value returns the tring value of a FormFieldTypeText or FormFieldTypeDropDown.
func (f FormField) Value() string {
	if f.x.TextInput != nil && f.textIC.T != nil {
		return f.textIC.T.Content
	} else if f.x.DdList != nil && f.x.DdList.Result != nil {
		pv := f.PossibleValues()
		idx := int(f.x.DdList.Result.ValAttr)
		if idx < len(pv) {
			return pv[idx]
		}
	} else if f.x.CheckBox != nil {
		if f.IsChecked() {
			return "true"
		}
		return "false"
	}
	return ""
}

// IsChecked returns true if a FormFieldTypeCheckBox is checked.
func (f FormField) IsChecked() bool {
	if f.x.CheckBox == nil {
		return false
	}
	if f.x.CheckBox.Checked != nil {
		return true
	}
	return false
}

// SetChecked marks a FormFieldTypeCheckBox as checked or unchecked.
func (f FormField) SetChecked(b bool) {
	if f.x.CheckBox == nil {
		return
	}
	if !b {
		f.x.CheckBox.Checked = nil
	} else {
		f.x.CheckBox.Checked = wml.NewCT_OnOff()
	}

}

// FindAllFields extracts all of the fields from a document.  They can then be
// manipulated via the methods on the field and the document can be saved.
func FindAllFields(d *Document) []FormField {
	ret := []FormField{}
	for _, p := range d.Paragraphs() {
		runs := p.Runs()
		for i, r := range runs {
			for _, ic := range r.x.EG_RunInnerContent {
				// skip non form fields
				if ic.FldChar == nil || ic.FldChar.FfData == nil {
					continue
				}

				// found a begin form field
				if ic.FldChar.FldCharTypeAttr == wml.ST_FldCharTypeBegin {
					// ensure it has a name
					if len(ic.FldChar.FfData.Name) == 0 || ic.FldChar.FfData.Name[0].ValAttr == nil {
						continue
					}

					field := FormField{x: ic.FldChar.FfData}
					// for text input boxes, we need a pointer to where to set
					// the text as well
					if ic.FldChar.FfData.TextInput != nil {

						// ensure we always have at lest two IC's
						for j := i + 1; j < len(runs)-1; j++ {
							if len(runs[j].x.EG_RunInnerContent) == 0 {
								continue
							}
							ic := runs[j].x.EG_RunInnerContent[0]
							// look for the 'separate' field
							if ic.FldChar != nil && ic.FldChar.FldCharTypeAttr == wml.ST_FldCharTypeSeparate {
								if len(runs[j+1].x.EG_RunInnerContent) == 0 {
									continue
								}
								// the value should be the text in the next inner content that is not a field char
								if runs[j+1].x.EG_RunInnerContent[0].FldChar == nil {
									field.textIC = runs[j+1].x.EG_RunInnerContent[0]
								}
							}
						}
					}
					ret = append(ret, field)
				}
			}
		}
	}
	return ret
}
