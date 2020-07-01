// Index extracts index number from a given string based on a preconfiguration.
// Copyright (C) 2020  Abdulrahman A Alsaedi
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package index_test

import (
	"testing"

	"github.com/begopher/index/v2"
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
		got := either.Get(test.value)
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
		got := either.Get(test.value)
		expected := test.expected
		if got != expected {
			t.Errorf("Expected value [%v], got [%v]", expected, got)
		}
	}

}
