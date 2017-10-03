// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml

// Consts for content types used throughout the package
const (
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

	// SML
	WorksheetType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	WorksheetContentType     = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
	SharedStingsType         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings"
	SharedStringsContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	SMLStyleSheetContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"
	TableType                = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/table"
	TableContentType         = "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml"
	ViewPropertiesType       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/viewProps"
	TableStylesType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"

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

	// VML
	VMLDrawingType        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/vmlDrawing"
	VMLDrawingContentType = "application/vnd.openxmlformats-officedocument.vmlDrawing"
)
