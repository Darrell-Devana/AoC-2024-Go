package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	mulRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	digitRegex := regexp.MustCompile(`\d+`)

	for _, line := range input {
		funcArr := mulRegex.FindAllString(line, -1)

		for _, data := range funcArr {
			foundDigits := convertToIntArr(digitRegex.FindAllString(data, -1))
			sum += foundDigits[0] * foundDigits[1]
		}
	}

	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0
	instruction := true
	mainRegex := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	digitRegex := regexp.MustCompile(`\d+`)

	for _, line := range input {
		funcArr := mainRegex.FindAllString(line, -1)

		for _, data := range funcArr {
			switch data {
			case "do()":
				instruction = true
			case "don't()":
				instruction = false
			default:
				if instruction {
					foundDigits := convertToIntArr(digitRegex.FindAllString(data, -1))

					if len(foundDigits) == 2 {
						sum += foundDigits[0] * foundDigits[1]
					}
				}
			}
		}
	}

	fmt.Println(sum)
}

func convertToIntArr(input []string) []int {
	var digits []int
	for _, num := range input {
		tmpStr, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		digits = append(digits, tmpStr)
	}

	return digits
}
