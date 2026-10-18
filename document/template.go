// Copyright 2026 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package document

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/unidoc/unioffice/v2"
)

// TokenType specifies the syntactic role of a template token.
type TokenType int

const (
	// TokenTypeExpression designates an evaluation token, such as {{ var }}.
	TokenTypeExpression TokenType = iota
	// TokenTypeStatement designates a control block token, such as {% if ... %}.
	TokenTypeStatement
	// TokenTypeComment designates a template comment, such as {# ... #}.
	TokenTypeComment
	// TokenTypeCustom designates a token matching custom delimiters.
	TokenTypeCustom
)

// String returns the human-readable description of a token type.
func (t TokenType) String() string {
	switch t {
	case TokenTypeExpression:
		return "Expression"
	case TokenTypeStatement:
		return "Statement"
	case TokenTypeComment:
		return "Comment"
	case TokenTypeCustom:
		return "Custom"
	default:
		return "Unknown"
	}
}

// TemplateToken holds details of an identified template token within document text.
type TemplateToken struct {
	Raw   string
	Tag   string
	Type  TokenType
	Start int
	End   int
}

// TemplateDelimiters defines the opening and closing delimiters for template tags.
type TemplateDelimiters struct {
	ExprStart    string
	ExprEnd      string
	StmtStart    string
	StmtEnd      string
	CommentStart string
	CommentEnd   string
}

// DefaultDelimiters returns the standard Jinja2 and Pongo2 delimiter set.
func DefaultDelimiters() TemplateDelimiters {
	return TemplateDelimiters{
		ExprStart:    "{{",
		ExprEnd:      "}}",
		StmtStart:    "{%",
		StmtEnd:      "%}",
		CommentStart: "{#",
		CommentEnd:   "#}",
	}
}

// normalize applies default delimiter values to any unassigned fields.
func (d TemplateDelimiters) normalize() TemplateDelimiters {
	def := DefaultDelimiters()
	if d.ExprStart == "" {
		d.ExprStart = def.ExprStart
	}
	if d.ExprEnd == "" {
		d.ExprEnd = def.ExprEnd
	}
	if d.StmtStart == "" {
		d.StmtStart = def.StmtStart
	}
	if d.StmtEnd == "" {
		d.StmtEnd = def.StmtEnd
	}
	if d.CommentStart == "" {
		d.CommentStart = def.CommentStart
	}
	if d.CommentEnd == "" {
		d.CommentEnd = def.CommentEnd
	}
	return d
}

// TemplateFilterFunc transforms an input value using optional arguments.
type TemplateFilterFunc func(val interface{}, args ...string) (interface{}, error)

// TemplateRenderer defines an interface for delegating tag evaluation to an engine.
type TemplateRenderer interface {
	Render(tag string, ctx map[string]interface{}) (string, error)
}

// TemplateRendererFunc adapts a plain function into a TemplateRenderer.
type TemplateRendererFunc func(tag string, ctx map[string]interface{}) (string, error)

// Render invokes the underlying function.
func (f TemplateRendererFunc) Render(tag string, ctx map[string]interface{}) (string, error) {
	return f(tag, ctx)
}

// TemplateEngine evaluates raw template content using a full template execution engine.
type TemplateEngine interface {
	RenderTemplate(content string, ctx map[string]interface{}) (string, error)
}

// TemplateOptions controls template scanning and substitution behavior.
type TemplateOptions struct {
	Delimiters      TemplateDelimiters
	Renderer        TemplateRenderer
	Engine          TemplateEngine
	Filters         map[string]TemplateFilterFunc
	RemoveUnmatched bool
	Strict          bool
}

// TemplateOption configures a TemplateOptions instance.
type TemplateOption func(*TemplateOptions)

// WithDelimiters sets custom opening and closing tag delimiters.
func WithDelimiters(delims TemplateDelimiters) TemplateOption {
	return func(o *TemplateOptions) {
		o.Delimiters = delims.normalize()
	}
}

// WithRenderer sets a custom tag renderer interface.
func WithRenderer(renderer TemplateRenderer) TemplateOption {
	return func(o *TemplateOptions) {
		o.Renderer = renderer
	}
}

// WithRendererFunc sets a custom tag rendering callback.
func WithRendererFunc(fn func(tag string, ctx map[string]interface{}) (string, error)) TemplateOption {
	return func(o *TemplateOptions) {
		o.Renderer = TemplateRendererFunc(fn)
	}
}

// WithEngine sets an external template engine implementation.
func WithEngine(engine TemplateEngine) TemplateOption {
	return func(o *TemplateOptions) {
		o.Engine = engine
	}
}

// WithFilter registers an expression filter.
func WithFilter(name string, fn TemplateFilterFunc) TemplateOption {
	return func(o *TemplateOptions) {
		if o.Filters == nil {
			o.Filters = make(map[string]TemplateFilterFunc)
		}
		o.Filters[strings.ToLower(name)] = fn
	}
}

// WithRemoveUnmatched controls whether unmapped tokens are replaced with empty text.
func WithRemoveUnmatched(remove bool) TemplateOption {
	return func(o *TemplateOptions) {
		o.RemoveUnmatched = remove
	}
}

// WithStrict causes template processing to fail on unmapped variables or evaluation errors.
func WithStrict(strict bool) TemplateOption {
	return func(o *TemplateOptions) {
		o.Strict = strict
	}
}

// TemplateProcessor manages template token extraction and document rendering.
type TemplateProcessor struct {
	options TemplateOptions
}

