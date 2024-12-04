package rules

import "math"

var Tolerance = 1

func MinDiff(input []int, min int) bool {
	for i := 1; i < len(input); i++ {
		if math.Abs(float64(input[i]-input[i-1])) < float64(min) {
			return false
		}
	}
	return true
}
