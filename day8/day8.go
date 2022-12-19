package day8

import (
	"strconv"
	"strings"

	"github.com/A-Siam/aoc-go/utils"
)

func Solution(inputPath string, afterTrendMatrices AfterTrendMatrices) {
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

	afterTrendMatrices(mat, matTopBottom, matBottomTop, matLeftRight, matRightLeft)
}
