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
		i, err := index.Get(anyValue)
		if err == nil {
			t.Errorf("Using get in illegalIndex must return error")
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
