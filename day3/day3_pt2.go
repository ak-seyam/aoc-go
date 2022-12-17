package day3

import (
	"fmt"

	"github.com/A-Siam/aoc-go/utils"
	"github.com/A-Siam/aoc-go/utils/datastructures"
)

func Solution_pt2(inputPath string) {
	batch := make([]datastructures.Set[byte], 0)
	res := 0
	utils.GetInput(inputPath, func(line string) {
		if len(batch) < 2 {
			batch = append(batch, datastructures.NewSet([]byte(line)))
		} else {
			batch = append(batch, datastructures.NewSet([]byte(line)))
			res += int(getCommonBadgeValue(batch))
			batch = make([]datastructures.Set[byte], 0)
		}
	})
	fmt.Println("result", res)
}

func getCommonBadgeValue(batch []datastructures.Set[byte]) (res byte) {
	var resSet = datastructures.SetIntersect(datastructures.SetIntersect(batch[0], batch[1]), batch[2])
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
