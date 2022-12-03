package main

import (
	"github.com/Toffee1347/adventofcode/challenges/year2015"
	"github.com/Toffee1347/adventofcode/challenges/year2022"
)

type challenge = func(string) [2]int

var challenges map[int](map[int]challenge) = map[int]map[int]challenge{
	2015: {
		1: year2015.Day01,
		2: year2015.Day02,
		3: year2015.Day03,
		4: year2015.Day04,
		5: year2015.Day05,
		6: year2015.Day06,
	},
	2022: {
		1: year2022.Day01,
		2: year2022.Day02,
		3: year2022.Day03,
	},
}
