package day3

import "github.com/A-Siam/aoc-go/utils"

func Solution_pt2(inputPath string) {
	batch := make([]utils.Set[byte], 0)
	res := 0
	utils.GetInput(inputPath, func(line string) {
		if len(batch) < 3 {
			batch = append(batch, utils.ToSet[byte]([]byte(line)))
		} else if len(batch) == 3 {
			res += int(getCommonBadgeValue(batch))
			batch = make([]utils.Set[byte], 0)
		}
	})
}

func getCommonBadgeValue(batch []utils.Set[byte]) byte {
	// TODO: fill this function
}
