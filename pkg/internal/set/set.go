package set

type stringSet map[string]struct{}

func NewStringSet() stringSet {
	return stringSet{}
}
