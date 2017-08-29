// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"fmt"
)

// ST_AnimationDgmBuildType is a union type
type ST_AnimationDgmBuildType struct {
	ST_AnimationBuildType        ST_AnimationBuildType
	ST_AnimationDgmOnlyBuildType ST_AnimationDgmOnlyBuildType
}

func (m *ST_AnimationDgmBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_AnimationDgmBuildType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		mems = append(mems, "ST_AnimationBuildType")
	}
	if m.ST_AnimationDgmOnlyBuildType != ST_AnimationDgmOnlyBuildTypeUnset {
		mems = append(mems, "ST_AnimationDgmOnlyBuildType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_AnimationDgmBuildType) String() string {
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		return m.ST_AnimationBuildType.String()
	}
	if m.ST_AnimationDgmOnlyBuildType != ST_AnimationDgmOnlyBuildTypeUnset {
		return m.ST_AnimationDgmOnlyBuildType.String()
	}
	return ""
}
