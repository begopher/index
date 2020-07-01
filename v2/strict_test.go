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
	"github.com/begopher/index/v2"
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
		index, err := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err == nil {
			t.Error("invalid indexOf value for Strict index should return error")
		} else {
			expected := test.errMsg
			got := err.Error()
			if expected != got {
				t.Errorf("Expected error message [%v], got [%v]", expected, got)
			}
		}
		i := index.Get(anyValue)
		if i != -1 {
			t.Errorf("Using get in illegalIndex must return -1")
		}
		if i != -1 {
			t.Error("Using illegalIndex must return -1")
		}
	}
}
func TestGetInStrictStrategy(t *testing.T) {
	testCases := strictIndexTestTable()
	for _, test := range testCases {
		index, err := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		if err != nil {
			t.Errorf("Valid value of StrictIndex should not return error")
		}
		got := index.Get(test.value)
		expected := test.expected
		if got < 0 {
			t.Error("Valid value for StrictIndex should not return value less than 0")
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
		got := index.Get(test.value)
		if got != -1 {
			t.Error("Pass empty value must return -1")
		}
	}
	testCases = indexNotFoundTestTable()
	for _, test := range testCases {
		index, _ := index.Strict(test.indexOf, test.occur, test.exclusive, test.ltr)
		got := index.Get(test.value)
		if got != -1 {
			t.Error("No match should return error")
		}
	}
}
