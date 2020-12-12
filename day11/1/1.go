package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
)

func countOccupiedSeats(matrix [][]string, startY, endY, startX, endX int, currentState string) int {
	occupied := 0
	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			if matrix[y][x] == "#" {
				occupied += 1
			}
		}
	}

	if currentState == "#" {
		occupied--
	}

	return occupied
}

func deepCopyMatrix(in [][]string) [][]string {
	out := make([][]string, len(in))

	for row, _ := range in {
		out[row] = append([]string{}, in[row]...)
	}

	return out
}

func main() {
	utils.Time(func() {
		input := utils.ReadInputMatrix(os.Args[1])

		didSwitcharoo := true
		for didSwitcharoo {
			didSwitcharoo = false

			newInput := deepCopyMatrix(input)

			for y := 0; y < len(input); y++ {
				for x := 0; x < len(input[y]); x++ {
					cell := input[y][x]

					if cell == "." {
						continue
					}

					startRow := utils.Max(y - 1, 0)
					endRow := utils.Min(y + 2, len(input))
					startCol := utils.Max(x - 1, 0)
					endCol := utils.Min(x + 2, len(input[y]))

					if cell == "L" {
						occupied := countOccupiedSeats(input, startRow, endRow, startCol, endCol, cell)
						if occupied == 0 {
							newInput[y][x] = "#"
							didSwitcharoo = true
						}
					} else if cell == "#" {
						occupied := countOccupiedSeats(input, startRow, endRow, startCol, endCol, cell)
						if occupied >= 4 {
							newInput[y][x] = "L"
							didSwitcharoo = true
						}
					}
				}
			}
			input = newInput
		}

		occupiedSeats := 0
		for row, _ := range input  {
			for col, _ := range input[row] {
				if input[row][col] == "#" {
					occupiedSeats += 1
				}
			}
		}

		fmt.Println("Occupied seats:", occupiedSeats)
	})
}
