// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

// Extracted from ECMA-376 Part 1 Section 18.8.30
/*
0 General
1 0
2 0.00
3 #,##0
4 #,##0.00
9 0%
10 0.00%
11 0.00E+00
12 # ?/?
13 # ??/??
14 mm-dd-yy
15 d-mmm-yy
16 d-mmm
17 mmm-yy
18 h:mm AM/PM
19 h:mm:ss AM/PM
20 h:mm
21 h:mm:ss
22 m/d/yy h:mm
37 #,##0 ;(#,##0)
38 #,##0 ;[Red](#,##0)
39 #,##0.00;(#,##0.00)
40 #,##0.00;[Red](#,##0.00)
45 mm:ss
46 [h]:mm:ss
47 mmss.0
48 ##0.0E+0
49 @
*/

// StandardFormat is a standard ECMA 376 number format.
type StandardFormat uint32

// StandardFormat constants
const (
	StandardFormatGeneral  StandardFormat = 0
	StandardFormat0        StandardFormat = 0
	StandardFormat1        StandardFormat = 1
	StandardFormat2        StandardFormat = 2
	StandardFormat3        StandardFormat = 3
	StandardFormat4        StandardFormat = 4
	StandardFormatPercent  StandardFormat = 9
	StandardFormat9        StandardFormat = 9
	StandardFormat10       StandardFormat = 10
	StandardFormat11       StandardFormat = 11
	StandardFormat12       StandardFormat = 12
	StandardFormat13       StandardFormat = 13
	StandardFormatDate     StandardFormat = 14
	StandardFormat14       StandardFormat = 14
	StandardFormat15       StandardFormat = 15
	StandardFormat16       StandardFormat = 16
	StandardFormat17       StandardFormat = 17
	StandardFormat18       StandardFormat = 18
	StandardFormatTime     StandardFormat = 19
	StandardFormat19       StandardFormat = 19
	StandardFormat20       StandardFormat = 20
	StandardFormat21       StandardFormat = 21
	StandardFormatDateTime StandardFormat = 22
	StandardFormat22       StandardFormat = 22
	StandardFormat37       StandardFormat = 37
	StandardFormat38       StandardFormat = 38
	StandardFormat39       StandardFormat = 39
	StandardFormat40       StandardFormat = 40
	StandardFormat45       StandardFormat = 45
	StandardFormat46       StandardFormat = 46
	StandardFormat47       StandardFormat = 47
	StandardFormat48       StandardFormat = 48
	StandardFormat49       StandardFormat = 49
)
