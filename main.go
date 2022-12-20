package main

import (
	"fmt"
	"math"

	"github.com/A-Siam/aoc-go/day1"
	"github.com/A-Siam/aoc-go/day2"
	"github.com/A-Siam/aoc-go/day3"
	"github.com/A-Siam/aoc-go/day4"
	"github.com/A-Siam/aoc-go/day5"
	"github.com/A-Siam/aoc-go/day6"
	"github.com/A-Siam/aoc-go/day7"
	"github.com/A-Siam/aoc-go/day8"
	"github.com/A-Siam/aoc-go/utils"
)

func main() {
	inputPrefix := "./input"
	fmt.Println("Day 1 Pt.1")
	day1.Solution_pt1(inputPrefix + "/day1")
	fmt.Println("Day 1 Pt.2")
	day1.Solution_pt2(inputPrefix + "/day1")
	fmt.Println("Day 2 Pt.1")
	day2.Solution_pt1(inputPrefix + "/day2")
	fmt.Println("Day 2 Pt.2")
	day2.Solution_pt2(inputPrefix + "/day2")
	fmt.Println("Day 3 Pt.1")
	day3.Solution_pt1(inputPrefix + "/day3")
	fmt.Println("Day 3 Pt.2")
	day3.Solution_pt2(inputPrefix + "/day3")
	fmt.Println("Day 4 Pt.1")
	day4.Solution_pt1(inputPrefix + "/day4")
	fmt.Println("Day 4 Pt.2")
	day4.Solution_pt2(inputPrefix + "/day4")
	fmt.Println("Day 5 Pt.1")
	day5.SolutionPt1(inputPrefix+"/day5/complx/puz", inputPrefix+"/day5/complx/steps", 9)
	fmt.Println("Day 5 Pt.2")
	day5.SolutionPt2(inputPrefix+"/day5/complx/puz", inputPrefix+"/day5/complx/steps", 9)
	fmt.Println("Day 6 Pt.1")
	day6.Solution(inputPrefix+"/day6", 4)
	fmt.Println("Day 6 Pt.2")
	day6.Solution(inputPrefix+"/day6", 14)
	fmt.Println("Day 7 Pt.1")
	var sum int64 = 0
	day7.Solution(inputPrefix+"/day7", func(f day7.File, root day7.File) bool {
		return f.Size <= 100000
	}, func(f day7.File) {
		sum += f.Size
	})
	fmt.Println("sum is", sum)
	fmt.Println("Day 7 Pt.2")
	var requiredSize int64 = 30000000
	var smallest int64 = 70000000
	day7.Solution(inputPrefix+"/day7",
		func(f day7.File, root day7.File) bool {
			availableSize := 70000000 - root.Size
			return f.Size+availableSize > requiredSize
		},
		func(f day7.File) {
			smallest = int64(math.Min(float64(f.Size), float64(smallest)))
		},
	)
	fmt.Println("Solution", smallest)
	fmt.Println("Day 8 Pt.1")
	day8.Solution(inputPrefix+"/day8", func(mat, matTopBottom, matBottomTop, matLeftRight, matRightLeft [][]int) {
		res := 2*len(mat) + 2*len(mat[0]) - 4
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
	})
	fmt.Println("Day 8 Pt.2")
	day8.Solution(inputPrefix+"/day8", func(mat, matTopBottom, matBottomTop, matLeftRight, matRightLeft [][]int) {
		res := 0
		for i := 1; i < len(mat)-1; i++ {
			for j := 1; j < len(mat[0])-1; j++ {
				// scan top
				valueFromTop := 0
				valueFromBottom := 0
				valueFromRight := 0
				valueFromLeft := 0
				for k := i + 1; k < len(mat); k++ {
					valueFromBottom++
					if mat[i][j] <= mat[k][j] {
						break
					}
				}
				for k := i - 1; k > -1; k-- {
					valueFromTop++
					if mat[i][j] <= mat[k][j] {
						break
					}
				}
				for k := j - 1; k > -1; k-- {
					valueFromLeft++
					if mat[i][j] <= mat[i][k] {
						break
					}
				}
				for k := j + 1; k < len(mat); k++ {
					valueFromRight++
					if mat[i][j] <= mat[i][k] {
						break
					}
				}
				scenicScore := valueFromBottom * valueFromLeft * valueFromRight * valueFromTop
				res = utils.Max(res, scenicScore)
			}
		}

		fmt.Println("result is", res)
	})

}
