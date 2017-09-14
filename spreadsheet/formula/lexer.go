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
	"log"
)

// line 20 "lexer.go"
var _formula_actions []byte = []byte{
	0, 1, 2, 1, 10, 1, 11, 1, 12,
	1, 13, 1, 14, 1, 15, 1, 16,
	1, 17, 1, 18, 1, 19, 1, 20,
	1, 21, 1, 22, 1, 23, 1, 24,
	1, 25, 1, 26, 1, 27, 1, 28,
	1, 29, 1, 30, 1, 31, 1, 32,
	1, 33, 1, 34, 1, 35, 1, 36,
	1, 37, 1, 38, 2, 0, 1, 2,
	3, 4, 2, 3, 5, 2, 3, 6,
	2, 3, 7, 2, 3, 8, 2, 3,
	9,
}

var _formula_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 61, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0,
}

var _formula_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0,
}

const formula_start int = 26
const formula_first_final int = 26
const formula_error int = 0

const formula_en_main int = 26

// line 103 "lexer.rl"

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

		// line 97 "lexer.go"
		{
			cs = formula_start
			ts = 0
			te = 0
			act = 0
		}

		// line 105 "lexer.go"
		{
			var _acts int
			var _nacts uint

			if p == pe {
				goto _test_eof
			}
			if cs == 0 {
				goto _out
			}
		_resume:
			_acts = int(_formula_from_state_actions[cs])
			_nacts = uint(_formula_actions[_acts])
			_acts++
			for ; _nacts > 0; _nacts-- {
				_acts++
				switch _formula_actions[_acts-1] {
				case 2:
					// line 1 "NONE"

					ts = p

					// line 127 "lexer.go"
				}
			}

			switch cs {
			case 26:
				switch data[p] {
				case 34:
					goto tr1
				case 35:
					goto tr31
				case 36:
					goto tr32
				case 37:
					goto tr4
				case 38:
					goto tr33
				case 39:
					goto tr34
				case 40:
					goto tr35
				case 41:
					goto tr36
				case 42:
					goto tr37
				case 43:
					goto tr38
				case 44:
					goto tr39
				case 45:
					goto tr40
				case 47:
					goto tr41
				case 58:
					goto tr43
				case 60:
					goto tr44
				case 61:
					goto tr45
				case 62:
					goto tr46
				case 70:
					goto tr48
				case 84:
					goto tr49
				case 94:
					goto tr50
				case 95:
					goto tr51
				case 123:
					goto tr52
				case 125:
					goto tr53
				}
				switch {
				case data[p] < 59:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr42
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 93 {
							goto tr4
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr4
				}
				goto tr30
			case 27:
				switch data[p] {
				case 47:
					goto tr54
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 37:
					if 34 <= data[p] && data[p] <= 35 {
						goto tr54
					}
				case data[p] > 45:
					switch {
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 58:
						goto tr54
					}
				default:
					goto tr54
				}
				goto tr30
			case 1:
				if data[p] == 34 {
					goto tr2
				}
				goto tr1
			case 28:
				if data[p] == 34 {
					goto tr1
				}
				goto tr55
			case 2:
				switch data[p] {
				case 78:
					goto tr3
				case 82:
					goto tr5
				}
				goto tr4
			case 0:
				goto _out
			case 3:
				switch data[p] {
				case 47:
					goto tr6
				case 85:
					goto tr7
				}
				goto tr4
			case 4:
				if data[p] == 65 {
					goto tr8
				}
				goto tr4
			case 5:
				switch data[p] {
				case 76:
					goto tr9
				case 77:
					goto tr10
				}
				goto tr4
			case 6:
				if data[p] == 76 {
					goto tr10
				}
				goto tr4
			case 7:
				if data[p] == 33 {
					goto tr8
				}
				goto tr4
			case 8:
				if data[p] == 69 {
					goto tr11
				}
				goto tr4
			case 9:
				if data[p] == 70 {
					goto tr12
				}
				goto tr4
			case 10:
				if data[p] == 33 {
					goto tr13
				}
				goto tr4
			case 29:
				switch data[p] {
				case 47:
					goto tr54
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 45 {
							goto tr54
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr54
						}
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					default:
						goto tr57
					}
				default:
					goto tr56
				}
				goto tr30
			case 30:
				switch data[p] {
				case 47:
					goto tr54
				case 58:
					goto tr58
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 45 {
							goto tr54
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 57:
					switch {
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 59:
						goto tr54
					}
				default:
					goto tr56
				}
				goto tr30
			case 11:
				if data[p] == 36 {
					goto tr14
				}
				if 48 <= data[p] && data[p] <= 57 {
					goto tr15
				}
				goto tr0
			case 12:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr15
				}
				goto tr0
			case 31:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr15
				}
				goto tr59
			case 32:
				switch data[p] {
				case 36:
					goto tr60
				case 47:
					goto tr54
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr57
					}
				default:
					goto tr54
				}
				goto tr30
			case 33:
				switch data[p] {
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
					goto tr61
				}
				goto tr30
			case 13:
				switch data[p] {
				case 39:
					goto tr4
				case 42:
					goto tr4
				case 47:
					goto tr4
				case 58:
					goto tr4
				case 63:
					goto tr4
				}
				if 91 <= data[p] && data[p] <= 93 {
					goto tr4
				}
				goto tr16
			case 14:
				switch data[p] {
				case 39:
					goto tr17
				case 42:
					goto tr4
				case 47:
					goto tr4
				case 58:
					goto tr4
				case 63:
					goto tr4
				}
				if 91 <= data[p] && data[p] <= 93 {
					goto tr4
				}
				goto tr16
			case 15:
				if data[p] == 33 {
					goto tr18
				}
				goto tr4
			case 34:
				switch data[p] {
				case 46:
					goto tr63
				case 58:
					goto tr58
				case 101:
					goto tr64
				case 123:
					goto tr62
				case 125:
					goto tr62
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 47 {
							goto tr62
						}
					case data[p] >= 34:
						goto tr62
					}
				case data[p] > 57:
					switch {
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr62
						}
					case data[p] >= 59:
						goto tr62
					}
				default:
					goto tr42
				}
				goto tr30
			case 35:
				switch data[p] {
				case 47:
					goto tr62
				case 101:
					goto tr64
				case 123:
					goto tr62
				case 125:
					goto tr62
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 45 {
							goto tr62
						}
					case data[p] >= 34:
						goto tr62
					}
				case data[p] > 57:
					switch {
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr62
						}
					case data[p] >= 58:
						goto tr62
					}
				default:
					goto tr63
				}
				goto tr30
			case 36:
				switch data[p] {
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
					goto tr65
				}
				goto tr30
			case 37:
				switch data[p] {
				case 61:
					goto tr67
				case 62:
					goto tr68
				}
				goto tr66
			case 38:
				if data[p] == 61 {
					goto tr70
				}
				goto tr69
			case 39:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr0
				case 123:
					goto tr0
				case 125:
					goto tr0
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr0
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr0
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr0
				}
				goto tr30
			case 40:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 65:
					goto tr72
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 66:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 41:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 76:
					goto tr73
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 42:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 83:
					goto tr74
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 43:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 69:
					goto tr75
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 44:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 79:
					goto tr76
				case 82:
					goto tr77
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 45:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 68:
					goto tr78
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 46:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 79:
					goto tr79
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 47:
				switch data[p] {
				case 36:
					goto tr60
				case 40:
					goto tr71
				case 47:
					goto tr54
				case 85:
					goto tr74
				case 123:
					goto tr54
				case 125:
					goto tr54
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] > 45:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr61
						}
					case data[p] >= 34:
						goto tr54
					}
				case data[p] > 63:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 95 {
							goto tr54
						}
					case data[p] >= 65:
						goto tr47
					}
				default:
					goto tr54
				}
				goto tr30
			case 16:
				if data[p] == 120 {
					goto tr19
				}
				goto tr4
			case 17:
				if data[p] == 108 {
					goto tr20
				}
				goto tr4
			case 18:
				switch data[p] {
				case 102:
					goto tr21
				case 110:
					goto tr22
				}
				goto tr4
			case 19:
				if data[p] == 110 {
					goto tr23
				}
				goto tr4
			case 20:
				if data[p] == 46 {
					goto tr24
				}
				goto tr4
			case 21:
				if data[p] == 95 {
					goto tr25
				}
				if 65 <= data[p] && data[p] <= 90 {
					goto tr25
				}
				goto tr4
			case 22:
				switch data[p] {
				case 40:
					goto tr26
				case 95:
					goto tr25
				}
				if 65 <= data[p] && data[p] <= 90 {
					goto tr25
				}
				goto tr4
			case 23:
				if data[p] == 109 {
					goto tr27
				}
				goto tr4
			case 24:
				if data[p] == 46 {
					goto tr28
				}
				goto tr4
			case 25:
				if data[p] == 95 {
					goto tr29
				}
				if 65 <= data[p] && data[p] <= 90 {
					goto tr29
				}
				goto tr4
			case 48:
				if data[p] == 95 {
					goto tr29
				}
				if 65 <= data[p] && data[p] <= 90 {
					goto tr29
				}
				goto tr80
			}

		tr4:
			cs = 0
			goto _again
		tr1:
			cs = 1
			goto _again
		tr31:
			cs = 2
			goto _again
		tr3:
			cs = 3
			goto _again
		tr6:
			cs = 4
			goto _again
		tr7:
			cs = 5
			goto _again
		tr9:
			cs = 6
			goto _again
		tr10:
			cs = 7
			goto _again
		tr5:
			cs = 8
			goto _again
		tr11:
			cs = 9
			goto _again
		tr12:
			cs = 10
			goto _again
		tr58:
			cs = 11
			goto _again
		tr14:
			cs = 12
			goto _again
		tr34:
			cs = 13
			goto _again
		tr16:
			cs = 14
			goto _again
		tr17:
			cs = 15
			goto _again
		tr51:
			cs = 16
			goto _again
		tr19:
			cs = 17
			goto _again
		tr20:
			cs = 18
			goto _again
		tr21:
			cs = 19
			goto _again
		tr23:
			cs = 20
			goto _again
		tr24:
			cs = 21
			goto _again
		tr25:
			cs = 22
			goto _again
		tr22:
			cs = 23
			goto _again
		tr27:
			cs = 24
			goto _again
		tr28:
			cs = 25
			goto _again
		tr0:
			cs = 26
			goto f0
		tr8:
			cs = 26
			goto f2
		tr13:
			cs = 26
			goto f3
		tr18:
			cs = 26
			goto f4
		tr26:
			cs = 26
			goto f5
		tr33:
			cs = 26
			goto f8
		tr35:
			cs = 26
			goto f9
		tr36:
			cs = 26
			goto f10
		tr37:
			cs = 26
			goto f11
		tr38:
			cs = 26
			goto f12
		tr39:
			cs = 26
			goto f13
		tr40:
			cs = 26
			goto f14
		tr41:
			cs = 26
			goto f15
		tr43:
			cs = 26
			goto f17
		tr45:
			cs = 26
			goto f18
		tr50:
			cs = 26
			goto f20
		tr52:
			cs = 26
			goto f21
		tr53:
			cs = 26
			goto f22
		tr54:
			cs = 26
			goto f23
		tr55:
			cs = 26
			goto f24
		tr59:
			cs = 26
			goto f25
		tr62:
			cs = 26
			goto f27
		tr66:
			cs = 26
			goto f28
		tr67:
			cs = 26
			goto f29
		tr68:
			cs = 26
			goto f30
		tr69:
			cs = 26
			goto f31
		tr70:
			cs = 26
			goto f32
		tr71:
			cs = 26
			goto f33
		tr80:
			cs = 26
			goto f36
		tr30:
			cs = 27
			goto _again
		tr2:
			cs = 28
			goto f1
		tr32:
			cs = 29
			goto _again
		tr56:
			cs = 30
			goto f19
		tr15:
			cs = 31
			goto _again
		tr57:
			cs = 32
			goto _again
		tr60:
			cs = 33
			goto f19
		tr61:
			cs = 33
			goto f26
		tr42:
			cs = 34
			goto f16
		tr63:
			cs = 35
			goto _again
		tr65:
			cs = 36
			goto f16
		tr64:
			cs = 36
			goto f19
		tr44:
			cs = 37
			goto _again
		tr46:
			cs = 38
			goto _again
		tr47:
			cs = 39
			goto f19
		tr75:
			cs = 39
			goto f34
		tr79:
			cs = 39
			goto f35
		tr48:
			cs = 40
			goto _again
		tr72:
			cs = 41
			goto _again
		tr73:
			cs = 42
			goto _again
		tr74:
			cs = 43
			goto _again
		tr49:
			cs = 44
			goto _again
		tr76:
			cs = 45
			goto _again
		tr78:
			cs = 46
			goto _again
		tr77:
			cs = 47
			goto _again
		tr29:
			cs = 48
			goto _again

		f2:
			_acts = 3
			goto execFuncs
		f3:
			_acts = 5
			goto execFuncs
		f4:
			_acts = 7
			goto execFuncs
		f33:
			_acts = 9
			goto execFuncs
		f5:
			_acts = 11
			goto execFuncs
		f8:
			_acts = 13
			goto execFuncs
		f21:
			_acts = 15
			goto execFuncs
		f22:
			_acts = 17
			goto execFuncs
		f9:
			_acts = 19
			goto execFuncs
		f10:
			_acts = 21
			goto execFuncs
		f12:
			_acts = 23
			goto execFuncs
		f14:
			_acts = 25
			goto execFuncs
		f11:
			_acts = 27
			goto execFuncs
		f15:
			_acts = 29
			goto execFuncs
		f20:
			_acts = 31
			goto execFuncs
		f18:
			_acts = 33
			goto execFuncs
		f29:
			_acts = 35
			goto execFuncs
		f32:
			_acts = 37
			goto execFuncs
		f30:
			_acts = 39
			goto execFuncs
		f17:
			_acts = 41
			goto execFuncs
		f13:
			_acts = 43
			goto execFuncs
		f27:
			_acts = 45
			goto execFuncs
		f25:
			_acts = 47
			goto execFuncs
		f23:
			_acts = 49
			goto execFuncs
		f36:
			_acts = 51
			goto execFuncs
		f24:
			_acts = 53
			goto execFuncs
		f28:
			_acts = 55
			goto execFuncs
		f31:
			_acts = 57
			goto execFuncs
		f0:
			_acts = 59
			goto execFuncs
		f34:
			_acts = 64
			goto execFuncs
		f16:
			_acts = 67
			goto execFuncs
		f26:
			_acts = 70
			goto execFuncs
		f35:
			_acts = 73
			goto execFuncs
		f19:
			_acts = 76
			goto execFuncs
		f1:
			_acts = 79
			goto execFuncs

		execFuncs:
			_nacts = uint(_formula_actions[_acts])
			_acts++
			for ; _nacts > 0; _nacts-- {
				_acts++
				switch _formula_actions[_acts-1] {
				case 3:
					// line 1 "NONE"

					te = p + 1

				case 4:
					// line 54 "lexer.rl"

					act = 1
				case 5:
					// line 55 "lexer.rl"

					act = 2
				case 6:
					// line 56 "lexer.rl"

					act = 3
				case 7:
					// line 57 "lexer.rl"

					act = 4
				case 8:
					// line 61 "lexer.rl"

					act = 8
				case 9:
					// line 69 "lexer.rl"

					act = 12
				case 10:
					// line 58 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenError, data[ts:te])
					}
				case 11:
					// line 59 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenErrorRef, data[ts:te])
					}
				case 12:
					// line 61 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenSheet, data[ts:te])
					}
				case 13:
					// line 65 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenFunctionBultin, data[ts:te-1])
					}
				case 14:
					// line 66 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenFunctionBultin, data[ts:te-1])
					}
				case 15:
					// line 72 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenAmpersand, data[ts:te])
					}
				case 16:
					// line 73 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenLBrace, data[ts:te])
					}
				case 17:
					// line 74 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenRBrace, data[ts:te])
					}
				case 18:
					// line 75 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenLParen, data[ts:te])
					}
				case 19:
					// line 76 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenRParen, data[ts:te])
					}
				case 20:
					// line 77 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenPlus, data[ts:te])
					}
				case 21:
					// line 78 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenMinus, data[ts:te])
					}
				case 22:
					// line 79 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenMult, data[ts:te])
					}
				case 23:
					// line 80 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenDiv, data[ts:te])
					}
				case 24:
					// line 81 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenExp, data[ts:te])
					}
				case 25:
					// line 84 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenEQ, data[ts:te])
					}
				case 26:
					// line 85 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenLEQ, data[ts:te])
					}
				case 27:
					// line 86 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenGEQ, data[ts:te])
					}
				case 28:
					// line 87 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenNE, data[ts:te])
					}
				case 29:
					// line 89 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenColon, data[ts:te])
					}
				case 30:
					// line 90 "lexer.rl"

					te = p + 1
					{
						l.emit(tokenComma, data[ts:te])
					}
				case 31:
					// line 55 "lexer.rl"

					te = p
					p--
					{
						l.emit(tokenNumber, data[ts:te])
					}
				case 32:
					// line 60 "lexer.rl"

					te = p
					p--
					{
						l.emit(tokenHorizontalRange, data[ts:te])
					}
				case 33:
					// line 61 "lexer.rl"

					te = p
					p--
					{
						l.emit(tokenSheet, data[ts:te])
					}
				case 34:
					// line 62 "lexer.rl"

					te = p
					p--
					{
						l.emit(tokenReservedName, data[ts:te])
					}
				case 35:
					// line 69 "lexer.rl"

					te = p
					p--
					{
						l.emit(tokenString, data[ts+1:te-1])
					}
				case 36:
					// line 82 "lexer.rl"

					te = p
					p--
					{
						l.emit(tokenLT, data[ts:te])
					}
				case 37:
					// line 83 "lexer.rl"

					te = p
					p--
					{
						l.emit(tokenGT, data[ts:te])
					}
				case 38:
					// line 1 "NONE"

					switch act {
					case 0:
						{
							cs = 0
							goto _again
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
					case 8:
						{
							p = (te) - 1
							l.emit(tokenSheet, data[ts:te])
						}
					case 12:
						{
							p = (te) - 1
							l.emit(tokenString, data[ts+1:te-1])
						}
					}

					// line 1344 "lexer.go"
				}
			}
			goto _again

		_again:
			_acts = int(_formula_to_state_actions[cs])
			_nacts = uint(_formula_actions[_acts])
			_acts++
			for ; _nacts > 0; _nacts-- {
				_acts++
				switch _formula_actions[_acts-1] {
				case 0:
					// line 1 "NONE"

					ts = 0

				case 1:
					// line 1 "NONE"

					act = 0

					// line 1365 "lexer.go"
				}
			}

			if cs == 0 {
				goto _out
			}
			if p++; p != pe {
				goto _resume
			}
		_test_eof:
			{
			}
			if p == eof {
				switch cs {
				case 27:
					goto tr54
				case 1:
					goto tr0
				case 28:
					goto tr55
				case 29:
					goto tr54
				case 30:
					goto tr54
				case 11:
					goto tr0
				case 12:
					goto tr0
				case 31:
					goto tr59
				case 32:
					goto tr54
				case 33:
					goto tr0
				case 34:
					goto tr62
				case 35:
					goto tr62
				case 36:
					goto tr0
				case 37:
					goto tr66
				case 38:
					goto tr69
				case 39:
					goto tr0
				case 40:
					goto tr54
				case 41:
					goto tr54
				case 42:
					goto tr54
				case 43:
					goto tr54
				case 44:
					goto tr54
				case 45:
					goto tr54
				case 46:
					goto tr54
				case 47:
					goto tr54
				case 48:
					goto tr80
				}
			}

		_out:
			{
			}
		}

		// line 139 "lexer.rl"

		if ts > 0 {
			// currently parsing a token, so shift it to the
			// beginning of the buffer
			copy(data[0:], data[ts:])
		}
	}

	_ = eof
	if cs == formula_error {
		log.Fatal("parse error")
	}
	close(l.nodes)
}
