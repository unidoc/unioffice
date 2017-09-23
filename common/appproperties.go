// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"baliance.com/gooxml/schema/soo/ofc/extended_properties"
)

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct {
	x *extended_properties.Properties
}

// NewAppProperties constructs a new AppProperties.
func NewAppProperties() AppProperties {
	p := AppProperties{x: extended_properties.NewProperties()}
	p.SetApplication("baliance.com/gooxml")
	p.SetApplicationVersion("0.1")
	return p
}

// Application returns the name of the application that created the document.
// For gooxml created documents, it defaults to baliance.com/gooxml
func (a AppProperties) Application() string {
	if a.x.Application != nil {
		return *a.x.Application
	}
	return ""
}

// SetApplication sets the name of the application that created the document.
func (a AppProperties) SetApplication(s string) {
	a.x.Application = &s
}

// ApplicationVersion returns the version of the application that created the
// document.
func (a AppProperties) ApplicationVersion() string {
	if a.x.AppVersion != nil {
		return *a.x.AppVersion
	}
	return ""
}

// SetApplicationVersion sets the version of the application that created the document.
func (a AppProperties) SetApplicationVersion(s string) {
	a.x.AppVersion = &s
}

// X returns the inner wrapped XML type.
func (a AppProperties) X() *extended_properties.Properties {
	return a.x
}

// Company returns the name of the company that created the document.
// For gooxml created documents, it defaults to baliance.com/gooxml
func (a AppProperties) Company() string {
	if a.x.Company != nil {
		return *a.x.Company
	}
	return ""
}

// SetCompany sets the name of the company that created the document.
func (a AppProperties) SetCompany(s string) {
	a.x.Company = &s
}
