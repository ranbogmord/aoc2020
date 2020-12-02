package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
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

func main() {
	utils.Time(func() {
		instr := utils.ReadInput(os.Args[1])

		var input BySize
		for _, row := range instr {
			rInt, err := strconv.Atoi(row)
			if err != nil {
				log.Fatalf("failed to parse number as int: %v", err)
			}
			input = append(input, rInt)
		}
		sort.Sort(input)

		var checkedInputs map[string]int

		val1 := -1
		val2 := -1

		for _, x := range input {
			maxVal := 2020 - x
			for _, y := range input {
				key := fmt.Sprintf("%f-%f", math.Min(float64(x), float64(y)), math.Max(float64(x), float64(y)))
				if _, ok := checkedInputs[key]; ok == true {
					continue
				}

				if y > maxVal {
					break
				}

				if x + y == 2020 {
					val1 = x
					val2 = y
					break
				}
			}

			if val1 > -1 && val2 > -1 {
				break
			}
		}

		if val1 == -1 || val2 == -1 {
			log.Fatalln("no matches found")
		}

		log.Printf("Answer is:\n%d", val1 * val2)
	})
}