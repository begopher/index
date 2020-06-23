package index

import "fmt"

type illegalIndex struct{}

func (i illegalIndex) Get(value string) (int, error) {
	return -1, fmt.Errorf("illegal ")
}
