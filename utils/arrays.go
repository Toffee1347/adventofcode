package utils

import "sort"

func ArrayGetBoundValues(input []int, count int, getBottomValues bool) []int {
	values := input

	if getBottomValues {
		sort.Ints(values)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(values)))
	}

	return values[0:count]
}

func ArraySum(input []int) (total int) {
	for _, max := range input {
		total += max
	}

	return
}

func ArrayContains(input []string, target string) bool {
	for _, part := range input {
		if part == target {
			return true
		}
	}

	return false
}
