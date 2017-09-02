// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/presentationml"
)

func TestNotesMasterConstructor(t *testing.T) {
	v := presentationml.NewNotesMaster()
	if v == nil {
		t.Errorf("presentationml.NewNotesMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed presentationml.NotesMaster should validate: %s", err)
	}
}

func TestNotesMasterMarshalUnmarshal(t *testing.T) {
	v := presentationml.NewNotesMaster()
	buf, _ := xml.Marshal(v)
	v2 := presentationml.NewNotesMaster()
	xml.Unmarshal(buf, v2)
}
