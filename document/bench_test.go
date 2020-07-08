package document_test

import (
	"bytes"
	"testing"

	"github.com/unidoc/unioffice/document"
)

func BenchmarkAddPara(b *testing.B) {
	doc := document.New()
	defer doc.Close()
	for i := 0; i < b.N; i++ {
		doc.AddParagraph()
	}
}

func BenchmarkAddRuns(b *testing.B) {
	doc := document.New()
	defer doc.Close()
	para := doc.AddParagraph()
	for i := 0; i < b.N; i++ {
		run := para.AddRun()
		run.AddText("test")
	}
}

func BenchmarkSave(b *testing.B) {
	doc := document.New()
	defer doc.Close()
	for i := 0; i < 100; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText("test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := bytes.Buffer{}
		doc.Save(&buf)
	}
}

func BenchmarkOpen(b *testing.B) {
	doc := document.New()
	defer doc.Close()
	for i := 0; i < 100; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText("test")
	}
	buf := bytes.Buffer{}
	doc.Save(&buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = document.Read(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	}
}
