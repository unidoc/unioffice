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

// FormField is a form within a document. It references the document, so changes
// to the form field wil be reflected in the document if it is saved.
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
