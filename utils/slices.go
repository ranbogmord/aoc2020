package utils

import (
	"sort"
)

func SliceMax(slice []int) int {
	sort.Ints(slice)

	return slice[len(slice) - 1]
}

func SliceMin(slice []int) int {
	sort.Ints(slice)
	return slice[0]
}
