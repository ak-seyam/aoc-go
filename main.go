package main

import (
	"fmt"

	"github.com/A-Siam/aoc-go/day1"
	"github.com/A-Siam/aoc-go/day2"
)

func main() {
	inputPrefix := "./input"
	fmt.Println("Day 1")
	day1.Solution(inputPrefix + "/day1")
	fmt.Println("Day 2 Pt.1")
	day2.Solution_pt1(inputPrefix + "/day2")
	fmt.Println("Day 2 Pt.2")
	day2.Solution_pt2(inputPrefix + "/day2")
}
