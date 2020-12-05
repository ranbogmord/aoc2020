package main

import (
	"aoc2020/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		var matrix [128][8]int
		for i := 0; i < 128; i++ {
			for j := 0; j < 8; j++ {
				matrix[i][j] = 0
			}
		}

		maxId := 0

		for _, code := range input {
			min := 0.0
			max := 127.0

			rMin := 0.0
			rMax := 7.0

			for _, char := range strings.Split(code, "") {
				if char == "F" {
					max = math.Floor(max - ((max - min) / 2))
				} else if char == "B" {
					min = math.Ceil(max - ((max - min) / 2))
				} else if char == "L" {
					rMax = math.Floor(rMax - ((rMax - rMin) / 2))
				} else if char == "R" {
					rMin = math.Ceil(rMax - ((rMax - rMin) / 2))
				}
			}

			// fmt.Println(min, rMin, (min * 8) + rMin)
			id := int((min * 8) + rMin)
			if id > maxId {
				maxId = id
			}
			matrix[int(min)][int(rMin)] = 1
		}

		for i := 0; i < 128; i++ {
			emptySeats := 0
			for j := 0; j < 8; j++ {
				if matrix[i][j] == 0 {
					emptySeats += 1
				}
			}

			if emptySeats == 1 {
				fmt.Println("Found row:", i)

				for jj := 0; jj < 8; jj++ {
					if matrix[i][jj] == 0 {
						fmt.Println("Possible Seat: ", i, jj, (i * 8) + jj)
					}
				}
			}
		}
	})
}
