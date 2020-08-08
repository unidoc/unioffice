// Copyright 2018 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
