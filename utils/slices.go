package utils

import "math"

func SliceMax(slice []int) int {
	max := 0
	for _, val := range slice {
		if val > max {
			max = val
		}
	}

	return max
}

func SliceMin(slice []int) int {
	min := math.MaxInt32
	for _, val := range slice {
		if val < min {
			min = val
		}
	}

	return min
}
