package day1

import (
	"fmt"
	"strconv"

	"github.com/A-Siam/aoc-go/utils"
)

func Day1(inputPath string) {
	var max int64 = 0
	var res int64 = 0
	utils.GetInput(inputPath, func(line string) {
		n, err := strconv.Atoi(line)
		if err == nil {
			res += int64(n)
		} else {
			max = utils.Max(max, res)
			res = 0
		}
	})
	fmt.Println("Highest is", max)
}
