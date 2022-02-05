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

package formatutils ;import (_b "fmt";_da "github.com/unidoc/unioffice/schema/soo/wml";_d "strconv";_ab "strings";);func _ad (_gf string )(_aa []string ){for _gb :=0;_gb < len (_gf )-2;_gb ++{if string (_gf [_gb ])=="\u0025"{if !_ab .Contains (string (_gf [_gb +2:]),"\u0025"){_aa =append (_aa ,_b .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_gf [_gb ]),string (_gf [_gb +1]),string (_gf [_gb +2:])));}else {_aa =append (_aa ,_b .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_gf [_gb ]),string (_gf [_gb +1]),string (_gf [_gb +2])));};};};return ;};var (_gg =[]string {"","\u0049","\u0049\u0049","\u0049\u0049\u0049","\u0049\u0056","\u0056","\u0056\u0049","\u0056\u0049\u0049","\u0056\u0049\u0049\u0049","\u0049\u0058"};_gd =[]string {"","\u0058","\u0058\u0058","\u0058\u0058\u0058","\u0058\u004c","\u004c","\u004c\u0058","\u004c\u0058\u0058","\u004c\u0058\u0058\u0058","\u0058\u0043"};_f =[]string {"","\u0043","\u0043\u0043","\u0043\u0043\u0043","\u0043\u0044","\u0044","\u0044\u0043","\u0044\u0043\u0043","\u0044\u0043\u0043\u0043","\u0043\u004d","\u004d"};_ace =[]string {"","\u004d","\u004d\u004d","\u004d\u004d\u004d","\u004d\u004d\u004d\u004d","\u004d\u004d\u004dM\u004d","\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"};_dg =[]string {"\u006f\u006e\u0065","\u0074\u0077\u006f","\u0074\u0068\u0072e\u0065","\u0066\u006f\u0075\u0072","\u0066\u0069\u0076\u0065","\u0073\u0069\u0078","\u0073\u0065\u0076e\u006e","\u0065\u0069\u0067h\u0074","\u006e\u0069\u006e\u0065","\u0074\u0065\u006e","\u0065\u006c\u0065\u0076\u0065\u006e","\u0074\u0077\u0065\u006c\u0076\u0065","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e","\u0066i\u0066\u0074\u0065\u0065\u006e","\u0073i\u0078\u0074\u0065\u0065\u006e","\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"};_fb =[]string {"\u0074\u0065\u006e","\u0074\u0077\u0065\u006e\u0074\u0079","\u0074\u0068\u0069\u0072\u0074\u0079","\u0066\u006f\u0072t\u0079","\u0066\u0069\u0066t\u0079","\u0073\u0069\u0078t\u0079","\u0073e\u0076\u0065\u006e\u0074\u0079","\u0065\u0069\u0067\u0068\u0074\u0079","\u006e\u0069\u006e\u0065\u0074\u0079"};_fa =[]string {"\u0066\u0069\u0072s\u0074","\u0073\u0065\u0063\u006f\u006e\u0064","\u0074\u0068\u0069r\u0064","\u0066\u006f\u0075\u0072\u0074\u0068","\u0066\u0069\u0066t\u0068","\u0073\u0069\u0078t\u0068","\u0073e\u0076\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0068","\u006e\u0069\u006et\u0068","\u0074\u0065\u006et\u0068","\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068","\u0074w\u0065\u006c\u0066\u0074\u0068","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h","\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h","s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"};_dbe =[]string {"\u0074\u0065\u006et\u0068","\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h","\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h","\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068","\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068","\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068","\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068","\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h","\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"};_cd ="\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a";);func _gba (_ga int64 ,_e *_da .CT_NumFmt )(_eg string ){if _e ==nil {return ;};_aab :=_e .ValAttr ;switch _aab {case _da .ST_NumberFormatNone :_eg ="";case _da .ST_NumberFormatDecimal :_eg =_d .Itoa (int (_ga ));case _da .ST_NumberFormatDecimalZero :_eg =_d .Itoa (int (_ga ));if _ga < 10{_eg ="\u0030"+_eg ;};case _da .ST_NumberFormatUpperRoman :var (_fc =_ga %10;_egf =(_ga %100)/10;_df =(_ga %1000)/100;_ea =_ga /1000;);_eg =_ace [_ea ]+_f [_df ]+_gd [_egf ]+_gg [_fc ];case _da .ST_NumberFormatLowerRoman :var (_fd =_ga %10;_gbe =(_ga %100)/10;_gge =(_ga %1000)/100;_gbg =_ga /1000;);_eg =_ace [_gbg ]+_f [_gge ]+_gd [_gbe ]+_gg [_fd ];_eg =_ab .ToLower (_eg );case _da .ST_NumberFormatUpperLetter :_dgf :=_ga %780;if _dgf ==0{_dgf =780;};_dga :=(_dgf -1)/26;_eb :=(_dgf -1)%26;_ed :=_cd [_dga +_eb ];_eg =string (_ed );case _da .ST_NumberFormatLowerLetter :_caf :=_ga %780;if _caf ==0{_caf =780;};_de :=(_caf -1)/26;_bc :=(_caf -1)%26;_dfe :=_cd [_de +_bc ];_eg =_ab .ToLower (string (_dfe ));default:_eg ="";};return ;};func FormatNumberingText (currentNumber int64 ,ilvl int64 ,lvlText string ,numFmt *_da .CT_NumFmt ,levelNumbers map[int64 ]int64 )string {_dd :=_ad (lvlText );_c :=_gba (currentNumber ,numFmt );_ac :=int64 (0);for _g ,_ca :=range _dd {_ce :=_b .Sprintf ("\u0025\u0025\u0025\u0064",_g +1);if len (_dd )==1{_ce =_b .Sprintf ("\u0025\u0025\u0025\u0064",ilvl +1);_dd [_g ]=_ab .Replace (_ca ,_ce ,_c ,1);break ;};if ilvl > 0&&ilvl > _ac &&_g < (len (_dd )-1){_db :=_gba (levelNumbers [_ac ],numFmt );_dd [_g ]=_ab .Replace (_ca ,_ce ,_db ,1);_ac ++;}else {_dd [_g ]=_ab .Replace (_ca ,_ce ,_c ,1);};};return _ab .Join (_dd ,"");};