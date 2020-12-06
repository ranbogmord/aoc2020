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

		currentRow := ""
		var rows []string

		for _, in := range input {
			if in == "" {
				rows = append(rows, currentRow)
				currentRow = ""
			}

			currentRow += in
		}

		sum := 0
		for _, row := range rows {
			seenChars := ""
			count := 0
			for _, char := range strings.Split(row, "") {
				if !strings.Contains(seenChars, char) {
					count += 1
					seenChars = seenChars + char
				}
			}
			sum += count
			fmt.Println(row)
			fmt.Println(count)
		}

		fmt.Println("Sum is:", sum)
	})
}
