package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

var dirs = []string{"V", "^", "<", ">"}

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

	byteInput := convertToByteArr(input)
	sum := 0

	start := time.Now()
	// fmt.Println(part1(byteInput, sum))
	fmt.Println(part2(byteInput, input, sum))
	fmt.Println("Part 2 took", time.Since(start))
}

func part1(input [][]byte, sum int) int {
	for i, line := range input {
		for j, char := range line {
			if !slices.Contains(dirs, string(char)) {
				continue
			}

			switch string(char) {
			case "^":
				fmt.Println("Up")
				nextY := i - 1
				if nextY < 0 {
					return sum
				}

				if string(input[nextY][j]) == "#" {
					input[i][j] = '>'
					return part1(input, sum)
				}

				if string(input[nextY][j]) == "X" {
					fmt.Println("X detected on the next element")
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					return part1(input, sum)
				} else {
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					input[i][j] = 'X'
					sum++
					fmt.Println(sum)
					return part1(input, sum)
				}
			case "V":
				fmt.Println("Down")
				nextY := i + 1
				if nextY > len(input)-1 {
					return sum
				}

				if string(input[nextY][j]) == "#" {
					input[i][j] = '<'
					return part1(input, sum)
				}

				if string(input[nextY][j]) == "X" {
					fmt.Println("X detected on the next element")
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					return part1(input, sum)
				} else {
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					input[i][j] = 'X'
					sum++
					fmt.Println(sum)
					return part1(input, sum)
				}
			case ">":
				fmt.Println("Right")
				nextX := j + 1
				if nextX > len(line)-1 {
					return sum
				}

				if string(input[i][nextX]) == "#" {
					input[i][j] = 'V'
					return part1(input, sum)
				}

				if string(input[i][nextX]) == "X" {
					fmt.Println("X detected on the next element")
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]
					return part1(input, sum)
				} else {
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]
					input[i][j] = 'X'
					sum++
					fmt.Println(sum)
					return part1(input, sum)
				}
			case "<":
				fmt.Println("Left")
				nextX := j - 1
				if nextX < 0 {
					return sum
				}

				if string(input[i][nextX]) == "#" {
					input[i][j] = '^'

					return part1(input, sum)
				}

				if string(input[i][nextX]) == "X" {
					fmt.Println("X detected on the next element")
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]

					return part1(input, sum)
				} else {
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]
					input[i][j] = 'X'
					sum++
					fmt.Println(sum)

					return part1(input, sum)
				}
			}
		}
	}

	return sum
}

func part2(input [][]byte, backup []string, sum int) int {
	locs := make(map[string]bool)
	locs = recordAllSteps(input, locs)
	input = convertToByteArr(backup)

	for i := range input {
		for j := range input[i] {
			coords := fmt.Sprintf("{%v,%v}", i, j)
			if !locs[coords] {
				continue
			}

			fmt.Println("new loop")
			locs2 := make(map[string]bool)
			input[i][j] = 'O'

			if checkForLoop(input, locs2) {
				fmt.Println("this is a loop")

				for i := range input {
					inputStr := string(input[i])
					fmt.Println(inputStr)
				}

				sum++
			}

			input = convertToByteArr(backup)
		}
	}

	return sum
}

