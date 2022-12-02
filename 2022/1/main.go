package main

import (
	"fmt"
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

func main() {
	data := utils.GetAsset("2022/1.txt")
	maxValues := getMax(data, 3)

	total := 0
	for _, max := range maxValues {
		total += max
	}

	utils.FinalOutput(maxValues[0], total)
}

func getMax(data string, topValues int) []int {
	splitData := strings.Split(data, "\r\n\r\n")
	if topValues > len(splitData) {
		fmt.Println("topValues can not be more than the count of elves")
		return []int{}
	}

	maxValues := []int{}
	for i := 0; i < topValues; i++ {
		maxValues = append(maxValues, 0)
	}

	for _, elfData := range splitData {
		total := 0
		singleElfData := strings.Split(elfData, "\r\n")

		for _, count := range singleElfData {
			total += utils.StrToInt(count)
		}

		for i, max := range maxValues {
			if total > max {
				maxValues = append(maxValues[:i+1], maxValues[i:len(maxValues)-1]...)
				maxValues[i] = total
				break
			}
		}
	}

	return maxValues
}
