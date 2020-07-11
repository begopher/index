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

	"github.com/begopher/index"
)

func TestExactCreatation(t *testing.T) {
	testCases := []struct {
		number uint
	}{
		{number: 0},
		{number: 1},
		{number: 2},
		{number: 3},
		{number: 4},
		{number: 5},
		{number: 6},
		{number: 7},
		{number: 8},
		{number: 9},
		{number: 11},
		{number: 100},
	}
	for _, test := range testCases {
		index.Exact(test.number)
	}
}
func TestGetInExactStrategy(t *testing.T) {
	testCases := []struct {
		number   uint
		value    string
		expected int
	}{
		{number: 0, expected: 0, value: "any non-empty value"},
		{number: 1, expected: 1, value: "filetargz"},
		{number: 2, expected: 2, value: "filetar.gz"},
		{number: 3, expected: 3, value: "file.tar.gz"},
		{number: 4, expected: 4, value: "begopher"},
		{number: 5, expected: 5, value: "begopher.com"},
		{number: 6, expected: 6, value: "Abdulrahman A Alsaedi"},
		// exact
		{number: 0, expected: 0, value: ""},
		{number: 1, expected: 1, value: "a"},
		{number: 2, expected: 2, value: "ab"},
		{number: 3, expected: 3, value: "abc"},
		{number: 4, expected: 4, value: "abcd"},
	}
	for _, test := range testCases {
		index := index.Exact(test.number)
		got, err := index.Get(test.value)
		if err != nil {
			t.Errorf("Valid value should not return error")
		}
		expected := test.expected
		if expected != got {
			t.Errorf("Expected index [%v], got [%v]", expected, got)
		}

	}
}
func TestGetInExactStrategy_Error(t *testing.T) {
	testCases := []struct {
		number   uint
		value    string
		expected int
	}{
		{number: 1, expected: -1, value: ""},
		{number: 2, expected: -1, value: ""},
		{number: 3, expected: -1, value: ""},
		{number: 4, expected: -1, value: ""},
		{number: 5, expected: -1, value: ""},
		{number: 6, expected: -1, value: ""},
		{number: 2, expected: -1, value: "a"},
		{number: 3, expected: -1, value: "ab"},
		{number: 4, expected: -1, value: "abc"},
		{number: 5, expected: -1, value: "abcd"},
	}
	for _, test := range testCases {
		index := index.Exact(test.number)
		got, err := index.Get(test.value)
		if err == nil {
			t.Errorf("invalid value should return error")
		}
		expected := test.expected
		if got != expected {
			t.Errorf("Expected index [%v], got [%v]", expected, got)
		}
	}
}
