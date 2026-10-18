// Copyright 2026 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package document

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	_ebgec = true
}

type userProfile struct {
	City    string
	Country string
}

type testUser struct {
	Name    string
	Age     int
	Active  bool
	Profile userProfile
}

func (u testUser) FullTitle() string {
	return "Mr. " + u.Name
}

func TestTemplateExtractTokens(t *testing.T) {
	tp := NewTemplateProcessor()

	tests := []struct {
		name     string
		input    string
		expected []TemplateToken
	}{
		{
			name:     "empty text",
			input:    "",
			expected: nil,
		},
		{
			name:     "no tokens",
			input:    "Plain document text without any template tags.",
			expected: nil,
		},
		{
			name:  "single expression",
			input: "Hello {{ user.name }}, welcome!",
			expected: []TemplateToken{
				{
					Raw:   "{{ user.name }}",
					Tag:   "user.name",
					Type:  TokenTypeExpression,
					Start: 6,
					End:   21,
				},
			},
		},
		{
			name:  "multiple mixed tokens",
			input: "Dear {{ name }}: {% if is_admin %}Admin Panel{% endif %} {# private note #}",
			expected: []TemplateToken{
				{
					Raw:   "{{ name }}",
					Tag:   "name",
					Type:  TokenTypeExpression,
					Start: 5,
					End:   15,
				},
				{
					Raw:   "{% if is_admin %}",
					Tag:   "if is_admin",
					Type:  TokenTypeStatement,
					Start: 17,
					End:   34,
				},
				{
					Raw:   "{% endif %}",
					Tag:   "endif",
					Type:  TokenTypeStatement,
					Start: 45,
					End:   56,
				},
				{
					Raw:   "{# private note #}",
					Tag:   "private note",
					Type:  TokenTypeComment,
					Start: 57,
					End:   75,
				},
			},
		},
		{
			name:     "unclosed delimiter ignored cleanly",
			input:    "Unclosed tag {{ name without closing delimiter",
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens := tp.ExtractTokens(tc.input)
			assert.Equal(t, len(tc.expected), len(tokens))
			for i, exp := range tc.expected {
				assert.Equal(t, exp.Raw, tokens[i].Raw)
				assert.Equal(t, exp.Tag, tokens[i].Tag)
				assert.Equal(t, exp.Type, tokens[i].Type)
				assert.Equal(t, exp.Start, tokens[i].Start)
				assert.Equal(t, exp.End, tokens[i].End)
			}
		})
	}
}

func TestTemplateCustomDelimiters(t *testing.T) {
	customDelims := TemplateDelimiters{
		ExprStart: "[[",
		ExprEnd:   "]]",
	}
	tp := NewTemplateProcessor(WithDelimiters(customDelims))

	tokens := tp.ExtractTokens("Hello [[ user.name ]] from [[ location ]]")
	require.Len(t, tokens, 2)
	assert.Equal(t, "user.name", tokens[0].Tag)
	assert.Equal(t, "location", tokens[1].Tag)

	rendered, err := tp.RenderText("Hello [[ user.name ]] from [[ location ]]!", map[string]interface{}{
		"user": map[string]interface{}{
			"name": "David",
		},
		"location": "Stockholm",
	})
	require.NoError(t, err)
	assert.Equal(t, "Hello David from Stockholm!", rendered)
}

