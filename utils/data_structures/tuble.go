package data_structures

import "fmt"

type Int int

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

type Tuple[LEFT, RIGHT fmt.Stringer] struct {
	left  LEFT
	right RIGHT
}

func NewTuple[LEFT, RIGHT fmt.Stringer](left LEFT, right RIGHT) Tuple[LEFT, RIGHT] {
	return Tuple[LEFT, RIGHT]{
		left:  left,
		right: right,
	}
}

func (t *Tuple[LEFT, RIGHT]) GetLeft() LEFT {
	return t.left
}

func (t *Tuple[LEFT, RIGHT]) GetRight() RIGHT {
	return t.right
}

func (t *Tuple[LEFT, RIGHT]) String() string {
	return fmt.Sprintf("%s,%s", t.left, t.right)
}

func (t *Tuple[LEFT, RIGHT]) ToComparable() string {
	return t.String()
}
