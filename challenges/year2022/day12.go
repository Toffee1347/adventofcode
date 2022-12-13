package year2022

import (
	"fmt"
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

func Day12(input string) [2]any {
	elevationMap, startPoint, endPoint := processHeatmapData(strings.Split(input, "\r\n"))
	fmt.Println(elevationMap)

	partOneValue := findBestRouteLength(elevationMap, startPoint, endPoint, map[string]bool{}) - 2

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
				value = 0
			} else if letter == 'E' {
				endPoint = utils.Coordinate{X: x, Y: y}
				value = 0
			}

			elevation[y] = append(elevation[y], value)
		}
	}

	return
}

func findBestRouteLength(elevationMap [][]int, startPoint utils.Coordinate, endPoint utils.Coordinate, visited map[string]bool) int {
	currentPoint := elevationMap[startPoint.Y][startPoint.X]
	validRouteLengths := []int{}

	for _, direction := range utils.Directions {
		newPointCoord := utils.Coordinate{
			X: startPoint.X + direction[0],
			Y: startPoint.Y + direction[1],
		}

		if newPointCoord.X < 0 || newPointCoord.X >= len(elevationMap[0]) || newPointCoord.Y < 0 || newPointCoord.Y >= len(elevationMap) {
			continue
		}

		visitedKey := utils.IntToStr(newPointCoord.X) + ":" + utils.IntToStr(newPointCoord.Y)
		if _, exists := visited[visitedKey]; exists {
			continue
		}
		visited[visitedKey] = true

		if utils.AxisCoordinatesMatch(newPointCoord, endPoint) {
			return 1
		}

		newPoint := elevationMap[newPointCoord.Y][newPointCoord.X]
		if currentPoint == 0 || newPoint <= currentPoint+1 {
			routeLength := findBestRouteLength(elevationMap, newPointCoord, endPoint, visited)

			if routeLength != 0 {
				fmt.Println("Found end", routeLength+1)
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
