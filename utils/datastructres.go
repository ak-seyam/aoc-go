package utils

import (
	"errors"
)

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

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Stack[T interface{}] struct {
	topIdx    int
	innerList []T
}

func NewStack[T interface{}](inner []T) Stack[T] {
	return Stack[T]{
		innerList: inner,
		topIdx:    len(inner) - 1,
	}
}

/*
return reference to element or nil if error
*/
func (s *Stack[T]) Pop() (*T, error) {
	if s.topIdx == -1 {
		return nil, errors.New("empty stack")
	}
	top := s.innerList[s.topIdx]
	s.topIdx -= 1
	return &top, nil
}

/*
add element to stack
*/
func (s *Stack[T]) Push(e T) {
	if s.topIdx+1 == len(s.innerList) {
		s.innerList = append(s.innerList, e)
	} else {
		s.innerList[s.topIdx+1] = e
	}
	s.topIdx += 1
}

func (s *Stack[T]) Reverse() Stack[T] {
	var result Stack[T] = NewStack([]T{})
	for s.HasNext() {
		nxt, _ := s.Pop()
		result.Push(*nxt)
	}
	return result
}

func (s *Stack[T]) HasNext() bool {
	return s.topIdx != -1
}
