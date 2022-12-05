package utils

type number interface {
	int | float32 | float64 | int64 | int32
}

func Max[T number](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}
