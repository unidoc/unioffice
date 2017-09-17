// line 1 "lexer.rl"
// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"io"
)

// line 20 "lexer.go"
const formula_start int = 27
const formula_first_final int = 27
const formula_error int = 0

const formula_en_main int = 27

// line 107 "lexer.rl"

func (l *Lexer) lex(r io.Reader) {
	cs, p, pe := 0, 0, 0
	eof := -1
	ts, te, act := 0, 0, 0
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

		// line 63 "lexer.go"
		{
			cs = formula_start
			ts = 0
			te = 0
			act = 0
		}

		// line 71 "lexer.go"
		{
			if p == pe {
				goto _test_eof
			}
			switch cs {
			case 27:
				goto st_case_27
			case 1:
				goto st_case_1
			case 0:
				goto st_case_0
			case 2:
				goto st_case_2
			case 28:
				goto st_case_28
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
			case 13:
				goto st_case_13
			case 14:
				goto st_case_14
			case 15:
				goto st_case_15
			case 29:
				goto st_case_29
			case 16:
				goto st_case_16
			case 17:
				goto st_case_17
			case 30:
				goto st_case_30
			case 18:
				goto st_case_18
			case 19:
				goto st_case_19
			case 20:
				goto st_case_20
			case 31:
				goto st_case_31
			case 32:
				goto st_case_32
			case 21:
				goto st_case_21
			case 33:
				goto st_case_33
			case 34:
				goto st_case_34
			case 35:
				goto st_case_35
			case 22:
				goto st_case_22
			case 36:
				goto st_case_36
			case 37:
				goto st_case_37
			case 38:
				goto st_case_38
			case 39:
				goto st_case_39
			case 40:
				goto st_case_40
			case 23:
				goto st_case_23
			case 41:
				goto st_case_41
			case 42:
				goto st_case_42
			case 43:
				goto st_case_43
			case 24:
				goto st_case_24
			case 44:
				goto st_case_44
			case 45:
				goto st_case_45
			case 46:
				goto st_case_46
			case 25:
				goto st_case_25
			case 26:
				goto st_case_26
			case 47:
				goto st_case_47
			case 48:
				goto st_case_48
			case 49:
				goto st_case_49
			case 50:
				goto st_case_50
			case 51:
				goto st_case_51
			case 52:
				goto st_case_52
			case 53:
				goto st_case_53
			case 54:
				goto st_case_54
			case 55:
				goto st_case_55
			}
			goto st_out
		tr0:
			// line 1 "NONE"

			switch act {
			case 0:
				{
					{
						goto st0
					}
				}
			case 1:
				{
					p = (te) - 1
					l.emit(tokenBool, data[ts:te])
				}
			case 2:
				{
					p = (te) - 1
					l.emit(tokenNumber, data[ts:te])
				}
			case 3:
				{
					p = (te) - 1
					l.emit(tokenCell, data[ts:te])
				}
			case 4:
				{
					p = (te) - 1
					l.emit(tokenDDECall, data[ts:te])
				}
			case 10:
				{
					p = (te) - 1
					l.emit(tokenReservedName, data[ts:te])
				}
			case 13:
				{
					p = (te) - 1
					l.emit(tokenNamedRange, data[ts:te])
				}
			case 14:
				{
					p = (te) - 1
					l.emit(tokenString, data[ts+1:te-1])
				}
			}

			goto st27
		tr2:
			// line 62 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenSheet, data[ts:te-1])
			}
			goto st27
		tr10:
			// line 59 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenError, data[ts:te])
			}
			goto st27
		tr15:
			// line 60 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenErrorRef, data[ts:te])
			}
			goto st27
		tr25:
			// line 63 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenSheet, data[ts+1:te-2])
			}
			goto st27
		tr26:
			// line 56 "lexer.rl"

			p = (te) - 1
			{
				l.emit(tokenNumber, data[ts:te])
			}
			goto st27
		tr28:
			// line 67 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenFunctionBuiltin, data[ts:te-1])
			}
			goto st27
		tr40:
			// line 75 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenAmpersand, data[ts:te])
			}
			goto st27
		tr42:
			// line 78 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenLParen, data[ts:te])
			}
			goto st27
		tr43:
			// line 79 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenRParen, data[ts:te])
			}
			goto st27
		tr44:
			// line 82 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenMult, data[ts:te])
			}
			goto st27
		tr45:
			// line 80 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenPlus, data[ts:te])
			}
			goto st27
		tr46:
			// line 94 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenComma, data[ts:te])
			}
			goto st27
		tr47:
			// line 81 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenMinus, data[ts:te])
			}
			goto st27
		tr48:
			// line 83 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenDiv, data[ts:te])
			}
			goto st27
		tr50:
			// line 92 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenColon, data[ts:te])
			}
			goto st27
		tr51:
			// line 93 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenSemi, data[ts:te])
			}
			goto st27
		tr53:
			// line 87 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenEQ, data[ts:te])
			}
			goto st27
		tr59:
			// line 84 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenExp, data[ts:te])
			}
			goto st27
		tr61:
			// line 76 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenLBrace, data[ts:te])
			}
			goto st27
		tr62:
			// line 77 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenRBrace, data[ts:te])
			}
			goto st27
		tr63:
			// line 72 "lexer.rl"

			te = p
			p--
			{
				l.emit(tokenString, data[ts+1:te-1])
			}
			goto st27
		tr64:
			// line 61 "lexer.rl"

			te = p
			p--
			{
				l.emit(tokenHorizontalRange, data[ts:te])
			}
			goto st27
		tr65:
			// line 57 "lexer.rl"

			te = p
			p--
			{
				l.emit(tokenCell, data[ts:te])
			}
			goto st27
		tr66:
			// line 56 "lexer.rl"

			te = p
			p--
			{
				l.emit(tokenNumber, data[ts:te])
			}
			goto st27
		tr69:
			// line 85 "lexer.rl"

			te = p
			p--
			{
				l.emit(tokenLT, data[ts:te])
			}
			goto st27
		tr70:
			// line 88 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenLEQ, data[ts:te])
			}
			goto st27
		tr71:
			// line 90 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenNE, data[ts:te])
			}
			goto st27
		tr72:
			// line 86 "lexer.rl"

			te = p
			p--
			{
				l.emit(tokenGT, data[ts:te])
			}
			goto st27
		tr73:
			// line 89 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenGEQ, data[ts:te])
			}
			goto st27
		tr74:
			// line 70 "lexer.rl"

			te = p
			p--
			{
				l.emit(tokenNamedRange, data[ts:te])
			}
			goto st27
		tr86:
			// line 68 "lexer.rl"

			te = p + 1
			{
				l.emit(tokenFunctionBuiltin, data[ts:te-1])
			}
			goto st27
		st27:
			// line 1 "NONE"

			ts = 0

			// line 1 "NONE"

			act = 0

			if p++; p == pe {
				goto _test_eof27
			}
		st_case_27:
			// line 1 "NONE"

			ts = p

			// line 431 "lexer.go"
			switch data[p] {
			case 34:
				goto st2
			case 35:
				goto st3
			case 36:
				goto st12
			case 38:
				goto tr40
			case 39:
				goto st18
			case 40:
				goto tr42
			case 41:
				goto tr43
			case 42:
				goto tr44
			case 43:
				goto tr45
			case 44:
				goto tr46
			case 45:
				goto tr47
			case 47:
				goto tr48
			case 58:
				goto tr50
			case 59:
				goto tr51
			case 60:
				goto st34
			case 61:
				goto tr53
			case 62:
				goto st35
			case 63:
				goto st0
			case 70:
				goto st23
			case 84:
				goto st24
			case 92:
				goto st25
			case 94:
				goto tr59
			case 95:
				goto st26
			case 123:
				goto tr61
			case 125:
				goto tr62
			}
			switch {
			case data[p] < 48:
				if 33 <= data[p] && data[p] <= 37 {
					goto st0
				}
			case data[p] > 57:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 93 {
						goto st0
					}
				case data[p] >= 65:
					goto st22
				}
			default:
				goto tr49
			}
			goto st1
		st1:
			if p++; p == pe {
				goto _test_eof1
			}
		st_case_1:
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto tr0
			case 123:
				goto tr0
			case 125:
				goto tr0
			}
			switch {
			case data[p] < 37:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr0
				}
			case data[p] > 45:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto tr0
					}
				case data[p] >= 58:
					goto tr0
				}
			default:
				goto tr0
			}
			goto st1
		st_case_0:
		st0:
			cs = 0
			goto _out
		st2:
			if p++; p == pe {
				goto _test_eof2
			}
		st_case_2:
			if data[p] == 34 {
				goto tr4
			}
			goto st2
		tr4:
			// line 1 "NONE"

			te = p + 1

			// line 72 "lexer.rl"

			act = 14
			goto st28
		st28:
			if p++; p == pe {
				goto _test_eof28
			}
		st_case_28:
			// line 562 "lexer.go"
			if data[p] == 34 {
				goto st2
			}
			goto tr63
		st3:
			if p++; p == pe {
				goto _test_eof3
			}
		st_case_3:
			switch data[p] {
			case 78:
				goto st4
			case 82:
				goto st9
			}
			goto st0
		st4:
			if p++; p == pe {
				goto _test_eof4
			}
		st_case_4:
			switch data[p] {
			case 47:
				goto st5
			case 85:
				goto st6
			}
			goto st0
		st5:
			if p++; p == pe {
				goto _test_eof5
			}
		st_case_5:
			if data[p] == 65 {
				goto tr10
			}
			goto st0
		st6:
			if p++; p == pe {
				goto _test_eof6
			}
		st_case_6:
			switch data[p] {
			case 76:
				goto st7
			case 77:
				goto st8
			}
			goto st0
		st7:
			if p++; p == pe {
				goto _test_eof7
			}
		st_case_7:
			if data[p] == 76 {
				goto st8
			}
			goto st0
		st8:
			if p++; p == pe {
				goto _test_eof8
			}
		st_case_8:
			if data[p] == 33 {
				goto tr10
			}
			goto st0
		st9:
			if p++; p == pe {
				goto _test_eof9
			}
		st_case_9:
			if data[p] == 69 {
				goto st10
			}
			goto st0
		st10:
			if p++; p == pe {
				goto _test_eof10
			}
		st_case_10:
			if data[p] == 70 {
				goto st11
			}
			goto st0
		st11:
			if p++; p == pe {
				goto _test_eof11
			}
		st_case_11:
			if data[p] == 33 {
				goto tr15
			}
			goto st0
		st12:
			if p++; p == pe {
				goto _test_eof12
			}
		st_case_12:
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto st0
			case 123:
				goto st0
			case 125:
				goto st0
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 45 {
						goto st0
					}
				case data[p] >= 34:
					goto st0
				}
			case data[p] > 57:
				switch {
				case data[p] < 65:
					if 58 <= data[p] && data[p] <= 63 {
						goto st0
					}
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 95 {
						goto st0
					}
				default:
					goto st16
				}
			default:
				goto st13
			}
			goto st1
		st13:
			if p++; p == pe {
				goto _test_eof13
			}
		st_case_13:
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto st0
			case 58:
				goto st14
			case 123:
				goto st0
			case 125:
				goto st0
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 45 {
						goto st0
					}
				case data[p] >= 34:
					goto st0
				}
			case data[p] > 57:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto st0
					}
				case data[p] >= 59:
					goto st0
				}
			default:
				goto st13
			}
			goto st1
		st14:
			if p++; p == pe {
				goto _test_eof14
			}
		st_case_14:
			if data[p] == 36 {
				goto st15
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
			goto tr0
		st15:
			if p++; p == pe {
				goto _test_eof15
			}
		st_case_15:
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
			goto tr0
		st29:
			if p++; p == pe {
				goto _test_eof29
			}
		st_case_29:
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
			goto tr64
		st16:
			if p++; p == pe {
				goto _test_eof16
			}
		st_case_16:
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 47:
				goto st0
			case 123:
				goto st0
			case 125:
				goto st0
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 45:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr22
					}
				case data[p] >= 34:
					goto st0
				}
			case data[p] > 63:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 95 {
						goto st0
					}
				case data[p] >= 65:
					goto st16
				}
			default:
				goto st0
			}
			goto st1
		st17:
			if p++; p == pe {
				goto _test_eof17
			}
		st_case_17:
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto tr0
			case 123:
				goto tr0
			case 125:
				goto tr0
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 45 {
						goto tr0
					}
				case data[p] >= 34:
					goto tr0
				}
			case data[p] > 57:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto tr0
					}
				case data[p] >= 58:
					goto tr0
				}
			default:
				goto tr22
			}
			goto st1
		tr22:
			// line 1 "NONE"

			te = p + 1

			// line 57 "lexer.rl"

			act = 3
			goto st30
		st30:
			if p++; p == pe {
				goto _test_eof30
			}
		st_case_30:
			// line 861 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto tr65
			case 123:
				goto tr65
			case 125:
				goto tr65
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 45 {
						goto tr65
					}
				case data[p] >= 34:
					goto tr65
				}
			case data[p] > 57:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto tr65
					}
				case data[p] >= 58:
					goto tr65
				}
			default:
				goto tr22
			}
			goto st1
		st18:
			if p++; p == pe {
				goto _test_eof18
			}
		st_case_18:
			switch data[p] {
			case 39:
				goto st0
			case 42:
				goto st0
			case 47:
				goto st0
			case 58:
				goto st0
			case 63:
				goto st0
			}
			if 91 <= data[p] && data[p] <= 93 {
				goto st0
			}
			goto st19
		st19:
			if p++; p == pe {
				goto _test_eof19
			}
		st_case_19:
			switch data[p] {
			case 39:
				goto st20
			case 42:
				goto st0
			case 47:
				goto st0
			case 58:
				goto st0
			case 63:
				goto st0
			}
			if 91 <= data[p] && data[p] <= 93 {
				goto st0
			}
			goto st19
		st20:
			if p++; p == pe {
				goto _test_eof20
			}
		st_case_20:
			if data[p] == 33 {
				goto tr25
			}
			goto st0
		tr49:
			// line 1 "NONE"

			te = p + 1

			// line 56 "lexer.rl"

			act = 2
			goto st31
		st31:
			if p++; p == pe {
				goto _test_eof31
			}
		st_case_31:
			// line 960 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 46:
				goto tr67
			case 58:
				goto st14
			case 101:
				goto st21
			case 123:
				goto tr66
			case 125:
				goto tr66
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 47 {
						goto tr66
					}
				case data[p] >= 34:
					goto tr66
				}
			case data[p] > 57:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto tr66
					}
				case data[p] >= 59:
					goto tr66
				}
			default:
				goto tr49
			}
			goto st1
		tr67:
			// line 1 "NONE"

			te = p + 1

			// line 56 "lexer.rl"

			act = 2
			goto st32
		st32:
			if p++; p == pe {
				goto _test_eof32
			}
		st_case_32:
			// line 1012 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto tr66
			case 101:
				goto st21
			case 123:
				goto tr66
			case 125:
				goto tr66
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 45 {
						goto tr66
					}
				case data[p] >= 34:
					goto tr66
				}
			case data[p] > 57:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto tr66
					}
				case data[p] >= 58:
					goto tr66
				}
			default:
				goto tr67
			}
			goto st1
		st21:
			if p++; p == pe {
				goto _test_eof21
			}
		st_case_21:
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto tr26
			case 123:
				goto tr26
			case 125:
				goto tr26
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 45 {
						goto tr26
					}
				case data[p] >= 34:
					goto tr26
				}
			case data[p] > 57:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto tr26
					}
				case data[p] >= 58:
					goto tr26
				}
			default:
				goto tr27
			}
			goto st1
		tr27:
			// line 1 "NONE"

			te = p + 1

			// line 56 "lexer.rl"

			act = 2
			goto st33
		st33:
			if p++; p == pe {
				goto _test_eof33
			}
		st_case_33:
			// line 1100 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 47:
				goto tr66
			case 123:
				goto tr66
			case 125:
				goto tr66
			}
			switch {
			case data[p] < 48:
				switch {
				case data[p] > 35:
					if 37 <= data[p] && data[p] <= 45 {
						goto tr66
					}
				case data[p] >= 34:
					goto tr66
				}
			case data[p] > 57:
				switch {
				case data[p] > 63:
					if 91 <= data[p] && data[p] <= 95 {
						goto tr66
					}
				case data[p] >= 58:
					goto tr66
				}
			default:
				goto tr27
			}
			goto st1
		st34:
			if p++; p == pe {
				goto _test_eof34
			}
		st_case_34:
			switch data[p] {
			case 61:
				goto tr70
			case 62:
				goto tr71
			}
			goto tr69
		st35:
			if p++; p == pe {
				goto _test_eof35
			}
		st_case_35:
			if data[p] == 61 {
				goto tr73
			}
			goto tr72
		st22:
			if p++; p == pe {
				goto _test_eof22
			}
		st_case_22:
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto st0
			case 125:
				goto st0
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto st0
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto st0
				}
			default:
				goto st0
			}
			goto st1
		tr29:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st36
		st36:
			if p++; p == pe {
				goto _test_eof36
			}
		st_case_36:
			// line 1219 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] < 37:
					if 34 <= data[p] && data[p] <= 35 {
						goto tr74
					}
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr29
					}
				default:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr29
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		st37:
			if p++; p == pe {
				goto _test_eof37
			}
		st_case_37:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st37
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto tr74
		tr33:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st38
		st38:
			if p++; p == pe {
				goto _test_eof38
			}
		st_case_38:
			// line 1307 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 46:
				goto tr33
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] < 37:
					if 34 <= data[p] && data[p] <= 35 {
						goto tr74
					}
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr33
					}
				default:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr33
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		tr30:
			// line 1 "NONE"

			te = p + 1

			// line 57 "lexer.rl"

			act = 3
			goto st39
		st39:
			if p++; p == pe {
				goto _test_eof39
			}
		st_case_39:
			// line 1367 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr65
			case 125:
				goto tr65
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] < 37:
					if 34 <= data[p] && data[p] <= 35 {
						goto tr65
					}
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				default:
					goto tr65
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr29
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr65
				}
			default:
				goto tr65
			}
			goto st1
		tr31:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st40
		tr77:
			// line 1 "NONE"

			te = p + 1

			// line 55 "lexer.rl"

			act = 1
			goto st40
		tr79:
			// line 1 "NONE"

			te = p + 1

			// line 58 "lexer.rl"

			act = 4
			goto st40
		st40:
			if p++; p == pe {
				goto _test_eof40
			}
		st_case_40:
			// line 1447 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr0
			case 125:
				goto tr0
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto tr0
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr0
				}
			default:
				goto tr0
			}
			goto st1
		st23:
			if p++; p == pe {
				goto _test_eof23
			}
		st_case_23:
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 65:
				goto tr34
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto st0
			case 125:
				goto st0
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto st0
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 66 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto st0
				}
			default:
				goto st0
			}
			goto st1
		tr34:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st41
		st41:
			if p++; p == pe {
				goto _test_eof41
			}
		st_case_41:
			// line 1559 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 76:
				goto tr75
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		tr75:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st42
		st42:
			if p++; p == pe {
				goto _test_eof42
			}
		st_case_42:
			// line 1621 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 83:
				goto tr76
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		tr76:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st43
		st43:
			if p++; p == pe {
				goto _test_eof43
			}
		st_case_43:
			// line 1683 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 69:
				goto tr77
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		st24:
			if p++; p == pe {
				goto _test_eof24
			}
		st_case_24:
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 79:
				goto tr35
			case 82:
				goto tr36
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto st0
			case 125:
				goto st0
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto st0
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto st0
				}
			default:
				goto st0
			}
			goto st1
		tr35:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st44
		st44:
			if p++; p == pe {
				goto _test_eof44
			}
		st_case_44:
			// line 1799 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 68:
				goto tr78
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		tr78:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st45
		st45:
			if p++; p == pe {
				goto _test_eof45
			}
		st_case_45:
			// line 1861 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 79:
				goto tr79
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		tr36:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st46
		st46:
			if p++; p == pe {
				goto _test_eof46
			}
		st_case_46:
			// line 1923 "lexer.go"
			switch data[p] {
			case 33:
				goto tr2
			case 36:
				goto st17
			case 40:
				goto tr28
			case 46:
				goto tr29
			case 85:
				goto tr76
			case 92:
				goto st37
			case 95:
				goto st37
			case 123:
				goto tr74
			case 125:
				goto tr74
			}
			switch {
			case data[p] < 58:
				switch {
				case data[p] > 47:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr30
					}
				case data[p] >= 34:
					goto tr74
				}
			case data[p] > 63:
				switch {
				case data[p] < 91:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr31
					}
				case data[p] > 94:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr33
					}
				default:
					goto tr74
				}
			default:
				goto tr74
			}
			goto st1
		st25:
			if p++; p == pe {
				goto _test_eof25
			}
		st_case_25:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st37
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto st0
		st26:
			if p++; p == pe {
				goto _test_eof26
			}
		st_case_26:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st37
			case 120:
				goto st47
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto st0
		st47:
			if p++; p == pe {
				goto _test_eof47
			}
		st_case_47:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st37
			case 108:
				goto st48
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto tr74
		st48:
			if p++; p == pe {
				goto _test_eof48
			}
		st_case_48:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st37
			case 102:
				goto st49
			case 110:
				goto st53
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto tr74
		st49:
			if p++; p == pe {
				goto _test_eof49
			}
		st_case_49:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st37
			case 110:
				goto st50
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto tr74
		st50:
			if p++; p == pe {
				goto _test_eof50
			}
		st_case_50:
			switch data[p] {
			case 46:
				goto st51
			case 92:
				goto st37
			case 95:
				goto st37
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto tr74
		st51:
			if p++; p == pe {
				goto _test_eof51
			}
		st_case_51:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st52
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st52
			}
			goto tr74
		st52:
			if p++; p == pe {
				goto _test_eof52
			}
		st_case_52:
			switch data[p] {
			case 40:
				goto tr86
			case 46:
				goto st52
			case 92:
				goto st37
			case 95:
				goto st37
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st52
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st52
			}
			goto tr74
		st53:
			if p++; p == pe {
				goto _test_eof53
			}
		st_case_53:
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto st37
			case 109:
				goto st54
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto tr74
		st54:
			if p++; p == pe {
				goto _test_eof54
			}
		st_case_54:
			switch data[p] {
			case 46:
				goto tr88
			case 92:
				goto st37
			case 95:
				goto st37
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto st37
			}
			goto tr74
		tr88:
			// line 1 "NONE"

			te = p + 1

			// line 70 "lexer.rl"

			act = 13
			goto st55
		tr89:
			// line 1 "NONE"

			te = p + 1

			// line 64 "lexer.rl"

			act = 10
			goto st55
		st55:
			if p++; p == pe {
				goto _test_eof55
			}
		st_case_55:
			// line 2268 "lexer.go"
			switch data[p] {
			case 46:
				goto st37
			case 92:
				goto st37
			case 95:
				goto tr89
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			default:
				goto tr89
			}
			goto tr0
		st_out:
		_test_eof27:
			cs = 27
			goto _test_eof
		_test_eof1:
			cs = 1
			goto _test_eof
		_test_eof2:
			cs = 2
			goto _test_eof
		_test_eof28:
			cs = 28
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
		_test_eof13:
			cs = 13
			goto _test_eof
		_test_eof14:
			cs = 14
			goto _test_eof
		_test_eof15:
			cs = 15
			goto _test_eof
		_test_eof29:
			cs = 29
			goto _test_eof
		_test_eof16:
			cs = 16
			goto _test_eof
		_test_eof17:
			cs = 17
			goto _test_eof
		_test_eof30:
			cs = 30
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
		_test_eof31:
			cs = 31
			goto _test_eof
		_test_eof32:
			cs = 32
			goto _test_eof
		_test_eof21:
			cs = 21
			goto _test_eof
		_test_eof33:
			cs = 33
			goto _test_eof
		_test_eof34:
			cs = 34
			goto _test_eof
		_test_eof35:
			cs = 35
			goto _test_eof
		_test_eof22:
			cs = 22
			goto _test_eof
		_test_eof36:
			cs = 36
			goto _test_eof
		_test_eof37:
			cs = 37
			goto _test_eof
		_test_eof38:
			cs = 38
			goto _test_eof
		_test_eof39:
			cs = 39
			goto _test_eof
		_test_eof40:
			cs = 40
			goto _test_eof
		_test_eof23:
			cs = 23
			goto _test_eof
		_test_eof41:
			cs = 41
			goto _test_eof
		_test_eof42:
			cs = 42
			goto _test_eof
		_test_eof43:
			cs = 43
			goto _test_eof
		_test_eof24:
			cs = 24
			goto _test_eof
		_test_eof44:
			cs = 44
			goto _test_eof
		_test_eof45:
			cs = 45
			goto _test_eof
		_test_eof46:
			cs = 46
			goto _test_eof
		_test_eof25:
			cs = 25
			goto _test_eof
		_test_eof26:
			cs = 26
			goto _test_eof
		_test_eof47:
			cs = 47
			goto _test_eof
		_test_eof48:
			cs = 48
			goto _test_eof
		_test_eof49:
			cs = 49
			goto _test_eof
		_test_eof50:
			cs = 50
			goto _test_eof
		_test_eof51:
			cs = 51
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

		_test_eof:
			{
			}
			if p == eof {
				switch cs {
				case 1:
					goto tr0
				case 2:
					goto tr0
				case 28:
					goto tr63
				case 14:
					goto tr0
				case 15:
					goto tr0
				case 29:
					goto tr64
				case 17:
					goto tr0
				case 30:
					goto tr65
				case 31:
					goto tr66
				case 32:
					goto tr66
				case 21:
					goto tr26
				case 33:
					goto tr66
				case 34:
					goto tr69
				case 35:
					goto tr72
				case 36:
					goto tr74
				case 37:
					goto tr74
				case 38:
					goto tr74
				case 39:
					goto tr65
				case 40:
					goto tr0
				case 41:
					goto tr74
				case 42:
					goto tr74
				case 43:
					goto tr74
				case 44:
					goto tr74
				case 45:
					goto tr74
				case 46:
					goto tr74
				case 47:
					goto tr74
				case 48:
					goto tr74
				case 49:
					goto tr74
				case 50:
					goto tr74
				case 51:
					goto tr74
				case 52:
					goto tr74
				case 53:
					goto tr74
				case 54:
					goto tr74
				case 55:
					goto tr0
				}
			}

		_out:
			{
			}
		}

		// line 143 "lexer.rl"

		if ts > 0 {
			// currently parsing a token, so shift it to the
			// beginning of the buffer
			copy(data[0:], data[ts:])
		}
	}

	_ = eof
	if cs == formula_error {
		l.emit(tokenLexError, nil)
	}
	close(l.nodes)
}
