//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package algo ;import _d "strconv";func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_eeb :=make ([]byte ,len (s )*cnt );_ab :=[]byte (s );for _dd :=0;_dd < cnt ;_dd ++{copy (_eeb [_dd :],_ab );};return string (_eeb );};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_ee ,_b :=0,0;for _ee < len (lhs )&&_b < len (rhs ){_ec :=lhs [_ee ];_f :=rhs [_b ];_ba :=_df (_ec );_ga :=_df (_f );switch {case _ba &&!_ga :return true ;case !_ba &&_ga :return false ;case !_ba &&!_ga :if _ec !=_f {return _ec < _f ;
};_ee ++;_b ++;default:_c :=_ee +1;_eec :=_b +1;for _c < len (lhs )&&_df (lhs [_c ]){_c ++;};for _eec < len (rhs )&&_df (rhs [_eec ]){_eec ++;};_fc ,_ :=_d .ParseUint (lhs [_ee :_c ],10,64);_ge ,_ :=_d .ParseUint (rhs [_ee :_eec ],10,64);if _fc !=_ge {return _fc < _ge ;
};_ee =_c ;_b =_eec ;};};return len (lhs )< len (rhs );};func _df (_e byte )bool {return _e >='0'&&_e <='9'};