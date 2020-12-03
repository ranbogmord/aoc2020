package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
	"strings"
)

func getHits(input []string, dx, dy int) int {
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
		//fmt.Println(strings.Join(matrix[y], " "))
		x = ((x + dx) % len(matrix[0]))
		y += dy

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

	return hitTrees
}

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		var hits []int
		hits = append(hits, getHits(input, 1, 1))
		hits = append(hits, getHits(input, 3, 1))
		hits = append(hits, getHits(input, 5, 1))
		hits = append(hits, getHits(input, 7, 1))
		hits = append(hits, getHits(input, 1, 2))

		res := hits[0] * hits[1]
		for i := 2; i < len(hits); i += 1 {
			res = res * hits[i]
		}

		fmt.Println(res)
	})
}
