package index_test

import (
	"github.com/begopher/index"
	"testing"
)

func TestGetInZeroStrategy(t *testing.T) {
	testCases := zeroTestTable()
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
