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

// ST_TransitionEightDirectionType is a union type
type ST_TransitionEightDirectionType struct {
	ST_TransitionSideDirectionType   ST_TransitionSideDirectionType
	ST_TransitionCornerDirectionType ST_TransitionCornerDirectionType
}

func (m *ST_TransitionEightDirectionType) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_TransitionEightDirectionType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TransitionSideDirectionType != ST_TransitionSideDirectionTypeUnset {
		mems = append(mems, "ST_TransitionSideDirectionType")
	}
	if m.ST_TransitionCornerDirectionType != ST_TransitionCornerDirectionTypeUnset {
		mems = append(mems, "ST_TransitionCornerDirectionType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_TransitionEightDirectionType) String() string {
	if m.ST_TransitionSideDirectionType != ST_TransitionSideDirectionTypeUnset {
		return m.ST_TransitionSideDirectionType.String()
	}
	if m.ST_TransitionCornerDirectionType != ST_TransitionCornerDirectionTypeUnset {
		return m.ST_TransitionCornerDirectionType.String()
	}
	return ""
}
