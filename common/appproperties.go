// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/extended_properties"
)

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct {
	x *extended_properties.Properties
}

// NewAppProperties constructs a new AppProperties.
func NewAppProperties() AppProperties {
	p := AppProperties{x: extended_properties.NewProperties()}
	p.x.Application = gooxml.String("baliance.com/gooxml")
	p.x.AppVersion = gooxml.String("1.0")
	return p
}

// X returns the inner wrapped XML type.
func (a AppProperties) X() *extended_properties.Properties {
	return a.x
}
