// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"strings"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/pkg/content_types"
)

// ContentTypes is the top level "[Content_Types].xml" in a zip package.
type ContentTypes struct {
	x *content_types.Types
}

// NewContentTypes returns a wrapper around a newly constructed content-types.
func NewContentTypes() ContentTypes {
	ct := ContentTypes{x: content_types.NewTypes()}
	// add content type defaults
	ct.AddDefault("xml", "application/xml")
	ct.AddDefault("rels", "application/vnd.openxmlformats-package.relationships+xml")
	ct.AddDefault("png", "image/png")
	ct.AddDefault("jpeg", "image/jpeg")
	ct.AddDefault("jpg", "image/jpg")
	ct.AddDefault("wmf", "image/x-wmf")

	ct.AddOverride("/docProps/core.xml", "application/vnd.openxmlformats-package.core-properties+xml")
	ct.AddOverride("/docProps/app.xml", "application/vnd.openxmlformats-officedocument.extended-properties+xml")

	return ct
}

// X returns the inner raw content types.
func (c ContentTypes) X() *content_types.Types {
	return c.x
}

// AddDefault registers a default content type for a given file extension.
func (c ContentTypes) AddDefault(fileExtension string, contentType string) {
	def := content_types.NewDefault()
	def.ExtensionAttr = fileExtension
	def.ContentTypeAttr = contentType
	c.x.Default = append(c.x.Default, def)
}

// AddOverride adds an override content type for a given path name.
func (c ContentTypes) AddOverride(path, contentType string) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if strings.HasPrefix(contentType, "http") {
		gooxml.Log("content type '%s' is incorrect, must not start with http", contentType)
	}
	or := content_types.NewOverride()
	or.PartNameAttr = path
	or.ContentTypeAttr = contentType
	c.x.Override = append(c.x.Override, or)
}

// EnsureDefault esnures that an extension and default content type exist,
// adding it if necessary.
func (c ContentTypes) EnsureDefault(ext, contentType string) {
	for _, def := range c.x.Default {
		if def.ExtensionAttr == ext {
			def.ContentTypeAttr = contentType
			return
		}
	}

	def := &content_types.Default{}
	def.ContentTypeAttr = contentType
	def.ExtensionAttr = ext
	c.x.Default = append(c.x.Default, def)
}

// EnsureOverride ensures that an override for the given path exists, adding it if necessary
func (c ContentTypes) EnsureOverride(path, contentType string) {
	for _, ovr := range c.x.Override {
		// found one, so just ensure the content type matches and bail
		if ovr.PartNameAttr == path {
			if strings.HasPrefix(contentType, "http") {
				gooxml.Log("content type '%s' is incorrect, must not start with http", contentType)
			}
			ovr.ContentTypeAttr = contentType
			return
		}
	}

	// Didn't find a matching override for the target path, so add one
	c.AddOverride(path, contentType)
}

// RemoveOverride removes an override given a path.
func (c ContentTypes) RemoveOverride(path string) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	for i, ovr := range c.x.Override {
		if ovr.PartNameAttr == path {
			copy(c.x.Override[i:], c.x.Override[i+1:])
			c.x.Override = c.x.Override[0 : len(c.x.Override)-1]
		}
	}
}
