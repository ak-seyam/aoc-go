package day2

import (
	"fmt"
	"strings"

	"github.com/A-Siam/aoc-go/utils"
)

func Solution_pt1(inputPath string) {
	res := 0
	utils.GetInput(inputPath, func(line string) {
		inputs := strings.Split(line, " ")
		opInp := inputs[0]
		yourInp := inputs[1]
		res += score(conv(opInp), conv(yourInp))
	})
	fmt.Println("your score=", res)
}
