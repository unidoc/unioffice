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


%%{
  machine format;
  write data;
  
  dquote = '"';
  not_dquote = [^"];
  AMPM = "AM/PM" | "A/P";
  INTLColor = "Black" | "Blue" | "Cyan" | "Green" | "Magenta" | "Red" | "White" | "Yellow";
  INTLCHARDECIMALSEP = '.';
  INTLCHARNUMGRPSEP = ',';
  INTLCHARDATESEP = '/';
  INTLCHARTIMESEP = ':';
  NFPartStrColor = "Color";
  NFPartNumber1To9 = [1-9];
  NFPartNumber1To6 = [1-6];
  NFPartNumber1To4 = [1-4];
  NFPart1To56 = (NFPartNumber1To9 | (NFPartNumber1To4 [0-9]) | '5' ('0' | [1-6]));
  NFPartColor = '[' (INTLColor | NFPartStrColor NFPart1To56) ']';
  NFPartIntNum = digit+;
  NFPartNumToken1 = '#' | '?' | '0';
  NFPartNumToken2 = NFPartNumToken1 | '.'  | ',';
  NFPartSign = '+' | '-';
  NFPartMonth = 'm'{1,5}; 
  NFPartDay = 'd'{1,4};
  NFPartHour = 'h'{1,2};
  NFPartSecond = 's'{1,2};
  NFPartSecondFrac = 's'{1,2} ('.' '0'{1,3})?;
  NFPartMinute = 'm'{1,2};
  NFPartYear = 'yy' | 'yyyy';
  NFPartAbsHour = '[h]';
  NFPartAbsMinute = '[m]'; 
  NFPartAbsSecond = '[s]';
  NFPartExponential = 'E' NFPartSign;
  NFGeneral = 'General';
  NFPartCompOper = '<' [=>]? | '=' | '>' '='?;
  NFPartCondNum = '-'? NFPartIntNum ('.' NFPartIntNum)? (NFPartExponential NFPartIntNum)?;
  NFPartCond = '[' NFPartCompOper NFPartCondNum ']';
  NFPartLocaleID = '[' '$' any+ ('-' xdigit{3,8})? ']';
  NFPartNum = NFPartNumToken2+ (NFPartNumToken2 | '%')* | 
  (NFPartNumToken2 | '%') NFPartNumToken2+;
  NFNumber = NFPartNum (NFPartExponential NFPartNum)? ','* AMPM*;
  NFPartFraction = (NFPartIntNum+ (NFPartIntNum | '%')) |
    ((NFPartIntNum | '%')* NFPartIntNum+) | 
    (NFPartNumToken1+ (NFPartNumToken1 | '%')*) | 
    ((NFPartNumToken1 | '%') NFPartNumToken1+);
  NFFraction = NFPartFraction '/' NFPartFraction NFPartNum? AMPM*;
  NFText = '@'+ ('@' | AMPM)* | ('@' | AMPM)* '@'*;
  NFAbsTimeToken = NFPartAbsHour | NFPartAbsSecond | NFPartAbsMinute;
  
  NFDateToken = NFPartYear | NFPartMonth | NFPartDay ;
  NFTimeToken = NFPartHour | NFPartMinute | NFPartSecond | NFPartSecondFrac | AMPM;
  NFDate = (NFDateToken  | '/' )+;
  NFTime = (NFTimeToken | ':')+;
  cond = '[' any+ ']';
main := |*
  '0' => { l.fmt.AddPlaceholder(FmtTypeDigit,nil) };
  '#' => { l.fmt.AddPlaceholder(FmtTypeDigitOpt,nil) };
  '?' => { }; # ignore for now
  '.' => { l.fmt.AddPlaceholder(FmtTypeDecimal,nil) };
  ',' => { l.fmt.AddPlaceholder(FmtTypeComma,nil) };
  '%' => { l.fmt.AddPlaceholder(FmtTypePercent,nil) };
  '$' => { l.fmt.AddPlaceholder(FmtTypeDollar,nil) };
  '_' => { l.fmt.AddPlaceholder(FmtTypeUnderscore,nil) };
  ';' => { l.nextFmt() };
  NFGeneral => { l.fmt.isGeneral = true };
  #NFFraction => {fmt.Println("FRACTION",string(data[ts:te]))}; 
  # we have to keep date/time separate as 'mm' is both minutes and month
  NFDate => { l.fmt.AddPlaceholder(FmtTypeDate,data[ts:te]) };
  NFTime => { l.fmt.AddPlaceholder(FmtTypeTime,data[ts:te]) };
  NFAbsTimeToken => { l.fmt.AddPlaceholder(FmtTypeTime,data[ts:te]) };
  NFPartExponential => { l.fmt.IsExponential = true };
  cond => {}; # ignoring 
  # escaped
  '\\' any => { l.fmt.AddPlaceholder(FmtTypeLiteral,data[ts+1:te]) };
  any => { l.fmt.AddPlaceholder(FmtTypeLiteral,data[ts:te]) };
  dquote ( not_dquote | dquote dquote)* dquote => { l.fmt.AddPlaceholder(FmtTypeLiteral,data[ts+1:te-1])};

*|;

}%%
 func(l *Lexer) Lex(r io.Reader)  {
  cs, p, pe := 0, 0, 0
  eof := -1
  ts, te,act := 0,0,0
  _ = te
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
  if cs == format_error {
     //l.emit(tokenLexError,nil)
     log.Panic("ERROR LEXING")
  }
}
