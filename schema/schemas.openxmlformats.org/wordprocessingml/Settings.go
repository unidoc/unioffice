// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/schemaLibrary"
)

type Settings struct {
	CT_Settings
}

func NewSettings() *Settings {
	ret := &Settings{}
	ret.CT_Settings = *NewCT_Settings()
	return ret
}

func (m *Settings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:settings"
	return m.CT_Settings.MarshalXML(e, start)
}

func (m *Settings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Settings = *NewCT_Settings()
lSettings:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "writeProtection":
				m.WriteProtection = NewCT_WriteProtection()
				if err := d.DecodeElement(m.WriteProtection, &el); err != nil {
					return err
				}
			case "view":
				m.View = NewCT_View()
				if err := d.DecodeElement(m.View, &el); err != nil {
					return err
				}
			case "zoom":
				m.Zoom = NewCT_Zoom()
				if err := d.DecodeElement(m.Zoom, &el); err != nil {
					return err
				}
			case "removePersonalInformation":
				m.RemovePersonalInformation = NewCT_OnOff()
				if err := d.DecodeElement(m.RemovePersonalInformation, &el); err != nil {
					return err
				}
			case "removeDateAndTime":
				m.RemoveDateAndTime = NewCT_OnOff()
				if err := d.DecodeElement(m.RemoveDateAndTime, &el); err != nil {
					return err
				}
			case "doNotDisplayPageBoundaries":
				m.DoNotDisplayPageBoundaries = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDisplayPageBoundaries, &el); err != nil {
					return err
				}
			case "displayBackgroundShape":
				m.DisplayBackgroundShape = NewCT_OnOff()
				if err := d.DecodeElement(m.DisplayBackgroundShape, &el); err != nil {
					return err
				}
			case "printPostScriptOverText":
				m.PrintPostScriptOverText = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintPostScriptOverText, &el); err != nil {
					return err
				}
			case "printFractionalCharacterWidth":
				m.PrintFractionalCharacterWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFractionalCharacterWidth, &el); err != nil {
					return err
				}
			case "printFormsData":
				m.PrintFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFormsData, &el); err != nil {
					return err
				}
			case "embedTrueTypeFonts":
				m.EmbedTrueTypeFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedTrueTypeFonts, &el); err != nil {
					return err
				}
			case "embedSystemFonts":
				m.EmbedSystemFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedSystemFonts, &el); err != nil {
					return err
				}
			case "saveSubsetFonts":
				m.SaveSubsetFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveSubsetFonts, &el); err != nil {
					return err
				}
			case "saveFormsData":
				m.SaveFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveFormsData, &el); err != nil {
					return err
				}
			case "mirrorMargins":
				m.MirrorMargins = NewCT_OnOff()
				if err := d.DecodeElement(m.MirrorMargins, &el); err != nil {
					return err
				}
			case "alignBordersAndEdges":
				m.AlignBordersAndEdges = NewCT_OnOff()
				if err := d.DecodeElement(m.AlignBordersAndEdges, &el); err != nil {
					return err
				}
			case "bordersDoNotSurroundHeader":
				m.BordersDoNotSurroundHeader = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundHeader, &el); err != nil {
					return err
				}
			case "bordersDoNotSurroundFooter":
				m.BordersDoNotSurroundFooter = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundFooter, &el); err != nil {
					return err
				}
			case "gutterAtTop":
				m.GutterAtTop = NewCT_OnOff()
				if err := d.DecodeElement(m.GutterAtTop, &el); err != nil {
					return err
				}
			case "hideSpellingErrors":
				m.HideSpellingErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideSpellingErrors, &el); err != nil {
					return err
				}
			case "hideGrammaticalErrors":
				m.HideGrammaticalErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideGrammaticalErrors, &el); err != nil {
					return err
				}
			case "activeWritingStyle":
				tmp := NewCT_WritingStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ActiveWritingStyle = append(m.ActiveWritingStyle, tmp)
			case "proofState":
				m.ProofState = NewCT_Proof()
				if err := d.DecodeElement(m.ProofState, &el); err != nil {
					return err
				}
			case "formsDesign":
				m.FormsDesign = NewCT_OnOff()
				if err := d.DecodeElement(m.FormsDesign, &el); err != nil {
					return err
				}
			case "attachedTemplate":
				m.AttachedTemplate = NewCT_Rel()
				if err := d.DecodeElement(m.AttachedTemplate, &el); err != nil {
					return err
				}
			case "linkStyles":
				m.LinkStyles = NewCT_OnOff()
				if err := d.DecodeElement(m.LinkStyles, &el); err != nil {
					return err
				}
			case "stylePaneFormatFilter":
				m.StylePaneFormatFilter = NewCT_StylePaneFilter()
				if err := d.DecodeElement(m.StylePaneFormatFilter, &el); err != nil {
					return err
				}
			case "stylePaneSortMethod":
				m.StylePaneSortMethod = NewCT_StyleSort()
				if err := d.DecodeElement(m.StylePaneSortMethod, &el); err != nil {
					return err
				}
			case "documentType":
				m.DocumentType = NewCT_DocType()
				if err := d.DecodeElement(m.DocumentType, &el); err != nil {
					return err
				}
			case "mailMerge":
				m.MailMerge = NewCT_MailMerge()
				if err := d.DecodeElement(m.MailMerge, &el); err != nil {
					return err
				}
			case "revisionView":
				m.RevisionView = NewCT_TrackChangesView()
				if err := d.DecodeElement(m.RevisionView, &el); err != nil {
					return err
				}
			case "trackRevisions":
				m.TrackRevisions = NewCT_OnOff()
				if err := d.DecodeElement(m.TrackRevisions, &el); err != nil {
					return err
				}
			case "doNotTrackMoves":
				m.DoNotTrackMoves = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackMoves, &el); err != nil {
					return err
				}
			case "doNotTrackFormatting":
				m.DoNotTrackFormatting = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackFormatting, &el); err != nil {
					return err
				}
			case "documentProtection":
				m.DocumentProtection = NewCT_DocProtect()
				if err := d.DecodeElement(m.DocumentProtection, &el); err != nil {
					return err
				}
			case "autoFormatOverride":
				m.AutoFormatOverride = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoFormatOverride, &el); err != nil {
					return err
				}
			case "styleLockTheme":
				m.StyleLockTheme = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockTheme, &el); err != nil {
					return err
				}
			case "styleLockQFSet":
				m.StyleLockQFSet = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockQFSet, &el); err != nil {
					return err
				}
			case "defaultTabStop":
				m.DefaultTabStop = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DefaultTabStop, &el); err != nil {
					return err
				}
			case "autoHyphenation":
				m.AutoHyphenation = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoHyphenation, &el); err != nil {
					return err
				}
			case "consecutiveHyphenLimit":
				m.ConsecutiveHyphenLimit = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.ConsecutiveHyphenLimit, &el); err != nil {
					return err
				}
			case "hyphenationZone":
				m.HyphenationZone = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.HyphenationZone, &el); err != nil {
					return err
				}
			case "doNotHyphenateCaps":
				m.DoNotHyphenateCaps = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotHyphenateCaps, &el); err != nil {
					return err
				}
			case "showEnvelope":
				m.ShowEnvelope = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowEnvelope, &el); err != nil {
					return err
				}
			case "summaryLength":
				m.SummaryLength = NewCT_DecimalNumberOrPrecent()
				if err := d.DecodeElement(m.SummaryLength, &el); err != nil {
					return err
				}
			case "clickAndTypeStyle":
				m.ClickAndTypeStyle = NewCT_String()
				if err := d.DecodeElement(m.ClickAndTypeStyle, &el); err != nil {
					return err
				}
			case "defaultTableStyle":
				m.DefaultTableStyle = NewCT_String()
				if err := d.DecodeElement(m.DefaultTableStyle, &el); err != nil {
					return err
				}
			case "evenAndOddHeaders":
				m.EvenAndOddHeaders = NewCT_OnOff()
				if err := d.DecodeElement(m.EvenAndOddHeaders, &el); err != nil {
					return err
				}
			case "bookFoldRevPrinting":
				m.BookFoldRevPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldRevPrinting, &el); err != nil {
					return err
				}
			case "bookFoldPrinting":
				m.BookFoldPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldPrinting, &el); err != nil {
					return err
				}
			case "bookFoldPrintingSheets":
				m.BookFoldPrintingSheets = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.BookFoldPrintingSheets, &el); err != nil {
					return err
				}
			case "drawingGridHorizontalSpacing":
				m.DrawingGridHorizontalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalSpacing, &el); err != nil {
					return err
				}
			case "drawingGridVerticalSpacing":
				m.DrawingGridVerticalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalSpacing, &el); err != nil {
					return err
				}
			case "displayHorizontalDrawingGridEvery":
				m.DisplayHorizontalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayHorizontalDrawingGridEvery, &el); err != nil {
					return err
				}
			case "displayVerticalDrawingGridEvery":
				m.DisplayVerticalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayVerticalDrawingGridEvery, &el); err != nil {
					return err
				}
			case "doNotUseMarginsForDrawingGridOrigin":
				m.DoNotUseMarginsForDrawingGridOrigin = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseMarginsForDrawingGridOrigin, &el); err != nil {
					return err
				}
			case "drawingGridHorizontalOrigin":
				m.DrawingGridHorizontalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalOrigin, &el); err != nil {
					return err
				}
			case "drawingGridVerticalOrigin":
				m.DrawingGridVerticalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalOrigin, &el); err != nil {
					return err
				}
			case "doNotShadeFormData":
				m.DoNotShadeFormData = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotShadeFormData, &el); err != nil {
					return err
				}
			case "noPunctuationKerning":
				m.NoPunctuationKerning = NewCT_OnOff()
				if err := d.DecodeElement(m.NoPunctuationKerning, &el); err != nil {
					return err
				}
			case "characterSpacingControl":
				m.CharacterSpacingControl = NewCT_CharacterSpacing()
				if err := d.DecodeElement(m.CharacterSpacingControl, &el); err != nil {
					return err
				}
			case "printTwoOnOne":
				m.PrintTwoOnOne = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintTwoOnOne, &el); err != nil {
					return err
				}
			case "strictFirstAndLastChars":
				m.StrictFirstAndLastChars = NewCT_OnOff()
				if err := d.DecodeElement(m.StrictFirstAndLastChars, &el); err != nil {
					return err
				}
			case "noLineBreaksAfter":
				m.NoLineBreaksAfter = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksAfter, &el); err != nil {
					return err
				}
			case "noLineBreaksBefore":
				m.NoLineBreaksBefore = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksBefore, &el); err != nil {
					return err
				}
			case "savePreviewPicture":
				m.SavePreviewPicture = NewCT_OnOff()
				if err := d.DecodeElement(m.SavePreviewPicture, &el); err != nil {
					return err
				}
			case "doNotValidateAgainstSchema":
				m.DoNotValidateAgainstSchema = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotValidateAgainstSchema, &el); err != nil {
					return err
				}
			case "saveInvalidXml":
				m.SaveInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveInvalidXml, &el); err != nil {
					return err
				}
			case "ignoreMixedContent":
				m.IgnoreMixedContent = NewCT_OnOff()
				if err := d.DecodeElement(m.IgnoreMixedContent, &el); err != nil {
					return err
				}
			case "alwaysShowPlaceholderText":
				m.AlwaysShowPlaceholderText = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysShowPlaceholderText, &el); err != nil {
					return err
				}
			case "doNotDemarcateInvalidXml":
				m.DoNotDemarcateInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDemarcateInvalidXml, &el); err != nil {
					return err
				}
			case "saveXmlDataOnly":
				m.SaveXmlDataOnly = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveXmlDataOnly, &el); err != nil {
					return err
				}
			case "useXSLTWhenSaving":
				m.UseXSLTWhenSaving = NewCT_OnOff()
				if err := d.DecodeElement(m.UseXSLTWhenSaving, &el); err != nil {
					return err
				}
			case "saveThroughXslt":
				m.SaveThroughXslt = NewCT_SaveThroughXslt()
				if err := d.DecodeElement(m.SaveThroughXslt, &el); err != nil {
					return err
				}
			case "showXMLTags":
				m.ShowXMLTags = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowXMLTags, &el); err != nil {
					return err
				}
			case "alwaysMergeEmptyNamespace":
				m.AlwaysMergeEmptyNamespace = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysMergeEmptyNamespace, &el); err != nil {
					return err
				}
			case "updateFields":
				m.UpdateFields = NewCT_OnOff()
				if err := d.DecodeElement(m.UpdateFields, &el); err != nil {
					return err
				}
			case "hdrShapeDefaults":
				m.HdrShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.HdrShapeDefaults, &el); err != nil {
					return err
				}
			case "footnotePr":
				m.FootnotePr = NewCT_FtnDocProps()
				if err := d.DecodeElement(m.FootnotePr, &el); err != nil {
					return err
				}
			case "endnotePr":
				m.EndnotePr = NewCT_EdnDocProps()
				if err := d.DecodeElement(m.EndnotePr, &el); err != nil {
					return err
				}
			case "compat":
				m.Compat = NewCT_Compat()
				if err := d.DecodeElement(m.Compat, &el); err != nil {
					return err
				}
			case "docVars":
				m.DocVars = NewCT_DocVars()
				if err := d.DecodeElement(m.DocVars, &el); err != nil {
					return err
				}
			case "rsids":
				m.Rsids = NewCT_DocRsids()
				if err := d.DecodeElement(m.Rsids, &el); err != nil {
					return err
				}
			case "mathPr":
				m.MathPr = math.NewMathPr()
				if err := d.DecodeElement(m.MathPr, &el); err != nil {
					return err
				}
			case "attachedSchema":
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AttachedSchema = append(m.AttachedSchema, tmp)
			case "themeFontLang":
				m.ThemeFontLang = NewCT_Language()
				if err := d.DecodeElement(m.ThemeFontLang, &el); err != nil {
					return err
				}
			case "clrSchemeMapping":
				m.ClrSchemeMapping = NewCT_ColorSchemeMapping()
				if err := d.DecodeElement(m.ClrSchemeMapping, &el); err != nil {
					return err
				}
			case "doNotIncludeSubdocsInStats":
				m.DoNotIncludeSubdocsInStats = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotIncludeSubdocsInStats, &el); err != nil {
					return err
				}
			case "doNotAutoCompressPictures":
				m.DoNotAutoCompressPictures = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotAutoCompressPictures, &el); err != nil {
					return err
				}
			case "forceUpgrade":
				m.ForceUpgrade = NewCT_Empty()
				if err := d.DecodeElement(m.ForceUpgrade, &el); err != nil {
					return err
				}
			case "captions":
				m.Captions = NewCT_Captions()
				if err := d.DecodeElement(m.Captions, &el); err != nil {
					return err
				}
			case "readModeInkLockDown":
				m.ReadModeInkLockDown = NewCT_ReadingModeInkLockDown()
				if err := d.DecodeElement(m.ReadModeInkLockDown, &el); err != nil {
					return err
				}
			case "smartTagType":
				tmp := NewCT_SmartTagType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SmartTagType = append(m.SmartTagType, tmp)
			case "schemaLibrary":
				m.SchemaLibrary = schemaLibrary.NewSchemaLibrary()
				if err := d.DecodeElement(m.SchemaLibrary, &el); err != nil {
					return err
				}
			case "shapeDefaults":
				m.ShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.ShapeDefaults, &el); err != nil {
					return err
				}
			case "doNotEmbedSmartTags":
				m.DoNotEmbedSmartTags = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotEmbedSmartTags, &el); err != nil {
					return err
				}
			case "decimalSymbol":
				m.DecimalSymbol = NewCT_String()
				if err := d.DecodeElement(m.DecimalSymbol, &el); err != nil {
					return err
				}
			case "listSeparator":
				m.ListSeparator = NewCT_String()
				if err := d.DecodeElement(m.ListSeparator, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lSettings
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Settings and its children
func (m *Settings) Validate() error {
	return m.ValidateWithPath("Settings")
}

// ValidateWithPath validates the Settings and its children, prefixing error messages with path
func (m *Settings) ValidateWithPath(path string) error {
	if err := m.CT_Settings.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
