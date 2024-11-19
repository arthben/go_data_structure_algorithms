package latihan

import (
	"sort"
)

func penjadwalan(room int, occupancies [][]int) int {
	sort.SliceStable(occupancies, func(i, j int) bool {
		return occupancies[i][0] < occupancies[j][0]
	})

	count := 0
	prevEnd := occupancies[0][1]
	for _, occupan := range occupancies[1:] {
		if prevEnd < occupan[0] {
			count++
			prevEnd = occupan[1]
		}
	}

	if count == 0 && room > prevEnd {
		count++
	}

	return count
}
