package document

import (
	"testing"
)

func TestParseMergeField(t *testing.T) {
	td := []struct {
		instruction string
		exp         mergeFieldInfo
	}{
		{
			`MERGEFIELD FirstName \f " "`,
			mergeFieldInfo{
				fieldName:  "FirstName",
				followText: " ",
			},
		},
		{
			`MERGEFIELD FirstName`,
			mergeFieldInfo{
				fieldName: "FirstName",
			},
		},
		{
			`MERGEFIELD FirstName \f "after" \b "before"`,
			mergeFieldInfo{
				fieldName:  "FirstName",
				followText: "after",
				beforeText: "before",
			},
		},
		{
			`MERGEFIELD FirstName \* Upper`,
			mergeFieldInfo{
				fieldName: "FirstName",
				upper:     true,
			},
		},
		{
			`MERGEFIELD FirstName \* Lower`,
			mergeFieldInfo{
				fieldName: "FirstName",
				lower:     true,
			},
		},
		{
			`MERGEFIELD FirstName \* Caps`,
			mergeFieldInfo{
				fieldName: "FirstName",
				caps:      true,
			},
		},
		{
			`MERGEFIELD FirstName \* FirstCap`,
			mergeFieldInfo{
				fieldName: "FirstName",
				firstCap:  true,
			},
		},
	}
	for _, tc := range td {
		t.Run(tc.instruction, func(t *testing.T) {
			parsed := parseField(tc.instruction)
			if parsed.fieldName != tc.exp.fieldName {
				t.Errorf("expected fieldName = %s, got %s", tc.exp.fieldName, parsed.fieldName)
			}
			if parsed.followText != tc.exp.followText {
				t.Errorf("expected followText = %s, got %s", tc.exp.followText, parsed.followText)
			}
			if parsed.beforeText != tc.exp.beforeText {
				t.Errorf("expected beforeText = %s, got %s", tc.exp.beforeText, parsed.beforeText)
			}
			if parsed.upper != tc.exp.upper {
				t.Errorf("expected upper = %v, got %v", tc.exp.upper, parsed.upper)
			}
			if parsed.lower != tc.exp.lower {
				t.Errorf("expected lower = %v, got %v", tc.exp.lower, parsed.lower)
			}
			if parsed.caps != tc.exp.caps {
				t.Errorf("expected caps = %v, got %v", tc.exp.caps, parsed.caps)
			}
			if parsed.firstCap != tc.exp.firstCap {
				t.Errorf("expected firstCap = %v, got %v", tc.exp.firstCap, parsed.firstCap)
			}
		})
	}
}
