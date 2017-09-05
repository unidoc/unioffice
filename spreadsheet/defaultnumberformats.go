// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

// StandardFormat is a standard ECMA 376 number format.
type StandardFormat uint32

// StandardFormat constants, extracted from ECMA-376 Part 1 Section 18.8.30
const (
	StandardFormatGeneral  StandardFormat = 0  // General
	StandardFormat0        StandardFormat = 0  // General
	StandardFormat1        StandardFormat = 1  // 0
	StandardFormat2        StandardFormat = 2  // 0.00
	StandardFormat3        StandardFormat = 3  // #,##0
	StandardFormat4        StandardFormat = 4  // #,##0.00
	StandardFormatPercent  StandardFormat = 9  // 0%
	StandardFormat9        StandardFormat = 9  // 0%
	StandardFormat10       StandardFormat = 10 // 0.00%
	StandardFormat11       StandardFormat = 11 // 0.00E+00
	StandardFormat12       StandardFormat = 12 // # ?/?
	StandardFormat13       StandardFormat = 13 // # ??/??
	StandardFormatDate     StandardFormat = 14 // mm-dd-yy
	StandardFormat14       StandardFormat = 14 // mm-dd-yy
	StandardFormat15       StandardFormat = 15 // d-mmm-yy
	StandardFormat16       StandardFormat = 16 // d-mmm
	StandardFormat17       StandardFormat = 17 // mmm-yy
	StandardFormat18       StandardFormat = 18 // h:mm AM/PM
	StandardFormatTime     StandardFormat = 19 // h:mm:ss AM/PM
	StandardFormat19       StandardFormat = 19 // h:mm:ss AM/PM
	StandardFormat20       StandardFormat = 20 // h:mm
	StandardFormat21       StandardFormat = 21 // h:mm:ss
	StandardFormatDateTime StandardFormat = 22 // m/d/yy h:mm
	StandardFormat22       StandardFormat = 22 // m/d/yy h:mm
	StandardFormat37       StandardFormat = 37 // #,##0 ;(#,##0)
	StandardFormat38       StandardFormat = 38 // #,##0 ;[Red](#,##0)
	StandardFormat39       StandardFormat = 39 // #,##0.00;(#,##0.00)
	StandardFormat40       StandardFormat = 40 // #,##0.00;[Red](#,##0.00)
	StandardFormat45       StandardFormat = 45 // mm:ss
	StandardFormat46       StandardFormat = 46 // [h]:mm:ss
	StandardFormat47       StandardFormat = 47 // mmss.0
	StandardFormat48       StandardFormat = 48 // ##0.0E+0
	StandardFormat49       StandardFormat = 49 // @
)
