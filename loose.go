package index

import "fmt"

type looseImmutableIndex struct {
	ltr       bool
	indexOf   string
	occur     uint
	exclusive bool
}

func (index looseImmutableIndex) Get(value string) (int, error) {
	if value == "" {
		return -1, fmt.Errorf("Value cannot be empty")
	}

	var i int
	var err error
	if index.ltr {
		i, err = ltrIndex(value, index.indexOf, index.occur, index.exclusive)
	} else {
		i, err = rtlIndex(value, index.indexOf, index.occur, index.exclusive)
	}

	if err != nil {
		return len(value), nil
	}
	return i, err
}

func Loose(indexOf string, occur uint, exclusive bool, ltr bool) (Index, error) {
	if indexOf == "" {
		return illegalIndex{}, fmt.Errorf("indexOf cannot be empty")
	}
	if occur < 1 {
		return illegalIndex{}, fmt.Errorf("occur value must be bigger than zero")
	}
	return looseImmutableIndex{
		indexOf:   indexOf,
		occur:     occur,
		exclusive: exclusive,
		ltr:       ltr,
	}, nil
}
