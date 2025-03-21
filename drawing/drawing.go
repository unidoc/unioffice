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

package drawing ;import (_c "github.com/unidoc/unioffice/v2";_cc "github.com/unidoc/unioffice/v2/color";_g "github.com/unidoc/unioffice/v2/measurement";_e "github.com/unidoc/unioffice/v2/schema/soo/dml";);

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_cce LineProperties )SetWidth (w _g .Distance ){_cce ._gb .WAttr =_c .Int32 (int32 (w /_g .EMU ))};func (_fd LineProperties )SetNoFill (){_fd .clearFill ();_fd ._gb .LineFillPropertiesChoice .NoFill =_e .NewCT_NoFillProperties ();};

// AddBreak adds a new line break to a paragraph.
func (_gf Paragraph )AddBreak (){_b :=_e .NewEG_TextRun ();_b .TextRunChoice .Br =_e .NewCT_TextLineBreak ();_gf ._ccf .EG_TextRun =append (_gf ._ccf .EG_TextRun ,_b );};

// X returns the inner wrapped XML type.
func (_ga Paragraph )X ()*_e .CT_TextParagraph {return _ga ._ccf };

// SetPosition sets the position of the shape.
func (_cb ShapeProperties )SetPosition (x ,y _g .Distance ){_cb .ensureXfrm ();if _cb ._afe .Xfrm .Off ==nil {_cb ._afe .Xfrm .Off =_e .NewCT_Point2D ();};_cb ._afe .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_c .Int64 (int64 (x /_g .EMU ));_cb ._afe .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_c .Int64 (int64 (y /_g .EMU ));
};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// SetNumbered controls if bullets are numbered or not.
func (_af ParagraphProperties )SetNumbered (scheme _e .ST_TextAutonumberScheme ){if scheme ==_e .ST_TextAutonumberSchemeUnset {_af ._eg .TextBulletChoice .BuAutoNum =nil ;}else {_af ._eg .TextBulletChoice .BuAutoNum =_e .NewCT_TextAutonumberBullet ();_af ._eg .TextBulletChoice .BuAutoNum .TypeAttr =scheme ;
};};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_eg *_e .CT_TextParagraphProperties ;};

// SetBulletChar sets the bullet character for the paragraph.
func (_gff ParagraphProperties )SetBulletChar (c string ){if c ==""{_gff ._eg .TextBulletChoice .BuChar =nil ;}else {_gff ._eg .TextBulletChoice .BuChar =_e .NewCT_TextCharBullet ();_gff ._eg .TextBulletChoice .BuChar .CharAttr =c ;};};

// AddRun adds a new run to a paragraph.
func (_cg Paragraph )AddRun ()Run {_gg :=MakeRun (_e .NewEG_TextRun ());_cg ._ccf .EG_TextRun =append (_cg ._ccf .EG_TextRun ,_gg .X ());return _gg ;};

// Paragraph is a paragraph within a document.
type Paragraph struct{_ccf *_e .CT_TextParagraph };

// SetSize sets the font size of the run text
func (_beb RunProperties )SetSize (sz _g .Distance ){_beb ._fe .SzAttr =_c .Int32 (int32 (sz /_g .HundredthPoint ));};func (_cd ShapeProperties )SetNoFill (){_cd .clearFill ();_cd ._afe .FillPropertiesChoice .NoFill =_e .NewCT_NoFillProperties ();};

// Properties returns the run's properties.
func (_bc Run )Properties ()RunProperties {if _bc ._bf .TextRunChoice .R ==nil {_bc ._bf .TextRunChoice .R =_e .NewCT_RegularTextRun ();};if _bc ._bf .TextRunChoice .R .RPr ==nil {_bc ._bf .TextRunChoice .R .RPr =_e .NewCT_TextCharacterProperties ();};
return RunProperties {_bc ._bf .TextRunChoice .R .RPr };};

