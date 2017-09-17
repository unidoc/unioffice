// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import __yyfmt__ "fmt"

type yySymType struct {
	yys  int
	node *node
	expr Expression
	args []Expression
	rows [][]Expression
}

const tokenHorizontalRange = 57346
const tokenReservedName = 57347
const tokenDDECall = 57348
const tokenLexError = 57349
const tokenNamedRange = 57350
const tokenBool = 57351
const tokenNumber = 57352
const tokenString = 57353
const tokenError = 57354
const tokenErrorRef = 57355
const tokenSheet = 57356
const tokenCell = 57357
const tokenFunctionBuiltin = 57358
const tokenLBrace = 57359
const tokenRBrace = 57360
const tokenLParen = 57361
const tokenRParen = 57362
const tokenPlus = 57363
const tokenMinus = 57364
const tokenMult = 57365
const tokenDiv = 57366
const tokenExp = 57367
const tokenEQ = 57368
const tokenLT = 57369
const tokenGT = 57370
const tokenLEQ = 57371
const tokenGEQ = 57372
const tokenNE = 57373
const tokenColon = 57374
const tokenComma = 57375
const tokenAmpersand = 57376
const tokenSemi = 57377

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokenHorizontalRange",
	"tokenReservedName",
	"tokenDDECall",
	"tokenLexError",
	"tokenNamedRange",
	"tokenBool",
	"tokenNumber",
	"tokenString",
	"tokenError",
	"tokenErrorRef",
	"tokenSheet",
	"tokenCell",
	"tokenFunctionBuiltin",
	"tokenLBrace",
	"tokenRBrace",
	"tokenLParen",
	"tokenRParen",
	"tokenPlus",
	"tokenMinus",
	"tokenMult",
	"tokenDiv",
	"tokenExp",
	"tokenEQ",
	"tokenLT",
	"tokenGT",
	"tokenLEQ",
	"tokenGEQ",
	"tokenNE",
	"tokenColon",
	"tokenComma",
	"tokenAmpersand",
	"tokenSemi",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 179

var yyAct = [...]int{

	43, 3, 42, 30, 18, 38, 68, 44, 45, 28,
	29, 30, 37, 46, 47, 30, 50, 41, 24, 13,
	37, 19, 66, 52, 48, 23, 21, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 67,
	51, 65, 76, 11, 9, 1, 20, 10, 2, 8,
	73, 71, 70, 26, 27, 28, 29, 30, 35, 31,
	32, 33, 34, 36, 72, 0, 37, 0, 0, 75,
	74, 0, 0, 77, 69, 26, 27, 28, 29, 30,
	35, 31, 32, 33, 34, 36, 0, 0, 37, 26,
	27, 28, 29, 30, 35, 31, 32, 33, 34, 36,
	0, 0, 37, 26, 27, 28, 29, 30, 0, 0,
	0, 24, 14, 15, 16, 17, 37, 25, 23, 22,
	39, 0, 12, 0, 6, 7, 0, 0, 0, 40,
	24, 14, 15, 16, 17, 0, 25, 23, 22, 5,
	0, 12, 0, 6, 7, 0, 0, 0, 4, 24,
	14, 15, 16, 17, 0, 25, 23, 22, 39, 0,
	12, 49, 6, 7, 24, 14, 15, 16, 17, 0,
	25, 23, 22, 39, 0, 12, 0, 6, 7,
}
var yyPact = [...]int{

	122, -1000, -1000, 68, 156, 103, 156, 156, -1000, -1000,
	-1000, -1000, 156, -1000, -1000, -1000, -1000, -1000, -18, 10,
	-1000, -1000, 141, -1000, -1000, -1000, 156, 156, 156, 156,
	156, 156, 156, 156, 156, 156, 156, 156, 68, 156,
	156, 4, -27, 68, -14, -14, 54, 10, -1000, -1000,
	31, -1000, 68, -14, -14, -22, -22, -1000, 82, 82,
	82, 82, 82, 82, -10, 32, -1000, 156, 156, -1000,
	-1000, -1000, 156, -1000, -27, 68, -1000, 68,
}
var yyPgo = [...]int{

	0, 0, 49, 48, 47, 4, 46, 45, 44, 43,
	42, 40, 26, 21, 19, 17, 16, 2,
}
var yyR1 = [...]int{

	0, 7, 3, 3, 3, 8, 8, 8, 8, 1,
	1, 1, 2, 2, 2, 2, 2, 14, 15, 15,
	17, 17, 4, 4, 4, 13, 5, 5, 6, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 9, 9, 9, 16, 16, 11, 10, 10,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 4, 1, 1, 1, 1, 2,
	2, 1, 1, 1, 1, 3, 1, 3, 1, 3,
	1, 3, 1, 2, 1, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 2, 3, 1, 3, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -7, -3, -1, 26, 17, 21, 22, -2, -8,
	-4, -9, 19, -14, 9, 10, 11, 12, -5, -13,
	-6, -12, 16, 15, 8, 14, 21, 22, 23, 24,
	25, 27, 28, 29, 30, 26, 31, 34, -1, 17,
	26, -15, -17, -1, -1, -1, -1, 32, -5, 20,
	-16, -11, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, 18, 35, 33, 20,
	-5, 20, 33, 18, -17, -1, -10, -1,
}
var yyDef = [...]int{

	0, -2, 1, 2, 0, 0, 0, 0, 11, 12,
	13, 14, 0, 16, 5, 6, 7, 8, 22, 0,
	24, 41, 0, 26, 27, 25, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 18, 20, 9, 10, 0, 0, 23, 42,
	0, 44, 46, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 0, 17, 0, 0, 15,
	28, 43, 48, 4, 19, 21, 45, 47,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yylex.(*plex).result = yyVAL.expr
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = NewBool(yyDollar[1].node.val)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = NewNumber(yyDollar[1].node.val)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = NewString(yyDollar[1].node.val)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = NewError(yyDollar[1].node.val)
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = NewNegate(yyDollar[2].expr)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewConstArrayExpr(yyDollar[2].rows)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = append(yyVAL.rows, yyDollar[1].args)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].args)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.args = append(yyVAL.args, yyDollar[1].expr)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.args = append(yyDollar[1].args, yyDollar[3].expr)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = NewPrefixExpr(yyDollar[1].expr, yyDollar[2].expr)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = NewSheetPrefixExpr(yyDollar[1].node.val)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = NewCellRef(yyDollar[1].node.val)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = NewNamedRangeRef(yyDollar[1].node.val)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewRange(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypePlus, yyDollar[3].expr)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeMinus, yyDollar[3].expr)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeMult, yyDollar[3].expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeDiv, yyDollar[3].expr)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeExp, yyDollar[3].expr)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeLT, yyDollar[3].expr)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeGT, yyDollar[3].expr)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeLEQ, yyDollar[3].expr)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeGEQ, yyDollar[3].expr)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeEQ, yyDollar[3].expr)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeNE, yyDollar[3].expr)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewBinaryExpr(yyDollar[1].expr, BinOpTypeConcat, yyDollar[3].expr)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = NewFunction(yyDollar[1].node.val, nil)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = NewFunction(yyDollar[1].node.val, yyDollar[2].args)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.args = append(yyVAL.args, yyDollar[1].expr)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.args = append(yyDollar[1].args, yyDollar[3].expr)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = NewEmptyExpr()
		}
	}
	goto yystack /* stack new state and value */
}
