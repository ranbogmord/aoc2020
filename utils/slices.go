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

func SliceFilterString(slice []string, f func(thing string) bool) []string {
	var newSlice []string
	for _, item := range slice {
		if f(item) {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}
