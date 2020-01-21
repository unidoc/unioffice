//line lexer.rl:1
// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"io"
)

//line lexer.go:20
var _formula_actions []byte = []byte{
	0, 1, 2, 1, 12, 1, 13, 1, 14,
	1, 15, 1, 16, 1, 17, 1, 18,
	1, 19, 1, 20, 1, 21, 1, 22,
	1, 23, 1, 24, 1, 25, 1, 26,
	1, 27, 1, 28, 1, 29, 1, 30,
	1, 31, 1, 32, 1, 33, 1, 34,
	1, 35, 1, 36, 1, 37, 1, 38,
	1, 39, 1, 40, 1, 41, 1, 42,
	1, 43, 1, 44, 2, 0, 1, 2,
	3, 4, 2, 3, 5, 2, 3, 6,
	2, 3, 7, 2, 3, 8, 2, 3,
	9, 2, 3, 10, 2, 3, 11,
}

var _formula_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0,
}

var _formula_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0,
}

const formula_start int = 27
const formula_first_final int = 27
const formula_error int = 0

const formula_en_main int = 27

//line lexer.rl:107

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

//line lexer.go:100
		{
			cs = formula_start
			ts = 0
			te = 0
			act = 0
		}

//line lexer.go:108
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
//line NONE:1
					ts = p

