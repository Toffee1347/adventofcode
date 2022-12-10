package year2022

import (
	"fmt"
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

type ropeGridPoint struct {
	tailHasVisited bool
}

var ropeInstructionDirections map[string][]int = map[string][]int{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

func Day09(input string) [2]any {
	splitInput := strings.Split(input, "\r\n")
	grid := processRopeInstructions(splitInput)
	fmt.Println(grid)
	tailVisitedCount := getTailVisitedCount(grid)

	return [2]any{tailVisitedCount}
}

func processRopeInstructions(instructions []string) (grid [][]ropeGridPoint) {
	grid = append(grid, []ropeGridPoint{{
		tailHasVisited: true,
	}})

	headLocation := utils.Coordinate{
		X: 0,
		Y: 0,
	}
	tailLocation := utils.Coordinate{
		X: 0,
		Y: 0,
	}
	for _, instruction := range instructions {
		instructionData := strings.Split(instruction, " ")
		direction := ropeInstructionDirections[instructionData[0]]

		for i := 0; i < utils.StrToInt(instructionData[1]); i++ {
			headLocation.X += direction[0]
			headLocation.Y += direction[1]

			tailLocation.X += direction[0]
			tailLocation.Y += direction[1]

			if tailShouldMove(headLocation, tailLocation) {
				if headLocation.X != tailLocation.X && headLocation.Y != tailLocation.Y {
					if direction[0] == 0 {
						tailLocation.X = headLocation.X
					} else if direction[1] == 0 {
						tailLocation.Y = headLocation.Y
					}
				}

				grid[tailLocation.X][tailLocation.Y].tailHasVisited = true
			}
		}
	}

	return
}

func tailShouldMove(head utils.Coordinate, tail utils.Coordinate) bool {
	length := utils.Coordinate{
		X: head.X - tail.X,
		Y: head.Y - tail.Y,
	}

	lengthSquared := length.X*length.Y + length.Y*length.Y

	return lengthSquared > 2
}

func getTailVisitedCount(grid [][]ropeGridPoint) (count int) {
	for _, row := range grid {
		for _, part := range row {
			if part.tailHasVisited {
				count++
			}
		}
	}
	return
}
