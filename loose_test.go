package index_test

import (
	"testing"

	"github.com/begopher/index"
)

func TestLooseCreatation(t *testing.T) {
	testCases := indexCreationTestTable()
	for _, test := range testCases {
		_, err := index.Loose(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err != nil {
			t.Errorf("Valid value of Loose index should not return error")
		}
	}
}
func TestLooseCreatation_errors(t *testing.T) {
	testCases := indexCreationErrorTestTable()
	for _, test := range testCases {
		_, err := index.Loose(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err == nil {
			t.Error("invalid indexOf value for Loose index should return error")
		} else {
			expected := test.errMsg
			got := err.Error()
			if expected != got {
				t.Errorf("Expected error message [%v], got [%v]", expected, got)
			}
		}
	}
}
func TestGetInLooseStrategy(t *testing.T) {
	testCases := strictIndexTestTable()
	testCases = append(looseIndexTestTable(), testCases...)
	for _, test := range testCases {
		index, err := index.Loose(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err != nil {
			t.Errorf("Valid value of Loose index should not return error")
		}
		got, err := index.Get(test.value)
		expected := test.expected
		if err != nil {
			t.Error("Valid value for Loose index should not return error")
		} else {
			if got != expected {
				t.Errorf("Expected index is [%d], got [%d]", expected, got)
			}
		}
	}
}
func TestGetInLooseStrategy_errors(t *testing.T) {
	testCases := emptyValueTestTable()
	for _, test := range testCases {
		index, _ := index.Loose(test.indexOf, test.occur, test.exclusive, test.ltr)
		_, err := index.Get(test.value)
		if err == nil {
			t.Error("Pass empty value must return error")
		} else {
			got := err.Error()
			expected := test.errMsg
			if got != expected {
				t.Errorf("Expected error message [%v],got [%v]", expected, got)
			}
		}
	}
}
