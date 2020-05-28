// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package unioffice

// Consts for content types used throughout the package
const (
	// Common strict
	OfficeDocumentTypeStrict     = "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument"
	StylesTypeStrict             = "http://purl.oclc.org/ooxml/officeDocument/relationships/styles"
	ThemeTypeStrict              = "http://purl.oclc.org/ooxml/officeDocument/relationships/theme"
	SettingsTypeStrict           = "http://purl.oclc.org/ooxml/officeDocument/relationships/settings"
	ImageTypeStrict              = "http://purl.oclc.org/ooxml/officeDocument/relationships/image"
	CommentsTypeStrict           = "http://purl.oclc.org/ooxml/officeDocument/relationships/comments"
	ThumbnailTypeStrict          = "http://purl.oclc.org/ooxml/officeDocument/relationships/metadata/thumbnail"
	DrawingTypeStrict            = "http://purl.oclc.org/ooxml/officeDocument/relationships/drawing"
	ChartTypeStrict              = "http://purl.oclc.org/ooxml/officeDocument/relationships/chart"
	ExtendedPropertiesTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/extendedProperties"
	CustomXMLTypeStrict          = "http://purl.oclc.org/ooxml/officeDocument/relationships/customXml"

	// SML strict
	WorksheetTypeStrict     = "http://purl.oclc.org/ooxml/officeDocument/relationships/worksheet"
	SharedStringsTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/sharedStrings"
	// Deprecated: Renamed to SharedStringsTypeStrict, will be removed in next major version.
	SharedStingsTypeStrict = SharedStringsTypeStrict
	TableTypeStrict        = "http://purl.oclc.org/ooxml/officeDocument/relationships/table"

	// WML strict
	HeaderTypeStrict      = "http://purl.oclc.org/ooxml/officeDocument/relationships/header"
	FooterTypeStrict      = "http://purl.oclc.org/ooxml/officeDocument/relationships/footer"
	NumberingTypeStrict   = "http://purl.oclc.org/ooxml/officeDocument/relationships/numbering"
	FontTableTypeStrict   = "http://purl.oclc.org/ooxml/officeDocument/relationships/fontTable"
	WebSettingsTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/webSettings"
	FootNotesTypeStrict   = "http://purl.oclc.org/ooxml/officeDocument/relationships/footnotes"
	EndNotesTypeStrict    = "http://purl.oclc.org/ooxml/officeDocument/relationships/endnotes"

	// PML strict
	SlideTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/slide"

	// VML strict
	VMLDrawingTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/vmlDrawing"

	// Common
	OfficeDocumentType     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	StylesType             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
	ThemeType              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"
	ThemeContentType       = "application/vnd.openxmlformats-officedocument.theme+xml"
	SettingsType           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings"
	ImageType              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
	CommentsType           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
	CommentsContentType    = "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml"
	ThumbnailType          = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/thumbnail"
	DrawingType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/drawing"
	DrawingContentType     = "application/vnd.openxmlformats-officedocument.drawing+xml"
	ChartType              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	ChartContentType       = "application/vnd.openxmlformats-officedocument.drawingml.chart+xml"
	HyperLinkType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
	ExtendedPropertiesType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	CorePropertiesType     = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
	CustomPropertiesType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/custom-properties"
	CustomXMLType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXml"
	TableStylesType        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"
	ViewPropertiesType     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/viewProps"

	// SML
	WorksheetType        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	WorksheetContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
	SharedStringsType    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings"
	// Deprecated: Renamed to SharedStringsType, will be removed in next major version.
	SharedStingsType         = SharedStringsType
	SharedStringsContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	SMLStyleSheetContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"
	TableType                = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/table"
	TableContentType         = "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml"

	// WML
	HeaderType      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header"
	FooterType      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer"
	NumberingType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering"
	FontTableType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable"
	WebSettingsType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings"
	FootNotesType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footnotes"
	EndNotesType    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/endnotes"

	// PML
	SlideType                  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"
	SlideContentType           = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"
	SlideMasterType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"
	SlideMasterContentType     = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"
	SlideLayoutType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout"
	SlideLayoutContentType     = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"
	PresentationPropertiesType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/presProps"
	HandoutMasterType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/handoutMaster"
	NotesMasterType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesMaster"

	// VML
	VMLDrawingType        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/vmlDrawing"
	VMLDrawingContentType = "application/vnd.openxmlformats-officedocument.vmlDrawing"
)
