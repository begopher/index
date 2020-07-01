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

import "fmt"

type strictImmutableIndex struct {
	ltr       bool
	indexOf   string
	occur     uint
	exclusive bool
}

func (index strictImmutableIndex) Get(value string) (int, error) {
	if value == "" {
		return -1, fmt.Errorf("Value cannot be empty")
	}
	if index.ltr {
		return ltrIndex(value, index.indexOf, index.occur, index.exclusive)
	} else {
		return rtlIndex(value, index.indexOf, index.occur, index.exclusive)
	}
}

// Strict is responsible for extracting the correct index number based on a given configuration.
//
// indexOf is the substring that searching algorightem is searching for.
// occur allows the searching algorithem to skip occur-1 match.
// exclusive determines which index will be returned:
// 	if ture  it will return the index of the first rune of the match
// 	if false it will return the index of the last  rune of the match + 1
// ltr determines the direction of the searching algorithem
func Strict(indexOf string, occur uint, exclusive bool, ltr bool) (Index, error) {
	if indexOf == "" {
		errMsg := "indexOf cannot be empty"
		return illegalIndex{}, fmt.Errorf(errMsg)
	}
	if occur < 1 {
		errMsg := "occur value must be bigger than zero"
		return illegalIndex{}, fmt.Errorf(errMsg)
	}
	return strictImmutableIndex{
		indexOf:   indexOf,
		occur:     occur,
		exclusive: exclusive,
		ltr:       ltr,
	}, nil
}
