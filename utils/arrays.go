package utils

import "sort"

func GetBoundValues(input []int, count int, getBottomValues bool) []int {
	values := input

	if getBottomValues {
		sort.Ints(values)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(values)))
	}

	return values[0:count]
}

func Sum(input []int) (total int) {
	for _, max := range input {
		total += max
	}

	return
}
