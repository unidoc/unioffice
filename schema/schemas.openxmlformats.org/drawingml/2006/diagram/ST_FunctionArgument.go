// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"fmt"
)

// ST_FunctionArgument is a union type
type ST_FunctionArgument struct {
	ST_VariableType ST_VariableType
}

func (m *ST_FunctionArgument) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_FunctionArgument) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_VariableType != ST_VariableTypeUnset {
		mems = append(mems, "ST_VariableType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_FunctionArgument) String() string {
	if m.ST_VariableType != ST_VariableTypeUnset {
		return m.ST_VariableType.String()
	}
	return ""
}
