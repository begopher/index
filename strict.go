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
		return ltrStrictIndex(value, index.indexOf, index.occur, index.exclusive)
	} else {
		return rtlStrictIndex(value, index.indexOf, index.occur, index.exclusive)
	}
}
func Strict(indexOf string, occur uint, exclusive bool, ltr bool) (Index, error) {
	if indexOf == "" {
		return illegalIndex{}, fmt.Errorf("indexOf cannot be empty")
	}
	if occur < 1 {
		return illegalIndex{}, fmt.Errorf("occur value must be bigger than zero")
	}
	return strictImmutableIndex{
		indexOf:   indexOf,
		occur:     occur,
		exclusive: exclusive,
		ltr:       ltr,
	}, nil
}
