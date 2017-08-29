// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements

import "baliance.com/gooxml"

func init() {
	gooxml.RegisterConstructor("http://purl.org/dc/elements/1.1/", "SimpleLiteral", NewSimpleLiteral)
	gooxml.RegisterConstructor("http://purl.org/dc/elements/1.1/", "elementContainer", NewElementContainer)
	gooxml.RegisterConstructor("http://purl.org/dc/elements/1.1/", "any", NewAny)
	gooxml.RegisterConstructor("http://purl.org/dc/elements/1.1/", "elementsGroup", NewElementsGroup)
}
