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

func TestGetInLengthStrategy(t *testing.T) {
	testCases := lengthTestTable()
	length := index.Length()
	for _, test := range testCases {

		expected := test.expected
		got := length.Get(test.value)
		if got < 0 {
			t.Error("Length Index should never return value less than 0")
		}
		if got != expected {
			t.Errorf("Expected index is [%d], got [%d]", expected, got)
		}
	}
}
