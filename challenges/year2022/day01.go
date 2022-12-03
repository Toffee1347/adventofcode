package year2022

import (
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

func Day01(input string) [2]int {
	parsedInput := parseFoodData(input)
	maxValues := utils.ArrayGetBoundValues(parsedInput, 3, false)
	total := utils.ArraySum(maxValues)

	return [2]int{maxValues[0], total}
}

func parseFoodData(input string) (totals []int) {
	splitInput := strings.Split(input, "\r\n\r\n")

	for _, splitData := range splitInput {
		convertedData := utils.SplitStrToInt(splitData, "\r\n")
		totals = append(totals, utils.ArraySum(convertedData))
	}

	return
}
