// Copyright 2018 FoxyUtils ehf. All rights reserved.

package document

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/unidoc/unioffice/schema/soo/wml"
)

type mergeFieldInfo struct {
	fieldName  string
	followText string
	beforeText string
	upper      bool
	lower      bool
	firstCap   bool
	caps       bool

	// normal merge fields
	para                   Paragraph
	begIdx, sepIdx, endIdx int

	// simple fields
	pc       *wml.EG_PContent
	isSimple bool
}

func parseField(mf string) mergeFieldInfo {
	//fields := strings.Fields(mf)
	fields := []string{}
	sb := bytes.Buffer{}
	pq := -1
	for i, c := range mf {
		switch c {
		case ' ':
			if sb.Len() != 0 {
				fields = append(fields, sb.String())
			}
			sb.Reset()
		case '"':
			if pq != -1 {
				// this doesn't handled nested quotes, but it's good enough
				// for the field info I've seen
				fields = append(fields, mf[pq+1:i])
				pq = -1
			} else {
				pq = i
			}
		default:
			sb.WriteRune(c)
		}
	}
	if sb.Len() != 0 {
		fields = append(fields, sb.String())
	}

	field := mergeFieldInfo{}
	for i := 0; i < len(fields)-1; i++ {
		k := fields[i]
		switch k {
		case "MERGEFIELD":
			field.fieldName = fields[i+1]
			i++
		case "\\f":
			field.followText = fields[i+1]
			i++
		case "\\b":
			field.beforeText = fields[i+1]
			i++
		case "\\*":
			switch fields[i+1] {
			case "Upper":
				field.upper = true
			case "Lower":
				field.lower = true
			case "Caps":
				field.caps = true
			case "FirstCap":
				field.firstCap = true
			}
			i++
		}
	}
	return field
}

// MergeFields returns the list of all mail merge fields found in the document.
func (d Document) mergeFields() []mergeFieldInfo {
	paragraphs := []Paragraph{}
	mf := []mergeFieldInfo{}
	for _, t := range d.Tables() {
		for _, r := range t.Rows() {
			for _, c := range r.Cells() {
				paragraphs = append(paragraphs, c.Paragraphs()...)
			}
		}
	}
	paragraphs = append(paragraphs, d.Paragraphs()...)
	for _, p := range paragraphs {
		runs := p.Runs()
		begIdx := -1
		sepIdx := -1
		endIdx := -1
		mergeField := mergeFieldInfo{}
		for _, pc := range p.x.EG_PContent {
			for _, fs := range pc.FldSimple {
				if strings.Contains(fs.InstrAttr, "MERGEFIELD") {
					f := parseField(fs.InstrAttr)
					f.isSimple = true
					f.para = p
					f.pc = pc
					mf = append(mf, f)
				}
			}
		}
		for i := 0; i < len(runs); i++ {
			r := runs[i]
			for _, ic := range r.X().EG_RunInnerContent {
				if ic.FldChar != nil {
					switch ic.FldChar.FldCharTypeAttr {
					case wml.ST_FldCharTypeBegin:
						begIdx = i
					case wml.ST_FldCharTypeSeparate:
						sepIdx = i
					case wml.ST_FldCharTypeEnd:
						endIdx = i
						if mergeField.fieldName != "" {
							mergeField.para = p
							mergeField.begIdx = begIdx
							mergeField.endIdx = endIdx
							mergeField.sepIdx = sepIdx
							mf = append(mf, mergeField)
						}
						begIdx = -1
						sepIdx = -1
						endIdx = -1
						mergeField = mergeFieldInfo{}
					}
				} else if ic.InstrText != nil && strings.Contains(ic.InstrText.Content, "MERGEFIELD") {
					if begIdx != -1 && endIdx == -1 {
						mergeField = parseField(ic.InstrText.Content)
					}
				}
			}
		}
	}
	return mf
}

// MergeFields returns the list of all mail merge fields found in the document.
func (d Document) MergeFields() []string {
	flds := map[string]struct{}{}
	for _, mf := range d.mergeFields() {
		flds[mf.fieldName] = struct{}{}
	}
	ret := []string{}
	for k := range flds {
		ret = append(ret, k)
	}
	return ret
}

// MailMerge finds mail merge fields and replaces them with the text provided.  It also removes
// the mail merge source info from the document settings.
func (d *Document) MailMerge(mergeContent map[string]string) {
	fields := d.mergeFields()
	remove := map[Paragraph][]Run{}
	for _, v := range fields {
		repText, ok := mergeContent[v.fieldName]
		if ok {
			if v.upper {
				repText = strings.ToUpper(repText)
			} else if v.lower {
				repText = strings.ToLower(repText)
			} else if v.caps {
				repText = strings.Title(repText)
			} else if v.firstCap {
				sb := bytes.Buffer{}
				for i, v := range repText {
					if i == 0 {
						sb.WriteRune(unicode.ToUpper(v))
					} else {
						sb.WriteRune(v)
					}
				}
				repText = sb.String()
			}

			if repText != "" && v.beforeText != "" {
				repText = v.beforeText + repText
			}
			if repText != "" && v.followText != "" {
				repText = repText + v.followText
			}
		}

		if v.isSimple {
			// simple field replacement, just promote the run
			if len(v.pc.FldSimple) == 1 &&
				len(v.pc.FldSimple[0].EG_PContent) == 1 &&
				len(v.pc.FldSimple[0].EG_PContent[0].EG_ContentRunContent) == 1 {
				rc := &wml.EG_ContentRunContent{}
				rc.R = v.pc.FldSimple[0].EG_PContent[0].EG_ContentRunContent[0].R
				v.pc.FldSimple = nil
				run := Run{d, rc.R}
				run.ClearContent()
				run.AddText(repText)
				v.pc.EG_ContentRunContent = append(v.pc.EG_ContentRunContent, rc)
			}
		} else {
			// non-simple so we'll remove the extra stuff
			runs := v.para.Runs()
			for i := v.begIdx; i <= v.endIdx; i++ {
				if i == v.sepIdx+1 {
					runs[i].ClearContent()
					runs[i].AddText(repText)
				} else {
					remove[v.para] = append(remove[v.para], runs[i])
				}
			}
		}
	}
	// remove any of the mail merge field runs except for the one immediately after 'separate'
	for p, runs := range remove {
		for _, r := range runs {
			p.RemoveRun(r)
		}
	}

	d.Settings.RemoveMailMerge()
}
