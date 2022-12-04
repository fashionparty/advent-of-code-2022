package day_4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run() {
	fmt.Println("Day 4")
	input, err := os.Open("day_4\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	var fullyContained int
	var overlapping int
	for _, v := range values {
		first, second := splitString(v, ',')
		firstRangeFloor, firstRangeCeiling := splitString(first, '-')
		secondRangeFloor, secondRangeCeiling := splitString(second, '-')

		if checkIfFullyContained(stringToInt(firstRangeFloor), stringToInt(firstRangeCeiling),
			stringToInt(secondRangeFloor), stringToInt(secondRangeCeiling)) {
			fullyContained++
		}
		if checkIfOverlap(stringToInt(firstRangeFloor), stringToInt(firstRangeCeiling),
			stringToInt(secondRangeFloor), stringToInt(secondRangeCeiling)) {
			overlapping++
		}
	}
	fmt.Println("Answer 1:", fullyContained)
	fmt.Println("Answer 2:", overlapping)
	fmt.Println()
}

func stringToInt(value string) int {
	convertedValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return convertedValue
}

func checkIfOverlap(firstFloor int, firstCeiling int, secondFloor int, secondCeiling int) bool {
	isOverlapping := false
	if checkIfFullyContained(firstFloor, firstCeiling, secondFloor, secondCeiling) {
		isOverlapping = true
	}
	if firstFloor < secondFloor && firstCeiling >= secondFloor {
		isOverlapping = true
	}
	if firstCeiling > secondCeiling && firstFloor <= secondCeiling {
		isOverlapping = true
	}
	return isOverlapping
}

func checkIfFullyContained(firstFloor int, firstCeiling int, secondFloor int, secondCeiling int) bool {
	isFullyContained := false
	if firstFloor >= secondFloor && firstCeiling <= secondCeiling {
		isFullyContained = true
	}
	if secondFloor >= firstFloor && secondCeiling <= firstCeiling {
		isFullyContained = true
	}
	return isFullyContained
}

func splitString(pair string, symbol int32) (string, string) {
	symbolPosition := findSymbolPosition(pair, symbol)
	firstValue := pair[:symbolPosition]
	secondValue := pair[symbolPosition+1:]
	return firstValue, secondValue
}

func findSymbolPosition(pair string, symbol int32) int {
	var position int
	for i, v := range pair {
		if v == symbol {
			position = i
		}
	}
	return position
}
