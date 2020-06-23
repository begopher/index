package index_test

import (
	"testing"

	"github.com/begopher/index"
)

func TestGetInEitherStrategyWithStrictAndLength(t *testing.T) {
	testCases := strictIndexTestTable()
	testCases = append(testCases, lengthTestTable()...)
	for _, test := range testCases {
		strict, err := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err != nil {
			t.Errorf("should not return error [%v]", err.Error())
		}
		length := index.Length()
		either := index.Either(strict, length)
		got, _ := either.Get(test.value)
		expected := test.expected
		if got != expected {
			t.Errorf("Expected value [%v], got [%v]", expected, got)
		}
	}
}
func TestGetInEitherStrategyWithStrictAndZero(t *testing.T) {
	testCases := strictIndexTestTable()
	testCases = append(testCases, zeroTestTable()...)
	for _, test := range testCases {
		strict, err := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err != nil {
			t.Errorf("should not return error [%v]", err.Error())
		}
		zero := index.Zero()
		either := index.Either(strict, zero)
		got, _ := either.Get(test.value)
		expected := test.expected
		if got != expected {
			t.Errorf("Expected value [%v], got [%v]", expected, got)
		}
	}

}
