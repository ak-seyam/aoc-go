package day3

import (
	"fmt"

	"github.com/A-Siam/aoc-go/utils"
)

func Solution_pt2(inputPath string) {
	batch := make([]utils.Set[byte], 0)
	res := 0
	utils.GetInput(inputPath, func(line string) {
		if len(batch) < 2 {
			batch = append(batch, utils.ToSet([]byte(line)))
		} else {
			batch = append(batch, utils.ToSet([]byte(line)))
			res += int(getCommonBadgeValue(batch))
			batch = make([]utils.Set[byte], 0)
		}
	})
	fmt.Println("result", res)
}

func getCommonBadgeValue(batch []utils.Set[byte]) (res byte) {
	var resSet = utils.SetIntersect(utils.SetIntersect(batch[0], batch[1]), batch[2])
	for k := range resSet {
		if k >= 97 && k <= 122 {
			res = k - 96
		} else {
			res = k - 38
		}
		break
	}
	return
}
