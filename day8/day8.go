package day8

import "github.com/A-Siam/aoc-go/utils"

type State struct {
	HTop    int // H = highest
	HBottom int
	HLeft   int
	HRight  int
}

type OnPointSelected func(point int, state State)

func Solution(inputPath string, onPointSelected OnPointSelected) {
	utils.GetInput(inputPath, func(line string) {

	})
}
