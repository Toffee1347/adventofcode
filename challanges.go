package main

import (
	"github.com/Toffee1347/adventofcode/challenges/year2015"
	"github.com/Toffee1347/adventofcode/challenges/year2022"
)

type challenge = func(string) [2]int

var challenges map[int](map[int]challenge) = map[int]map[int]challenge{
	2015: {
		1: year2015.Day01,
	},
	2022: {
		1: year2022.Day01,
		2: year2022.Day02,
	},
}
