package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	var sum int
	var list1 []int
	var list2 []int

	for _, line := range input {
		tmpNumbers := strings.Split(line, "   ")

		fmt.Println(tmpNumbers)
		number1, _ := strconv.Atoi(tmpNumbers[0])
		number2, _ := strconv.Atoi(tmpNumbers[1])
		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		distance := float64(list1[i] - list2[i])
		sum += int(math.Abs(distance))
	}

	fmt.Println(sum)
}

func part2(input []string) {
	var sum int
	var list1 []int
	map2 := make(map[int]int)

	for _, line := range input {
		tmpNumbers := strings.Split(line, "   ")

		fmt.Println(tmpNumbers)
		number1, _ := strconv.Atoi(tmpNumbers[0])
		number2, _ := strconv.Atoi(tmpNumbers[1])

		list1 = append(list1, number1)
		map2[number2]++
	}

	for _, num := range list1 {
		if val := map2[num]; val != 0 {
			sum += num * val
		}
	}

	fmt.Println(sum)
}
