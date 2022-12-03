package year2022

import "strings"

func Day03(input string) [2]int {
	rucksacks := strings.Split(input, "\r\r\n")

	for _, rucksack := range rucksacks {
		size := len(rucksack)
		firstComp := rucksack[0:size/2]
		secondComp := rucksack[size/2:size]
	}
}
