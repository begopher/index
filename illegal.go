package index

import "fmt"

type illegalIndex struct {
	message string
}

func (i illegalIndex) Get(value string) (int, error) {
	return -1, fmt.Errorf("Illegal index")
}
