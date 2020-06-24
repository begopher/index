package index_test

import (
	"testing"

	"github.com/begopher/index"
)

func TestGetInLengthStrategy(t *testing.T) {
	testCases := lengthTestTable()
	length := index.Length()
	for _, test := range testCases {

		expected := test.expected
		got, err := length.Get(test.value)
		if err != nil {
			t.Error("Last Index shout never return error")
		}
		if got != expected {
			t.Errorf("Expected index is [%d], got [%d]", expected, got)
		}
	}
}