// NewTemplateProcessor creates a TemplateProcessor configured with the specified options.
func NewTemplateProcessor(opts ...TemplateOption) *TemplateProcessor {
	options := TemplateOptions{
		Delimiters: DefaultDelimiters(),
		Filters:    defaultTemplateFilters(),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	options.Delimiters = options.Delimiters.normalize()
	return &TemplateProcessor{options: options}
}

// SetDelimiters updates the tag delimiters on the processor.
func (tp *TemplateProcessor) SetDelimiters(delims TemplateDelimiters) {
	tp.options.Delimiters = delims.normalize()
}

// SetRenderer configures an external tag renderer.
func (tp *TemplateProcessor) SetRenderer(renderer TemplateRenderer) {
	tp.options.Renderer = renderer
}

// RegisterFilter adds or replaces a named template filter.
func (tp *TemplateProcessor) RegisterFilter(name string, fn TemplateFilterFunc) {
	if tp.options.Filters == nil {
		tp.options.Filters = make(map[string]TemplateFilterFunc)
	}
	tp.options.Filters[strings.ToLower(name)] = fn
}

// ExtractTokens locates all template tokens within a string.
func (tp *TemplateProcessor) ExtractTokens(text string) []TemplateToken {
	delims := tp.options.Delimiters
	var tokens []TemplateToken

	i := 0
	textLen := len(text)
	for i < textLen {
		if strings.HasPrefix(text[i:], delims.CommentStart) {
			endIdx := strings.Index(text[i+len(delims.CommentStart):], delims.CommentEnd)
			if endIdx >= 0 {
				totalEnd := i + len(delims.CommentStart) + endIdx + len(delims.CommentEnd)
				tokens = append(tokens, TemplateToken{
					Raw:   text[i:totalEnd],
					Tag:   strings.TrimSpace(text[i+len(delims.CommentStart) : totalEnd-len(delims.CommentEnd)]),
					Type:  TokenTypeComment,
					Start: i,
					End:   totalEnd,
				})
				i = totalEnd
				continue
			}
		}

		if strings.HasPrefix(text[i:], delims.StmtStart) {
			endIdx := strings.Index(text[i+len(delims.StmtStart):], delims.StmtEnd)
			if endIdx >= 0 {
				totalEnd := i + len(delims.StmtStart) + endIdx + len(delims.StmtEnd)
				tokens = append(tokens, TemplateToken{
					Raw:   text[i:totalEnd],
					Tag:   strings.TrimSpace(text[i+len(delims.StmtStart) : totalEnd-len(delims.StmtEnd)]),
					Type:  TokenTypeStatement,
					Start: i,
					End:   totalEnd,
				})
				i = totalEnd
				continue
			}
		}

		if strings.HasPrefix(text[i:], delims.ExprStart) {
			endIdx := strings.Index(text[i+len(delims.ExprStart):], delims.ExprEnd)
			if endIdx >= 0 {
				totalEnd := i + len(delims.ExprStart) + endIdx + len(delims.ExprEnd)
				tokens = append(tokens, TemplateToken{
					Raw:   text[i:totalEnd],
					Tag:   strings.TrimSpace(text[i+len(delims.ExprStart) : totalEnd-len(delims.ExprEnd)]),
					Type:  TokenTypeExpression,
					Start: i,
					End:   totalEnd,
				})
				i = totalEnd
				continue
			}
		}

		i++
	}

	return tokens
}

// ExtractDocumentTokens gathers all template tokens present in the document.
func (tp *TemplateProcessor) ExtractDocumentTokens(doc *Document) []TemplateToken {
	if doc == nil {
		return nil
	}
	var allTokens []TemplateToken
	for _, p := range allDocumentParagraphs(doc) {
		text := paragraphText(p)
		if text == "" {
			continue
		}
		tokens := tp.ExtractTokens(text)
		allTokens = append(allTokens, tokens...)
	}
	return allTokens
}

// RenderText evaluates template syntax against the provided context and returns the resulting string.
func (tp *TemplateProcessor) RenderText(text string, ctx map[string]interface{}) (string, error) {
	if tp.options.Engine != nil {
		return tp.options.Engine.RenderTemplate(text, ctx)
	}

	text = tp.stripComments(text)

	var err error
	text, err = tp.renderStatements(text, ctx)
	if err != nil {
		return "", err
	}

	tokens := tp.ExtractTokens(text)
	if len(tokens) == 0 {
		return text, nil
	}

	// Substitute expressions from right to left to preserve offsets
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if token.Type != TokenTypeExpression && token.Type != TokenTypeCustom {
			continue
		}

		replacement, repErr := tp.evaluateTag(token.Tag, ctx)
		if repErr != nil {
			if tp.options.Strict {
				return "", repErr
			}
			if !tp.options.RemoveUnmatched {
				continue
			}
			replacement = ""
		}
		text = text[:token.Start] + replacement + text[token.End:]
	}

	return text, nil
}

// stripComments removes {# ... #} comments from source text.
func (tp *TemplateProcessor) stripComments(text string) string {
	delims := tp.options.Delimiters
	for {
		start := strings.Index(text, delims.CommentStart)
		if start < 0 {
			break
		}
		end := strings.Index(text[start+len(delims.CommentStart):], delims.CommentEnd)
		if end < 0 {
			break
		}
		totalEnd := start + len(delims.CommentStart) + end + len(delims.CommentEnd)
		text = text[:start] + text[totalEnd:]
	}
	return text
}

