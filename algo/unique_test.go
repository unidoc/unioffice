package algo_test

import (
	"bytes"
	"testing"

	"baliance.com/gooxml/algo"
)

func TestUnique(t *testing.T) {
	td := []struct {
		Inp []byte
		Exp []byte
	}{
		{[]byte{1, 2, 2, 3, 4, 4, 4, 4, 5, 6, 6},
			[]byte{1, 2, 3, 4, 5, 6}},
		{[]byte{1, 2, 2, 3, 4},
			[]byte{1, 2, 3, 4}},
		{[]byte{2, 3},
			[]byte{2, 3}},
		{[]byte{2, 2},
			[]byte{2}},
		{[]byte{2},
			[]byte{2}},
	}

	for _, tc := range td {
		newLen := algo.Unique(tc.Inp, func(i, j int) bool {
			return tc.Inp[i] == tc.Inp[j]
		})
		tc.Inp = tc.Inp[0:newLen]
		if !bytes.Equal(tc.Inp, tc.Exp) {
			t.Errorf("got %v, expted %v", tc.Inp, tc.Exp)
		}
	}
}
