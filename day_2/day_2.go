package day_2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pointsTotalFirstRound int
var pointsTotalSecondRound int

func Run() {
	input, err := os.Open("day_2\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	for _, v := range values {
		opponentMove := string(v[0])
		playerMove := string(v[2])
		if checkIfWon(opponentMove, playerMove) {
			pointsTotalFirstRound += 6
		} else if checkIfDraw(opponentMove, playerMove) {
			pointsTotalFirstRound += 3
		}
		pointsTotalFirstRound += pointsForShape(playerMove)
	}
	for _, v := range values {
		opponentMove := string(v[0])
		strategy := string(v[2])
		pointsTotalSecondRound += calculateOutcome(opponentMove, strategy)
	}
	fmt.Println("Day 2")
	fmt.Println("Answer 1:", pointsTotalFirstRound)
	fmt.Println("Answer 2:", pointsTotalSecondRound)
	fmt.Println()
}

func calculateOutcome(opponent string, strategy string) int {
	if strategy == "X" {
		if opponent == "A" {
			return 3
		} else if opponent == "B" {
			return 1
		} else if opponent == "C" {
			return 2
		} else {
			return 0
		}
	} else if strategy == "Y" {
		if opponent == "A" {
			return 4
		} else if opponent == "B" {
			return 5
		} else if opponent == "C" {
			return 6
		} else {
			return 0
		}
	} else if strategy == "Z" {
		if opponent == "A" {
			return 8
		} else if opponent == "B" {
			return 9
		} else if opponent == "C" {
			return 7
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func pointsForShape(shape string) int {
	if shape == "X" {
		return 1
	} else if shape == "Y" {
		return 2
	} else if shape == "Z" {
		return 3
	} else {
		return 0
	}
}

func checkIfDraw(opponent string, player string) bool {
	if opponent == "A" && player == "X" {
		return true
	} else if opponent == "B" && player == "Y" {
		return true
	} else if opponent == "C" && player == "Z" {
		return true
	} else {
		return false
	}
}

func checkIfWon(opponent string, player string) bool {
	if opponent == "A" && player == "Y" {
		return true
	} else if opponent == "B" && player == "Z" {
		return true
	} else if opponent == "C" && player == "X" {
		return true
	} else {
		return false
	}
}
