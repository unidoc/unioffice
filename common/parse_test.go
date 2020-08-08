// Copyright 2018 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package common_test

import (
	"testing"

	"github.com/unidoc/unioffice/schema/soo/wml"
)

// Fields of these types must be integers per the spec, but Google doc
// writes out documents with floating point values. Relax our parsing
// somewhat so we can read these documents, truncating at the decimal
// point back to an integer.

// Issue #196

func TestParseGoogleDocsST_TwipsMeasure(t *testing.T) {
	ms, err := wml.ParseUnionST_TwipsMeasure("123.4")
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
	if *ms.ST_UnsignedDecimalNumber != 123 {
		t.Errorf("expected 123, got %#v", ms)
	}
}

func TestParseGoogleDocsST_MeasurementOrPercent(t *testing.T) {
	mp, err := wml.ParseUnionST_MeasurementOrPercent("123.4")
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
	if *mp.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage != 123 {
		t.Errorf("expected 123, got %#v", mp)
	}
}
