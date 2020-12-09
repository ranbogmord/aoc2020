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

func findMinMaxSumForInvalidNumber(nums []int, number int) {
	minIdx := -1
	maxIdx := -1
	found := false

	windowSize := 1
	for windowSize < len(nums) {
		windowSize += 1

		for i := 0; i <= len(nums) - windowSize; i++ {
			window := nums[i:i+windowSize]
			windowSum := 0
			for _, num := range window {
				windowSum += num
			}

			if windowSum == number {
				minIdx = i
				maxIdx = i + (windowSize - 1)
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	if !found {
		log.Fatalln("failed to find window")
	}

	minFound := -1
	maxFound := -1

	for _, val := range nums[minIdx:maxIdx] {
		if minFound == -1 || val < minFound {
			minFound = val
		}

		if maxFound == -1 || val > maxFound {
			maxFound = val
		}
	}

	sum :=  minFound + maxFound

	fmt.Println("Weakness:", sum)
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
		findMinMaxSumForInvalidNumber(nums, invalid)
	})
}
