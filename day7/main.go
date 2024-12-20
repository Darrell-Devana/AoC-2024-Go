package main

import (
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

	part2(input)
}

func part1(input []string) {
	sum := 0
	equations := make(map[int]bool)

	for _, line := range input {
		splits := strings.Split(line, ":")
		result, _ := strconv.Atoi(strings.TrimSpace(splits[0]))
		nums := convertToIntArr(strings.Split(strings.TrimSpace(splits[1]), " "))

		configs := generateConfigurations(len(nums) - 1)

		fmt.Println(configs)

		for _, config := range configs {
			tmp := nums[0]
			for i, op := range config {
				if op == 0 {
					tmp = tmp + nums[i+1]
				} else {
					tmp = tmp * nums[i+1]
				}
			}

			if tmp == result {
				if !equations[tmp] {
					sum += tmp
					equations[result] = true
				}
			}
		}
	}

	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0
	equations := make(map[int]bool)

	for _, line := range input {
		splits := strings.Split(line, ":")
		result, _ := strconv.Atoi(strings.TrimSpace(splits[0]))
		nums := convertToIntArr(strings.Split(strings.TrimSpace(splits[1]), " "))

		possibleConfigs := math.Pow(float64(3), float64(len(nums)-1))

		configs := generateConfigurations2(int(possibleConfigs), len(nums)-1)

		fmt.Println(configs)

		for _, config := range configs {
			tmp := nums[0]
			for i, op := range config {
				switch op {
				case 0:
					tmp = tmp * nums[i+1]
				case 1:
					tmp = tmp + nums[i+1]
				default:
					tmpStr := strconv.Itoa(tmp) + strconv.Itoa(nums[i+1])
					tmp, _ = strconv.Atoi(tmpStr)
				}
			}

			if tmp == result {
				if !equations[tmp] {
					sum += tmp
					equations[result] = true
				}
			}
		}
	}

	fmt.Println(sum)
}

func generateConfigurations2(n int, space int) [][]int {
	configs := [][]int{}

	for i := 0; i < n; i++ {
		conf := decimalToTernary(i)
		config := []int{}

		for _, ch := range conf {
			num, _ := strconv.Atoi(string(ch))
			config = append(config, num)
		}

		for len(config) < space {
			config = append([]int{0}, config...)
		}

		configs = append(configs, config)
	}

	return configs
}

func decimalToTernary(decimal int) string {
	if decimal == 0 {
		return "0"
	}

	ternary := ""

	for decimal > 0 {
		remainder := decimal % 3
		ternary = fmt.Sprintf("%d", remainder) + ternary
		decimal = decimal / 3
	}

	return ternary
}

func generateConfigurations(n int) [][]int {
	configs := [][]int{}

	for i := 0; i < (1 << n); i++ {
		conf := fmt.Sprintf("%0*b", n, i)
		var config []int

		for _, ch := range conf {
			num, _ := strconv.Atoi(string(ch))
			config = append(config, num)
		}

		configs = append(configs, config)
	}

	return configs
}

func convertToIntArr(input []string) []int {
	var intArr []int

	for _, num := range input {
		num, _ := strconv.Atoi(num)
		intArr = append(intArr, num)
	}

	return intArr
}
