//line lexer.rl:1

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format

import (
	"io"
	"log"
)

//line lexer.go:20
const format_start int = 15
const format_first_final int = 15
const format_error int = -1

const format_en_main int = 15

//line lexer.rl:95
func (l *Lexer) Lex(r io.Reader) {
	cs, p, pe := 0, 0, 0
	eof := -1
	ts, te, act := 0, 0, 0
	_ = te
	_ = act
	curline := 1
	_ = curline
	data := make([]byte, 4096)

	done := false
	for !done {
		// p - index of next character to process
		// pe - index of the end of the data
		// eof - index of the end of the file
		// ts - index of the start of the current token
		// te - index of the end of the current token

		// still have a partial token
		rem := 0
		if ts > 0 {
			rem = p - ts
		}
		p = 0
		n, err := r.Read(data[rem:])
		if n == 0 || err != nil {
			done = true
		}
		pe = n + rem
		if pe < len(data) {
			eof = pe
		}

//line lexer.go:64
		{
			cs = format_start
			ts = 0
			te = 0
			act = 0
		}

//line lexer.go:72
		{
			if p == pe {
				goto _test_eof
			}
			switch cs {
			case 15:
				goto st_case_15
			case 16:
				goto st_case_16
			case 0:
				goto st_case_0
			case 17:
				goto st_case_17
			case 18:
				goto st_case_18
			case 1:
				goto st_case_1
			case 19:
				goto st_case_19
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
			case 20:
				goto st_case_20
			case 7:
				goto st_case_7
			case 21:
				goto st_case_21
			case 22:
				goto st_case_22
			case 23:
				goto st_case_23
			case 24:
				goto st_case_24
			case 25:
				goto st_case_25
			case 8:
				goto st_case_8
			case 9:
				goto st_case_9
			case 10:
				goto st_case_10
			case 11:
				goto st_case_11
			case 12:
				goto st_case_12
			case 26:
				goto st_case_26
			case 13:
				goto st_case_13
			case 27:
				goto st_case_27
			case 14:
				goto st_case_14
			case 28:
				goto st_case_28
			case 29:
				goto st_case_29
			case 30:
				goto st_case_30
			}
			goto st_out
		tr0:
//line NONE:1
			switch act {
			case 11:
				{
					p = (te) - 1
					l.fmt.AddPlaceholder(FmtTypeDate, data[ts:te])
				}
			case 12:
				{
					p = (te) - 1
					l.fmt.AddPlaceholder(FmtTypeTime, data[ts:te])
				}
			case 13:
				{
					p = (te) - 1
					l.fmt.AddPlaceholder(FmtTypeTime, data[ts:te])
				}
			case 15:
				{
					p = (te) - 1
				}
			case 17:
				{
					p = (te) - 1
					l.fmt.AddPlaceholder(FmtTypeLiteral, data[ts:te])
				}
			case 18:
				{
					p = (te) - 1
					l.fmt.AddPlaceholder(FmtTypeLiteral, data[ts+1:te-1])
				}
			}

			goto st15
		tr3:
//line lexer.rl:83
			p = (te) - 1
			{
				l.fmt.AddPlaceholder(FmtTypeDate, data[ts:te])
			}
			goto st15
		tr10:
//line lexer.rl:84
			p = (te) - 1
			{
				l.fmt.AddPlaceholder(FmtTypeTime, data[ts:te])
			}
			goto st15
		tr12:
//line lexer.rl:90
			p = (te) - 1
			{
				l.fmt.AddPlaceholder(FmtTypeLiteral, data[ts:te])
			}
			goto st15
		tr17:
//line lexer.rl:80
			te = p + 1
			{
				l.fmt.isGeneral = true
			}
			goto st15
		tr21:
//line lexer.rl:90
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeLiteral, data[ts:te])
			}
			goto st15
		tr23:
//line lexer.rl:72
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeDigitOpt, nil)
			}
			goto st15
		tr24:
//line lexer.rl:77
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeDollar, nil)
			}
			goto st15
		tr25:
//line lexer.rl:76
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypePercent, nil)
			}
			goto st15
		tr26:
//line lexer.rl:75
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeComma, nil)
			}
			goto st15
		tr27:
//line lexer.rl:74
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeDecimal, nil)
			}
			goto st15
		tr28:
//line lexer.rl:71
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeDigit, nil)
			}
			goto st15
		tr29:
//line lexer.rl:79
			te = p + 1
			{
				l.nextFmt()
			}
			goto st15
		tr30:
//line lexer.rl:73
			te = p + 1
			{
			}
			goto st15
		tr36:
//line lexer.rl:78
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeUnderscore, nil)
			}
			goto st15
		tr40:
//line lexer.rl:90
			te = p
			p--
			{
				l.fmt.AddPlaceholder(FmtTypeLiteral, data[ts:te])
			}
			goto st15
		tr41:
//line lexer.rl:91
			te = p
			p--
			{
				l.fmt.AddPlaceholder(FmtTypeLiteral, data[ts+1:te-1])
			}
			goto st15
		tr42:
//line lexer.rl:83
			te = p
			p--
			{
				l.fmt.AddPlaceholder(FmtTypeDate, data[ts:te])
			}
			goto st15
		tr44:
//line lexer.rl:84
			te = p
			p--
			{
				l.fmt.AddPlaceholder(FmtTypeTime, data[ts:te])
			}
			goto st15
		tr48:
//line lexer.rl:86
			te = p + 1
			{
				l.fmt.IsExponential = true
			}
			goto st15
		tr51:
//line lexer.rl:89
			te = p + 1
			{
				l.fmt.AddPlaceholder(FmtTypeLiteral, data[ts+1:te])
			}
			goto st15
		st15:
//line NONE:1
			ts = 0

			if p++; p == pe {
				goto _test_eof15
			}
		st_case_15:
//line NONE:1
			ts = p

//line lexer.go:281
			switch data[p] {
			case 34:
				goto tr22
			case 35:
				goto tr23
			case 36:
				goto tr24
			case 37:
				goto tr25
			case 44:
				goto tr26
			case 46:
				goto tr27
			case 47:
				goto tr4
			case 48:
				goto tr28
			case 58:
				goto tr7
			case 59:
				goto tr29
			case 63:
				goto tr30
			case 65:
				goto tr31
			case 69:
				goto st24
			case 71:
				goto tr33
			case 91:
				goto tr34
			case 92:
				goto st28
			case 95:
				goto tr36
			case 100:
				goto tr4
			case 104:
				goto tr7
			case 109:
				goto tr37
			case 115:
				goto tr38
			case 121:
				goto st30
			}
			goto tr21
		tr22:
//line NONE:1
			te = p + 1

//line lexer.rl:90
			act = 17
			goto st16
		st16:
			if p++; p == pe {
				goto _test_eof16
			}
		st_case_16:
//line lexer.go:341
			if data[p] == 34 {
				goto tr2
			}
			goto st0
		st0:
			if p++; p == pe {
				goto _test_eof0
			}
		st_case_0:
			if data[p] == 34 {
				goto tr2
			}
			goto st0
		tr2:
//line NONE:1
			te = p + 1

//line lexer.rl:91
			act = 18
			goto st17
		st17:
			if p++; p == pe {
				goto _test_eof17
			}
		st_case_17:
//line lexer.go:367
			if data[p] == 34 {
				goto st0
			}
			goto tr41
		tr4:
//line NONE:1
			te = p + 1

			goto st18
		st18:
			if p++; p == pe {
				goto _test_eof18
			}
		st_case_18:
//line lexer.go:382
			switch data[p] {
			case 47:
				goto tr4
			case 100:
				goto tr4
			case 109:
				goto tr4
			case 121:
				goto st1
			}
			goto tr42
		st1:
			if p++; p == pe {
				goto _test_eof1
			}
		st_case_1:
			if data[p] == 121 {
				goto tr4
			}
			goto tr3
		tr7:
//line NONE:1
			te = p + 1

//line lexer.rl:84
			act = 12
			goto st19
		st19:
			if p++; p == pe {
				goto _test_eof19
			}
		st_case_19:
