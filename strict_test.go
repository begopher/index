package index_test

import (
	"github.com/begopher/index"
	"testing"
)

func TestStrictCreatation(t *testing.T) {
	testCases := indexCreationTestTable()
	for _, test := range testCases {
		_, err := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err != nil {
			t.Errorf("Valid value of Strict index should not return error")
		}
	}
}
func TestStrictCreatation_errors(t *testing.T) {
	testCases := indexCreationErrorTestTable()
	for _, test := range testCases {
		_, err := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err == nil {
			t.Error("invalid indexOf value for Strict index should return error")
		} else {
			expected := test.errMsg
			got := err.Error()
			if expected != got {
				t.Errorf("Expected error message [%v], got [%v]", expected, got)
			}
		}
	}
}
func TestGetInStrictStrategy(t *testing.T) {
	testCases := strictIndexTestTable()
	for _, test := range testCases {
		index, err := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err != nil {
			t.Errorf("Valid value of StrictIndex shout not return error")
		}
		got, err := index.Get(test.value)
		expected := test.expected
		if err != nil {
			t.Error("Valid value for StrictIndex should not return error")
		} else {
			if got != expected {
				t.Errorf("Expected index is [%d], got [%d]", expected, got)
			}
		}
	}
}
func TestGetInStrictStrategy_errors(t *testing.T) {
	testCases := emptyValueTestTable()
	for _, test := range testCases {
		index, _ := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
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
	testCases = indexNotFoundTestTable()
	for _, test := range testCases {
		index, _ := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		_, err := index.Get(test.value)
		if err == nil {
			t.Error("No match should return error")
		} else {
			expected := test.errMsg
			got := err.Error()
			if expected != got {
				t.Errorf("Expected error message [%v], got [%v]", expected, got)
			}
		}
	}
}
