package utils

type Set[T comparable] map[T]bool

func ToSet[T comparable](col []T) Set[T] {
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
