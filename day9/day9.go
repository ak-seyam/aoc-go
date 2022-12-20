package day9

import (
	"fmt"

	"github.com/A-Siam/aoc-go/utils"
	"github.com/A-Siam/aoc-go/utils/data_structures"
)

func Solution(inputPath string) {
	utils.GetInput(inputPath, func(line string) {
		var x data_structures.Int = 0
		t := data_structures.NewTuple(x, x)
		fmt.Println(t)
	})
}
