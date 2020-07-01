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

// Package index extracts inclusive or exclusive index from a given string based on a preconfiguration.
package index

// Index is the interface implemented by an onject that can extract index from a given string
//
// Get extracts the correct index number from a given value
type Index interface {
	Get(value string) int
}

func findAllIndexes(value, sub string) []int {
	valueRunes := []rune(value)
	subRunes := []rune(sub)
	var indexes []int
	if len(subRunes) != 0 && len(valueRunes) != 0 {
		for i := 0; i < len(valueRunes); i++ {
			if i+len(subRunes) <= len(valueRunes) {
				match := true
				var counter int
				for j := 0; j < len(subRunes); j++ {
					if valueRunes[i+j] != subRunes[j] {
						match = false
						counter += j
						break
					}
				}
				if match {
					indexes = append(indexes, i)
					i += len(subRunes) - 1 // when loop finish i will increase by 1
				} else {
					i += counter
				}
			}
		}
	}
	return indexes
}
func ltrIndex(name, sub string, occur uint, exclusive bool) int {
	matches := findAllIndexes(name, sub)
	if int(occur) <= len(matches) {
		i := matches[occur-1]
		if !exclusive {
			i += len([]rune(sub))
		}
		return i

	}
	return -1
}
func rtlIndex(name, sub string, occur uint, exclusive bool) int {
	matches := findAllIndexes(name, sub)
	if int(occur) <= len(matches) {
		i := matches[len(matches)-int(occur)]
		if !exclusive {
			i += len([]rune(sub))
		}
		return i

	}
	return -1
}