// X returns the inner wrapped XML type.
func (_d LineProperties )X ()*_e .CT_LineProperties {return _d ._gb };type LineProperties struct{_gb *_e .CT_LineProperties };func (_bcc ShapeProperties )LineProperties ()LineProperties {if _bcc ._afe .Ln ==nil {_bcc ._afe .Ln =_e .NewCT_LineProperties ();
};return LineProperties {_bcc ._afe .Ln };};

// SetBulletFont controls the font for the bullet character.
func (_gge ParagraphProperties )SetBulletFont (f string ){if f ==""{_gge ._eg .TextBulletTypefaceChoice .BuFont =nil ;}else {_gge ._eg .TextBulletTypefaceChoice .BuFont =_e .NewCT_TextFont ();_gge ._eg .TextBulletTypefaceChoice .BuFont .TypefaceAttr =f ;
};};

// Run is a run within a paragraph.
type Run struct{_bf *_e .EG_TextRun };

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_e .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// SetSolidFill controls the text color of a run.
func (_ec RunProperties )SetSolidFill (c _cc .Color ){_ec ._fe .FillPropertiesChoice =_e .NewEG_FillPropertiesChoice ();_ec ._fe .FillPropertiesChoice .SolidFill =_e .NewCT_SolidColorFillProperties ();_ec ._fe .FillPropertiesChoice .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();
_ec ._fe .FillPropertiesChoice .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetJoin sets the line join style.
func (_fb LineProperties )SetJoin (e LineJoin ){_fb ._gb .LineJoinPropertiesChoice =_e .NewEG_LineJoinPropertiesChoice ();switch e {case LineJoinRound :_fb ._gb .LineJoinPropertiesChoice .Round =_e .NewCT_LineJoinRound ();case LineJoinBevel :_fb ._gb .LineJoinPropertiesChoice .Bevel =_e .NewCT_LineJoinBevel ();
case LineJoinMiter :_fb ._gb .LineJoinPropertiesChoice .Miter =_e .NewCT_LineJoinMiterProperties ();};};func (_dda ShapeProperties )ensureXfrm (){if _dda ._afe .Xfrm ==nil {_dda ._afe .Xfrm =_e .NewCT_Transform2D ();};};

// SetSize sets the width and height of the shape.
func (_ggb ShapeProperties )SetSize (w ,h _g .Distance ){_ggb .SetWidth (w );_ggb .SetHeight (h )};func (_ge LineProperties )SetSolidFill (c _cc .Color ){_ge .clearFill ();_ge ._gb .LineFillPropertiesChoice .SolidFill =_e .NewCT_SolidColorFillProperties ();
_ge ._gb .LineFillPropertiesChoice .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();_ge ._gb .LineFillPropertiesChoice .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_e .EG_TextRun )Run {return Run {x }};

// SetAlign controls the paragraph alignment
func (_dg ParagraphProperties )SetAlign (a _e .ST_TextAlignType ){_dg ._eg .AlgnAttr =a };

// X returns the inner wrapped XML type.
func (_bg Run )X ()*_e .EG_TextRun {return _bg ._bf };func (_a LineProperties )clearFill (){_a ._gb .LineFillPropertiesChoice =_e .NewEG_LineFillPropertiesChoice ();};

// SetBold controls the bolding of a run.
func (_dd RunProperties )SetBold (b bool ){_dd ._fe .BAttr =_c .Bool (b )};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_bgb ShapeProperties )SetFlipHorizontal (b bool ){_bgb .ensureXfrm ();if !b {_bgb ._afe .Xfrm .FlipHAttr =nil ;}else {_bgb ._afe .Xfrm .FlipHAttr =_c .Bool (true );};};type ShapeProperties struct{_afe *_e .CT_ShapeProperties };

