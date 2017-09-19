// Package format provides support for parsing and evaluating
// spreadsheetml/Excel number formats.
//
// Internally spreadsheets store numbers and dates values as a text
// representation of a floating point number (e.g. 1.2345).  This number is then
// displayed in Excel or another spreadsheet viewer differently depending on the
// number fornat of the cell style applied to the cell.
//
// As an example, the same value of 1.2345 can be displayed as:
// - "1" with format "0"
// - "1.2" with format "0.0"
// - "1.23" with format "0.00"
// - "1.235" with format "0.000"
// - "123%" with format "0%"
// - "1 23/100" with fornat "0 0/100"
// - "1.23E+00" with format "0.00E+00"
// - "29:37:41s" with format `[h]:mm:ss"s"`
package format
