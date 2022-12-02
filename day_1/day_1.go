package day_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Run() {
	input, err := os.Open("day_1\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	var sums []int
	singleSum := 0
	for _, v := range values {
		if v != "" {
			convertedValue, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			singleSum += convertedValue
		} else {
			sums = append(sums, singleSum)
			singleSum = 0
		}
	}
	sort.Ints(sums)
	answer1 := sums[len(sums)-1]
	fmt.Println("Day 1")
	fmt.Println("Answer 1:", answer1)
	answer2 := sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
	fmt.Println("Answer 2:", answer2)
	fmt.Println()
}