// renderStatements evaluates statement blocks (for and if) in document order.
func (tp *TemplateProcessor) renderStatements(text string, ctx map[string]interface{}) (string, error) {
	delims := tp.options.Delimiters

	for {
		start := strings.Index(text, delims.StmtStart)
		if start < 0 {
			break
		}
		openTagEnd := strings.Index(text[start+len(delims.StmtStart):], delims.StmtEnd)
		if openTagEnd < 0 {
			break
		}
		tagContent := strings.TrimSpace(text[start+len(delims.StmtStart) : start+len(delims.StmtStart)+openTagEnd])

		if strings.HasPrefix(tagContent, "for ") {
			itemVar, collectionExpr, blockEnd, innerBody, err := tp.parseForStructure(text, start)
			if err != nil {
				if tp.options.Strict {
					return "", err
				}
				break
			}

			val, lErr := tp.resolvePath(ctx, collectionExpr)
			if lErr != nil && tp.options.Strict {
				return "", lErr
			}

			var items []interface{}
			if val != nil {
				rv := reflect.ValueOf(val)
				for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
					rv = rv.Elem()
				}
				if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
					items = make([]interface{}, rv.Len())
					for k := 0; k < rv.Len(); k++ {
						items[k] = rv.Index(k).Interface()
					}
				}
			}

			var renderedBuilder strings.Builder
			count := len(items)
			for idx, item := range items {
				subCtx := make(map[string]interface{}, len(ctx)+2)
				for k, v := range ctx {
					subCtx[k] = v
				}
				subCtx[itemVar] = item
				subCtx["loop"] = map[string]interface{}{
					"index":  idx,
					"index1": idx + 1,
					"first":  idx == 0,
					"last":   idx == count-1,
					"length": count,
				}

				iterText, iterErr := tp.RenderText(innerBody, subCtx)
				if iterErr != nil {
					return "", iterErr
				}
				renderedBuilder.WriteString(iterText)
			}

			text = text[:start] + renderedBuilder.String() + text[blockEnd:]
			continue
		}

		if strings.HasPrefix(tagContent, "if ") || strings.HasPrefix(tagContent, "if(") {
			blockEnd, branches, err := tp.parseIfStructure(text, start)
			if err != nil {
				if tp.options.Strict {
					return "", err
				}
				break
			}

			selectedBody := ""
			for _, b := range branches {
				if b.condition == "" {
					selectedBody = b.body
					break
				}
				condVal, cErr := tp.evaluateCondition(b.condition, ctx)
				if cErr != nil && tp.options.Strict {
					return "", cErr
				}
				if condVal {
					selectedBody = b.body
					break
				}
			}

			renderedBody, rErr := tp.RenderText(selectedBody, ctx)
			if rErr != nil {
				return "", rErr
			}

			text = text[:start] + renderedBody + text[blockEnd:]
			continue
		}

		// Advance past unrecognized statement
		nextSearch := start + len(delims.StmtStart) + openTagEnd + len(delims.StmtEnd)
		sub, err := tp.renderStatements(text[nextSearch:], ctx)
		if err != nil {
			return "", err
		}
		return text[:nextSearch] + sub, nil
	}

	return text, nil
}

type ifBranch struct {
	condition string
	body      string
}

func (tp *TemplateProcessor) parseIfStructure(text string, start int) (int, []ifBranch, error) {
	delims := tp.options.Delimiters
	curr := start
	depth := 0

	var branches []ifBranch
	var currentCondition string
	branchBodyStart := -1

	for curr < len(text) {
		tagStart := strings.Index(text[curr:], delims.StmtStart)
		if tagStart < 0 {
			return -1, nil, errors.New("unclosed if block: missing endif")
		}
		tagStart += curr
		tagEndRel := strings.Index(text[tagStart+len(delims.StmtStart):], delims.StmtEnd)
		if tagEndRel < 0 {
			return -1, nil, errors.New("unclosed statement tag")
		}
		tagEnd := tagStart + len(delims.StmtStart) + tagEndRel + len(delims.StmtEnd)
		content := strings.TrimSpace(text[tagStart+len(delims.StmtStart) : tagStart+len(delims.StmtStart)+tagEndRel])

		if strings.HasPrefix(content, "if ") || strings.HasPrefix(content, "if(") {
			if depth == 0 {
				currentCondition = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(content, "if"), "("))
				branchBodyStart = tagEnd
			}
			depth++
		} else if (strings.HasPrefix(content, "elif ") || strings.HasPrefix(content, "elif(")) && depth == 1 {
			branches = append(branches, ifBranch{
				condition: currentCondition,
				body:      text[branchBodyStart:tagStart],
			})
			currentCondition = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(content, "elif"), "("))
			branchBodyStart = tagEnd
		} else if content == "else" && depth == 1 {
			branches = append(branches, ifBranch{
				condition: currentCondition,
				body:      text[branchBodyStart:tagStart],
			})
			currentCondition = ""
			branchBodyStart = tagEnd
		} else if content == "endif" {
			depth--
			if depth == 0 {
				branches = append(branches, ifBranch{
					condition: currentCondition,
					body:      text[branchBodyStart:tagStart],
				})
				return tagEnd, branches, nil
			}
		}
		curr = tagEnd
	}

	return -1, nil, errors.New("unclosed if block: missing endif")
}

func (tp *TemplateProcessor) parseForStructure(text string, start int) (string, string, int, string, error) {
	delims := tp.options.Delimiters
	openEndRel := strings.Index(text[start+len(delims.StmtStart):], delims.StmtEnd)
	if openEndRel < 0 {
		return "", "", -1, "", errors.New("unclosed for tag")
	}
	firstTagEnd := start + len(delims.StmtStart) + openEndRel + len(delims.StmtEnd)
	header := strings.TrimSpace(text[start+len(delims.StmtStart) : start+len(delims.StmtStart)+openEndRel])
	header = strings.TrimPrefix(header, "for ")

	parts := strings.SplitN(header, " in ", 2)
	if len(parts) != 2 {
		return "", "", -1, "", fmt.Errorf("invalid for statement header: %s", header)
	}
	itemVar := strings.TrimSpace(parts[0])
	collectionExpr := strings.TrimSpace(parts[1])

	depth := 1
	curr := firstTagEnd
	for curr < len(text) {
		tagStart := strings.Index(text[curr:], delims.StmtStart)
		if tagStart < 0 {
			return "", "", -1, "", errors.New("unclosed for block: missing endfor")
		}
		tagStart += curr
		tagEndRel := strings.Index(text[tagStart+len(delims.StmtStart):], delims.StmtEnd)
		if tagEndRel < 0 {
			return "", "", -1, "", errors.New("unclosed statement tag")
		}
		tagEnd := tagStart + len(delims.StmtStart) + tagEndRel + len(delims.StmtEnd)
		content := strings.TrimSpace(text[tagStart+len(delims.StmtStart) : tagStart+len(delims.StmtStart)+tagEndRel])

		if strings.HasPrefix(content, "for ") {
			depth++
		} else if content == "endfor" {
			depth--
			if depth == 0 {
				body := text[firstTagEnd:tagStart]
				return itemVar, collectionExpr, tagEnd, body, nil
			}
		}
		curr = tagEnd
	}

	return "", "", -1, "", errors.New("unclosed for block: missing endfor")
}

