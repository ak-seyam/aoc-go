package day3

import (
	"fmt"

	"github.com/A-Siam/aoc-go/utils"
	"github.com/A-Siam/aoc-go/utils/data_structures"
)

func Solution_pt2(inputPath string) {
	batch := make([]data_structures.Set[byte], 0)
	res := 0
	utils.GetInput(inputPath, func(line string) {
		if len(batch) < 2 {
			batch = append(batch, data_structures.NewSet([]byte(line)))
		} else {
			batch = append(batch, data_structures.NewSet([]byte(line)))
			res += int(getCommonBadgeValue(batch))
			batch = make([]data_structures.Set[byte], 0)
		}
	})
	fmt.Println("result", res)
}

func getCommonBadgeValue(batch []data_structures.Set[byte]) (res byte) {
	var resSet = data_structures.SetIntersect(data_structures.SetIntersect(batch[0], batch[1]), batch[2])
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
