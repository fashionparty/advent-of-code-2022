package day_3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var sumOfPrioritiesFirst = 0
var sumOfPrioritiesSecond = 0

func Run() {
	fmt.Println("Day 3")
	input, err := os.Open("day_3\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	for _, v := range values {
		firstCompartment, secondCompartment := divideIntoCompartments(v)
		commonItem := findCommonItem(firstCompartment, secondCompartment)
		sumOfPrioritiesFirst += int(getPriority(commonItem))
	}

	for i := 0; i < len(values); i += 3 {
		commonItem := findCommonItemElves(values[i : i+3])
		sumOfPrioritiesSecond += int(getPriority(commonItem))
	}
	fmt.Println("Answer 1:", sumOfPrioritiesFirst)
	fmt.Println("Answer 2:", sumOfPrioritiesSecond)
	fmt.Println()
}

func findCommonItemElves(elves []string) int32 {
	var commonItem int32
	for _, val1 := range elves[0] {
		for _, val2 := range elves[1] {
			for _, val3 := range elves[2] {
				if val1 == val2 && val1 == val3 {
					commonItem = val1
				}
			}
		}
	}
	return commonItem
}

func getPriority(commonItem int32) int32 {
	var priority int32
	if commonItem < 97 {
		priority = commonItem - 65 + 27
	} else {
		priority = commonItem - 96
	}
	return priority
}

func findCommonItem(first string, second string) int32 {
	var commonItem int32
	for _, firstValue := range first {
		for _, secondValue := range second {
			if firstValue == secondValue {
				commonItem = firstValue
			}
		}
	}
	return commonItem
}

func divideIntoCompartments(items string) (string, string) {
	firstCompartment := items[:(len(items))/2]
	secondCompartment := items[len(items)/2:]
	return firstCompartment, secondCompartment
}
