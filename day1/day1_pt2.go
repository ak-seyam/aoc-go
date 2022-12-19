package day1

import (
	"fmt"
	"strconv"

	"github.com/A-Siam/aoc-go/utils"
	"github.com/A-Siam/aoc-go/utils/data_structures"
)

func Solution_pt2(inputPath string) {
	var top3 = [3]int64{0, 0, 0}
	var res int64 = 0
	utils.GetInput(inputPath, func(line string) {
		n, err := strconv.Atoi(line)
		if err == nil {
			res += int64(n)
		} else {
			updateTop3(res, &top3)
			res = 0
		}
	})
	updateTop3(res, &top3)
	fmt.Println("Highest is", top3[0]+top3[1]+top3[2])
}

func updateTop3[T data_structures.Number](res T, top3 *[3]T) {
	if res > top3[0] {
		top3[2] = top3[1]
		top3[1] = top3[0]
		top3[0] = res
	} else if res > top3[1] {
		top3[2] = top3[1]
		top3[1] = res
	} else if res > top3[2] {
		top3[2] = res
	}
}
