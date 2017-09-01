// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"fmt"
)

// ST_TLTime is a union type
type ST_TLTime struct {
	Uint32              *uint32
	ST_TLTimeIndefinite ST_TLTimeIndefinite
}

func (m *ST_TLTime) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_TLTime) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Uint32 != nil {
		mems = append(mems, "Uint32")
	}
	if m.ST_TLTimeIndefinite != ST_TLTimeIndefiniteUnset {
		mems = append(mems, "ST_TLTimeIndefinite")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TLTime) String() string {
	if m.Uint32 != nil {
		return fmt.Sprintf("%v", *m.Uint32)
	}
	if m.ST_TLTimeIndefinite != ST_TLTimeIndefiniteUnset {
		return m.ST_TLTimeIndefinite.String()
	}
	return ""
}
