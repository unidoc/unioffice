package format

import (
	"strings"
)

type Lexer struct {
	fmt     Format
	formats []Format
}

func (l *Lexer) nextFmt() {
	l.formats = append(l.formats, l.fmt)
	l.fmt = Format{}
}

func Parse(s string) []Format {
	l := Lexer{}
	l.Lex(strings.NewReader(s))
	l.formats = append(l.formats, l.fmt)
	return l.formats
}
