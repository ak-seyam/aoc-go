package day6

import (
	"fmt"

	"github.com/A-Siam/aoc-go/utils"
)

func Solution(inputPath string, windowSize int) {
	utils.GetInput(inputPath, func(line string) {
		// we have only a single line
		window := make(map[string]int)
		for i := 0; i < windowSize; i++ {
			toTheWindow := string(line[i])
			addToWindow(toTheWindow, window)
		}
		if len(window) == windowSize {
			fmt.Println(windowSize)
			return
		}
		for i := windowSize; i < len(line); i++ {
			outOfWindowStr := string(line[i-windowSize])
			window[outOfWindowStr]--
			if window[outOfWindowStr] == 0 {
				delete(window, outOfWindowStr)
			}
			addToWindow(string(line[i]), window)
			if len(window) == windowSize {
				fmt.Println(i + 1)
				return
			}
		}
	})
}

func addToWindow(toTheWindow string, window map[string]int) {
	_, found := window[toTheWindow]
	if !found {
		window[toTheWindow] = 1
	} else {
		window[toTheWindow]++
	}
}
