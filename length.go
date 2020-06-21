package index

type length struct{}

func (z length) Get(value string) (int, error) { return int(len(value)), nil }

func Length() Index {
	return length{}
}
