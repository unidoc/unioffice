//line isnumber.rl:1

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format

//line isnumber.go:18
const isnumber_start int = 0
const isnumber_first_final int = 0
const isnumber_error int = -1

const isnumber_en_main int = 0

//line isnumber.rl:27
func IsNumber(data string) (isNumber bool) {
	cs, p, pe := 0, 0, len(data)
	eof := len(data)
	ts, te, act := 0, 0, 0
	_ = te
	_ = act
	_ = ts

//line isnumber.go:38
	{
		cs = isnumber_start
		ts = 0
		te = 0
		act = 0
	}

//line isnumber.go:46
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		}
		goto st_out
	tr3:
//line isnumber.rl:25
		te = p
		p--
		{
			isNumber = false
		}
		goto st0
	tr4:
//line isnumber.rl:23
		te = p
		p--
		{
			isNumber = te == len(data)
		}
		goto st0
	tr7:
//line isnumber.rl:24
		te = p
		p--
		{
			isNumber = te == len(data)
		}
		goto st0
	tr10:
//line NONE:1
		switch act {
		case 2:
			{
				p = (te) - 1
				isNumber = te == len(data)
			}
		case 3:
			{
				p = (te) - 1
				isNumber = false
			}
		}

		goto st0
	st0:
//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
//line NONE:1
		ts = p

//line isnumber.go:111
		switch data[p] {
		case 43:
			goto st2
		case 45:
			goto st2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		goto st1
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st1
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 46 {
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st1
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto st1
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 69 {
			goto st6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto st1
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 43:
			goto tr9
		case 45:
			goto tr9
		}
		goto st1
	tr9:
//line NONE:1
		te = p + 1

//line isnumber.rl:25
		act = 3
		goto st7
	tr11:
//line NONE:1
		te = p + 1

//line isnumber.rl:24
		act = 2
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line isnumber.go:201
		if 48 <= data[p] && data[p] <= 57 {
			goto tr11
		}
		goto st1
	st_out:
	_test_eof0:
		cs = 0
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1:
				goto tr3
			case 2:
				goto tr3
			case 3:
				goto tr4
			case 4:
				goto tr3
			case 5:
				goto tr7
			case 6:
				goto tr3
			case 7:
				goto tr10
			}
		}

	}

//line isnumber.rl:40
	if cs == format_error {
		return false
	}
	return
}
