package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dirs = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var diagonals = [][]int{
	{-1, -1},
	{-1, 1},
}

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

	input = addPadding(input, ".", 3)

	for i, line := range input {
		for j, char := range line {
			if string(char) != "X" {
				continue
			}

			for _, dir := range dirs {
				word := "X"
				xNext := j
				yNext := i

				for count := 0; count < 3; count++ {
					xNext += dir[0]
					yNext += dir[1]

					word += string(input[yNext][xNext])
				}

				if word == "XMAS" {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0

	input = addPadding(input, ".", 3)

	for i, line := range input {
		for j, char := range line {
			if string(char) != "A" {
				continue
			}

			diagCount := 0

			for _, diag := range diagonals {
				word := ""

				for count := 0; count < 3; count++ {
					switch count {
					case 0:
						xNext := j + diag[0]
						yNext := i + diag[1]
						word += string(input[yNext][xNext])
					case 1:
						word += string(input[i][j])
					case 2:
						xNext := j + (diag[0] * (-1))
						yNext := i + (diag[1] * (-1))
						word += string(input[yNext][xNext])
					}
				}

				if word == "MAS" || word == "SAM" {
					diagCount++
				}
			}

			if diagCount == 2 {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

// Dibantu gpt
func addPadding(input []string, padValue string, padWidth int) []string {
	// Determine the size of the original matrix
	rows := len(input)
	cols := len(input[0])

	// New dimensions after padding
	newRows := rows + 2*padWidth
	newCols := cols + 2*padWidth

	// Create a new matrix with the new dimensions and fill with padValue
	newMatrix := make([][]byte, newRows)
	for i := range newMatrix {
		newMatrix[i] = make([]byte, newCols)
		for j := range newMatrix[i] {
			newMatrix[i][j] = padValue[0]
		}
	}

	// Copy the original matrix to the center of the new matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newMatrix[i+padWidth][j+padWidth] = input[i][j]
		}
	}

	return convertToStrings(newMatrix)
}

func convertToStrings(byteSlices [][]byte) []string {
	strings := make([]string, len(byteSlices))
	for i, b := range byteSlices {
		strings[i] = string(b)
	}
	return strings
}
