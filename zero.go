package index

type zero struct{}

func Zero() Index {
	return zero{}
}
func (z zero) Get(value string) (int, error) { return 0, nil }
