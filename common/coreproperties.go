// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/metadata/core_properties"
)

// CoreProperties contains document specific properties.
type CoreProperties struct {
	x *core_properties.CoreProperties
}

// NewCoreProperties constructs a new CoreProperties.
func NewCoreProperties() CoreProperties {
	return CoreProperties{x: core_properties.NewCoreProperties()}
}

// X returns the inner wrapped XML type.
func (c CoreProperties) X() *core_properties.CoreProperties {
	return c.x
}
