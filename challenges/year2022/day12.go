package year2022

import (
	"fmt"
	"math"
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

func Day12(input string) [2]any {
	elevationMap, startPoint, endPoint := processHeatmapData(strings.Split(input, "\r\n"))
	fmt.Println(elevationMap)

	partOneValue := findBestRouteLength(elevationMap, startPoint, endPoint)

	return [2]any{partOneValue}
}

func processHeatmapData(rows []string) (elevation [][]int, startPoint utils.Coordinate, endPoint utils.Coordinate) {
	for y, row := range rows {
		elevation = append(elevation, []int{})
		for x := 0; x < len(row); x++ {
			letter := row[x]
			value := int(letter)

			if letter == 'S' {
				startPoint = utils.Coordinate{X: x, Y: y}
				value = -2
			} else if letter == 'E' {
				endPoint = utils.Coordinate{X: x, Y: y}
				value = -2
			}

			elevation[y] = append(elevation[y], value)
		}
	}

	return
}

func findBestRouteLength(elevationMap [][]int, startPoint utils.Coordinate, endPoint utils.Coordinate) int {
	currentPoint := elevationMap[startPoint.Y][startPoint.X]
	validRouteLengths := []int{}

	for _, direction := range utils.Directions {
		newPointCoord := utils.Coordinate{
			X: startPoint.X + direction[0],
			Y: startPoint.Y + direction[1],
		}

		if newPointCoord.X < 0 || newPointCoord.X >= len(elevationMap[0]) || newPointCoord.Y < 0 || newPointCoord.Y >= len(elevationMap) {
			break
		}

		if newPointCoord.X == endPoint.X && newPointCoord.Y == endPoint.Y {
			return 1
		}

		newPoint := elevationMap[newPointCoord.Y][newPointCoord.X]
		if math.Abs(float64(currentPoint-newPoint)) <= 1 {
			routeLength := findBestRouteLength(elevationMap, newPointCoord, endPoint)

			if routeLength != 0 {
				validRouteLengths = append(validRouteLengths, routeLength+1)
			}
		}
	}

	if len(validRouteLengths) == 0 {
		return 0
	} else {
		return utils.ArrayGetBoundValues(validRouteLengths, 1, false)[0]
	}
}