func TestTemplateRenderTextExpressions(t *testing.T) {
	tp := NewTemplateProcessor()

	ctx := map[string]interface{}{
		"title": "Quarterly Report",
		"count": 42,
		"user": testUser{
			Name:   "Alice",
			Age:    30,
			Active: true,
			Profile: userProfile{
				City:    "Reykjavik",
				Country: "Iceland",
			},
		},
		"items":    []string{"Apple", "Banana", "Cherry"},
		"emptyVal": "",
	}

	tests := []struct {
		name     string
		template string
		expected string
	}{
		{
			name:     "simple variable",
			template: "Doc: {{ title }}",
			expected: "Doc: Quarterly Report",
		},
		{
			name:     "numeric variable",
			template: "Count: {{ count }}",
			expected: "Count: 42",
		},
		{
			name:     "nested struct field",
			template: "User: {{ user.Name }}, City: {{ user.Profile.City }}",
			expected: "User: Alice, City: Reykjavik",
		},
		{
			name:     "struct method invocation",
			template: "Greeting: {{ user.FullTitle }}",
			expected: "Greeting: Mr. Alice",
		},
		{
			name:     "slice indexing with brackets and dots",
			template: "First: {{ items[0] }}, Second: {{ items.1 }}",
			expected: "First: Apple, Second: Banana",
		},
		{
			name:     "filter upper and lower",
			template: "{{ user.Name | upper }} in {{ user.Profile.City | lower }}",
			expected: "ALICE in reykjavik",
		},
		{
			name:     "filter trim and length",
			template: "Items count: {{ items | length }}",
			expected: "Items count: 3",
		},
		{
			name:     "filter default fallback",
			template: "Val: {{ emptyVal | default(\"fallback-value\") }}",
			expected: "Val: fallback-value",
		},
		{
			name:     "ternary condition true",
			template: "Status: {{ user.Active ? \"ACTIVE\" : \"SUSPENDED\" }}",
			expected: "Status: ACTIVE",
		},
		{
			name:     "ternary condition comparison",
			template: "Check: {{ count > 40 ? \"High\" : \"Low\" }}",
			expected: "Check: High",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tp.RenderText(tc.template, ctx)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestTemplateControlStatements(t *testing.T) {
	tp := NewTemplateProcessor()

	t.Run("if true branch", func(t *testing.T) {
		res, err := tp.RenderText("Prefix {% if active %}User is Active{% else %}Inactive{% endif %} Suffix", map[string]interface{}{
			"active": true,
		})
		require.NoError(t, err)
		assert.Equal(t, "Prefix User is Active Suffix", res)
	})

	t.Run("if false branch", func(t *testing.T) {
		res, err := tp.RenderText("Prefix {% if active %}User is Active{% else %}Inactive{% endif %} Suffix", map[string]interface{}{
			"active": false,
		})
		require.NoError(t, err)
		assert.Equal(t, "Prefix Inactive Suffix", res)
	})

	t.Run("if elif else branching", func(t *testing.T) {
		tmpl := "Grade: {% if score >= 90 %}A{% elif score >= 80 %}B{% else %}C{% endif %}"
		resA, err := tp.RenderText(tmpl, map[string]interface{}{"score": 95})
		require.NoError(t, err)
		assert.Equal(t, "Grade: A", resA)

		resB, err := tp.RenderText(tmpl, map[string]interface{}{"score": 85})
		require.NoError(t, err)
		assert.Equal(t, "Grade: B", resB)

		resC, err := tp.RenderText(tmpl, map[string]interface{}{"score": 75})
		require.NoError(t, err)
		assert.Equal(t, "Grade: C", resC)
	})

	t.Run("for loop with loop variables", func(t *testing.T) {
		tmpl := "Items: {% for item in items %}{{ loop.index1 }}.{{ item }}{% if not loop.last %}, {% endif %}{% endfor %}"
		res, err := tp.RenderText(tmpl, map[string]interface{}{
			"items": []string{"Alpha", "Beta", "Gamma"},
		})
		require.NoError(t, err)
		assert.Equal(t, "Items: 1.Alpha, 2.Beta, 3.Gamma", res)
	})

	t.Run("comment stripping", func(t *testing.T) {
		tmpl := "Before {# secret comment #} After"
		res, err := tp.RenderText(tmpl, nil)
		require.NoError(t, err)
		assert.Equal(t, "Before  After", res)
	})
}

func TestTemplateCustomFilterRegistration(t *testing.T) {
	tp := NewTemplateProcessor()
	tp.RegisterFilter("shout", func(val interface{}, args ...string) (interface{}, error) {
		return fmt.Sprintf("%v!!!", val), nil
	})

	res, err := tp.RenderText("Say {{ message | shout }}", map[string]interface{}{
		"message": "hello",
	})
	require.NoError(t, err)
	assert.Equal(t, "Say hello!!!", res)
}

func TestTemplateCustomRenderer(t *testing.T) {
	tp := NewTemplateProcessor(WithRendererFunc(func(tag string, ctx map[string]interface{}) (string, error) {
		if tag == "custom_eval" {
			return "DELEGATED_VALUE", nil
		}
		return "UNKNOWN", nil
	}))

	res, err := tp.RenderText("Value: {{ custom_eval }}", nil)
	require.NoError(t, err)
	assert.Equal(t, "Value: DELEGATED_VALUE", res)
}

type mockEngine struct{}

func (m *mockEngine) RenderTemplate(content string, ctx map[string]interface{}) (string, error) {
	return "ENGINE_OVERRIDE: " + content, nil
}

func TestTemplateCustomEngine(t *testing.T) {
	tp := NewTemplateProcessor(WithEngine(&mockEngine{}))
	res, err := tp.RenderText("Raw input", nil)
	require.NoError(t, err)
	assert.Equal(t, "ENGINE_OVERRIDE: Raw input", res)
}

func TestTemplateStrictMode(t *testing.T) {
	tpStrict := NewTemplateProcessor(WithStrict(true))
	_, err := tpStrict.RenderText("Hello {{ missing_var }}", map[string]interface{}{})
	assert.Error(t, err)

	tpNonStrict := NewTemplateProcessor(WithStrict(false), WithRemoveUnmatched(true))
	res, err := tpNonStrict.RenderText("Hello {{ missing_var }}!", map[string]interface{}{})
	require.NoError(t, err)
	assert.Equal(t, "Hello !", res)
}

func TestTemplateParagraphSplitRunsSubstitution(t *testing.T) {
	doc := New()
	para := doc.AddParagraph()

	// Artificially split a template tag across 4 distinct runs to simulate Word run fragmentation
	r1 := para.AddRun()
	r1.AddText("Greeting: {{ ")
	r2 := para.AddRun()
	r2.AddText("user.")
	r3 := para.AddRun()
	r3.AddText("name")
	r4 := para.AddRun()
	r4.AddText(" }}!")

	ctx := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Alexander",
		},
	}

	err := doc.RenderTemplate(ctx)
	require.NoError(t, err)

	// Verify concatenated paragraph text
	assert.Equal(t, "Greeting: Alexander!", paragraphText(para))
}

func TestTemplateParagraphMultipleTokensAndRuns(t *testing.T) {
	doc := New()
	para := doc.AddParagraph()

	r1 := para.AddRun()
	r1.AddText("Report for {{ dept }} by {{ author }} on {{ date }}")

	ctx := map[string]interface{}{
		"dept":   "Engineering",
		"author": "Marcus",
		"date":   "2026-10-18",
	}

	err := doc.RenderTemplate(ctx)
	require.NoError(t, err)

	assert.Equal(t, "Report for Engineering by Marcus on 2026-10-18", paragraphText(para))
}

func TestTemplateDocumentFullTemplateRendering(t *testing.T) {
	doc := New()

	// Body paragraph
	p1 := doc.AddParagraph()
	r1 := p1.AddRun()
	r1.AddText("Welcome {{ customer.name }} to {{ service }}.")

	// Table with template tags
	table := doc.AddTable()
	row := table.AddRow()
	cell1 := row.AddCell()
	cell1P := cell1.AddParagraph()
	cell1P.AddRun().AddText("Invoice #{{ invoice.id }}")

	cell2 := row.AddCell()
	cell2P := cell2.AddParagraph()
	cell2P.AddRun().AddText("Total: ${{ invoice.amount }}")

	// Header
	header := doc.AddHeader()
	hP := header.AddParagraph()
	hP.AddRun().AddText("Confidential - {{ company }}")

	// Footer
	footer := doc.AddFooter()
	fP := footer.AddParagraph()
	fP.AddRun().AddText("Page 1 of {{ total_pages }}")

	ctx := map[string]interface{}{
		"customer": map[string]interface{}{
			"name": "Beatrice",
		},
		"service": "Cloud Sync",
		"invoice": map[string]interface{}{
			"id":     "INV-9901",
			"amount": 249.50,
		},
		"company":     "Acme Global",
		"total_pages": 5,
	}

	// Verify token extraction across document
	extracted := doc.ExtractTemplateTokens()
	assert.NotEmpty(t, extracted)

	err := doc.RenderTemplate(ctx)
	require.NoError(t, err)

	assert.Equal(t, "Welcome Beatrice to Cloud Sync.", paragraphText(p1))
	assert.Equal(t, "Invoice #INV-9901", paragraphText(cell1P))
	assert.Equal(t, "Total: $249.5", paragraphText(cell2P))
	assert.Equal(t, "Confidential - Acme Global", paragraphText(hP))
	assert.Equal(t, "Page 1 of 5", paragraphText(fP))
}

func TestTemplateSubstituteTokensDirect(t *testing.T) {
	doc := New()
	p := doc.AddParagraph()
	p.AddRun().AddText("Hello {{ first_name }} {{ last_name }}!")

	replacements := map[string]string{
		"first_name": "Elena",
		"last_name":  "Rostova",
	}

	err := doc.SubstituteTokens(replacements)
	require.NoError(t, err)
	assert.Equal(t, "Hello Elena Rostova!", paragraphText(p))
}

func TestTemplatePongoConvertAndFileRendering(t *testing.T) {
	tempDir := t.TempDir()
	templateFile := filepath.Join(tempDir, "base_template.docx")
	outputFile := filepath.Join(tempDir, "output_rendered.docx")

	doc := New()
	p := doc.AddParagraph()
	p.AddRun().AddText("Template Invoice: {{ invoice_id }} for {{ client }}")
	err := doc.SaveToFile(templateFile)
	require.NoError(t, err)

	ctx := map[string]interface{}{
		"invoice_id": "INV-2026-X",
		"client":     "Nordic Enterprises",
	}

	// Test RenderTemplateFile
	err = RenderTemplateFile(templateFile, outputFile, ctx)
	require.NoError(t, err)

	readDoc, err := Open(outputFile)
	require.NoError(t, err)
	paras := readDoc.Paragraphs()
	require.NotEmpty(t, paras)
	assert.Equal(t, "Template Invoice: INV-2026-X for Nordic Enterprises", paragraphText(paras[0]))

	// Test PongoConvert with explicit output path
	outPath, err := PongoConvert(templateFile, ctx, outputFile)
	require.NoError(t, err)
	assert.Equal(t, outputFile, outPath)

	// Test PongoConvert with automatic temp file
	autoPath, err := PongoConvert(templateFile, ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, autoPath)
	defer os.Remove(autoPath)

	autoDoc, err := Open(autoPath)
	require.NoError(t, err)
	autoParas := autoDoc.Paragraphs()
	require.NotEmpty(t, autoParas)
	assert.Equal(t, "Template Invoice: INV-2026-X for Nordic Enterprises", paragraphText(autoParas[0]))

	// Test error cases
	_, err = PongoConvert("", ctx)
	assert.Error(t, err)

	_, err = PongoConvert("/nonexistent/path/template.docx", ctx)
	assert.Error(t, err)
}