// evaluateTag evaluates an expression tag with filters or delegates to a configured renderer.
func (tp *TemplateProcessor) evaluateTag(tag string, ctx map[string]interface{}) (string, error) {
	if tp.options.Renderer != nil {
		return tp.options.Renderer.Render(tag, ctx)
	}

	val, err := tp.evaluateExpression(tag, ctx)
	if err != nil {
		return "", err
	}
	if val == nil {
		return "", nil
	}
	return formatValue(val), nil
}

// evaluateExpression evaluates a variable expression, ternary operator, or pipe filters.
func (tp *TemplateProcessor) evaluateExpression(expr string, ctx map[string]interface{}) (interface{}, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return "", nil
	}

	// Check for ternary operator: condition ? trueExpr : falseExpr
	ternaryIdx := indexOutsideQuotes(expr, '?')
	if ternaryIdx > 0 {
		condStr := strings.TrimSpace(expr[:ternaryIdx])
		rest := strings.TrimSpace(expr[ternaryIdx+1:])
		colonIdx := indexOutsideQuotes(rest, ':')
		if colonIdx > 0 {
			trueExpr := strings.TrimSpace(rest[:colonIdx])
			falseExpr := strings.TrimSpace(rest[colonIdx+1:])
			condVal, err := tp.evaluateCondition(condStr, ctx)
			if err != nil && tp.options.Strict {
				return nil, err
			}
			if condVal {
				return tp.evaluateExpression(trueExpr, ctx)
			}
			return tp.evaluateExpression(falseExpr, ctx)
		}
	}

	// Parse pipe segments: baseExpr | filter1 | filter2:arg
	segments := splitPipes(expr)
	base := strings.TrimSpace(segments[0])

	var current interface{}
	var err error

	if isQuotedString(base) {
		current = unquote(base)
	} else if num, nErr := strconv.ParseInt(base, 10, 64); nErr == nil {
		current = num
	} else if fNum, fErr := strconv.ParseFloat(base, 64); fErr == nil {
		current = fNum
	} else if base == "true" {
		current = true
	} else if base == "false" {
		current = false
	} else if base == "nil" || base == "null" {
		current = nil
	} else {
		current, err = tp.resolvePath(ctx, base)
		if err != nil && tp.options.Strict {
			return nil, err
		}
	}

	// Apply filter segments in sequence
	for _, seg := range segments[1:] {
		filterName, args := parseFilterCall(seg)
		fn, exists := tp.options.Filters[strings.ToLower(filterName)]
		if !exists {
			if tp.options.Strict {
				return nil, fmt.Errorf("unknown filter: %s", filterName)
			}
			continue
		}
		current, err = fn(current, args...)
		if err != nil {
			return nil, err
		}
	}

	return current, nil
}

// evaluateCondition computes the truth value of a condition expression.
func (tp *TemplateProcessor) evaluateCondition(cond string, ctx map[string]interface{}) (bool, error) {
	cond = strings.TrimSpace(cond)
	if cond == "" {
		return false, nil
	}

	if strings.HasPrefix(cond, "not ") {
		val, err := tp.evaluateCondition(strings.TrimSpace(cond[4:]), ctx)
		return !val, err
	}
	if strings.HasPrefix(cond, "!") {
		val, err := tp.evaluateCondition(strings.TrimSpace(cond[1:]), ctx)
		return !val, err
	}

	// Check binary comparison operators
	operators := []string{"==", "!=", "<=", ">=", "<", ">"}
	for _, op := range operators {
		idx := indexOutsideQuotes(cond, op[0])
		if idx >= 0 && strings.HasPrefix(cond[idx:], op) {
			leftExpr := strings.TrimSpace(cond[:idx])
			rightExpr := strings.TrimSpace(cond[idx+len(op):])

			leftVal, lErr := tp.evaluateExpression(leftExpr, ctx)
			if lErr != nil && tp.options.Strict {
				return false, lErr
			}
			rightVal, rErr := tp.evaluateExpression(rightExpr, ctx)
			if rErr != nil && tp.options.Strict {
				return false, rErr
			}

			return compareValues(leftVal, rightVal, op), nil
		}
	}

	val, err := tp.evaluateExpression(cond, ctx)
	if err != nil {
		return false, err
	}
	return isTruthy(val), nil
}

