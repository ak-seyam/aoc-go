package day2

import (
	"fmt"
	"strings"

	"github.com/A-Siam/aoc-go/utils"
)

func Solution_pt2(inputPath string) {
	res := 0
	utils.GetInput(inputPath, func(line string) {
		inputs := strings.Split(line, " ")
		opInp := conv(inputs[0])
		yourInp := getSuitableShape(opInp, convState(inputs[1]))
		res += score(opInp, yourInp)
	})
	fmt.Println("output should be", res)
}

func convState(inp string) state {
	switch inp {
	case "X":
		return lose
	case "Y":
		return draw
	default:
		return win
	}
}

func getSuitableShape(opInp shape, result state) shape {
	if opInp == paper {
		switch result {
		case win:
			return scissors
		case draw:
			return paper
		default:
			return rock
		}
	} else if opInp == rock {
		switch result {
		case win:
			return paper
		case draw:
			return rock
		default:
			return scissors
		}
	} else { // scissors
		switch result {
		case win:
			return rock
		case draw:
			return scissors
		default:
			return paper
		}
	}
}
