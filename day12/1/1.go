package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"math"
	"strconv"
)

func forward(amt, direction, x, y int) (int, int) {
	if direction == 0 || direction == 360 {
		y -= amt
	} else if direction == 90 {
		x += amt
	} else if direction == 180 {
		y += amt
	} else if direction == 270 {
		x -= amt
	}

	return x, y
}

func manhattanDistance(x, y int) float64 {
	return math.Abs(float64(x)) + math.Abs(float64(y))
}

func pythonMod(val, modulus int) int {
	return ((val % modulus) + modulus) % modulus
}

func main() {
	utils.Time(func() {
		//input := utils.ReadInput("day12/input-test.txt")
		input := utils.ReadInput("day12/input.txt")

		x := 0
		y := 0
		facing := 90
		for _, inst := range input {
			mvmtType := inst[:1]
			mvmtAmt, err := strconv.Atoi(inst[1:])
			if err != nil {
				log.Fatalf("failed  to parse int: %v", err)
			}

			switch mvmtType {
			case "N":
				y -= mvmtAmt
				break
			case "S":
				y += mvmtAmt
				break
			case "E":
				x += mvmtAmt
				break
			case "W":
				x -= mvmtAmt
				break
			case "L":
				facing = pythonMod(facing - mvmtAmt, 360)
				break
			case "R":
				facing = pythonMod(facing + mvmtAmt, 360)
				break
			case "F":
				x, y = forward(mvmtAmt, facing, x, y)
				break
			}
		}

		fmt.Println("W-E:", x, "N-S:", y)
		fmt.Println("Distance:", manhattanDistance(x, y))
	})
}