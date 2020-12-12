package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
)

func countOccupiedSeats(matrix [][]string, row, col int) int {
	occupied := 0

	changes := [][]int{
		[]int{+1, 0}, // down
		[]int{-1, 0}, // up
		[]int{0, -1}, // left
		[]int{0, +1}, // right
		[]int{-1, -1}, // diag left up
		[]int{+1, -1}, // diag left down
		[]int{-1, +1}, // diag right up
		[]int{+1, +1}, // diag right down
	}

	for _, change := range changes {
		cRow, cCol := row + change[0], col + change[1]
		for cRow >= 0 && cCol >= 0 && cRow < len(matrix) && cCol < len(matrix[0]) {
			cell := matrix[cRow][cCol]
			cRow, cCol = cRow + change[0], cCol + change[1]
			if cell == "." {
				continue
			} else {
				if cell == "#" {
					occupied++
				}
				break
			}
		}
	}

	return occupied
}

func main() {
	utils.Time(func() {
		input := utils.ReadInputMatrix(os.Args[1])

		didSwitcharoo := true
		for didSwitcharoo {
			didSwitcharoo = false

			newInput := utils.DeepCopyMatrix(input)

			for y := 0; y < len(input); y++ {
				for x := 0; x < len(input[y]); x++ {
					cell := input[y][x]

					if cell == "L" {
						occupied := countOccupiedSeats(input, y, x)
						if occupied == 0 {
							newInput[y][x] = "#"
							didSwitcharoo = true
						}
					} else if cell == "#" {
						occupied := countOccupiedSeats(input, y, x)
						if occupied >= 5 {
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
