package day8

func buildMatrix(i1, i2 int) [][]int {
	res := make([][]int, i1)
	for i := 0; i < i1; i++ {
		res[i] = make([]int, i2)
	}
	return res
}

type AfterTrendMatrices func(
	mat,
	matTopBottom,
	matBottomTop,
	matLeftRight,
	matRightLeft [][]int,
)
