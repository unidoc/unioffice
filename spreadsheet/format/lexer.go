//line lexer.rl:1

// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format

import (
	"io"

	"github.com/unidoc/unioffice"
)

//line lexer.go:20
const format_start int = 34
const format_first_final int = 34
const format_error int = -1

const format_en_main int = 34

//line lexer.rl:97
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
			case 34:
				goto st_case_34
			case 35:
				goto st_case_35
			case 0:
				goto st_case_0
			case 36:
				goto st_case_36
			case 37:
				goto st_case_37
			case 1:
				goto st_case_1
			case 2:
				goto st_case_2
			case 38:
				goto st_case_38
			case 3:
				goto st_case_3
			case 4:
				goto st_case_4
			case 39:
				goto st_case_39
			case 5:
				goto st_case_5
			case 6:
				goto st_case_6
			case 7:
				goto st_case_7
			case 8:
				goto st_case_8
			case 40:
				goto st_case_40
			case 9:
				goto st_case_9
			case 41:
				goto st_case_41
			case 10:
				goto st_case_10
			case 42:
				goto st_case_42
			case 11:
				goto st_case_11
			case 43:
				goto st_case_43
			case 44:
				goto st_case_44
			case 45:
				goto st_case_45
			case 12:
				goto st_case_12
			case 46:
				goto st_case_46
			case 13:
				goto st_case_13
			case 14:
				goto st_case_14
			case 15:
				goto st_case_15
			case 16:
				goto st_case_16
			case 47:
				goto st_case_47
			case 17:
				goto st_case_17
			case 48:
				goto st_case_48
			case 18:
				goto st_case_18
			case 19:
				goto st_case_19
			case 20:
				goto st_case_20
			case 49:
				goto st_case_49
			case 50:
				goto st_case_50
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
			case 51:
				goto st_case_51
			case 26:
				goto st_case_26
			case 52:
				goto st_case_52
			case 53:
				goto st_case_53
			case 54:
				goto st_case_54
			case 55:
				goto st_case_55
			case 56:
				goto st_case_56
			case 57:
				goto st_case_57
			case 27:
				goto st_case_27
			case 28:
				goto st_case_28
			case 29:
				goto st_case_29
			case 30:
				goto st_case_30
			case 31:
				goto st_case_31
			case 58:
				goto st_case_58
			case 32:
				goto st_case_32
			case 59:
				goto st_case_59
			case 33:
				goto st_case_33
			case 60:
				goto st_case_60
			case 61:
				goto st_case_61
			case 62:
				goto st_case_62
			}
			goto st_out
		tr0:
//line NONE:1
			switch act {
			case 2:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeDigit, nil)
				}
			case 3:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeDigitOpt, nil)
				}
			case 5:
				{
					p = (te) - 1
				}
			case 8:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypePercent, nil)
				}
			case 13:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeFraction, data[ts:te])
				}
			case 14:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeDate, data[ts:te])
				}
			case 15:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeTime, data[ts:te])
				}
			case 16:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeTime, data[ts:te])
				}
			case 18:
				{
					p = (te) - 1
				}
			case 20:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeLiteral, data[ts:te])
				}
			case 21:
				{
					p = (te) - 1
					l.fmt.AddToken(FmtTypeLiteral, data[ts+1:te-1])
				}
			}

			goto st34
		tr9:
//line lexer.rl:83
			p = (te) - 1
			{
				l.fmt.AddToken(FmtTypeFraction, data[ts:te])
			}
			goto st34
		tr19:
//line lexer.rl:73
			p = (te) - 1
			{
				l.fmt.AddToken(FmtTypeDigitOpt, nil)
			}
			goto st34
		tr20:
//line lexer.rl:71
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeDigitOptThousands, nil)
			}
			goto st34
		tr21:
//line lexer.rl:78
			p = (te) - 1
			{
				l.fmt.AddToken(FmtTypePercent, nil)
			}
			goto st34
		tr26:
//line lexer.rl:85
			p = (te) - 1
			{
				l.fmt.AddToken(FmtTypeDate, data[ts:te])
			}
			goto st34
		tr28:
//line lexer.rl:72
			p = (te) - 1
			{
				l.fmt.AddToken(FmtTypeDigit, nil)
			}
			goto st34
		tr37:
//line lexer.rl:86
			p = (te) - 1
			{
				l.fmt.AddToken(FmtTypeTime, data[ts:te])
			}
			goto st34
		tr39:
