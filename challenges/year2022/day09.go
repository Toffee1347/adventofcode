package year2022

import (
	"fmt"
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

var ropeInstructionDirections map[string][]int = map[string][]int{
	"L": {0, 1},
	"R": {0, -1},
	"U": {-1, 0},
	"D": {1, 0},
}

func Day09(input string) [2]any {
	splitInput := strings.Split(input, "\r\n")

	partOneGrid := processRopeInstructions(splitInput, 2)
	partOneCount := len(partOneGrid[1])

	partTwoGrid := processRopeInstructions(splitInput, 10)
	partTwoCount := len(partTwoGrid[9])

	return [2]any{partOneCount, partTwoCount}
}

func processRopeInstructions(instructions []string, knotCount int) (grids []map[string]bool) {
	knotLocations := []utils.Coordinate{}
	for i := 0; i < knotCount; i++ {
		grids = append(grids, map[string]bool{"0:0": true})
		knotLocations = append(knotLocations, utils.Coordinate{})
	}

	instructionData := [][]int{}
	for _, instructionString := range instructions {
		splitInstruction := strings.Split(instructionString, " ")
		direction := ropeInstructionDirections[splitInstruction[0]]

		for i := 0; i < utils.StrToInt(splitInstruction[1]); i++ {
			instructionData = append(instructionData, direction)
		}
	}

	for _, instruction := range instructionData {
		knotLocations[0].X += instruction[0]
		knotLocations[0].Y += instruction[1]
		grids[0][utils.IntToStr(knotLocations[0].X)+":"+utils.IntToStr(knotLocations[0].Y)] = true

		for knotI := 1; knotI < knotCount; knotI++ {
			knotLocations[knotI] = getNewTailLocation(knotLocations[knotI-1], knotLocations[knotI])
			grids[knotI][utils.IntToStr(knotLocations[knotI].X)+":"+utils.IntToStr(knotLocations[knotI].Y)] = true
		}

		fmt.Println(makePrintGrid(grids))
	}

	return
}

func getNewTailLocation(head utils.Coordinate, tail utils.Coordinate) (newTail utils.Coordinate) {
	if !tailShouldMove(head, tail) {
		return tail
	}

	newTail = tail

	targetPoints := []utils.Coordinate{
		{X: head.X, Y: head.Y + 1},
		{X: head.X, Y: head.Y - 1},
		{X: head.X + 1, Y: head.Y},
		{X: head.X - 1, Y: head.Y},
	}
	shortestDistance := -1
	shortestDistanceI := -1

	for i, point := range targetPoints {
		distance := int(utils.AxisGetDistanceBetweenPoints(point, tail))

		if shortestDistance == -1 || distance < shortestDistance {
			shortestDistance = distance
			shortestDistanceI = i
		}
	}

	return targetPoints[shortestDistanceI]
}

func tailShouldMove(head utils.Coordinate, tail utils.Coordinate) bool {
	length := utils.Coordinate{
		X: head.X - tail.X,
		Y: head.Y - tail.Y,
	}

	lengthSquared := length.X*length.X + length.Y*length.Y

	return lengthSquared > 2
}

func makePrintGrid(grids []map[string]bool) (out string) {
	largeGrid := [][]string{
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
	}

	for gridI, grid := range grids {
		for coord := range grid {
			coords := utils.SplitStrToInt(coord, ":")
			largeGrid[coords[0]][coords[1]] = utils.IntToStr(gridI)
		}
	}

	for i := len(largeGrid) - 1; i >= 0; i-- {
		out += strings.Join(largeGrid[i], "") + "\n"
	}

	return
}
