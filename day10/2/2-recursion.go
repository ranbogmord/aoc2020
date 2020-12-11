package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
)

func findChain(current []int, result int, tree utils.BTS, max int) int {
	if len(current) > 0 && current[len(current) - 1] == max {
		return result + 1
	}

	if len(current) > 0 && current[len(current) - 1] > max {
		return result
	}

	target := 0
	if len(current) > 0 {
		target = current[len(current) - 1]
	}

	added := false
	for i := 1; i < 4; i++ {
		exists := tree.Find(target + i)
		if exists != nil {
			newChain := append(current, target + i)
			result = findChain(newChain, result, tree, max)
			added = true
		}
	}

	if added {
		return result
	}

	return result
}

func main() {
	utils.Time(func() {
		input := utils.ReadInputInts(os.Args[1])

		max := utils.SliceMax(input)
		tree := utils.BTS{}
		for _, val := range input {
			tree.Insert(val)
		}

		result := findChain([]int{0}, 0, tree, max)

		fmt.Println(result)
	})
}
