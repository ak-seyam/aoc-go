package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/A-Siam/aoc-go/utils"
)

func Solution_pt1(inputPath string) {
	res := 0
	utils.GetInput(inputPath, func(line string) {
		paris := strings.Split(line, ",")
		fPair, sPair := paris[0], paris[1]
		if fullyContained(fPair, sPair) {
			res += 1
		}
	})
	fmt.Println("res is", res)
}

func fullyContained(fPair, sPair string) bool {
	fBoundries := strings.Split(fPair, "-")
	fStart, _ := strconv.Atoi(fBoundries[0])
	fEnd, _ := strconv.Atoi(fBoundries[1])
	sBoundries := strings.Split(sPair, "-")
	sStart, _ := strconv.Atoi(sBoundries[0])
	sEnd, _ := strconv.Atoi(sBoundries[1])
	fInS := fStart >= sStart && fEnd <= sEnd
	if fInS {
		return true
	}
	SInf := sStart >= fStart && sEnd <= fEnd
	if SInf {
		return true
	}
	return false
}
