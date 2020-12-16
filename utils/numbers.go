package utils

import (
	"log"
	"strconv"
)

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func AbsInt(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func ToInt(x string) int {
	val, err := strconv.Atoi(x)
	if err != nil {
		log.Fatalf("failed to parse int: %v", err)
	}

	return val
}
