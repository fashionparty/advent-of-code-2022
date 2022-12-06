package day_5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run() {
	fmt.Println("Day 5")
	input, err := os.Open("day_5\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	stack, movesList := separateStacksFromMoves(values)
	stacksOrganized := separateIndividualStacks(stack)
	stacksOrganizedCopy := make([]string, len(stacksOrganized))
	copy(stacksOrganizedCopy, stacksOrganized)
	moves := createMoveStructs(movesList)
	finishedStacks := moveCrates(stacksOrganized, moves)
	finishedStacks9001 := moveCrates9001(stacksOrganizedCopy, moves)

	fmt.Println("Answer 1:", getTopCrates(finishedStacks))
	fmt.Println("Answer 2:", getTopCrates(finishedStacks9001))
	fmt.Println()
	fmt.Println(reverseString("ABC"))
}

func moveCrates9001(stacks []string, moves []Move) []string {
	stacksCopy := make([]string, len(stacks))
	copy(stacksCopy, stacks)
	for i := 0; i < len(moves); i++ {
		quantity := moves[i].quantity
		target := moves[i].target - 1
		destination := moves[i].destination - 1
		if len(stacksCopy[target]) > 0 {
			crateSubStackCopy := ""
			if quantity > len(stacksCopy[target]) {
				crateSubStackCopy = stacksCopy[target][0:len(stacksCopy[target])]
			} else {
				crateSubStackCopy = stacksCopy[target][0:quantity]
			}
			stacksCopy[target] = stacksCopy[target][quantity:]
			stacksCopy[destination] = addCrate(stacksCopy[destination], crateSubStackCopy)
		}
	}
	return stacksCopy
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func moveCrates(stacks []string, moves []Move) []string {
	stacksCopy := make([]string, len(stacks))
	copy(stacksCopy, stacks)
	for i := 0; i < len(moves); i++ {
		quantity := moves[i].quantity
		target := moves[i].target - 1
		destination := moves[i].destination - 1
		for q := 0; q < quantity; q++ {
			if len(stacksCopy[target]) > 0 {
				crateCopy := ""
				crateCopy = string(stacksCopy[target][0])
				stacksCopy[target] = stacksCopy[target][1:]
				stacksCopy[destination] = addCrate(stacksCopy[destination], crateCopy)
			}
		}
	}
	return stacksCopy
}

func getTopCrates(stacks []string) string {
	answer := ""
	for i := 0; i < len(stacks); i++ {
		if len(stacks[i]) > 0 {
			answer += string(stacks[i][0])
		}
	}
	return answer
}

func addCrate(stack string, crate string) string {
	newStack := crate + stack
	return newStack
}

func createMoveStructs(moves []string) []Move {
	moveStructs := make([]Move, len(moves))
	for i := 0; i < len(moveStructs); i++ {
		moveStructs[i] = Move{
			quantity:    readQuantity(moves[i]),
			target:      readTarget(moves[i]),
			destination: readDestination(moves[i]),
		}
	}
	return moveStructs
}

func readDestination(move string) int {
	indexOfO := 0
	for i := 0; i < len(move); i++ {
		if string(move[i]) == "o" {
			indexOfO = i
		}
	}
	substring := move[indexOfO+2:]
	target, err := strconv.Atoi(substring)
	if err != nil {
		log.Panic(err)
	}
	return target
}

func readTarget(move string) int {
	indexOfM := 0
	indexOfT := 0
	for i := 0; i < len(move); i++ {
		if string(move[i]) == "m" {
			indexOfM = i
		}
		if string(move[i]) == "t" {
			indexOfT = i
		}
	}
	substring := move[indexOfM+2 : indexOfT-1]
	target, err := strconv.Atoi(substring)
	if err != nil {
		log.Panic(err)
	}
	return target
}

func readQuantity(move string) int {
	indexOfE := 0
	indexOfF := 0
	for i := 0; i < len(move); i++ {
		if string(move[i]) == "e" {
			indexOfE = i
		}
		if string(move[i]) == "f" {
			indexOfF = i
		}
	}
	substring := move[indexOfE+2 : indexOfF-1]
	quantity, err := strconv.Atoi(substring)
	if err != nil {
		log.Panic(err)
	}
	return quantity
}

func separateIndividualStacks(stacks []string) []string {
	numberOfColumns := getNumberOfColumns(stacks)
	columnsWithIndexes := make([]ColumnsWithIndexes, numberOfColumns)
	for i := 0; i < numberOfColumns; i++ {
		columnsWithIndexes[i].column = i
		columnsWithIndexes[i].index = getColumnIndex(i)
	}
	organizedStacks := getAllValuesFromGivenColumn(columnsWithIndexes, stacks[:len(stacks)-1])
	return organizedStacks
}

func getColumnIndex(column int) int {
	return (column)*4 + 2
}

func getAllValuesFromGivenColumn(columnsWithIndexes []ColumnsWithIndexes, stacks []string) []string {
	organizedStacks := make([]string, len(columnsWithIndexes))
	for i := 0; i < len(stacks); i++ {
		for k := 0; k < len(stacks[i]); k++ {
			for j := 0; j < len(columnsWithIndexes); j++ {
				if k == columnsWithIndexes[j].index {
					if string(stacks[i][k-1]) != " " {
						organizedStacks[columnsWithIndexes[j].column] += string(stacks[i][k-1])
					}
				}
			}
		}
	}
	return organizedStacks
}

func getNumberOfColumns(stack []string) int {
	lastRow := stack[len(stack)-1]
	numberOfColumns := lastRow[len(lastRow)-2]
	numberOfColumnsString := string(numberOfColumns)
	numberOfColumnsInt, err := strconv.Atoi(numberOfColumnsString)
	if err != nil {
		log.Panic(err)
	}
	return numberOfColumnsInt
}

func separateStacksFromMoves(input []string) ([]string, []string) {
	indexOfEmptyLine := getIndexOfEmptyLine(input)
	stacks := input[:indexOfEmptyLine]
	moves := input[indexOfEmptyLine+1:]
	return stacks, moves
}

func getIndexOfEmptyLine(input []string) int {
	index := 0
	for i, v := range input {
		if v == "" {
			index = i
		}
	}
	return index
}
