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

package index

type immutableExactStrategy struct {
	number uint
}

func (index immutableExactStrategy) Get(value string) int {
	runes := []rune(value)
	if int(index.number) > len(runes) {
		return -1
	}
	return int(index.number)
}

// Exact is used when index number has known ahead
//
// number will be returned if value's length is bigger or equals to number
// else -1 will be returned
func Exact(number uint) Index {
	return immutableExactStrategy{number: number}
}
