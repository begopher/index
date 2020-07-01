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
	"fmt"

	"github.com/begopher/index/v2"
)

func Example_toFileExtension() {
	firstIndex := index.Zero()
	dot, _ := index.Strict(".", 1, true, true)
	lastIndex := index.Either(dot, index.Length())

	fileName := "name.tar.gz"
	fi := firstIndex.Get(fileName)
	li := lastIndex.Get(fileName)
	fmt.Println(fileName[fi:li])
	// Output:
	// name

}
func Example_ifExtensionNotExistThenFileLength() {
	firstIndex := index.Zero()
	dot, _ := index.Strict(".", 1, true, true)
	lastIndex := index.Either(dot, index.Length())

	fileName := "nametargz"
	fi := firstIndex.Get(fileName)
	li := lastIndex.Get(fileName)
	if fi > -1 && li > -1 {
		fmt.Println(fileName[fi:li])
		// Output:
		// nametargz
	}
}
func Example_findSecondGopherLTR() {
	firstIndex, _ := index.Strict("gopher", 2, true, true)
	lastIndex, _ := index.Strict("gopher", 2, false, true)

	gophers := "gopher gopher gopher gopher"
	fi := firstIndex.Get(gophers)
	li := lastIndex.Get(gophers)
	if fi > -1 && li > -1 {
		fmt.Println(gophers[fi:li], fi, li)
		// Output:
		// gopher 7 13
	}
}
func Example_findSecondGopherRTL() {
	firstIndex, _ := index.Strict("gopher", 2, true, false)
	lastIndex, _ := index.Strict("gopher", 2, false, false)

	gophers := "gopher gopher gopher gopher"
	fi := firstIndex.Get(gophers)
	li := lastIndex.Get(gophers)
	if fi > -1 && li > -1 {
		fmt.Println(gophers[fi:li], fi, li)
		// Output:
		// gopher 14 20
	}
}
