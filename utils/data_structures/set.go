package data_structures

type Set[T comparable] map[T]bool

func (s *Set[T]) Add(e T) {
	(*s)[e] = true
}

func NewSet[T comparable](col []T) Set[T] {
	res := make(Set[T])
	for _, e := range col {
		res[e] = true
	}
	return res
}

func SetIntersect[T comparable](set1, set2 Set[T]) Set[T] {
	res := make(Set[T])
	for e := range set1 {
		if set2[e] {
			res[e] = true
		}
	}
	return res
}
