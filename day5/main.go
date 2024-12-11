package main

import (
	"bufio"
	"fmt"
	"log"
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

	var rules [][]int
	var pages [][]int
	rulesFlag := true

	for _, line := range input {
		if line == "" {
			rulesFlag = false
			continue
		}

		if rulesFlag {
			numArr := convertToIntArr(strings.Split(line, "|"))
			rules = append(rules, numArr)
		} else {
			pageArr := convertToIntArr(strings.Split(line, ","))
			pages = append(pages, pageArr)
		}
	}

	part1(rules, pages)
	part2(rules, pages)
}

func part1(rules [][]int, pages [][]int) {
	sum := 0

	for _, page := range pages {
		safe := true

		for _, rule := range rules {
			if !slices.Contains(page, rule[0]) || !slices.Contains(page, rule[1]) {
				continue
			}

			if !checkRule(page, rule) {
				safe = false
				break
			}
		}

		if safe {
			sum += page[len(page)/2]
		}
	}

	fmt.Println(sum)
}

func part2(rules [][]int, pages [][]int) {
	sum := 0

	for _, page := range pages {
		safe := true
		var tmpPage []int
		var rulesBroken [][]int

		for _, rule := range rules {
			if !slices.Contains(page, rule[0]) || !slices.Contains(page, rule[1]) {
				continue
			}

			if !checkRule(page, rule) {
				safe = false
				rulesBroken = append(rulesBroken, rule)
				break
			}
		}

		if !safe {
			tmpPage = fixPages(page, rules, rulesBroken)
			sum += tmpPage[len(tmpPage)/2]
		}
	}

	fmt.Println(sum)
}

func fixPages(page []int, rules [][]int, rulesBroken [][]int) []int {
	for _, rule := range rulesBroken {
		loc1 := slices.Index(page, rule[0])
		loc2 := slices.Index(page, rule[1])

		page[loc1], page[loc2] = page[loc2], page[loc1]
	}

	rulesBroken = [][]int{}
	for _, rule := range rules {
		if !slices.Contains(page, rule[0]) || !slices.Contains(page, rule[1]) {
			continue
		}

		if !checkRule(page, rule) {
			rulesBroken = append(rulesBroken, rule)
			break
		}
	}

	if len(rulesBroken) == 0 {
		return page
	} else {
		return fixPages(page, rules, rulesBroken)
	}
}

func checkRule(page []int, rule []int) bool {
	return slices.Index(page, rule[0]) < slices.Index(page, rule[1])
}

func convertToIntArr(input []string) []int {
	var intArr []int
	for _, num := range input {
		num, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}

		intArr = append(intArr, num)
	}

	return intArr
}
