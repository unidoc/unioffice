// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.


package format
import (
)


%%{
  machine isnumber;
  write data;
  
  sign  = '+' | '-';
  integer = sign? [0-9]+;
  float   = sign? [0-9]+ '.' [0-9]+ ('E' sign [0-9]+)?;

  main := |*
  integer => { isNumber = te == len(data); };
  float => { isNumber = te == len(data); };
  any* => { isNumber = false; };
  *|;
}%%
func IsNumber(data string) (isNumber bool) {
  cs, p, pe := 0, 0, len(data)
  eof := len(data)
  ts, te,act := 0,0,0
  _ = te
  _ = act
  _ = ts
  
  
  %%{
    write init;
    write exec;
  }%%
   if cs == format_error {
     return false
   }
   return
}
