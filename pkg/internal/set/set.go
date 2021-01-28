package set

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return StringSet{}
}
