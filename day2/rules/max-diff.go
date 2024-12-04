package rules

import "math"

func MaxDiff(input []int, max int) bool {
	for i := 1; i < len(input); i++ {
		if math.Abs(float64(input[i]-input[i-1])) > float64(max) {
			return false
		}
	}
	return true
}
