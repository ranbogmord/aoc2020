package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"math"
	"strconv"
)

func forward(amt, x, y, wx, wy int) (int, int) {
	x += wx * amt
	y += wy * amt

	return x, y
}

func manhattanDistance(x, y int) float64 {
	return math.Abs(float64(x)) + math.Abs(float64(y))
}

func pythonMod(val, modulus int) int {
	return ((val % modulus) + modulus) % modulus
}

func rotateWaypoint(mvmtType string, amt, wx, wy int) (int, int) {
	if mvmtType == "L" {
		amt = 360 - amt
	}

	for i := 0; i < (amt / 90) % 4; i++ {
		wx, wy = wy * -1, wx
	}

	return wx, wy
}

func main() {
	utils.Time(func() {
		//input := utils.ReadInput("day12/input-test.txt")
		input := utils.ReadInput("day12/input.txt")

		x := 0
		y := 0
		wx := 10
		wy := -1

		for _, inst := range input {
			mvmtType := inst[:1]
			mvmtAmt, err := strconv.Atoi(inst[1:])
			if err != nil {
				log.Fatalf("failed  to parse int: %v", err)
			}

			switch mvmtType {
			case "N":
				wy -= mvmtAmt
				break
			case "S":
				wy += mvmtAmt
				break
			case "E":
				wx += mvmtAmt
				break
			case "W":
				wx -= mvmtAmt
				break
			case "L":
				wx, wy = rotateWaypoint(mvmtType, mvmtAmt, wx, wy)
				break
			case "R":
				wx, wy = rotateWaypoint(mvmtType, mvmtAmt, wx, wy)
				break
			case "F":
				x, y = forward(mvmtAmt, x, y, wx, wy)
				break
			}
		}

		fmt.Println("W-E:", x, "N-S:", y)
		fmt.Println("Distance:", manhattanDistance(x, y))
	})
}