//line lexer.rl:92
			p = (te) - 1
			{
				l.fmt.AddToken(FmtTypeLiteral, data[ts:te])
			}
			goto st34
		tr44:
//line lexer.rl:82
			te = p + 1
			{
				l.fmt.isGeneral = true
			}
			goto st34
		tr48:
//line lexer.rl:92
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeLiteral, data[ts:te])
			}
			goto st34
		tr51:
//line lexer.rl:79
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeDollar, nil)
			}
			goto st34
		tr53:
//line lexer.rl:77
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeComma, nil)
			}
			goto st34
		tr54:
//line lexer.rl:76
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeDecimal, nil)
			}
			goto st34
		tr57:
//line lexer.rl:81
			te = p + 1
			{
				l.nextFmt()
			}
			goto st34
		tr59:
//line lexer.rl:74
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeText, nil)
			}
			goto st34
		tr65:
//line lexer.rl:80
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeUnderscore, nil)
			}
			goto st34
		tr69:
//line lexer.rl:92
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypeLiteral, data[ts:te])
			}
			goto st34
		tr70:
//line lexer.rl:93
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypeLiteral, data[ts+1:te-1])
			}
			goto st34
		tr71:
//line lexer.rl:73
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypeDigitOpt, nil)
			}
			goto st34
		tr73:
//line lexer.rl:83
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypeFraction, data[ts:te])
			}
			goto st34
		tr78:
//line lexer.rl:78
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypePercent, nil)
			}
			goto st34
		tr79:
//line lexer.rl:85
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypeDate, data[ts:te])
			}
			goto st34
		tr81:
//line lexer.rl:72
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypeDigit, nil)
			}
			goto st34
		tr82:
//line lexer.rl:86
			te = p
			p--
			{
				l.fmt.AddToken(FmtTypeTime, data[ts:te])
			}
			goto st34
		tr86:
//line lexer.rl:75
			te = p
			p--
			{
			}
			goto st34
		tr87:
//line lexer.rl:88
			te = p + 1
			{
				l.fmt.IsExponential = true
			}
			goto st34
		tr90:
//line lexer.rl:91
			te = p + 1
			{
				l.fmt.AddToken(FmtTypeLiteral, data[ts+1:te])
			}
			goto st34
		st34:
//line NONE:1
			ts = 0

			if p++; p == pe {
				goto _test_eof34
			}
		st_case_34:
//line NONE:1
			ts = p

//line lexer.go:400
			switch data[p] {
			case 34:
				goto tr49
			case 35:
				goto tr50
			case 36:
				goto tr51
			case 37:
				goto tr52
			case 44:
				goto tr53
			case 46:
				goto tr54
			case 47:
				goto tr27
			case 48:
				goto tr55
			case 58:
				goto tr34
			case 59:
				goto tr57
			case 63:
				goto tr58
			case 64:
				goto tr59
			case 65:
				goto tr60
			case 69:
				goto st56
			case 71:
				goto tr62
			case 91:
				goto tr63
			case 92:
				goto st60
			case 95:
				goto tr65
			case 100:
				goto tr27
			case 104:
				goto tr34
			case 109:
				goto tr66
			case 115:
				goto tr67
			case 121:
				goto st62
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto tr56
			}
			goto tr48
		tr49:
//line NONE:1
			te = p + 1

//line lexer.rl:92
			act = 20
			goto st35
		st35:
			if p++; p == pe {
				goto _test_eof35
			}
		st_case_35:
//line lexer.go:465
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

//line lexer.rl:93
			act = 21
			goto st36
		st36:
			if p++; p == pe {
				goto _test_eof36
			}
		st_case_36:
//line lexer.go:491
			if data[p] == 34 {
				goto st0
			}
			goto tr70
		tr50:
//line NONE:1
			te = p + 1

//line lexer.rl:73
			act = 3
			goto st37
		st37:
			if p++; p == pe {
				goto _test_eof37
			}
		st_case_37:
//line lexer.go:508
			switch data[p] {
			case 35:
				goto st1
			case 37:
				goto st1
			case 44:
				goto st12
			case 47:
				goto st2
			case 48:
				goto st1
			case 63:
				goto st1
			}
			goto tr71
		st1:
			if p++; p == pe {
				goto _test_eof1
			}
		st_case_1:
			switch data[p] {
			case 35:
				goto st1
			case 37:
				goto st1
			case 47:
				goto st2
			case 48:
				goto st1
			case 63:
				goto st1
			}
			goto tr0
		st2:
			if p++; p == pe {
				goto _test_eof2
			}
		st_case_2:
			switch data[p] {
			case 35:
				goto tr5
			case 37:
				goto st8
			case 48:
				goto tr7
			case 63:
				goto tr5
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto tr8
			}
			goto tr0
		tr5:
//line NONE:1
			te = p + 1

			goto st38
		st38:
			if p++; p == pe {
				goto _test_eof38
			}
		st_case_38:
//line lexer.go:571
			switch data[p] {
			case 35:
				goto tr5
			case 37:
				goto tr5
			case 44:
				goto tr5
			case 46:
				goto tr5
			case 48:
				goto tr5
			case 63:
				goto tr5
			case 65:
				goto st3
			}
			goto tr73
		st3:
			if p++; p == pe {
				goto _test_eof3
			}
		st_case_3:
			switch data[p] {
			case 47:
				goto st4
			case 77:
				goto st5
			}
			goto tr9
		st4:
			if p++; p == pe {
				goto _test_eof4
			}
		st_case_4:
			if data[p] == 80 {
				goto tr12
			}
			goto tr9
		tr12:
//line NONE:1
			te = p + 1

			goto st39
		st39:
			if p++; p == pe {
				goto _test_eof39
			}
		st_case_39:
//line lexer.go:620
			if data[p] == 65 {
				goto st3
			}
			goto tr73
		st5:
			if p++; p == pe {
				goto _test_eof5
			}
		st_case_5:
			if data[p] == 47 {
				goto st6
			}
			goto tr9
		st6:
			if p++; p == pe {
				goto _test_eof6
			}
		st_case_6:
			if data[p] == 80 {
				goto st7
			}
			goto tr9
		st7:
			if p++; p == pe {
				goto _test_eof7
			}
		st_case_7:
			if data[p] == 77 {
				goto tr12
			}
			goto tr9
		st8:
			if p++; p == pe {
				goto _test_eof8
			}
		st_case_8:
			switch data[p] {
			case 35:
				goto tr15
			case 37:
				goto st10
			case 63:
				goto tr15
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr17
			}
			goto tr0
		tr15:
//line NONE:1
			te = p + 1

			goto st40
		st40:
			if p++; p == pe {
				goto _test_eof40
			}
		st_case_40:
//line lexer.go:679
			switch data[p] {
			case 35:
				goto tr5
			case 37:
				goto st9
			case 44:
				goto tr5
			case 46:
				goto tr5
			case 48:
				goto tr5
			case 63:
				goto tr5
			case 65:
				goto st3
			}
			goto tr73
		st9:
			if p++; p == pe {
				goto _test_eof9
			}
		st_case_9:
			switch data[p] {
			case 35:
				goto tr18
			case 44:
				goto tr18
			case 46:
				goto tr18
			case 48:
				goto tr18
			case 63:
				goto tr18
			}
			goto tr9
		tr18:
//line NONE:1
			te = p + 1

			goto st41
		st41:
			if p++; p == pe {
				goto _test_eof41
			}
		st_case_41:
//line lexer.go:725
			switch data[p] {
			case 35:
				goto tr18
			case 44:
				goto tr18
			case 46:
				goto tr18
			case 48:
				goto tr18
			case 63:
				goto tr18
			case 65:
				goto st3
			}
			goto tr73
		st10:
			if p++; p == pe {
				goto _test_eof10
			}
		st_case_10:
			if data[p] == 37 {
				goto st10
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr17
			}
			goto tr0
		tr17:
//line NONE:1
			te = p + 1

//line lexer.rl:83
			act = 13
			goto st42
		st42:
			if p++; p == pe {
				goto _test_eof42
			}
		st_case_42:
//line lexer.go:765
			switch data[p] {
			case 35:
				goto tr5
			case 37:
				goto st11
			case 44:
				goto tr5
			case 46:
				goto tr5
			case 48:
				goto tr77
			case 63:
				goto tr5
			case 65:
				goto st3
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto tr17
			}
			goto tr73
		st11:
			if p++; p == pe {
				goto _test_eof11
			}
		st_case_11:
			switch data[p] {
			case 35:
				goto tr18
			case 37:
				goto st10
			case 44:
				goto tr18
			case 46:
				goto tr18
			case 63:
				goto tr18
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr17
			}
			goto tr9
		tr77:
//line NONE:1
			te = p + 1

			goto st43
		st43:
			if p++; p == pe {
				goto _test_eof43
			}
		st_case_43:
//line lexer.go:817
			switch data[p] {
			case 35:
				goto tr5
			case 37:
				goto tr77
			case 44:
				goto tr5
			case 46:
				goto tr5
			case 48:
				goto tr77
			case 63:
				goto tr5
			case 65:
				goto st3
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto tr17
			}
			goto tr73
		tr7:
//line NONE:1
			te = p + 1

			goto st44
		st44:
			if p++; p == pe {
				goto _test_eof44
			}
		st_case_44:
//line lexer.go:848
			switch data[p] {
			case 35:
				goto tr5
			case 37:
				goto tr77
			case 44:
				goto tr5
			case 46:
				goto tr5
			case 48:
				goto tr7
			case 63:
				goto tr5
			case 65:
				goto st3
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto tr8
			}
			goto tr73
		tr8:
//line NONE:1
			te = p + 1

			goto st45
		st45:
			if p++; p == pe {
				goto _test_eof45
			}
		st_case_45:
//line lexer.go:879
			switch data[p] {
			case 35:
				goto tr5
			case 37:
				goto tr17
			case 44:
				goto tr5
			case 46:
				goto tr5
			case 48:
				goto tr7
			case 63:
				goto tr5
			case 65:
				goto st3
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto tr8
			}
			goto tr73
		st12:
			if p++; p == pe {
				goto _test_eof12
			}
		st_case_12:
			if data[p] == 35 {
				goto tr20
			}
			goto tr19
		tr52:
//line NONE:1
			te = p + 1

//line lexer.rl:78
			act = 8
			goto st46
		st46:
			if p++; p == pe {
				goto _test_eof46
			}
		st_case_46:
//line lexer.go:921
			switch data[p] {
			case 35:
				goto st13
			case 37:
				goto st14
			case 48:
				goto st16
			case 63:
				goto st13
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto st15
			}
			goto tr78
		st13:
			if p++; p == pe {
				goto _test_eof13
			}
		st_case_13:
			switch data[p] {
			case 35:
				goto st13
			case 47:
				goto st2
			case 48:
				goto st13
			case 63:
				goto st13
			}
			goto tr21
		st14:
			if p++; p == pe {
				goto _test_eof14
			}
		st_case_14:
			if data[p] == 37 {
				goto st14
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
			goto tr0
		st15:
			if p++; p == pe {
				goto _test_eof15
			}
		st_case_15:
			switch data[p] {
			case 37:
				goto st14
			case 47:
				goto st2
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
			goto tr0
		st16:
			if p++; p == pe {
				goto _test_eof16
			}
		st_case_16:
			switch data[p] {
			case 35:
				goto st13
			case 37:
				goto st14
			case 47:
				goto st2
			case 48:
				goto st16
			case 63:
				goto st13
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto st15
			}
			goto tr21
		tr27:
//line NONE:1
			te = p + 1

			goto st47
		st47:
			if p++; p == pe {
				goto _test_eof47
			}
		st_case_47:
//line lexer.go:1010
			switch data[p] {
			case 47:
				goto tr27
			case 100:
				goto tr27
			case 109:
				goto tr27
			case 121:
				goto st17
			}
			goto tr79
		st17:
			if p++; p == pe {
				goto _test_eof17
			}
		st_case_17:
			if data[p] == 121 {
				goto tr27
			}
			goto tr26
		tr55:
//line NONE:1
			te = p + 1

//line lexer.rl:72
			act = 2
			goto st48
		st48:
			if p++; p == pe {
				goto _test_eof48
			}
		st_case_48:
//line lexer.go:1043
			switch data[p] {
			case 35:
				goto st1
			case 37:
				goto st18
			case 47:
				goto st2
			case 48:
				goto st19
			case 63:
				goto st1
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto st20
			}
			goto tr81
		st18:
			if p++; p == pe {
				goto _test_eof18
			}
		st_case_18:
			switch data[p] {
			case 35:
				goto st1
			case 37:
				goto st18
			case 47:
				goto st2
			case 48:
				goto st18
			case 63:
				goto st1
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto st15
			}
			goto tr28
		st19:
			if p++; p == pe {
				goto _test_eof19
			}
		st_case_19:
			switch data[p] {
			case 35:
				goto st1
			case 37:
				goto st18
			case 47:
				goto st2
			case 48:
				goto st19
			case 63:
				goto st1
			}
			if 49 <= data[p] && data[p] <= 57 {
				goto st20
			}
			goto tr28
		st20:
			if p++; p == pe {
				goto _test_eof20
			}
		st_case_20:
			switch data[p] {
			case 37:
				goto st15
			case 47:
				goto st2
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto st20
			}
			goto tr0
		tr56:
//line NONE:1
			te = p + 1

//line lexer.rl:92
			act = 20
			goto st49
		st49:
			if p++; p == pe {
				goto _test_eof49
			}
		st_case_49:
//line lexer.go:1129
			switch data[p] {
			case 37:
				goto st15
			case 47:
				goto st2
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto st20
			}
			goto tr69
		tr34:
//line NONE:1
			te = p + 1

//line lexer.rl:86
			act = 15
			goto st50
		st50:
			if p++; p == pe {
				goto _test_eof50
			}
		st_case_50:
//line lexer.go:1152
			switch data[p] {
			case 58:
				goto tr34
			case 65:
				goto st21
			case 104:
				goto tr34
			case 109:
				goto tr34
			case 115:
				goto tr67
			}
			goto tr82
		st21:
			if p++; p == pe {
				goto _test_eof21
			}
		st_case_21:
			switch data[p] {
			case 47:
				goto st22
			case 77:
				goto st23
			}
			goto tr0
		st22:
			if p++; p == pe {
				goto _test_eof22
			}
		st_case_22:
			if data[p] == 80 {
				goto tr34
			}
			goto tr0
		st23:
			if p++; p == pe {
				goto _test_eof23
			}
		st_case_23:
			if data[p] == 47 {
				goto st24
			}
			goto tr0
		st24:
			if p++; p == pe {
				goto _test_eof24
			}
		st_case_24:
			if data[p] == 80 {
				goto st25
			}
			goto tr0
		st25:
			if p++; p == pe {
				goto _test_eof25
			}
		st_case_25:
			if data[p] == 77 {
				goto tr34
			}
			goto tr0
		tr67:
//line NONE:1
			te = p + 1

//line lexer.rl:86
			act = 15
			goto st51
		st51:
			if p++; p == pe {
				goto _test_eof51
			}
		st_case_51:
//line lexer.go:1226
			switch data[p] {
			case 46:
				goto st26
			case 58:
				goto tr34
			case 65:
				goto st21
			case 104:
				goto tr34
			case 109:
				goto tr34
			case 115:
				goto tr67
			}
			goto tr82
		st26:
			if p++; p == pe {
				goto _test_eof26
			}
		st_case_26:
			if data[p] == 48 {
				goto tr38
			}
			goto tr37
		tr38:
//line NONE:1
			te = p + 1

//line lexer.rl:86
			act = 15
			goto st52
		st52:
			if p++; p == pe {
				goto _test_eof52
			}
		st_case_52:
//line lexer.go:1263
			switch data[p] {
			case 48:
				goto tr85
			case 58:
				goto tr34
			case 65:
				goto st21
			case 104:
				goto tr34
			case 109:
				goto tr34
			case 115:
				goto tr67
			}
			goto tr82
		tr85:
//line NONE:1
			te = p + 1

//line lexer.rl:86
			act = 15
			goto st53
		st53:
			if p++; p == pe {
				goto _test_eof53
			}
		st_case_53:
//line lexer.go:1291
			switch data[p] {
			case 48:
				goto tr34
			case 58:
				goto tr34
			case 65:
				goto st21
			case 104:
				goto tr34
			case 109:
				goto tr34
			case 115:
				goto tr67
			}
			goto tr82
		tr58:
//line NONE:1
			te = p + 1

//line lexer.rl:75
			act = 5
			goto st54
		st54:
			if p++; p == pe {
				goto _test_eof54
			}
		st_case_54:
//line lexer.go:1319
			switch data[p] {
			case 35:
				goto st1
			case 37:
				goto st1
			case 47:
				goto st2
			case 48:
				goto st1
			case 63:
				goto st1
			}
			goto tr86
		tr60:
//line NONE:1
			te = p + 1

//line lexer.rl:92
			act = 20
			goto st55
		st55:
			if p++; p == pe {
				goto _test_eof55
			}
		st_case_55:
//line lexer.go:1345
			switch data[p] {
			case 47:
				goto st22
			case 77:
				goto st23
			}
			goto tr69
		st56:
			if p++; p == pe {
				goto _test_eof56
			}
		st_case_56:
			switch data[p] {
			case 43:
				goto tr87
			case 45:
				goto tr87
			}
			goto tr69
		tr62:
//line NONE:1
			te = p + 1

			goto st57
		st57:
			if p++; p == pe {
				goto _test_eof57
			}
		st_case_57:
//line lexer.go:1375
			if data[p] == 101 {
				goto st27
			}
			goto tr69
		st27:
			if p++; p == pe {
				goto _test_eof27
			}
		st_case_27:
			if data[p] == 110 {
				goto st28
			}
			goto tr39
		st28:
			if p++; p == pe {
				goto _test_eof28
			}
		st_case_28:
			if data[p] == 101 {
				goto st29
			}
			goto tr39
		st29:
			if p++; p == pe {
				goto _test_eof29
			}
		st_case_29:
			if data[p] == 114 {
				goto st30
			}
			goto tr39
		st30:
			if p++; p == pe {
				goto _test_eof30
			}
		st_case_30:
			if data[p] == 97 {
				goto st31
			}
			goto tr39
		st31:
			if p++; p == pe {
				goto _test_eof31
			}
		st_case_31:
			if data[p] == 108 {
				goto tr44
			}
			goto tr39
		tr63:
//line NONE:1
			te = p + 1

//line lexer.rl:92
			act = 20
			goto st58
		st58:
			if p++; p == pe {
				goto _test_eof58
			}
		st_case_58:
//line lexer.go:1437
			switch data[p] {
			case 104:
				goto st33
			case 109:
				goto st33
			case 115:
				goto st33
			}
			goto st32
		st32:
			if p++; p == pe {
				goto _test_eof32
			}
		st_case_32:
			if data[p] == 93 {
				goto tr46
			}
			goto st32
		tr46:
//line NONE:1
			te = p + 1

//line lexer.rl:89
			act = 18
			goto st59
		tr47:
//line NONE:1
			te = p + 1

//line lexer.rl:87
			act = 16
			goto st59
		st59:
			if p++; p == pe {
				goto _test_eof59
			}
		st_case_59:
//line lexer.go:1475
			if data[p] == 93 {
				goto tr46
			}
			goto st32
		st33:
			if p++; p == pe {
				goto _test_eof33
			}
		st_case_33:
			if data[p] == 93 {
				goto tr47
			}
			goto st32
		st60:
			if p++; p == pe {
				goto _test_eof60
			}
		st_case_60:
			goto tr90
		tr66:
//line NONE:1
			te = p + 1

//line lexer.rl:85
			act = 14
			goto st61
		st61:
			if p++; p == pe {
				goto _test_eof61
			}
		st_case_61:
//line lexer.go:1507
			switch data[p] {
			case 47:
				goto tr27
			case 58:
				goto tr34
			case 65:
				goto st21
			case 100:
				goto tr27
			case 104:
				goto tr34
			case 109:
				goto tr66
			case 115:
				goto tr67
			case 121:
				goto st17
			}
			goto tr79
		st62:
			if p++; p == pe {
				goto _test_eof62
			}
		st_case_62:
			if data[p] == 121 {
				goto tr27
			}
			goto tr69
		st_out:
		_test_eof34:
			cs = 34
			goto _test_eof
		_test_eof35:
			cs = 35
			goto _test_eof
		_test_eof0:
			cs = 0
			goto _test_eof
		_test_eof36:
			cs = 36
			goto _test_eof
		_test_eof37:
			cs = 37
			goto _test_eof
		_test_eof1:
			cs = 1
			goto _test_eof
		_test_eof2:
			cs = 2
			goto _test_eof
		_test_eof38:
			cs = 38
			goto _test_eof
		_test_eof3:
			cs = 3
			goto _test_eof
		_test_eof4:
			cs = 4
			goto _test_eof
		_test_eof39:
			cs = 39
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
		_test_eof8:
			cs = 8
			goto _test_eof
		_test_eof40:
			cs = 40
			goto _test_eof
		_test_eof9:
			cs = 9
			goto _test_eof
		_test_eof41:
			cs = 41
			goto _test_eof
		_test_eof10:
			cs = 10
			goto _test_eof
		_test_eof42:
			cs = 42
			goto _test_eof
		_test_eof11:
			cs = 11
			goto _test_eof
		_test_eof43:
			cs = 43
			goto _test_eof
		_test_eof44:
			cs = 44
			goto _test_eof
		_test_eof45:
			cs = 45
			goto _test_eof
		_test_eof12:
			cs = 12
			goto _test_eof
		_test_eof46:
			cs = 46
			goto _test_eof
		_test_eof13:
			cs = 13
			goto _test_eof
		_test_eof14:
			cs = 14
			goto _test_eof
		_test_eof15:
			cs = 15
			goto _test_eof
		_test_eof16:
			cs = 16
			goto _test_eof
		_test_eof47:
			cs = 47
			goto _test_eof
		_test_eof17:
			cs = 17
			goto _test_eof
		_test_eof48:
			cs = 48
			goto _test_eof
		_test_eof18:
			cs = 18
			goto _test_eof
		_test_eof19:
			cs = 19
			goto _test_eof
		_test_eof20:
			cs = 20
			goto _test_eof
		_test_eof49:
			cs = 49
			goto _test_eof
		_test_eof50:
			cs = 50
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
		_test_eof51:
			cs = 51
			goto _test_eof
		_test_eof26:
			cs = 26
			goto _test_eof
		_test_eof52:
			cs = 52
			goto _test_eof
		_test_eof53:
			cs = 53
			goto _test_eof
		_test_eof54:
			cs = 54
			goto _test_eof
		_test_eof55:
			cs = 55
			goto _test_eof
		_test_eof56:
			cs = 56
			goto _test_eof
		_test_eof57:
			cs = 57
			goto _test_eof
		_test_eof27:
			cs = 27
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
		_test_eof31:
			cs = 31
			goto _test_eof
		_test_eof58:
			cs = 58
			goto _test_eof
		_test_eof32:
			cs = 32
			goto _test_eof
		_test_eof59:
			cs = 59
			goto _test_eof
		_test_eof33:
			cs = 33
			goto _test_eof
		_test_eof60:
			cs = 60
			goto _test_eof
		_test_eof61:
			cs = 61
			goto _test_eof
		_test_eof62:
			cs = 62
			goto _test_eof

		_test_eof:
			{
			}
			if p == eof {
				switch cs {
				case 35:
					goto tr69
				case 0:
					goto tr0
				case 36:
					goto tr70
				case 37:
					goto tr71
				case 1:
					goto tr0
				case 2:
					goto tr0
				case 38:
					goto tr73
				case 3:
					goto tr9
				case 4:
					goto tr9
				case 39:
					goto tr73
				case 5:
					goto tr9
				case 6:
					goto tr9
				case 7:
					goto tr9
				case 8:
					goto tr0
				case 40:
					goto tr73
				case 9:
					goto tr9
				case 41:
					goto tr73
				case 10:
					goto tr0
				case 42:
					goto tr73
				case 11:
					goto tr9
				case 43:
					goto tr73
				case 44:
					goto tr73
				case 45:
					goto tr73
				case 12:
					goto tr19
				case 46:
					goto tr78
				case 13:
					goto tr21
				case 14:
					goto tr0
				case 15:
					goto tr0
				case 16:
					goto tr21
				case 47:
					goto tr79
				case 17:
					goto tr26
				case 48:
					goto tr81
				case 18:
					goto tr28
				case 19:
					goto tr28
				case 20:
					goto tr0
				case 49:
					goto tr69
				case 50:
					goto tr82
				case 21:
					goto tr0
				case 22:
					goto tr0
				case 23:
					goto tr0
				case 24:
					goto tr0
				case 25:
					goto tr0
				case 51:
					goto tr82
				case 26:
					goto tr37
				case 52:
					goto tr82
				case 53:
					goto tr82
				case 54:
					goto tr86
				case 55:
					goto tr69
				case 56:
					goto tr69
				case 57:
					goto tr69
				case 27:
					goto tr39
				case 28:
					goto tr39
				case 29:
					goto tr39
				case 30:
					goto tr39
				case 31:
					goto tr39
				case 58:
					goto tr69
				case 32:
					goto tr0
				case 59:
					goto tr0
				case 33:
					goto tr39
				case 60:
					goto tr69
				case 61:
					goto tr79
				case 62:
					goto tr69
				}
			}

		}

//line lexer.rl:134
		if ts > 0 {
			// currently parsing a token, so shift it to the
			// beginning of the buffer
			copy(data[0:], data[ts:])
		}
	}
	_ = eof
	if cs == format_error {
		unioffice.Log("format parse error")
	}
}
