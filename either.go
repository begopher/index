package index

type immutableEitherStrategy struct {
	first Index
	last  Index
}

func Either(first, last Index) Index {
	return immutableEitherStrategy{
		first: first,
		last:  last,
	}
}
func (index immutableEitherStrategy) Get(value string) (int, error) {
	i, err := index.first.Get(value)
	if err != nil {
		i, err = index.last.Get(value)
	}
	return i, err
}
