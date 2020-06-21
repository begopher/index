package index_test

import (
	"github.com/begopher/index"
	"testing"
)

func TestZeroIndex(t *testing.T) {
	testCases := []struct {
		value    string
		expected int
	}{
		{value: "", expected: 0},
		{value: "begopher", expected: 0},
		{value: "begopher.tar", expected: 0},
		{value: "begopher.tar.gz", expected: 0},
	}
	indexZero := index.Zero()
	for _, test := range testCases {

		expected := test.expected
		got, err := indexZero.Get(test.value)
		if err != nil {
			t.Error("Last Index shout never return error")
		}

		if got != expected {
			t.Errorf("Expected index is [%d], got [%d]", expected, got)
		}
	}
}
