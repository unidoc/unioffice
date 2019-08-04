// Copyright 2017 FoxyUtils ehf. All rights reserved.
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


%%{
  machine formula;
  write data;

  s = [ \t];

  squote = "'";
  dquote = '"';
  not_dquote = [^"];
  

  bool = 'TRUE' | 'FALSE';
  cell = '$'? [A-Z]+ '$'? [0-9]+;
  ddeCall = 'TODO';
  errorLiteral = '#NULL!' | '#NUM!' | '#N/A' ;
  errorRef = '#REF!';
  fnName = 'TODO';
  file = 'TODO';
  horizontalRange = '$'? [0-9]+ ':' '$'? [0-9]+;

  # there is a function list at https://msdn.microsoft.com/en-us/library/dd906358(v=office.12).aspx
  builtinFunction = [A-Z] [A-Z0-9.]*  '(';
  excelFn = '_xlfn.' [A-Z_] [A-Z0-9.]* '(';  

  sheetChar = ^['%\[\]\\:/?();{}#"=<>&+\-*/^%,_!];
  enclosedSheetChar = ^['*\[\]\\:\/?];

  number = [0-9]+ '.'? [0-9]* ('e' [0-9]+)?;
  sheet = sheetChar+ '!';
  quotedSheet = (sheetChar+ | squote (enclosedSheetChar  | dquote)+ squote) '!';

  namedRange = [A-Z_\\][A-Za-z0-9\\_.]+;

  reservedName = '_xlnm.' [A-Z_]+;
  
  
  
  main := |*

  bool => { l.emit(tokenBool, data[ts:te])}; 
  number =>  {l.emit(tokenNumber, data[ts:te])}; 
  cell => { l.emit(tokenCell, data[ts:te])}; 
  ddeCall => {l.emit(tokenDDECall, data[ts:te])}; 
  errorLiteral => {l.emit(tokenError, data[ts:te])}; 
  errorRef => {l.emit(tokenErrorRef, data[ts:te])}; 
  horizontalRange => {l.emit(tokenHorizontalRange, data[ts:te])}; 
  sheet  =>  {l.emit(tokenSheet, data[ts:te-1])};  # chop '!'
  quotedSheet  =>  {l.emit(tokenSheet, data[ts+1:te-2])};  # chop leading quote and trailing quote & !
  reservedName => {l.emit(tokenReservedName, data[ts:te])};

  
  builtinFunction => {l.emit(tokenFunctionBuiltin, data[ts:te-1])}; # chop off the final '(' so we only pass the name
  excelFn => {l.emit(tokenFunctionBuiltin, data[ts:te-1])}; # chop off the final '(' so we only pass the name
  
  namedRange  => {l.emit(tokenNamedRange, data[ts:te])};

  dquote ( not_dquote | dquote dquote)* dquote => { l.emit(tokenString, data[ts+1:te-1])}; # chop off delimiters

  
  '&' { l.emit(tokenAmpersand,data[ts:te]) };
  '{' { l.emit(tokenLBrace,data[ts:te]) };
  '}' { l.emit(tokenRBrace,data[ts:te]) };
  '(' { l.emit(tokenLParen,data[ts:te]) };
  ')' { l.emit(tokenRParen,data[ts:te]) };
  '+' { l.emit(tokenPlus,data[ts:te]) };
  '-' { l.emit(tokenMinus,data[ts:te]) };
  '*' { l.emit(tokenMult,data[ts:te]) };
  '/' { l.emit(tokenDiv,data[ts:te]) };
  '^' { l.emit(tokenExp,data[ts:te]) };
  '<' { l.emit(tokenLT,data[ts:te]) };
  '>' { l.emit(tokenGT,data[ts:te]) };
  '=' { l.emit(tokenEQ,data[ts:te]) };
  '<=' { l.emit(tokenLEQ,data[ts:te]) };
  '>=' { l.emit(tokenGEQ,data[ts:te]) };
  '<>' { l.emit(tokenNE,data[ts:te]) };
  
  ':' { l.emit(tokenColon,data[ts:te]) };
  ';' { l.emit(tokenSemi,data[ts:te]) };
  ',' { l.emit(tokenComma,data[ts:te]) };
  
*|;


prepush {
  stack = append(stack,0)
}

postpop {
  stack = stack[0:len(stack)-1]
}

}%%
 func(l *Lexer) lex(r io.Reader)  {
  cs, p, pe := 0, 0, 0
  eof := -1
  ts, te,act := 0,0,0
  _ = act
  curline := 1
  _ = curline
  data := make([]byte,4096)

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
   pe = n+rem
   if pe < len(data) {
      eof = pe
   }

  %%{
    write init;
    write exec;
  }%%
  
  if ts > 0 {
      // currently parsing a token, so shift it to the
      // beginning of the buffer
      copy(data[0:],data[ts:])
    } 
  }
  
  _ = eof
  if cs == formula_error {
     l.emit(tokenLexError,nil)
  }
  close(l.nodes)
}