// SetWidth sets the width of the shape.
func (_ad ShapeProperties )SetWidth (w _g .Distance ){_ad .ensureXfrm ();if _ad ._afe .Xfrm .Ext ==nil {_ad ._afe .Xfrm .Ext =_e .NewCT_PositiveSize2D ();};_ad ._afe .Xfrm .Ext .CxAttr =int64 (w /_g .EMU );};

// X returns the inner wrapped XML type.
func (_ab ParagraphProperties )X ()*_e .CT_TextParagraphProperties {return _ab ._eg };

// SetHeight sets the height of the shape.
func (_cfb ShapeProperties )SetHeight (h _g .Distance ){_cfb .ensureXfrm ();if _cfb ._afe .Xfrm .Ext ==nil {_cfb ._afe .Xfrm .Ext =_e .NewCT_PositiveSize2D ();};_cfb ._afe .Xfrm .Ext .CyAttr =int64 (h /_g .EMU );};

// X returns the inner wrapped XML type.
func (_bgd ShapeProperties )X ()*_e .CT_ShapeProperties {return _bgd ._afe };

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_e .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};func (_dgd ShapeProperties )clearFill (){_dgd ._afe .FillPropertiesChoice =_e .NewEG_FillPropertiesChoice ();};

// SetFont controls the font of a run.
func (_ggd RunProperties )SetFont (s string ){_ggd ._fe .Latin =_e .NewCT_TextFont ();_ggd ._fe .Latin .TypefaceAttr =s ;};

// SetGeometry sets the shape type of the shape
func (_beg ShapeProperties )SetGeometry (g _e .ST_ShapeType ){if _beg ._afe .GeometryChoice .PrstGeom ==nil {_beg ._afe .GeometryChoice .PrstGeom =_e .NewCT_PresetGeometry2D ();};_beg ._afe .GeometryChoice .PrstGeom .PrstAttr =g ;};func (_bb ShapeProperties )SetSolidFill (c _cc .Color ){_bb .clearFill ();
_bb ._afe .FillPropertiesChoice .SolidFill =_e .NewCT_SolidColorFillProperties ();_bb ._afe .FillPropertiesChoice .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();_bb ._afe .FillPropertiesChoice .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetFlipVertical controls if the shape is flipped vertically.
func (_ecd ShapeProperties )SetFlipVertical (b bool ){_ecd .ensureXfrm ();if !b {_ecd ._afe .Xfrm .FlipVAttr =nil ;}else {_ecd ._afe .Xfrm .FlipVAttr =_c .Bool (true );};};

// SetText sets the run's text contents.
func (_be Run )SetText (s string ){_be ._bf .TextRunChoice .Br =nil ;_be ._bf .TextRunChoice .Fld =nil ;if _be ._bf .TextRunChoice .R ==nil {_be ._bf .TextRunChoice .R =_e .NewCT_RegularTextRun ();};_be ._bf .TextRunChoice .R .T =s ;};func MakeShapeProperties (x *_e .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};


// GetPosition gets the position of the shape in EMU.
func (_aa ShapeProperties )GetPosition ()(int64 ,int64 ){_aa .ensureXfrm ();if _aa ._afe .Xfrm .Off ==nil {_aa ._afe .Xfrm .Off =_e .NewCT_Point2D ();};return *_aa ._afe .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_aa ._afe .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;
};

// SetLevel sets the level of indentation of a paragraph.
func (_db ParagraphProperties )SetLevel (idx int32 ){_db ._eg .LvlAttr =_c .Int32 (idx )};

// RunProperties controls the run properties.
type RunProperties struct{_fe *_e .CT_TextCharacterProperties ;};

// LineJoin is the type of line join
type LineJoin byte ;

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_e .CT_TextParagraph )Paragraph {return Paragraph {x }};

// Properties returns the paragraph properties.
func (_cf Paragraph )Properties ()ParagraphProperties {if _cf ._ccf .PPr ==nil {_cf ._ccf .PPr =_e .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_cf ._ccf .PPr );};