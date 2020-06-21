package index

import "log"

type illegalIndex struct{}

func (i illegalIndex) Get(value string) (int, error) {
	log.Fatal("illegal index cannot be used")
	return 0, nil
}
