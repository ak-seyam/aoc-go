package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/A-Siam/aoc-go/utils"
)

type State struct {
	HTop    int // H = highest
	HBottom int
	HLeft   int
	HRight  int
}

type OnPointSelected func(point int, state State)

func Solution(inputPath string) {
	mat := [][]int{}
	utils.GetInput(inputPath, func(line string) {
		numsStr := strings.Split(line, "")
		lineSlice := []int{}
		for _, n := range numsStr {
			nn, _ := strconv.Atoi(n)
			lineSlice = append(lineSlice, nn)
		}
		mat = append(mat, lineSlice)
	})
	res := 2*len(mat) + 2*len(mat[0]) - 4
	// create matTopBottom
	matTopBottom, matBottomTop, matLeftRight, matRightLeft :=
		buildMatrix(len(mat), len(mat[0])),
		buildMatrix(len(mat), len(mat[0])),
		buildMatrix(len(mat), len(mat[0])),
		buildMatrix(len(mat), len(mat[0]))

	// top bottom
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if i == 0 || i == len(mat)-1 {
				matTopBottom[i][j] = mat[i][j]
				continue
			}
			val := mat[i][j]
			if val > matTopBottom[i-1][j] {
				matTopBottom[i][j] = val
			} else {
				matTopBottom[i][j] = matTopBottom[i-1][j]
			}
		}
	}

	// bottom top
	for i := len(mat) - 1; i > -1; i-- {
		for j := 0; j < len(mat[0]); j++ {
			if i == 0 || i == len(mat)-1 {
				matBottomTop[i][j] = mat[i][j]
				continue
			}
			val := mat[i][j]
			if val > matBottomTop[i+1][j] {
				matBottomTop[i][j] = val
			} else {
				matBottomTop[i][j] = matBottomTop[i+1][j]
			}
		}
	}

	// left right
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if j == 0 || j == len(mat)-1 {
				matLeftRight[i][j] = mat[i][j]
				continue
			}
			val := mat[i][j]
			if val > matLeftRight[i][j-1] {
				matLeftRight[i][j] = val
			} else {
				matLeftRight[i][j] = matLeftRight[i][j-1]
			}
		}
	}

	// right left
	for i := 0; i < len(mat); i++ {
		for j := len(mat[0]) - 1; j > -1; j-- {
			if j == 0 || j == len(mat)-1 {
				matRightLeft[i][j] = mat[i][j]
				continue
			}
			val := mat[i][j]
			if val > matRightLeft[i][j+1] {
				matRightLeft[i][j] = val
			} else {
				matRightLeft[i][j] = matRightLeft[i][j+1]
			}
		}
	}

	for i := 1; i < len(mat)-1; i++ {
		for j := 1; j < len(mat[0])-1; j++ {
			left := matLeftRight[i][j-1]
			right := matRightLeft[i][j+1]
			top := matTopBottom[i-1][j]
			bottom := matBottomTop[i+1][j]
			val := mat[i][j]
			if val > left || val > right || val > top || val > bottom {
				res++
				continue
			}
		}
	}
	fmt.Println("number of visible trees", res)
}

func buildMatrix(i1, i2 int) [][]int {
	res := make([][]int, i1)
	for i := 0; i < i1; i++ {
		res[i] = make([]int, i2)
	}
	return res
}
