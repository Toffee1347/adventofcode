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

func ArrayContains[T comparable](input []T, target T) bool {
	for _, part := range input {
		if part == target {
			return true
		}
	}

	return false
}

func ArrayGetValueCount[T comparable](input []T, target T) (count int) {
	for _, value := range input {
		if target == value {
			count++
		}
	}

	return count
}

func ArrayValuesAreUnique(input []string) bool {
	for _, value := range input {
		if ArrayGetValueCount(input, value) > 1 {
			return false
		}
	}

	return true
}
