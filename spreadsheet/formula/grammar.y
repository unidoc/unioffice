%{
// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

%}

%union {
	node *node
	expr Expression
	args []Expression
	rows [][]Expression
}

%type <expr> formula formula1 initial reference referenceItem refFunctionCall
%type <expr> start constant functionCall argument argument1
%type <expr> binOp prefix constArray
%type <rows> constArrayRows
%type <args> arguments  constArrayCols

%token <node> tokenHorizontalRange tokenReservedName tokenDDECall  tokenLexError tokenNamedRange
%token <node> tokenBool tokenNumber tokenString tokenError tokenErrorRef  tokenSheet tokenCell
%token <node> tokenFunctionBuiltin 

%token tokenLBrace tokenRBrace tokenLParen tokenRParen
%token tokenPlus tokenMinus tokenMult tokenDiv tokenExp tokenEQ tokenLT tokenGT tokenLEQ tokenGEQ  tokenNE 
%token tokenColon tokenComma tokenAmpersand tokenSemi

%left tokenEQ tokenLT tokenGT tokenLEQ tokenGEQ  tokenNE
%left tokenPlus tokenMinus
%left tokenMult tokenDiv
%left tokenAmpersand
%left tokenExp

%%

start: initial { yylex.(*plex).result = $$ }

initial:
	  formula
	| tokenEQ formula { $$ = $2} 
	| tokenLBrace tokenEQ formula tokenRBrace {};

constant: 
	  tokenBool { $$ = NewBool($1.val) } 
	| tokenNumber { $$ = NewNumber($1.val) } 
	| tokenString { $$ = NewString($1.val) } 
	| tokenError { $$ = NewError($1.val) } 

formula: 
	  tokenPlus formula { $$ = $2; } 
	| tokenMinus formula { $$ = NewNegate($2) } 
	| formula1
	;

formula1: 
	  constant 
    | reference
	| functionCall 
	| tokenLParen formula tokenRParen { $$ = $2 }
	| constArray
	;

constArray: tokenLBrace constArrayRows tokenRBrace { $$ = NewConstArrayExpr($2)} ;
constArrayRows: 
	  constArrayCols { $$ = append($$, $1) }
	| constArrayRows tokenSemi constArrayCols { $$ = append($1, $3)};

constArrayCols: 
	  formula { $$ = append($$,$1) }
	| constArrayCols tokenComma formula { $$ = append($1,$3)}


reference: 
	  referenceItem
    | prefix referenceItem { $$ = NewPrefixExpr($1,$2)}
	| refFunctionCall;

prefix: tokenSheet { $$ = NewSheetPrefixExpr($1.val) };

referenceItem: 
	  tokenCell { $$ = NewCellRef($1.val)}
	| tokenNamedRange { $$ = NewNamedRangeRef($1.val)}
	;

refFunctionCall:
	  referenceItem tokenColon referenceItem { $$ = NewRange($1,$3) };

           
binOp: 
      formula tokenPlus formula { $$ = NewBinaryExpr($1,BinOpTypePlus,$3); }
	| formula tokenMinus formula { $$ = NewBinaryExpr($1,BinOpTypeMinus,$3); }
	| formula tokenMult formula { $$ = NewBinaryExpr($1,BinOpTypeMult,$3); }
	| formula tokenDiv formula { $$ = NewBinaryExpr($1,BinOpTypeDiv,$3); }
	| formula tokenExp formula { $$ = NewBinaryExpr($1,BinOpTypeExp,$3); }
	| formula tokenLT formula { $$ = NewBinaryExpr($1,BinOpTypeLT,$3); }
	| formula tokenGT formula { $$ = NewBinaryExpr($1,BinOpTypeGT,$3); }
	| formula tokenLEQ formula { $$ = NewBinaryExpr($1,BinOpTypeLEQ,$3); }
	| formula tokenGEQ formula { $$ = NewBinaryExpr($1,BinOpTypeGEQ,$3); }
	| formula tokenEQ formula { $$ = NewBinaryExpr($1,BinOpTypeEQ,$3); }
	| formula tokenNE formula { $$ = NewBinaryExpr($1,BinOpTypeNE,$3); }
	| formula tokenAmpersand formula { $$ = NewBinaryExpr($1,BinOpTypeConcat,$3); }
	;

functionCall:  
	  binOp;
	| tokenFunctionBuiltin tokenRParen { $$ = NewFunction($1.val,nil)} ;
	| tokenFunctionBuiltin arguments tokenRParen { $$ = NewFunction($1.val,$2)} ;

arguments: 
	  argument1{ $$ = append($$, $1)  }
	| arguments tokenComma argument { $$ = append($1,$3) }
	;

argument1: formula ;

argument: 
	  formula 
	| /*empty*/ { $$ = NewEmptyExpr() } ;

%%