// resolvePath traverses dot-notation or index-based property paths.
func (tp *TemplateProcessor) resolvePath(target interface{}, path string) (interface{}, error) {
	if target == nil {
		return nil, nil
	}
	path = strings.TrimSpace(path)
	if path == "" {
		return target, nil
	}

	segments := splitPathSegments(path)
	current := target

	for _, seg := range segments {
		if current == nil {
			return nil, nil
		}

		rv := reflect.ValueOf(current)
		for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
			if rv.IsNil() {
				return nil, nil
			}
			rv = rv.Elem()
		}

		switch rv.Kind() {
		case reflect.Map:
			keyVal := reflect.ValueOf(seg)
			val := rv.MapIndex(keyVal)
			if !val.IsValid() {
				// Attempt numeric lookup if key is numeric
				if idx, err := strconv.Atoi(seg); err == nil {
					intVal := rv.MapIndex(reflect.ValueOf(idx))
					if intVal.IsValid() {
						current = intVal.Interface()
						continue
					}
				}
				if tp.options.Strict {
					return nil, fmt.Errorf("map key not found: %s", seg)
				}
				return nil, nil
			}
			current = val.Interface()

		case reflect.Slice, reflect.Array:
			idx, err := strconv.Atoi(seg)
			if err != nil || idx < 0 || idx >= rv.Len() {
				if tp.options.Strict {
					return nil, fmt.Errorf("index out of range or invalid: %s", seg)
				}
				return nil, nil
			}
			current = rv.Index(idx).Interface()

		case reflect.Struct:
			// Check exact field match
			field := rv.FieldByName(seg)
			if field.IsValid() && field.CanInterface() {
				current = field.Interface()
				continue
			}

			// Case-insensitive field search
			found := false
			for i := 0; i < rv.NumField(); i++ {
				if strings.EqualFold(rv.Type().Field(i).Name, seg) {
					field = rv.Field(i)
					if field.CanInterface() {
						current = field.Interface()
						found = true
						break
					}
				}
			}
			if found {
				continue
			}

			// Check exported zero-argument method
			method := rv.MethodByName(seg)
			if !method.IsValid() {
				// Check pointer method if struct was passed by value
				if rv.CanAddr() {
					method = rv.Addr().MethodByName(seg)
				}
			}
			if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() >= 1 {
				res := method.Call(nil)
				current = res[0].Interface()
				continue
			}

			if tp.options.Strict {
				return nil, fmt.Errorf("field or method not found: %s", seg)
			}
			return nil, nil

		default:
			if tp.options.Strict {
				return nil, fmt.Errorf("cannot navigate property %q on type %T", seg, current)
			}
			return nil, nil
		}
	}

	return current, nil
}

// RenderParagraph substitutes template expressions within a paragraph across its text runs.
func (tp *TemplateProcessor) RenderParagraph(p Paragraph, ctx map[string]interface{}) error {
	runs := p.Runs()
	if len(runs) == 0 {
		return nil
	}

	originalText := paragraphText(p)
	if originalText == "" {
		return nil
	}

	delims := tp.options.Delimiters
	hasDelim := strings.Contains(originalText, delims.ExprStart) ||
		strings.Contains(originalText, delims.StmtStart) ||
		strings.Contains(originalText, delims.CommentStart)
	if !hasDelim {
		return nil
	}

	// External engine delegation
	if tp.options.Engine != nil {
		rendered, err := tp.options.Engine.RenderTemplate(originalText, ctx)
		if err != nil {
			return err
		}
		if rendered != originalText {
			setRunText(runs[0], rendered)
			for _, r := range runs[1:] {
				setRunText(r, "")
			}
		}
		return nil
	}

	// If paragraph contains statement blocks (if/for/comments), perform full text render
	hasStatements := strings.Contains(originalText, delims.StmtStart) ||
		strings.Contains(originalText, delims.CommentStart)
	if hasStatements {
		rendered, err := tp.RenderText(originalText, ctx)
		if err != nil {
			return err
		}
		setRunText(runs[0], rendered)
		for _, r := range runs[1:] {
			setRunText(r, "")
		}
		return nil
	}

	// Expression substitution preserving individual run formatting
	tokens := tp.ExtractTokens(originalText)
	if len(tokens) == 0 {
		return nil
	}

	type replacementSpan struct {
		start int
		end   int
		text  string
	}
	var spans []replacementSpan

	for _, token := range tokens {
		if token.Type != TokenTypeExpression && token.Type != TokenTypeCustom {
			continue
		}
		rep, err := tp.evaluateTag(token.Tag, ctx)
		if err != nil {
			if tp.options.Strict {
				return err
			}
			if !tp.options.RemoveUnmatched {
				continue
			}
			rep = ""
		}
		spans = append(spans, replacementSpan{
			start: token.Start,
			end:   token.End,
			text:  rep,
		})
	}

	if len(spans) == 0 {
		return nil
	}

	// Process spans from right to left to avoid shifting offsets
	sort.Slice(spans, func(i, j int) bool {
		return spans[i].start > spans[j].start
	})

	for _, span := range spans {
		applySpanToRuns(runs, span.start, span.end, span.text)
	}

	return nil
}

// RenderDocument applies template rendering across all paragraphs, tables, headers, and footers.
func (tp *TemplateProcessor) RenderDocument(doc *Document, ctx map[string]interface{}) error {
	if doc == nil {
		return errors.New("document is nil")
	}
	for _, p := range allDocumentParagraphs(doc) {
		if err := tp.RenderParagraph(p, ctx); err != nil {
			return err
		}
	}
	return nil
}

