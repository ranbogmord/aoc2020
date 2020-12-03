package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		reps := 1

		var matrix [][]string
		for _, val := range input {
			var row []string

			for i := 0; i < reps; i++ {
				chars := strings.Split(val, "")
				for _, c := range chars {
					row = append(row, c)
				}
			}

			matrix = append(matrix, row)
		}

		hitTrees := 0
		x := 0
		y := 0

		for {
			fmt.Println(strings.Join(matrix[y], " "))
			x = ((x + 3) % len(matrix[0]))
			y += 1

			if y > len(matrix) - 1 {
				break
			}

			if matrix[y][x] == "#" {
				hitTrees += 1
				matrix[y][x] = "X"
			} else {
				matrix[y][x] = "O"
			}
		}

		fmt.Println(hitTrees)
	})
}
