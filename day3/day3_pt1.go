package day3

import (
	"fmt"

	"github.com/A-Siam/aoc-go/utils"
)

func Solution_pt1(inputPath string) {
	sum := 0
	utils.GetInput(inputPath, func(line string) {
		resSet := make(map[byte]bool)
		for i := 0; i < len(line)/2; i++ {
			resSet[line[i]] = true
		}
		for i := len(line) / 2; i < len(line); i++ {
			_, exist := resSet[line[i]]
			if exist {
				if line[i] >= 97 && line[i] <= 122 {
					sum += int(line[i]) - 96
				} else {
					sum += int(line[i]) - 38
				}
				break
			}
		}
	})
	fmt.Println("sum is", sum)
}
