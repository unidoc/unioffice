// Copyright 2018 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

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
