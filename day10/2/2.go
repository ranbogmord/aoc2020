package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
	"sort"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInputInts(os.Args[1])
		//input := utils.ReadInputInts("day10/input-test2.txt")

		input = append([]int{0}, input...)
		sort.Ints(input)
		input = append(input, input[len(input) - 1] + 3)

		steps := make(map[int]int)
		steps[0] = 1

		for i := 0; i < len(input); i++ {
			for j := i + 1; j < len(input); j++ {
				if input[j] > input[i] + 3 {
					break
				}

				if _, ok := steps[j]; ok {
					steps[j] = steps[j]  + steps[i]
				} else {
					steps[j] = steps[i]
				}
			}
		}

		fmt.Println(steps[len(input) - 1])
	})
}
