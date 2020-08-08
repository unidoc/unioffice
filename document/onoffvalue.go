// Copyright 2018 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import "github.com/unidoc/unioffice/schema/soo/wml"

// OnOffValue represents an on/off value that can also be unset
type OnOffValue byte

// OnOffValue constants
const (
	OnOffValueUnset OnOffValue = iota
	OnOffValueOff
	OnOffValueOn
)

func convertOnOff(v *wml.CT_OnOff) OnOffValue {
	if v == nil {
		return OnOffValueUnset
	}
	// set, but the value is set to false
	if v.ValAttr != nil && v.ValAttr.Bool != nil && *v.ValAttr.Bool == false {
		return OnOffValueOff
	}
	// element exists, which implies turned on (and boolean value can't be false)
	return OnOffValueOn
}