// SubstituteTokens performs direct replacement of string mappings across all runs.
func (tp *TemplateProcessor) SubstituteTokens(doc *Document, replacements map[string]string) error {
	if doc == nil {
		return errors.New("document is nil")
	}
	if len(replacements) == 0 {
		return nil
	}

	for _, p := range allDocumentParagraphs(doc) {
		runs := p.Runs()
		if len(runs) == 0 {
			continue
		}
		text := paragraphText(p)
		if text == "" {
			continue
		}

		type repMatch struct {
			start int
			end   int
			text  string
		}
		var matches []repMatch

		tokens := tp.ExtractTokens(text)
		if len(tokens) > 0 {
			// Delimited token matching
			for _, tok := range tokens {
				if val, ok := replacements[tok.Tag]; ok {
					matches = append(matches, repMatch{
						start: tok.Start,
						end:   tok.End,
						text:  val,
					})
				} else if val, ok := replacements[tok.Raw]; ok {
					matches = append(matches, repMatch{
						start: tok.Start,
						end:   tok.End,
						text:  val,
					})
				}
			}
		}

		// Plain text string replacements if no delimited tokens matched
		if len(matches) == 0 {
			// Sort keys by descending length to prevent partial prefix conflicts
			keys := make([]string, 0, len(replacements))
			for k := range replacements {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool {
				return len(keys[i]) > len(keys[j])
			})

			for _, target := range keys {
				val := replacements[target]
				idx := 0
				for {
					pos := strings.Index(text[idx:], target)
					if pos < 0 {
						break
					}
					actualStart := idx + pos
					actualEnd := actualStart + len(target)

					// Guarantee no overlap with already identified matches
					overlap := false
					for _, m := range matches {
						if !(actualEnd <= m.start || actualStart >= m.end) {
							overlap = true
							break
						}
					}
					if !overlap {
						matches = append(matches, repMatch{
							start: actualStart,
							end:   actualEnd,
							text:  val,
						})
					}
					idx = actualEnd
				}
			}
		}

		if len(matches) == 0 {
			continue
		}

		sort.Slice(matches, func(i, j int) bool {
			return matches[i].start > matches[j].start
		})

		for _, m := range matches {
			applySpanToRuns(runs, m.start, m.end, m.text)
		}
	}

	return nil
}

// RenderTemplate replaces all template tokens in the document with values from ctx.
func (doc *Document) RenderTemplate(ctx map[string]interface{}, opts ...TemplateOption) error {
	processor := NewTemplateProcessor(opts...)
	return processor.RenderDocument(doc, ctx)
}

// ExtractTemplateTokens returns all template tokens identified across the document.
func (doc *Document) ExtractTemplateTokens(opts ...TemplateOption) []TemplateToken {
	processor := NewTemplateProcessor(opts...)
	return processor.ExtractDocumentTokens(doc)
}

// SubstituteTokens performs direct key-value text substitution in all runs.
func (doc *Document) SubstituteTokens(substitutions map[string]string) error {
	processor := NewTemplateProcessor()
	return processor.SubstituteTokens(doc, substitutions)
}

// ExtractTemplateTokensFromText finds template tokens in a raw string.
func ExtractTemplateTokensFromText(text string, opts ...TemplateOption) []TemplateToken {
	processor := NewTemplateProcessor(opts...)
	return processor.ExtractTokens(text)
}

// RenderTemplateText renders a template string with the provided context.
func RenderTemplateText(text string, ctx map[string]interface{}, opts ...TemplateOption) (string, error) {
	processor := NewTemplateProcessor(opts...)
	return processor.RenderText(text, ctx)
}

// RenderTemplateFile opens a DOCX template, renders it with ctx, and saves to outputPath.
func RenderTemplateFile(inputPath string, outputPath string, ctx map[string]interface{}, opts ...TemplateOption) error {
	if inputPath == "" {
		return errors.New("input template path cannot be empty")
	}
	if outputPath == "" {
		return errors.New("output path cannot be empty")
	}

	doc, err := Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open template document %q: %w", inputPath, err)
	}

	if err := doc.RenderTemplate(ctx, opts...); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	if err := doc.SaveToFile(outputPath); err != nil {
		return fmt.Errorf("failed to save rendered document to %q: %w", outputPath, err)
	}

	return nil
}

// PongoConvert renders a DOCX template file using template variables provided in ctx.
// It loads the template from templatePath, substitutes all template expressions and statement blocks,
// and saves the rendered document to an output file.
// If outputPath is provided, it saves to that path and returns it.
// If outputPath is omitted, it generates a temporary file, saves to it, and returns the temporary file path.
func PongoConvert(templatePath string, ctx map[string]interface{}, outputPath ...string) (string, error) {
	if templatePath == "" {
		return "", errors.New("template path cannot be empty")
	}

	doc, err := Open(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to open template document %q: %w", templatePath, err)
	}

	if err := doc.RenderTemplate(ctx); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	var targetPath string
	if len(outputPath) > 0 && outputPath[0] != "" {
		targetPath = outputPath[0]
	} else {
		tmpFile, err := os.CreateTemp("", "unioffice_pongo_rendered_*.docx")
		if err != nil {
			return "", fmt.Errorf("failed to create temporary file: %w", err)
		}
		targetPath = tmpFile.Name()
		_ = tmpFile.Close()
	}

	if err := doc.SaveToFile(targetPath); err != nil {
		return "", fmt.Errorf("failed to save rendered document to %q: %w", targetPath, err)
	}

	return targetPath, nil
}

// allDocumentParagraphs gathers every paragraph from body, tables, headers, and footers.
func allDocumentParagraphs(doc *Document) []Paragraph {
	if doc == nil {
		return nil
	}
	var paras []Paragraph
	paras = append(paras, doc.Paragraphs()...)

	for _, h := range doc.Headers() {
		paras = append(paras, h.Paragraphs()...)
		for _, tbl := range h.Tables() {
			for _, row := range tbl.Rows() {
				for _, cell := range row.Cells() {
					paras = append(paras, cell.Paragraphs()...)
				}
			}
		}
	}

	for _, f := range doc.Footers() {
		paras = append(paras, f.Paragraphs()...)
		for _, tbl := range f.Tables() {
			for _, row := range tbl.Rows() {
				for _, cell := range row.Cells() {
					paras = append(paras, cell.Paragraphs()...)
				}
			}
		}
	}

	return paras
}

// paragraphText concatenates text across all runs in a paragraph.
func paragraphText(p Paragraph) string {
	runs := p.Runs()
	if len(runs) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, r := range runs {
		sb.WriteString(r.Text())
	}
	return sb.String()
}

