package main

import (
	"github.com/Toffee1347/adventofcode/utils"
)

func main() {
	data := utils.GetAsset("2015/1.txt")

	upCount := getCharacterCount(data, "(")
	downCount := getCharacterCount(data, ")")
	finalFloor := upCount - downCount

	pointAtBasement := getIndexInBasement(data)

	utils.FinalOutput(finalFloor, pointAtBasement)
}

func getCharacterCount(src string, targetString string) (count int) {
	for i := 0; i < len(src); i++ {
		if string(src[i]) == targetString {
			count++
		}
	}
	return
}

func getIndexInBasement(src string) int {
	currentFloor := 0

	for i := 0; i < len(src); i++ {
		action := string(src[i])
		if action == "(" {
			currentFloor++
		} else if action == ")" {
			currentFloor -= 1
		}

		if currentFloor == -1 {
			return i + 1
		}
	}

	return -1
}
