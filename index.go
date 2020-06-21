package index

import (
	"fmt"
)

type Index interface {
	Get(value string) (int, error)
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
func ltrStrictIndex(name, sub string, occur uint, exclusive bool) (int, error) {
	matches := findAllIndexes(name, sub)
	if int(occur) <= len(matches) {
		i := matches[occur-1]
		if !exclusive {
			i += len([]rune(sub))
		}
		return i, nil

	}
	return -1, fmt.Errorf("No match found")
}
func rtlStrictIndex(name, sub string, occur uint, exclusive bool) (int, error) {
	matches := findAllIndexes(name, sub)
	if int(occur) <= len(matches) {
		i := matches[len(matches)-int(occur)]
		if !exclusive {
			i += len([]rune(sub))
		}
		return i, nil

	}
	return -1, fmt.Errorf("No match found")
}
