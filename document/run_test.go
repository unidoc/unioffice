// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import (
	"testing"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func TestRunClear(t *testing.T) {
	doc := New()
	para := doc.AddParagraph()
	run := para.AddRun()
	if len(run.X().EG_RunInnerContent) != 0 {
		t.Errorf("expected inner content of length zero, had %d", len(run.X().EG_RunInnerContent))
	}
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			run.AddText("test")
		} else {
			run.AddTab()
		}
		if len(run.X().EG_RunInnerContent) != i+1 {
			t.Errorf("expected inner content of length %d, had %d", i+1, len(run.X().EG_RunInnerContent))
		}
	}
	run.Clear()
	if len(run.X().EG_RunInnerContent) != 0 {
		t.Errorf("expected inner content of length zero, had %d", len(run.X().EG_RunInnerContent))
	}
}

// Issue #204
func TestRunPropertiesBold(t *testing.T) {
	r := RunProperties{wml.NewCT_RPr()}
	if r.IsBold() {
		t.Errorf("expected IsBold = false with no bold element")
	}
	r.x.B = wml.NewCT_OnOff()
	r.x.B.ValAttr = &sharedTypes.ST_OnOff{}
	r.x.B.ValAttr.Bool = unioffice.Bool(false)

	if r.IsBold() {
		t.Errorf("expected IsBold = false with false bool value")
	}

	r.x.B.ValAttr.Bool = unioffice.Bool(true)
	if !r.IsBold() {
		t.Errorf("expected IsBold = true with true bool value")
	}

	r.x.B = wml.NewCT_OnOff()
	if !r.IsBold() {
		t.Errorf("expected IsBold = true with existence and no bool value")
	}
}
