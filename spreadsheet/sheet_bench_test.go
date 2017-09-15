package spreadsheet_test

import (
	"testing"

	"baliance.com/gooxml/spreadsheet"
)

func BenchmarkAddRow(b *testing.B) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	for r := 0; r < b.N; r++ {
		sheet.AddRow()
	}
}

func BenchmarkAddCell(b *testing.B) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	row := sheet.AddRow()

	for c := 0; c < b.N; c++ {
		row.AddCell()
	}
}
