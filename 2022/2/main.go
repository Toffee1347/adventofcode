package main

import (
	"strings"

	"github.com/Toffee1347/adventofcode/utils"
)

var scores map[string]int = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var playsToWin map[string]string = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
	"X": "C",
	"Y": "A",
	"Z": "B",
}
var playsToLoose map[string]string = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
	"X": "B",
	"Y": "C",
	"Z": "A",
}

func main() {
	data := utils.GetAsset("2022/2.txt")
	scores := getGameScore(data)

	utils.FinalOutput(scores[0], scores[1])
}

func getGameScore(data string) [2]int {
	splitData := strings.Split(data, "\r\n")
	totalScores := [2]int{}

	for _, play := range splitData {
		playData := strings.Split(play, " ")
		opponentPlay := playData[0]

		selfPlay := playData[1]
		totalScores[0] += scores[selfPlay]
		if scores[opponentPlay] == scores[selfPlay] {
			totalScores[0] += 3
		} else if playsToWin[selfPlay] == opponentPlay {
			totalScores[0] += 6
		}

		targetEnd := playData[1]
		if targetEnd == "X" {
			selfPlay = playsToWin[opponentPlay]
		} else if targetEnd == "Y" {
			selfPlay = opponentPlay
			totalScores[1] += 3
		} else if targetEnd == "Z" {
			selfPlay = playsToLoose[opponentPlay]
			totalScores[1] += 6
		}

		totalScores[1] += scores[selfPlay]
	}

	return totalScores
}
