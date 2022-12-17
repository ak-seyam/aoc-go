package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/A-Siam/aoc-go/utils"
)

func SolutionPt1(puzPath string, movPath string, puzSize int) {
	var world []utils.Stack[string]

	// init world and topIdx
	for i := 0; i < puzSize; i++ {
		world = append(world, utils.NewStack([]string{}))
	}

	utils.GetInput(puzPath, func(line string) {
		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				worldIdx := (i - 1) / 4
				world[worldIdx].Push(string(line[i]))
			}
		}
	})
	// reverse the stacks
	for i := 0; i < len(world); i++ {
		world[i] = world[i].Reverse()
	}
	utils.GetInput(movPath, func(line string) {
		splits := strings.Split(line, " ")
		amount, _ := strconv.Atoi(splits[1])
		from, _ := strconv.Atoi(splits[3])
		to, _ := strconv.Atoi(splits[5])
		from0Idx := from - 1
		to0Idx := to - 1
		for i := 0; i < amount; i++ {
			val, _ := world[from0Idx].Pop()
			if val != nil {
				world[to0Idx].Push(*val)
			}
		}
	})
	for i := 0; i < len(world); i++ {
		val, _ := world[i].Pop()
		if val != nil {
			fmt.Print(*val)
		}
	}
	fmt.Println()
}
