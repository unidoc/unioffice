// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"fmt"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/ofc/extended_properties"
)

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct {
	x *extended_properties.Properties
}

// NewAppProperties constructs a new AppProperties.
func NewAppProperties() AppProperties {
	p := AppProperties{x: extended_properties.NewProperties()}
	p.SetCompany("FoxyUtils ehf")
	p.SetApplication("github.com/unidoc/unioffice")
	p.SetDocSecurity(0)
	p.SetLinksUpToDate(false)

	var major, minor, patch int64
	fmt.Sscanf(Version, "%d.%d.%d", &major, &minor, &patch)
	f := float64(major) + float64(minor)/10000.0
	p.SetApplicationVersion(fmt.Sprintf("%07.4f", f))
	return p
}

// Application returns the name of the application that created the document.
// For gooxml created documents, it defaults to github.com/unidoc/unioffice
func (a AppProperties) Application() string {
	if a.x.Application != nil {
		return *a.x.Application
	}
	return ""
}

// SetLinksUpToDate sets the links up to date flag.
func (a AppProperties) SetLinksUpToDate(v bool) {
	a.x.LinksUpToDate = unioffice.Bool(v)
}

// SetDocSecurity sets the document security flag.
func (a AppProperties) SetDocSecurity(v int32) {
	a.x.DocSecurity = unioffice.Int32(v)
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

// SetApplicationVersion sets the version of the application that created the
// document.  Per MS, the verison string mut be in the form 'XX.YYYY'.
func (a AppProperties) SetApplicationVersion(s string) {
	a.x.AppVersion = &s
}

// X returns the inner wrapped XML type.
func (a AppProperties) X() *extended_properties.Properties {
	return a.x
}

// Company returns the name of the company that created the document.
// For gooxml created documents, it defaults to github.com/unidoc/unioffice
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
