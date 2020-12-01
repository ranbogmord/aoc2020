package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type BySize []int

func (b BySize) Len() int {
	return len(b)
}

func (b BySize) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b BySize) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func readInput() BySize {
	data, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	dataStr := string(data[:])
	rows := strings.Split(dataStr, "\n")

	var rowsInt []int
	for _, row := range rows {
		rInt, err := strconv.Atoi(row)
		if err != nil {
			log.Fatalf("failed to parse number as int: %v", err)
		}
		rowsInt = append(rowsInt, rInt)
	}

	sortedRows := BySize(rowsInt)
	sort.Sort(sortedRows)
	return sortedRows
}

func main() {
	input := readInput()

	val1 := -1
	val2 := -1
	val3 := -1

	for _, x := range input {
		for _, y := range input {
			if x + y  > 2020 {
				break
			}

			for _, z := range input {
				if x + y + z > 2020 {
					break
				}

				if x + y + z == 2020 {
					val1 = x
					val2 = y
					val3 = z
					break
				}
			}
		}

		if val1 > -1 && val2 > -1 && val3 > -1 {
			break
		}
	}

	if val1 == -1 || val2 == -1 || val3 == -1 {
		log.Fatalln("no matches found")
	}

	log.Printf("Answer is:\n%d", val1 * val2 * val3)
}