//line lexer.go:129
				}
			}

			switch cs {
			case 27:
				switch data[p] {
				case 34:
					goto tr3
				case 35:
					goto tr40
				case 36:
					goto tr41
				case 38:
					goto tr42
				case 39:
					goto tr23
				case 40:
					goto tr43
				case 41:
					goto tr44
				case 42:
					goto tr45
				case 43:
					goto tr46
				case 44:
					goto tr47
				case 45:
					goto tr48
				case 47:
					goto tr49
				case 58:
					goto tr51
				case 59:
					goto tr52
				case 60:
					goto tr53
				case 61:
					goto tr54
				case 62:
					goto tr55
				case 63:
					goto tr6
				case 70:
					goto tr57
				case 84:
					goto tr58
				case 92:
					goto tr59
				case 94:
					goto tr60
				case 95:
					goto tr61
				case 123:
					goto tr62
				case 125:
					goto tr63
				}
				switch {
				case data[p] < 48:
					if 33 <= data[p] && data[p] <= 37 {
						goto tr6
					}
				case data[p] > 57:
					switch {
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 93 {
							goto tr6
						}
					case data[p] >= 65:
						goto tr56
					}
				default:
					goto tr50
				}
				goto tr1
			case 1:
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
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr0
						}
					case data[p] >= 34:
						goto tr0
					}
				case data[p] > 45:
					switch {
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr0
						}
					case data[p] >= 58:
						goto tr0
					}
				default:
					goto tr0
				}
				goto tr1
			case 0:
				goto _out
			case 2:
				if data[p] == 34 {
					goto tr4
				}
				goto tr3
			case 28:
				if data[p] == 34 {
					goto tr3
				}
				goto tr64
			case 3:
				switch data[p] {
				case 78:
					goto tr5
				case 82:
					goto tr7
				}
				goto tr6
			case 4:
				switch data[p] {
				case 47:
					goto tr8
				case 85:
					goto tr9
				}
				goto tr6
			case 5:
				if data[p] == 65 {
					goto tr10
				}
				goto tr6
			case 6:
				switch data[p] {
				case 76:
					goto tr11
				case 77:
					goto tr12
				}
				goto tr6
			case 7:
				if data[p] == 76 {
					goto tr12
				}
				goto tr6
			case 8:
				if data[p] == 33 {
					goto tr10
				}
				goto tr6
			case 9:
				if data[p] == 69 {
					goto tr13
				}
				goto tr6
			case 10:
				if data[p] == 70 {
					goto tr14
				}
				goto tr6
			case 11:
				if data[p] == 33 {
					goto tr15
				}
				goto tr6
			case 12:
				switch data[p] {
				case 33:
					goto tr2
				case 47:
					goto tr6
				case 123:
					goto tr6
				case 125:
					goto tr6
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr6
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 45 {
							goto tr6
						}
					default:
						goto tr6
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr6
						}
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr6
						}
					default:
						goto tr17
					}
				default:
					goto tr16
				}
				goto tr1
			case 13:
				switch data[p] {
				case 33:
					goto tr2
				case 47:
					goto tr6
				case 58:
					goto tr18
				case 123:
					goto tr6
				case 125:
					goto tr6
				}
				switch {
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr6
						}
					case data[p] >= 34:
						goto tr6
					}
				case data[p] > 45:
					switch {
					case data[p] < 59:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr16
						}
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr6
						}
					default:
						goto tr6
					}
				default:
					goto tr6
				}
				goto tr1
			case 14:
				if data[p] == 36 {
					goto tr19
				}
				if 48 <= data[p] && data[p] <= 57 {
					goto tr20
				}
				goto tr0
			case 15:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr20
				}
				goto tr0
			case 29:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr20
				}
				goto tr65
			case 16:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 47:
					goto tr6
				case 123:
					goto tr6
				case 125:
					goto tr6
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 45 {
							goto tr6
						}
					case data[p] >= 34:
						goto tr6
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr6
						}
					case data[p] > 90:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr6
						}
					default:
						goto tr17
					}
				default:
					goto tr22
				}
				goto tr1
			case 17:
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
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr0
						}
					case data[p] >= 34:
						goto tr0
					}
				case data[p] > 45:
					switch {
					case data[p] < 58:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr22
						}
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr0
						}
					default:
						goto tr0
					}
				default:
					goto tr0
				}
				goto tr1
			case 30:
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
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr66
						}
					case data[p] >= 34:
						goto tr66
					}
				case data[p] > 45:
					switch {
					case data[p] < 58:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr22
						}
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr66
						}
					default:
						goto tr66
					}
				default:
					goto tr66
				}
				goto tr1
			case 18:
				switch data[p] {
				case 33:
					goto tr24
				case 39:
					goto tr1
				case 42:
					goto tr6
				case 47:
					goto tr6
				case 58:
					goto tr6
				case 63:
					goto tr6
				case 94:
					goto tr25
				case 123:
					goto tr25
				case 125:
					goto tr25
				}
				switch {
				case data[p] < 37:
					if 34 <= data[p] && data[p] <= 35 {
						goto tr25
					}
				case data[p] > 45:
					switch {
					case data[p] > 62:
						if 91 <= data[p] && data[p] <= 93 {
							goto tr6
						}
					case data[p] >= 59:
						goto tr25
					}
				default:
					goto tr25
				}
				goto tr23
			case 31:
				switch data[p] {
				case 39:
					goto tr26
				case 42:
					goto tr67
				case 47:
					goto tr67
				case 58:
					goto tr67
				case 63:
					goto tr67
				}
				if 91 <= data[p] && data[p] <= 93 {
					goto tr67
				}
				goto tr25
			case 19:
				switch data[p] {
				case 39:
					goto tr26
				case 42:
					goto tr0
				case 47:
					goto tr0
				case 58:
					goto tr0
				case 63:
					goto tr0
				}
				if 91 <= data[p] && data[p] <= 93 {
					goto tr0
				}
				goto tr25
			case 20:
				if data[p] == 33 {
					goto tr27
				}
				goto tr0
			case 32:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr69
				case 58:
					goto tr18
				case 101:
					goto tr70
				case 123:
					goto tr68
				case 125:
					goto tr68
				}
				switch {
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr68
						}
					case data[p] >= 34:
						goto tr68
					}
				case data[p] > 47:
					switch {
					case data[p] < 59:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr50
						}
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr68
						}
					default:
						goto tr68
					}
				default:
					goto tr68
				}
				goto tr1
			case 33:
				switch data[p] {
				case 33:
					goto tr2
				case 47:
					goto tr68
				case 101:
					goto tr70
				case 123:
					goto tr68
				case 125:
					goto tr68
				}
				switch {
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr68
						}
					case data[p] >= 34:
						goto tr68
					}
				case data[p] > 45:
					switch {
					case data[p] < 58:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr69
						}
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr68
						}
					default:
						goto tr68
					}
				default:
					goto tr68
				}
				goto tr1
			case 21:
				switch data[p] {
				case 33:
					goto tr2
				case 47:
					goto tr28
				case 123:
					goto tr28
				case 125:
					goto tr28
				}
				switch {
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr28
						}
					case data[p] >= 34:
						goto tr28
					}
				case data[p] > 45:
					switch {
					case data[p] < 58:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr29
						}
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr28
						}
					default:
						goto tr28
					}
				default:
					goto tr28
				}
				goto tr1
			case 34:
				switch data[p] {
				case 33:
					goto tr2
				case 47:
					goto tr68
				case 123:
					goto tr68
				case 125:
					goto tr68
				}
				switch {
				case data[p] < 40:
					switch {
					case data[p] > 35:
						if 37 <= data[p] && data[p] <= 38 {
							goto tr68
						}
					case data[p] >= 34:
						goto tr68
					}
				case data[p] > 45:
					switch {
					case data[p] < 58:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr29
						}
					case data[p] > 63:
						if 91 <= data[p] && data[p] <= 94 {
							goto tr68
						}
					default:
						goto tr68
					}
				default:
					goto tr68
				}
				goto tr1
			case 35:
				switch data[p] {
				case 61:
					goto tr72
				case 62:
					goto tr73
				}
				goto tr71
			case 36:
				if data[p] == 61 {
					goto tr75
				}
				goto tr74
			case 22:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr6
				case 125:
					goto tr6
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr6
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr6
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr6
					}
				default:
					goto tr6
				}
				goto tr1
			case 37:
				switch data[p] {
				case 33:
					goto tr2
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 41 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr31
					}
				default:
					goto tr31
				}
				goto tr1
			case 38:
				switch data[p] {
				case 46:
					goto tr34
				case 92:
					goto tr34
				case 95:
					goto tr34
				}
				switch {
				case data[p] < 65:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr34
					}
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr34
					}
				default:
					goto tr34
				}
				goto tr76
			case 39:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 40:
				switch data[p] {
				case 33:
					goto tr2
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr66
				case 125:
					goto tr66
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr66
						}
					case data[p] > 38:
						if 41 <= data[p] && data[p] <= 47 {
							goto tr66
						}
					default:
						goto tr66
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr66
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr66
						}
					default:
						goto tr31
					}
				default:
					goto tr32
				}
				goto tr1
			case 41:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr0
				case 125:
					goto tr0
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr0
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr0
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr0
					}
				default:
					goto tr0
				}
				goto tr1
			case 23:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 65:
					goto tr36
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr6
				case 125:
					goto tr6
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr6
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr6
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 66 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr6
					}
				default:
					goto tr6
				}
				goto tr1
			case 42:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 76:
					goto tr77
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr76
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr76
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr76
					}
				default:
					goto tr76
				}
				goto tr1
			case 43:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 83:
					goto tr78
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr76
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr76
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr76
					}
				default:
					goto tr76
				}
				goto tr1
			case 44:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 69:
					goto tr79
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr76
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr76
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr76
					}
				default:
					goto tr76
				}
				goto tr1
			case 24:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 79:
					goto tr37
				case 82:
					goto tr38
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr6
				case 125:
					goto tr6
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr6
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr6
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr6
					}
				default:
					goto tr6
				}
				goto tr1
			case 45:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 68:
					goto tr80
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr76
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr76
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr76
					}
				default:
					goto tr76
				}
				goto tr1
			case 46:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 79:
					goto tr81
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr76
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr76
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr76
					}
				default:
					goto tr76
				}
				goto tr1
			case 47:
				switch data[p] {
				case 33:
					goto tr2
				case 36:
					goto tr21
				case 40:
					goto tr30
				case 46:
					goto tr31
				case 85:
					goto tr78
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 58:
					switch {
					case data[p] < 41:
						if 34 <= data[p] && data[p] <= 38 {
							goto tr76
						}
					case data[p] > 47:
						if 48 <= data[p] && data[p] <= 57 {
							goto tr32
						}
					default:
						goto tr76
					}
				case data[p] > 63:
					switch {
					case data[p] < 91:
						if 65 <= data[p] && data[p] <= 90 {
							goto tr33
						}
					case data[p] > 94:
						if 97 <= data[p] && data[p] <= 122 {
							goto tr35
						}
					default:
						goto tr76
					}
				default:
					goto tr76
				}
				goto tr1
			case 25:
				switch data[p] {
				case 46:
					goto tr34
				case 92:
					goto tr34
				case 95:
					goto tr34
				}
				switch {
				case data[p] < 65:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr34
					}
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr34
					}
				default:
					goto tr34
				}
				goto tr6
			case 26:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 120:
					goto tr39
				case 123:
					goto tr6
				case 125:
					goto tr6
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr6
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr6
						}
					default:
						goto tr6
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr6
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr6
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 48:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 108:
					goto tr82
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 49:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 102:
					goto tr83
				case 110:
					goto tr84
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 50:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 110:
					goto tr85
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 51:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr86
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 52:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr87
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr87
					}
				default:
					goto tr35
				}
				goto tr1
			case 53:
				switch data[p] {
				case 33:
					goto tr2
				case 40:
					goto tr88
				case 46:
					goto tr87
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 41 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr87
					}
				default:
					goto tr87
				}
				goto tr1
			case 54:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 109:
					goto tr89
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 55:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr90
				case 92:
					goto tr34
				case 95:
					goto tr35
				case 123:
					goto tr76
				case 125:
					goto tr76
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr76
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr76
						}
					default:
						goto tr76
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr76
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr76
						}
					default:
						goto tr35
					}
				default:
					goto tr35
				}
				goto tr1
			case 56:
				switch data[p] {
				case 33:
					goto tr2
				case 46:
					goto tr35
				case 92:
					goto tr34
				case 95:
					goto tr91
				case 123:
					goto tr0
				case 125:
					goto tr0
				}
				switch {
				case data[p] < 48:
					switch {
					case data[p] < 37:
						if 34 <= data[p] && data[p] <= 35 {
							goto tr0
						}
					case data[p] > 38:
						if 40 <= data[p] && data[p] <= 47 {
							goto tr0
						}
					default:
						goto tr0
					}
				case data[p] > 57:
					switch {
					case data[p] < 65:
						if 58 <= data[p] && data[p] <= 63 {
							goto tr0
						}
					case data[p] > 90:
						switch {
						case data[p] > 94:
							if 97 <= data[p] && data[p] <= 122 {
								goto tr35
							}
						case data[p] >= 91:
							goto tr0
						}
					default:
						goto tr91
					}
				default:
					goto tr35
				}
				goto tr1
			}

		tr6:
			cs = 0
			goto _again
		tr1:
			cs = 1
			goto _again
		tr3:
			cs = 2
			goto _again
		tr40:
			cs = 3
			goto _again
		tr5:
			cs = 4
			goto _again
		tr8:
			cs = 5
			goto _again
		tr9:
			cs = 6
			goto _again
		tr11:
			cs = 7
			goto _again
		tr12:
			cs = 8
			goto _again
		tr7:
			cs = 9
			goto _again
		tr13:
			cs = 10
			goto _again
		tr14:
			cs = 11
			goto _again
		tr41:
			cs = 12
			goto _again
		tr16:
			cs = 13
			goto _again
		tr18:
			cs = 14
			goto _again
		tr19:
			cs = 15
			goto _again
		tr17:
			cs = 16
			goto _again
		tr21:
			cs = 17
			goto _again
		tr23:
			cs = 18
			goto _again
		tr25:
			cs = 19
			goto _again
		tr26:
			cs = 20
			goto _again
		tr70:
			cs = 21
			goto _again
		tr56:
			cs = 22
			goto _again
		tr57:
			cs = 23
			goto _again
		tr58:
			cs = 24
			goto _again
		tr59:
			cs = 25
			goto _again
		tr61:
			cs = 26
			goto _again
		tr0:
			cs = 27
			goto f0
		tr2:
			cs = 27
			goto f1
		tr10:
			cs = 27
			goto f3
		tr15:
			cs = 27
			goto f4
		tr27:
			cs = 27
			goto f7
		tr28:
			cs = 27
			goto f8
		tr30:
			cs = 27
			goto f10
		tr42:
			cs = 27
			goto f14
		tr43:
			cs = 27
			goto f15
		tr44:
			cs = 27
			goto f16
		tr45:
			cs = 27
			goto f17
		tr46:
			cs = 27
			goto f18
		tr47:
			cs = 27
			goto f19
		tr48:
			cs = 27
			goto f20
		tr49:
			cs = 27
			goto f21
		tr51:
			cs = 27
			goto f22
		tr52:
			cs = 27
			goto f23
		tr54:
			cs = 27
			goto f24
		tr60:
			cs = 27
			goto f25
		tr62:
			cs = 27
			goto f26
		tr63:
			cs = 27
			goto f27
		tr64:
			cs = 27
			goto f28
		tr65:
			cs = 27
			goto f29
		tr66:
			cs = 27
			goto f30
		tr67:
			cs = 27
			goto f31
		tr68:
			cs = 27
			goto f32
		tr71:
			cs = 27
			goto f33
		tr72:
			cs = 27
			goto f34
		tr73:
			cs = 27
			goto f35
		tr74:
			cs = 27
			goto f36
		tr75:
			cs = 27
			goto f37
		tr76:
			cs = 27
			goto f38
		tr88:
			cs = 27
			goto f41
		tr4:
			cs = 28
			goto f2
		tr20:
			cs = 29
			goto _again
		tr22:
			cs = 30
			goto f5
		tr24:
			cs = 31
			goto f6
		tr50:
			cs = 32
			goto f9
		tr69:
			cs = 33
			goto f9
		tr29:
			cs = 34
			goto f9
		tr53:
			cs = 35
			goto _again
		tr55:
			cs = 36
			goto _again
		tr31:
			cs = 37
			goto f11
		tr34:
			cs = 38
			goto _again
		tr35:
			cs = 39
			goto f11
		tr32:
			cs = 40
			goto f5
		tr33:
			cs = 41
			goto f11
		tr79:
			cs = 41
			goto f39
		tr81:
			cs = 41
			goto f40
		tr36:
			cs = 42
			goto f11
		tr77:
			cs = 43
			goto f11
		tr78:
			cs = 44
			goto f11
		tr37:
			cs = 45
			goto f11
		tr80:
			cs = 46
			goto f11
		tr38:
			cs = 47
			goto f11
		tr39:
			cs = 48
			goto f11
		tr82:
			cs = 49
			goto f11
		tr83:
			cs = 50
			goto f11
		tr85:
			cs = 51
			goto f11
		tr86:
			cs = 52
			goto f11
		tr87:
			cs = 53
			goto f11
		tr84:
			cs = 54
			goto f11
		tr89:
			cs = 55
			goto f11
		tr90:
			cs = 56
			goto f11
		tr91:
			cs = 56
			goto f42

		f3:
			_acts = 3
			goto execFuncs
		f4:
			_acts = 5
			goto execFuncs
		f1:
			_acts = 7
			goto execFuncs
		f7:
			_acts = 9
			goto execFuncs
		f10:
			_acts = 11
			goto execFuncs
		f41:
			_acts = 13
			goto execFuncs
		f14:
			_acts = 15
			goto execFuncs
		f26:
			_acts = 17
			goto execFuncs
		f27:
			_acts = 19
			goto execFuncs
		f15:
			_acts = 21
			goto execFuncs
		f16:
			_acts = 23
			goto execFuncs
		f18:
			_acts = 25
			goto execFuncs
		f20:
			_acts = 27
			goto execFuncs
		f17:
			_acts = 29
			goto execFuncs
		f21:
			_acts = 31
			goto execFuncs
		f25:
			_acts = 33
			goto execFuncs
		f24:
			_acts = 35
			goto execFuncs
		f34:
			_acts = 37
			goto execFuncs
		f37:
			_acts = 39
			goto execFuncs
		f35:
			_acts = 41
			goto execFuncs
		f22:
			_acts = 43
			goto execFuncs
		f23:
			_acts = 45
			goto execFuncs
		f19:
			_acts = 47
			goto execFuncs
		f32:
			_acts = 49
			goto execFuncs
		f30:
			_acts = 51
			goto execFuncs
		f29:
			_acts = 53
			goto execFuncs
		f31:
			_acts = 55
			goto execFuncs
		f38:
			_acts = 57
			goto execFuncs
		f28:
			_acts = 59
			goto execFuncs
		f33:
			_acts = 61
			goto execFuncs
		f36:
			_acts = 63
			goto execFuncs
		f8:
			_acts = 65
			goto execFuncs
		f0:
			_acts = 67
			goto execFuncs
		f39:
			_acts = 72
			goto execFuncs
		f9:
			_acts = 75
			goto execFuncs
		f5:
			_acts = 78
			goto execFuncs
		f40:
			_acts = 81
			goto execFuncs
		f6:
			_acts = 84
			goto execFuncs
		f42:
			_acts = 87
			goto execFuncs
		f11:
			_acts = 90
			goto execFuncs
		f2:
			_acts = 93
			goto execFuncs

		execFuncs:
			_nacts = uint(_formula_actions[_acts])
			_acts++
			for ; _nacts > 0; _nacts-- {
				_acts++
				switch _formula_actions[_acts-1] {
				case 3:
//line NONE:1
					te = p + 1

				case 4:
//line lexer.rl:55
					act = 1
				case 5:
//line lexer.rl:56
					act = 2
				case 6:
//line lexer.rl:57
					act = 3
				case 7:
//line lexer.rl:58
					act = 4
				case 8:
//line lexer.rl:62
					act = 8
				case 9:
//line lexer.rl:64
					act = 10
				case 10:
//line lexer.rl:70
					act = 13
				case 11:
//line lexer.rl:72
					act = 14
				case 12:
//line lexer.rl:59
					te = p + 1
					{
						l.emit(tokenError, data[ts:te])
					}
				case 13:
//line lexer.rl:60
					te = p + 1
					{
						l.emit(tokenErrorRef, data[ts:te])
					}
				case 14:
//line lexer.rl:62
					te = p + 1
					{
						l.emit(tokenSheet, data[ts:te-1])
					}
				case 15:
//line lexer.rl:63
					te = p + 1
					{
						l.emit(tokenSheet, data[ts+1:te-2])
					}
				case 16:
//line lexer.rl:67
					te = p + 1
					{
						l.emit(tokenFunctionBuiltin, data[ts:te-1])
					}
				case 17:
//line lexer.rl:68
					te = p + 1
					{
						l.emit(tokenFunctionBuiltin, data[ts:te-1])
					}
				case 18:
//line lexer.rl:75
					te = p + 1
					{
						l.emit(tokenAmpersand, data[ts:te])
					}
				case 19:
//line lexer.rl:76
					te = p + 1
					{
						l.emit(tokenLBrace, data[ts:te])
					}
				case 20:
//line lexer.rl:77
					te = p + 1
					{
						l.emit(tokenRBrace, data[ts:te])
					}
				case 21:
//line lexer.rl:78
					te = p + 1
					{
						l.emit(tokenLParen, data[ts:te])
					}
				case 22:
//line lexer.rl:79
					te = p + 1
					{
						l.emit(tokenRParen, data[ts:te])
					}
				case 23:
//line lexer.rl:80
					te = p + 1
					{
						l.emit(tokenPlus, data[ts:te])
					}
				case 24:
//line lexer.rl:81
					te = p + 1
					{
						l.emit(tokenMinus, data[ts:te])
					}
				case 25:
//line lexer.rl:82
					te = p + 1
					{
						l.emit(tokenMult, data[ts:te])
					}
				case 26:
//line lexer.rl:83
					te = p + 1
					{
						l.emit(tokenDiv, data[ts:te])
					}
				case 27:
//line lexer.rl:84
					te = p + 1
					{
						l.emit(tokenExp, data[ts:te])
					}
				case 28:
//line lexer.rl:87
					te = p + 1
					{
						l.emit(tokenEQ, data[ts:te])
					}
				case 29:
//line lexer.rl:88
					te = p + 1
					{
						l.emit(tokenLEQ, data[ts:te])
					}
				case 30:
//line lexer.rl:89
					te = p + 1
					{
						l.emit(tokenGEQ, data[ts:te])
					}
				case 31:
//line lexer.rl:90
					te = p + 1
					{
						l.emit(tokenNE, data[ts:te])
					}
				case 32:
//line lexer.rl:92
					te = p + 1
					{
						l.emit(tokenColon, data[ts:te])
					}
				case 33:
//line lexer.rl:93
					te = p + 1
					{
						l.emit(tokenSemi, data[ts:te])
					}
				case 34:
//line lexer.rl:94
					te = p + 1
					{
						l.emit(tokenComma, data[ts:te])
					}
				case 35:
//line lexer.rl:56
					te = p
					p--
					{
						l.emit(tokenNumber, data[ts:te])
					}
				case 36:
//line lexer.rl:57
					te = p
					p--
					{
						l.emit(tokenCell, data[ts:te])
					}
				case 37:
//line lexer.rl:61
					te = p
					p--
					{
						l.emit(tokenHorizontalRange, data[ts:te])
					}
				case 38:
//line lexer.rl:62
					te = p
					p--
					{
						l.emit(tokenSheet, data[ts:te-1])
					}
				case 39:
//line lexer.rl:70
					te = p
					p--
					{
						l.emit(tokenNamedRange, data[ts:te])
					}
				case 40:
//line lexer.rl:72
					te = p
					p--
					{
						l.emit(tokenString, data[ts+1:te-1])
					}
				case 41:
//line lexer.rl:85
					te = p
					p--
					{
						l.emit(tokenLT, data[ts:te])
					}
				case 42:
//line lexer.rl:86
					te = p
					p--
					{
						l.emit(tokenGT, data[ts:te])
					}
				case 43:
//line lexer.rl:56
					p = (te) - 1
					{
						l.emit(tokenNumber, data[ts:te])
					}
				case 44:
//line NONE:1
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
							l.emit(tokenSheet, data[ts:te-1])
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

//line lexer.go:2352
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
//line NONE:1
					ts = 0

				case 1:
//line NONE:1
					act = 0

//line lexer.go:2371
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
				case 1:
					goto tr0
				case 2:
					goto tr0
				case 28:
					goto tr64
				case 14:
					goto tr0
				case 15:
					goto tr0
				case 29:
					goto tr65
				case 17:
					goto tr0
				case 30:
					goto tr66
				case 31:
					goto tr67
				case 19:
					goto tr0
				case 20:
					goto tr0
				case 32:
					goto tr68
				case 33:
					goto tr68
				case 21:
					goto tr28
				case 34:
					goto tr68
				case 35:
					goto tr71
				case 36:
					goto tr74
				case 37:
					goto tr76
				case 38:
					goto tr76
				case 39:
					goto tr76
				case 40:
					goto tr66
				case 41:
					goto tr0
				case 42:
					goto tr76
				case 43:
					goto tr76
				case 44:
					goto tr76
				case 45:
					goto tr76
				case 46:
					goto tr76
				case 47:
					goto tr76
				case 48:
					goto tr76
				case 49:
					goto tr76
				case 50:
					goto tr76
				case 51:
					goto tr76
				case 52:
					goto tr76
				case 53:
					goto tr76
				case 54:
					goto tr76
				case 55:
					goto tr76
				case 56:
					goto tr0
				}
			}

		_out:
			{
			}
		}

//line lexer.rl:143

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