//line lexer.go:415
			switch data[p] {
			case 58:
				goto tr7
			case 65:
				goto st2
			case 104:
				goto tr7
			case 109:
				goto tr7
			case 115:
				goto tr38
			}
			goto tr44
		st2:
			if p++; p == pe {
				goto _test_eof2
			}
		st_case_2:
			switch data[p] {
			case 47:
				goto st3
			case 77:
				goto st4
			}
			goto tr0
		st3:
			if p++; p == pe {
				goto _test_eof3
			}
		st_case_3:
			if data[p] == 80 {
				goto tr7
			}
			goto tr0
		st4:
			if p++; p == pe {
				goto _test_eof4
			}
		st_case_4:
			if data[p] == 47 {
				goto st5
			}
			goto tr0
		st5:
			if p++; p == pe {
				goto _test_eof5
			}
		st_case_5:
			if data[p] == 80 {
				goto st6
			}
			goto tr0
		st6:
			if p++; p == pe {
				goto _test_eof6
			}
		st_case_6:
			if data[p] == 77 {
				goto tr7
			}
			goto tr0
		tr38:
//line NONE:1
			te = p + 1

//line lexer.rl:84
			act = 12
			goto st20
		st20:
			if p++; p == pe {
				goto _test_eof20
			}
		st_case_20:
//line lexer.go:489
			switch data[p] {
			case 46:
				goto st7
			case 58:
				goto tr7
			case 65:
				goto st2
			case 104:
				goto tr7
			case 109:
				goto tr7
			case 115:
				goto tr38
			}
			goto tr44
		st7:
			if p++; p == pe {
				goto _test_eof7
			}
		st_case_7:
			if data[p] == 48 {
				goto tr11
			}
			goto tr10
		tr11:
//line NONE:1
			te = p + 1

//line lexer.rl:84
			act = 12
			goto st21
		st21:
			if p++; p == pe {
				goto _test_eof21
			}
		st_case_21:
//line lexer.go:526
			switch data[p] {
			case 48:
				goto tr47
			case 58:
				goto tr7
			case 65:
				goto st2
			case 104:
				goto tr7
			case 109:
				goto tr7
			case 115:
				goto tr38
			}
			goto tr44
		tr47:
//line NONE:1
			te = p + 1

//line lexer.rl:84
			act = 12
			goto st22
		st22:
			if p++; p == pe {
				goto _test_eof22
			}
		st_case_22:
//line lexer.go:554
			switch data[p] {
			case 48:
				goto tr7
			case 58:
				goto tr7
			case 65:
				goto st2
			case 104:
				goto tr7
			case 109:
				goto tr7
			case 115:
				goto tr38
			}
			goto tr44
		tr31:
//line NONE:1
			te = p + 1

//line lexer.rl:90
			act = 17
			goto st23
		st23:
			if p++; p == pe {
				goto _test_eof23
			}
		st_case_23:
//line lexer.go:582
			switch data[p] {
			case 47:
				goto st3
			case 77:
				goto st4
			}
			goto tr40
		st24:
			if p++; p == pe {
				goto _test_eof24
			}
		st_case_24:
			switch data[p] {
			case 43:
				goto tr48
			case 45:
				goto tr48
			}
			goto tr40
		tr33:
//line NONE:1
			te = p + 1

			goto st25
		st25:
			if p++; p == pe {
				goto _test_eof25
			}
		st_case_25:
//line lexer.go:612
			if data[p] == 101 {
				goto st8
			}
			goto tr40
		st8:
			if p++; p == pe {
				goto _test_eof8
			}
		st_case_8:
			if data[p] == 110 {
				goto st9
			}
			goto tr12
		st9:
			if p++; p == pe {
				goto _test_eof9
			}
		st_case_9:
			if data[p] == 101 {
				goto st10
			}
			goto tr12
		st10:
			if p++; p == pe {
				goto _test_eof10
			}
		st_case_10:
			if data[p] == 114 {
				goto st11
			}
			goto tr12
		st11:
			if p++; p == pe {
				goto _test_eof11
			}
		st_case_11:
			if data[p] == 97 {
				goto st12
			}
			goto tr12
		st12:
			if p++; p == pe {
				goto _test_eof12
			}
		st_case_12:
			if data[p] == 108 {
				goto tr17
			}
			goto tr12
		tr34:
//line NONE:1
			te = p + 1

//line lexer.rl:90
			act = 17
			goto st26
		st26:
			if p++; p == pe {
				goto _test_eof26
			}
		st_case_26:
//line lexer.go:674
			switch data[p] {
			case 104:
				goto st14
			case 109:
				goto st14
			case 115:
				goto st14
			}
			goto st13
		st13:
			if p++; p == pe {
				goto _test_eof13
			}
		st_case_13:
			if data[p] == 93 {
				goto tr19
			}
			goto st13
		tr19:
//line NONE:1
			te = p + 1

