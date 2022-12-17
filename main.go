package main

import (
	"fmt"

	"github.com/A-Siam/aoc-go/day1"
	"github.com/A-Siam/aoc-go/day2"
	"github.com/A-Siam/aoc-go/day3"
	"github.com/A-Siam/aoc-go/day4"
	"github.com/A-Siam/aoc-go/day5"
	"github.com/A-Siam/aoc-go/day6"
)

func main() {
	inputPrefix := "./input"
	fmt.Println("Day 1 Pt.1")
	day1.Solution_pt1(inputPrefix + "/day1")
	fmt.Println("Day 1 Pt.2")
	day1.Solution_pt2(inputPrefix + "/day1")
	fmt.Println("Day 2 Pt.1")
	day2.Solution_pt1(inputPrefix + "/day2")
	fmt.Println("Day 2 Pt.2")
	day2.Solution_pt2(inputPrefix + "/day2")
	fmt.Println("Day 3 Pt.1")
	day3.Solution_pt1(inputPrefix + "/day3")
	fmt.Println("Day 3 Pt.2")
	day3.Solution_pt2(inputPrefix + "/day3")
	fmt.Println("Day 4 Pt.1")
	day4.Solution_pt1(inputPrefix + "/day4")
	fmt.Println("Day 4 Pt.2")
	day4.Solution_pt2(inputPrefix + "/day4")
	fmt.Println("Day 5 Pt.1")
	day5.SolutionPt1(inputPrefix+"/day5/complx/puz", inputPrefix+"/day5/complx/steps", 9)
	fmt.Println("Day 5 Pt.2")
	day5.SolutionPt2(inputPrefix+"/day5/complx/puz", inputPrefix+"/day5/complx/steps", 9)
	fmt.Println("Day 6 Pt.1")
	day6.Solution(inputPrefix+"/day6", 4)
	fmt.Println("Day 6 Pt.2")
	day6.Solution(inputPrefix+"/day6", 14)
}
