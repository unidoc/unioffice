// Package formula provides formula parsing and evaluation. The lexer is
// implemented with a ragel grammar while the the parser is implemented with
// goyacc. The entire formula grammar is not implemented and not all functions
// are supported yet. For compatibility sake, upon failure to parse or execute a
// formula, gooxml leaves cached formula results blank allowing Excel to compute
// formulas upon load. This is similar to what most other Excel libraries do
// which leave all cached results blank instead of attempting to execute
// formulas.
//
// The unit tests for this package are unique in that we can take advantage of
// "cached" formula results that Excel/LibreOffice write to the sheet.  These
// are the computed results of a formula in string form.  By comparing these
// values to the value computed by the gooxml evaluation of the formula, adding
// a new test means just adding a new formula to one of the reference sheets
// with Excel. During the unit test, we evaluate the formula and compare it to
// the value that Excel computed.  If they're the same, the test passes.
package formula
