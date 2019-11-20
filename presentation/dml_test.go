// Copyright 2018 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package presentation_test

import (
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml"
)

// Issue #207
func TestParseUnionST_AdjCoordinate(t *testing.T) {
	// this crashed due to a null pointer dereferences when not initializing the
	// returned value correctly
	dml.ParseUnionST_AdjCoordinate("123")
}