// setRunText sets text content for a run while supporting multiline break insertions.
func setRunText(r Run, text string) {
	if !strings.Contains(text, "\n") {
		// Update existing T element when present
		for _, ic := range r._cbefb.EG_RunInnerContent {
			if ic.RunInnerContentChoice != nil && ic.RunInnerContentChoice.T != nil {
				ic.RunInnerContentChoice.T.Content = text
				if unioffice.NeedsSpacePreserve(text) {
					preserve := "preserve"
					ic.RunInnerContentChoice.T.SpaceAttr = &preserve
				} else {
					ic.RunInnerContentChoice.T.SpaceAttr = nil
				}
				return
			}
		}
	}

	r.ClearContent()
	if text == "" {
		return
	}

	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if i > 0 {
			r.AddBreak()
		}
		if line != "" {
			r.AddText(line)
		}
	}
}

// applySpanToRuns replaces the character span [start, end) across the runs slice.
func applySpanToRuns(runs []Run, start, end int, repText string) {
	if len(runs) == 0 || start >= end {
		return
	}

	type runOffset struct {
		start int
		end   int
		text  string
	}
	offsets := make([]runOffset, len(runs))
	curr := 0
	for i, r := range runs {
		txt := r.Text()
		offsets[i] = runOffset{
			start: curr,
			end:   curr + len(txt),
			text:  txt,
		}
		curr += len(txt)
	}

	if start >= curr {
		return
	}
	if end > curr {
		end = curr
	}

	startRun := -1
	endRun := -1

	for i, o := range offsets {
		if startRun == -1 && o.start <= start && start < o.end {
			startRun = i
		}
		if endRun == -1 && o.start <= (end-1) && (end-1) < o.end {
			endRun = i
		}
	}

	if startRun == -1 || endRun == -1 || startRun > endRun {
		return
	}

	startOff := start - offsets[startRun].start
	endOff := end - offsets[endRun].start

	if startRun == endRun {
		orig := offsets[startRun].text
		newTxt := orig[:startOff] + repText + orig[endOff:]
		setRunText(runs[startRun], newTxt)
		return
	}

	// Span across multiple runs
	origStart := offsets[startRun].text
	newStart := origStart[:startOff] + repText
	setRunText(runs[startRun], newStart)

	for mid := startRun + 1; mid < endRun; mid++ {
		setRunText(runs[mid], "")
	}

	origEnd := offsets[endRun].text
	newEnd := origEnd[endOff:]
	setRunText(runs[endRun], newEnd)
}

func defaultTemplateFilters() map[string]TemplateFilterFunc {
	return map[string]TemplateFilterFunc{
		"upper": func(val interface{}, args ...string) (interface{}, error) {
			return strings.ToUpper(formatValue(val)), nil
		},
		"lower": func(val interface{}, args ...string) (interface{}, error) {
			return strings.ToLower(formatValue(val)), nil
		},
		"trim": func(val interface{}, args ...string) (interface{}, error) {
			return strings.TrimSpace(formatValue(val)), nil
		},
		"title": func(val interface{}, args ...string) (interface{}, error) {
			words := strings.Fields(formatValue(val))
			for i, w := range words {
				if len(w) > 0 {
					r := []rune(w)
					r[0] = unicode.ToUpper(r[0])
					words[i] = string(r)
				}
			}
			return strings.Join(words, " "), nil
		},
		"capitalize": func(val interface{}, args ...string) (interface{}, error) {
			s := formatValue(val)
			if s == "" {
				return "", nil
			}
			r := []rune(s)
			r[0] = unicode.ToUpper(r[0])
			for i := 1; i < len(r); i++ {
				r[i] = unicode.ToLower(r[i])
			}
			return string(r), nil
		},
		"length": func(val interface{}, args ...string) (interface{}, error) {
			if val == nil {
				return 0, nil
			}
			rv := reflect.ValueOf(val)
			for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
				rv = rv.Elem()
			}
			switch rv.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
				return rv.Len(), nil
			default:
				return 0, nil
			}
		},
		"count": func(val interface{}, args ...string) (interface{}, error) {
			if val == nil {
				return 0, nil
			}
			rv := reflect.ValueOf(val)
			for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
				rv = rv.Elem()
			}
			switch rv.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
				return rv.Len(), nil
			default:
				return 0, nil
			}
		},
		"default": func(val interface{}, args ...string) (interface{}, error) {
			if !isTruthy(val) && len(args) > 0 {
				return args[0], nil
			}
			return val, nil
		},
		"int": func(val interface{}, args ...string) (interface{}, error) {
			s := formatValue(val)
			n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
			if err != nil {
				return int64(0), nil
			}
			return n, nil
		},
		"float": func(val interface{}, args ...string) (interface{}, error) {
			s := formatValue(val)
			f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
			if err != nil {
				return float64(0), nil
			}
			return f, nil
		},
		"json": func(val interface{}, args ...string) (interface{}, error) {
			b, err := json.Marshal(val)
			if err != nil {
				return "", err
			}
			return string(b), nil
		},
		"first": func(val interface{}, args ...string) (interface{}, error) {
			if val == nil {
				return nil, nil
			}
			rv := reflect.ValueOf(val)
			for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
				rv = rv.Elem()
			}
			if (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array) && rv.Len() > 0 {
				return rv.Index(0).Interface(), nil
			}
			if rv.Kind() == reflect.String && rv.Len() > 0 {
				return string([]rune(rv.String())[:1]), nil
			}
			return nil, nil
		},
		"last": func(val interface{}, args ...string) (interface{}, error) {
			if val == nil {
				return nil, nil
			}
			rv := reflect.ValueOf(val)
			for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
				rv = rv.Elem()
			}
			if (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array) && rv.Len() > 0 {
				return rv.Index(rv.Len() - 1).Interface(), nil
			}
			if rv.Kind() == reflect.String && rv.Len() > 0 {
				runes := []rune(rv.String())
				return string(runes[len(runes)-1:]), nil
			}
			return nil, nil
		},
		"reverse": func(val interface{}, args ...string) (interface{}, error) {
			if val == nil {
				return nil, nil
			}
			rv := reflect.ValueOf(val)
			for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
				rv = rv.Elem()
			}
			if rv.Kind() == reflect.String {
				runes := []rune(rv.String())
				for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
					runes[i], runes[j] = runes[j], runes[i]
				}
				return string(runes), nil
			}
			if rv.Kind() == reflect.Slice {
				out := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Len())
				for i := 0; i < rv.Len(); i++ {
					out.Index(i).Set(rv.Index(rv.Len() - 1 - i))
				}
				return out.Interface(), nil
			}
			return val, nil
		},
	}
}

