package year2015

import (
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

type coordinate struct {
	x int
	y int
}

var directions map[string]coordinate = map[string]coordinate{
	"^": {x: 0, y: 1},
	">": {x: 1, y: 0},
	"v": {x: 0, y: -1},
	"<": {x: -1, y: 0},
}

func Day03(input string) [2]int {
	inputDirections := strings.Split(input, "")

	oneTrip, _ := visitedLocationCount(inputDirections, map[string]bool{})

	splitDirections := []([]string){
		[]string{},
		[]string{},
	}

	for i := 0; i < len(inputDirections); i++ {
		splitDirections[i%2] = append(splitDirections[i%2], inputDirections[i])
	}

	_, locations := visitedLocationCount(splitDirections[0], map[string]bool{})
	count, _ := visitedLocationCount(splitDirections[1], locations)

	return [2]int{oneTrip, count}
}

func visitedLocationCount(inputDirections []string, visitedLocations map[string]bool) (int, map[string]bool) {
	currentLocation := coordinate{x: 0, y: 0}
	if len(visitedLocations) == 0 {
		visitedLocations = map[string]bool{
			"0:0": true,
		}
	}

	for _, direction := range inputDirections {
		currentLocation = modifyCoordinate(currentLocation, directions[direction])

		visitedLocations[getCoordinateString(currentLocation)] = true
	}

	return len(visitedLocations), visitedLocations
}

func modifyCoordinate(orig coordinate, modifier coordinate) coordinate {
	orig.x += modifier.x
	orig.y += modifier.y
	return orig
}

func getCoordinateString(coord coordinate) string {
	x := utils.IntToStr(coord.x)
	y := utils.IntToStr(coord.y)

	return x + ":" + y
}
