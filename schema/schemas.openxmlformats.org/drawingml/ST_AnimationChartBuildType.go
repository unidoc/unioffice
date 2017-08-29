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

// ST_AnimationChartBuildType is a union type
type ST_AnimationChartBuildType struct {
	ST_AnimationBuildType          ST_AnimationBuildType
	ST_AnimationChartOnlyBuildType ST_AnimationChartOnlyBuildType
}

func (m *ST_AnimationChartBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_AnimationChartBuildType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		mems = append(mems, "ST_AnimationBuildType")
	}
	if m.ST_AnimationChartOnlyBuildType != ST_AnimationChartOnlyBuildTypeUnset {
		mems = append(mems, "ST_AnimationChartOnlyBuildType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_AnimationChartBuildType) String() string {
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		return m.ST_AnimationBuildType.String()
	}
	if m.ST_AnimationChartOnlyBuildType != ST_AnimationChartOnlyBuildTypeUnset {
		return m.ST_AnimationChartOnlyBuildType.String()
	}
	return ""
}
