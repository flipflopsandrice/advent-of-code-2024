package one

import (
	"slices"
)

func distanceSum(list1 []int, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0
	for i, v := range list1 {
		v2 := list2[i]

		if v < v2 {
			distance += v2 - v
		} else {
			distance += v - v2
		}
	}

	return distance
}
