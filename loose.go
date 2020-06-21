package index

type looseImmutableIndex struct {
	ltr       bool
	indexOf   string
	occur     uint
	exclusive bool
}

func (looseImmutableIndex) Get(name string) (int, error) {
	return 0, nil
}

func LoseIndex(indexOf string, occur uint, exclusive bool, ltr bool) Index {
	return looseImmutableIndex{
		indexOf:   indexOf,
		occur:     occur,
		exclusive: exclusive,
		ltr:       ltr,
	}
}
