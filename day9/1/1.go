package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findInvalidNumber(nums []int, preambleLength int) int {
	for i := preambleLength; i < len(nums); i++ {
		min := i - preambleLength
		max := i + preambleLength
		window := nums[min:max]
		target := nums[i]
		foundMatch := false

		for _, v1 := range window {
			for _, v2 := range window {
				if v1 == v2 {
					continue
				}

				if v1 + v2 == target {
					foundMatch = true
					break
				}
			}

			if foundMatch {
				break
			}
		}

		if !foundMatch {
			return target
		}
	}

	return -1
}

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])
		var nums []int

		for _, val := range input {
			n, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalf("failed to parse int: %v", err)
			}

			nums = append(nums, n)
		}

		invalid := findInvalidNumber(nums, 25)
		fmt.Println("Invalid number =", invalid)
	})
}
