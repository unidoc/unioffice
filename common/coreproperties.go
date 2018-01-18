// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"encoding/xml"
	"time"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/pkg/metadata/core_properties"
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

// Category returns the category of the document
func (c CoreProperties) Category() string {
	if c.x.Category != nil {
		return *c.x.Category
	}
	return ""
}

// SetCategory records the category of the document.
func (c CoreProperties) SetCategory(s string) {
	c.x.Category = &s
}

// ContentStatus returns the content status of the document (e.g. "Final", "Draft")
func (c CoreProperties) ContentStatus() string {
	if c.x.ContentStatus != nil {
		return *c.x.ContentStatus
	}
	return ""
}

// SetContentStatus records the content status of the document.
func (c CoreProperties) SetContentStatus(s string) {
	c.x.ContentStatus = &s
}

// Author returns the author of the document
func (c CoreProperties) Author() string {
	if c.x.Creator != nil {
		return string(c.x.Creator.Data)
	}
	return ""
}

// SetAuthor records the author of the document.
func (c CoreProperties) SetAuthor(s string) {
	if c.x.Creator == nil {
		c.x.Creator = &gooxml.XSDAny{XMLName: xml.Name{Local: "dc:creator"}}
	}
	c.x.Creator.Data = []byte(s)
}

// LastModifiedBy returns the name of the last person to modify the document
func (c CoreProperties) LastModifiedBy() string {
	if c.x.LastModifiedBy != nil {
		return *c.x.LastModifiedBy
	}
	return ""
}

// SetLastModifiedBy records the last person to modify the document.
func (c CoreProperties) SetLastModifiedBy(s string) {
	c.x.LastModifiedBy = &s
}

// SetLanguage records the language of the document.
func (c CoreProperties) SetLanguage(s string) {
	c.x.Language = &gooxml.XSDAny{XMLName: xml.Name{Local: "dc:language"}}
	c.x.Language.Data = []byte(s)
}

const cpTimeFormatW3CDTF = "2006-01-02T15:04:05Z"

func parseTime(x *gooxml.XSDAny) time.Time {
	if x == nil {
		return time.Time{}
	}

	// We should be checking the attributes as it can specify the format,
	// however I've only ever seen the W3CDTF format so I wouldn't know what code
	// to write to handle any other format anyway.  If you see another format
	// in the wild, please let me know.

	t, err := time.Parse(cpTimeFormatW3CDTF, string(x.Data))
	if err != nil {
		gooxml.Log("error parsing time from %s: %s", string(x.Data), err)
	}
	return t
}

// Created returns the time that the document was created.
func (c CoreProperties) Created() time.Time {
	return parseTime(c.x.Created)
}

func cpSetTime(t time.Time, name string) *gooxml.XSDAny {
	x := &gooxml.XSDAny{XMLName: xml.Name{Local: name}}
	x.Attrs = append(x.Attrs,
		xml.Attr{Name: xml.Name{Local: "xsi:type"}, Value: "dcterms:W3CDTF"})
	x.Attrs = append(x.Attrs,
		xml.Attr{Name: xml.Name{Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"})
	x.Attrs = append(x.Attrs,
		xml.Attr{Name: xml.Name{Local: "xmlns:dcterms"}, Value: "http://purl.org/dc/terms/"})
	x.Data = []byte(t.Format(cpTimeFormatW3CDTF))
	return x
}

// SetCreated sets the time that the document was created.
func (c CoreProperties) SetCreated(t time.Time) {
	c.x.Created = cpSetTime(t, "dcterms:created")
}

// Modified returns the time that the document was modified.
func (c CoreProperties) Modified() time.Time {
	return parseTime(c.x.Modified)
}

// SetModified sets the time that the document was modified.
func (c CoreProperties) SetModified(t time.Time) {
	c.x.Modified = cpSetTime(t, "dcterms:modified")
}

// Title returns the Title of the document
func (c CoreProperties) Title() string {
	if c.x.Title != nil {
		return string(c.x.Title.Data)
	}
	return ""
}

// SetTitle records the title of the document.
func (c CoreProperties) SetTitle(s string) {
	if c.x.Title == nil {
		c.x.Title = &gooxml.XSDAny{XMLName: xml.Name{Local: "dc:title"}}
	}
	c.x.Title.Data = []byte(s)
}

// Description returns the description of the document
func (c CoreProperties) Description() string {
	if c.x.Description != nil {
		return string(c.x.Description.Data)
	}
	return ""
}

// SetDescription records the description of the document.
func (c CoreProperties) SetDescription(s string) {
	if c.x.Description == nil {
		c.x.Description = &gooxml.XSDAny{XMLName: xml.Name{Local: "dc:description"}}
	}
	c.x.Description.Data = []byte(s)
}
