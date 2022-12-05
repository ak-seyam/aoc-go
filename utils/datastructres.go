package utils

type Set[T comparable] map[T]bool

func ToSet[T comparable](col []T) Set[T] {
	res := make(Set[T])
	for _, e := range col {
		res[e] = true
	}
	return res
}