func recordAllSteps(input [][]byte, locs map[string]bool) map[string]bool {
	for i, line := range input {
		for j, char := range line {
			if !slices.Contains(dirs, string(char)) {
				continue
			}

			coords := fmt.Sprintf("{%v,%v}", i, j)

			switch string(char) {
			case "^":
				fmt.Println("Up")
				nextY := i - 1
				if nextY < 0 {
					if !locs[coords] {
						locs[coords] = true
					}

					return locs
				}

				if string(input[nextY][j]) == "#" {
					input[i][j] = '>'
					return recordAllSteps(input, locs)
				}

				input[i][j], input[nextY][j] = input[nextY][j], input[i][j]

				if !locs[coords] {
					locs[coords] = true
				}

				return recordAllSteps(input, locs)
			case "V":
				fmt.Println("Down")
				nextY := i + 1
				if nextY > len(input)-1 {
					if !locs[coords] {
						locs[coords] = true
					}
					return locs
				}

				if string(input[nextY][j]) == "#" {
					input[i][j] = '<'
					return recordAllSteps(input, locs)
				}

				input[i][j], input[nextY][j] = input[nextY][j], input[i][j]

				if !locs[coords] {
					locs[coords] = true
				}

				return recordAllSteps(input, locs)
			case ">":
				fmt.Println("Right")
				nextX := j + 1
				if nextX > len(line)-1 {
					if !locs[coords] {
						locs[coords] = true
					}
					return locs
				}

				if string(input[i][nextX]) == "#" {
					input[i][j] = 'V'
					return recordAllSteps(input, locs)
				}

				input[i][j], input[i][nextX] = input[i][nextX], input[i][j]

				if !locs[coords] {
					locs[coords] = true
				}

				return recordAllSteps(input, locs)
			case "<":
				fmt.Println("Left")
				nextX := j - 1
				if nextX < 0 {
					if !locs[coords] {
						locs[coords] = true
					}
					return locs
				}

				if string(input[i][nextX]) == "#" {
					input[i][j] = '^'

					return recordAllSteps(input, locs)
				}

				input[i][j], input[i][nextX] = input[i][nextX], input[i][j]

				if !locs[coords] {
					locs[coords] = true
				}

				return recordAllSteps(input, locs)
			}
		}
	}

	return locs
}

func checkForLoop(tmpInput [][]byte, locs map[string]bool) bool {
	input := tmpInput
	for i, line := range input {
		for j, char := range line {
			if !slices.Contains(dirs, string(char)) {
				continue
			}

			switch string(char) {
			case "^":
				nextY := i - 1

				if nextY < 0 {
					return false
				}

				if string(input[nextY][j]) == "#" || string(input[nextY][j]) == "O" {
					input[i][j] = '>'
					return checkForLoop(input, locs)
				}

				if string(input[nextY][j]) == "X" {
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					coords := string(i) + string(j) + string(char)

					if locs[coords] {
						return true
					} else {
						locs[coords] = true
					}

					return checkForLoop(input, locs)
				} else {
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					input[i][j] = 'X'
					coords := string(i) + string(j) + string(char)
					locs[coords] = true

					return checkForLoop(input, locs)
				}
			case "V":
				nextY := i + 1

				if nextY > len(input)-1 {
					return false
				}

				if string(input[nextY][j]) == "#" || string(input[nextY][j]) == "O" {
					input[i][j] = '<'
					return checkForLoop(input, locs)
				}

				if string(input[nextY][j]) == "X" {
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					coords := string(i) + string(j) + string(char)

					if locs[coords] {
						return true
					} else {
						locs[coords] = true
					}

					return checkForLoop(input, locs)
				} else {
					input[i][j], input[nextY][j] = input[nextY][j], input[i][j]
					input[i][j] = 'X'
					coords := string(i) + string(j) + string(char)
					locs[coords] = true

					return checkForLoop(input, locs)
				}
			case ">":
				nextX := j + 1

				if nextX > len(line)-1 {
					return false
				}

				if string(input[i][nextX]) == "#" || string(input[i][nextX]) == "O" {
					input[i][j] = 'V'
					return checkForLoop(input, locs)
				}

				if string(input[i][nextX]) == "X" {
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]
					coords := string(i) + string(j) + string(char)

					if locs[coords] {
						return true
					} else {
						locs[coords] = true
					}

					return checkForLoop(input, locs)
				} else {
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]
					input[i][j] = 'X'
					coords := string(i) + string(j) + string(char)
					locs[coords] = true

					return checkForLoop(input, locs)
				}
			case "<":
				nextX := j - 1

				if nextX < 0 {
					return false
				}

				if string(input[i][nextX]) == "#" || string(input[i][nextX]) == "O" {
					input[i][j] = '^'
					return checkForLoop(input, locs)
				}

				if string(input[i][nextX]) == "X" {
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]
					coords := string(i) + string(j) + string(char)

					if locs[coords] {
						return true
					} else {
						locs[coords] = true
					}

					return checkForLoop(input, locs)
				} else {
					input[i][j], input[i][nextX] = input[i][nextX], input[i][j]
					input[i][j] = 'X'
					coords := string(i) + string(j) + string(char)
					locs[coords] = true

					return checkForLoop(input, locs)
				}
			}
		}
	}

	return false
}

func convertToByteArr(input []string) [][]byte {
	var output [][]byte

	for _, line := range input {
		byteLine := []byte(line)
		output = append(output, byteLine)
	}

	return output
}
