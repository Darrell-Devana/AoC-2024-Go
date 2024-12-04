package main

import (
	"aoc2024/day2/rules"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	part1(input)
	part2(input)
}

func part1(input []string) {
	sum := 0

	for _, line := range input {
		ascending := true
		levelsStr := strings.Split(line, " ")
		levelsInt := convertToIntArr(levelsStr)

		if levelsInt[1] < levelsInt[0] {
			ascending = false
		}

		if validate(levelsInt, ascending) {
			sum++
		}
	}

	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0

	for _, line := range input {
		ascending := true
		levelsStr := strings.Split(line, " ")
		levelsInt := convertToIntArr(levelsStr)

		if levelsInt[1] < levelsInt[0] {
			ascending = false
		}

		if validate(levelsInt, ascending) || validateDamp(levelsInt, ascending) {
			sum++
		}
	}

	fmt.Println(sum)
}

func validateGit(data []int) bool {
	desc := rules.AllDecreasing(data)
	incr := rules.AllIncreasing(data)
	max := rules.MaxDiff(data, 3)
	min := rules.MinDiff(data, 1)
	return (desc || incr) && max && min
}

func validateDampGit(data []int) bool {
	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		if validateGit(data1) {
			return true
		}
	}
	return false
}

func validate(input []int, ascending bool) bool {
	for i := 1; i < len(input); i++ {
		diff := math.Abs(float64(input[i] - input[i-1]))
		if (diff < 1 || diff > 3) || ascending != (input[i] > input[i-1]) {
			return false
		}
	}

	return true
}

func validateDamp(input []int, ascending bool) bool {
	for i := 0; i < len(input); i++ {
		tmpLevels := append([]int{}, input[:i]...)
		tmpLevels = append(tmpLevels, input[i+1:]...)

		if validate(tmpLevels, ascending) {
			return true
		}
	}

	return false
}

func convertToIntArr(input []string) []int {
	var output []int

	for _, level := range input {
		levelInt, err := strconv.Atoi(level)
		if err != nil {
			log.Fatal(err)
		}

		output = append(output, levelInt)
	}

	return output
}