func formatValue(val interface{}) string {
	if val == nil {
		return ""
	}
	switch v := val.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func isTruthy(val interface{}) bool {
	if val == nil {
		return false
	}
	switch v := val.(type) {
	case bool:
		return v
	case string:
		return len(strings.TrimSpace(v)) > 0 && v != "false" && v != "0"
	case int:
		return v != 0
	case int64:
		return v != 0
	case float64:
		return v != 0.0
	default:
		rv := reflect.ValueOf(val)
		for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
			if rv.IsNil() {
				return false
			}
			rv = rv.Elem()
		}
		switch rv.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map:
			return rv.Len() > 0
		default:
			return true
		}
	}
}

func compareValues(left, right interface{}, op string) bool {
	lf, isLeftNum := toFloat(left)
	rf, isRightNum := toFloat(right)

	if isLeftNum && isRightNum {
		switch op {
		case "==":
			return lf == rf
		case "!=":
			return lf != rf
		case "<":
			return lf < rf
		case "<=":
			return lf <= rf
		case ">":
			return lf > rf
		case ">=":
			return lf >= rf
		}
	}

	ls := formatValue(left)
	rs := formatValue(right)
	switch op {
	case "==":
		return ls == rs
	case "!=":
		return ls != rs
	case "<":
		return ls < rs
	case "<=":
		return ls <= rs
	case ">":
		return ls > rs
	case ">=":
		return ls >= rs
	default:
		return false
	}
}

func toFloat(val interface{}) (float64, bool) {
	if val == nil {
		return 0, false
	}
	switch v := val.(type) {
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case string:
		f, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		return f, err == nil
	default:
		return 0, false
	}
}

func splitPipes(expr string) []string {
	var segments []string
	var sb strings.Builder
	inQuote := rune(0)

	for _, ch := range expr {
		if inQuote != 0 {
			if ch == inQuote {
				inQuote = 0
			}
			sb.WriteRune(ch)
		} else if ch == '\'' || ch == '"' {
			inQuote = ch
			sb.WriteRune(ch)
		} else if ch == '|' {
			segments = append(segments, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(ch)
		}
	}
	if sb.Len() > 0 {
		segments = append(segments, sb.String())
	}
	return segments
}

func parseFilterCall(seg string) (string, []string) {
	seg = strings.TrimSpace(seg)
	colonIdx := indexOutsideQuotes(seg, ':')
	if colonIdx > 0 {
		name := strings.TrimSpace(seg[:colonIdx])
		rawArgs := strings.TrimSpace(seg[colonIdx+1:])
		args := splitFilterArgs(rawArgs)
		return name, args
	}

	parenIdx := strings.Index(seg, "(")
	if parenIdx > 0 && strings.HasSuffix(seg, ")") {
		name := strings.TrimSpace(seg[:parenIdx])
		rawArgs := strings.TrimSpace(seg[parenIdx+1 : len(seg)-1])
		args := splitFilterArgs(rawArgs)
		return name, args
	}

	return seg, nil
}

func splitFilterArgs(s string) []string {
	var args []string
	var sb strings.Builder
	inQuote := rune(0)

	for _, ch := range s {
		if inQuote != 0 {
			if ch == inQuote {
				inQuote = 0
			}
			sb.WriteRune(ch)
		} else if ch == '\'' || ch == '"' {
			inQuote = ch
			sb.WriteRune(ch)
		} else if ch == ',' {
			arg := strings.TrimSpace(sb.String())
			if isQuotedString(arg) {
				arg = unquote(arg)
			}
			args = append(args, arg)
			sb.Reset()
		} else {
			sb.WriteRune(ch)
		}
	}
	if sb.Len() > 0 {
		arg := strings.TrimSpace(sb.String())
		if isQuotedString(arg) {
			arg = unquote(arg)
		}
		args = append(args, arg)
	}
	return args
}

func splitPathSegments(path string) []string {
	var segments []string
	var sb strings.Builder
	inBracket := false

	for _, ch := range path {
		if ch == '[' {
			if sb.Len() > 0 {
				segments = append(segments, sb.String())
				sb.Reset()
			}
			inBracket = true
		} else if ch == ']' {
			inBracket = false
			seg := sb.String()
			if isQuotedString(seg) {
				seg = unquote(seg)
			}
			if seg != "" {
				segments = append(segments, seg)
			}
			sb.Reset()
		} else if ch == '.' && !inBracket {
			if sb.Len() > 0 {
				segments = append(segments, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteRune(ch)
		}
	}
	if sb.Len() > 0 {
		segments = append(segments, sb.String())
	}
	return segments
}

func indexOutsideQuotes(s string, target byte) int {
	inQuote := rune(0)
	for i, ch := range s {
		if inQuote != 0 {
			if ch == inQuote {
				inQuote = 0
			}
		} else if ch == '\'' || ch == '"' {
			inQuote = ch
		} else if byte(ch) == target {
			return i
		}
	}
	return -1
}

func isQuotedString(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return true
		}
	}
	return false
}

func unquote(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