//line lexer.rl:87
			act = 15
			goto st27
		tr20:
//line NONE:1
			te = p + 1

//line lexer.rl:85
			act = 13
			goto st27
		st27:
			if p++; p == pe {
				goto _test_eof27
			}
		st_case_27:
//line lexer.go:712
			if data[p] == 93 {
				goto tr19
			}
			goto st13
		st14:
			if p++; p == pe {
				goto _test_eof14
			}
		st_case_14:
			if data[p] == 93 {
				goto tr20
			}
			goto st13
		st28:
			if p++; p == pe {
				goto _test_eof28
			}
		st_case_28:
			goto tr51
		tr37:
//line NONE:1
			te = p + 1

//line lexer.rl:83
			act = 11
			goto st29
		st29:
			if p++; p == pe {
				goto _test_eof29
			}
		st_case_29:
//line lexer.go:744
			switch data[p] {
			case 47:
				goto tr4
			case 58:
				goto tr7
			case 65:
				goto st2
			case 100:
				goto tr4
			case 104:
				goto tr7
			case 109:
				goto tr37
			case 115:
				goto tr38
			case 121:
				goto st1
			}
			goto tr42
		st30:
			if p++; p == pe {
				goto _test_eof30
			}
		st_case_30:
			if data[p] == 121 {
				goto tr4
			}
			goto tr40
		st_out:
		_test_eof15:
			cs = 15
			goto _test_eof
		_test_eof16:
			cs = 16
			goto _test_eof
		_test_eof0:
			cs = 0
			goto _test_eof
		_test_eof17:
			cs = 17
			goto _test_eof
		_test_eof18:
			cs = 18
			goto _test_eof
		_test_eof1:
			cs = 1
			goto _test_eof
		_test_eof19:
			cs = 19
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
		_test_eof20:
			cs = 20
			goto _test_eof
		_test_eof7:
			cs = 7
			goto _test_eof
		_test_eof21:
			cs = 21
			goto _test_eof
		_test_eof22:
			cs = 22
			goto _test_eof
		_test_eof23:
			cs = 23
			goto _test_eof
		_test_eof24:
			cs = 24
			goto _test_eof
		_test_eof25:
			cs = 25
			goto _test_eof
		_test_eof8:
			cs = 8
			goto _test_eof
		_test_eof9:
			cs = 9
			goto _test_eof
		_test_eof10:
			cs = 10
			goto _test_eof
		_test_eof11:
			cs = 11
			goto _test_eof
		_test_eof12:
			cs = 12
			goto _test_eof
		_test_eof26:
			cs = 26
			goto _test_eof
		_test_eof13:
			cs = 13
			goto _test_eof
		_test_eof27:
			cs = 27
			goto _test_eof
		_test_eof14:
			cs = 14
			goto _test_eof
		_test_eof28:
			cs = 28
			goto _test_eof
		_test_eof29:
			cs = 29
			goto _test_eof
		_test_eof30:
			cs = 30
			goto _test_eof

		_test_eof:
			{
			}
			if p == eof {
				switch cs {
				case 16:
					goto tr40
				case 0:
					goto tr0
				case 17:
					goto tr41
				case 18:
					goto tr42
				case 1:
					goto tr3
				case 19:
					goto tr44
				case 2:
					goto tr0
				case 3:
					goto tr0
				case 4:
					goto tr0
				case 5:
					goto tr0
				case 6:
					goto tr0
				case 20:
					goto tr44
				case 7:
					goto tr10
				case 21:
					goto tr44
				case 22:
					goto tr44
				case 23:
					goto tr40
				case 24:
					goto tr40
				case 25:
					goto tr40
				case 8:
					goto tr12
				case 9:
					goto tr12
				case 10:
					goto tr12
				case 11:
					goto tr12
				case 12:
					goto tr12
				case 26:
					goto tr40
				case 13:
					goto tr0
				case 27:
					goto tr0
				case 14:
					goto tr12
				case 28:
					goto tr40
				case 29:
					goto tr42
				case 30:
					goto tr40
				}
			}

		}

//line lexer.rl:132
		if ts > 0 {
			// currently parsing a token, so shift it to the
			// beginning of the buffer
			copy(data[0:], data[ts:])
		}
	}

	_ = eof
	if cs == format_error {
		//l.emit(tokenLexError,nil)
		log.Panic("ERROR LEXING")
	}
}
