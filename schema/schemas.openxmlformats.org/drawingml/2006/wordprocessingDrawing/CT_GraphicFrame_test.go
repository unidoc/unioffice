// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
)

func TestCT_GraphicFrameConstructor(t *testing.T) {
	v := wordprocessingDrawing.NewCT_GraphicFrame()
	if v == nil {
		t.Errorf("wordprocessingDrawing.NewCT_GraphicFrame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingDrawing.CT_GraphicFrame should validate: %s", err)
	}
}

func TestCT_GraphicFrameMarshalUnmarshal(t *testing.T) {
	v := wordprocessingDrawing.NewCT_GraphicFrame()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingDrawing.NewCT_GraphicFrame()
	xml.Unmarshal(buf, v2)
}
