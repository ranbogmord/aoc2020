package main

import (
	"aoc2020/utils"
	"log"
	"os"
	"strconv"
)

func main() {
	utils.Time(func () {
		instr := utils.ReadInput(os.Args[1])

		var input []int
		for _, row := range instr {
			rInt, err := strconv.Atoi(row)
			if err != nil {
				log.Fatalf("failed to parse number as int: %v", err)
			}
			input = append(input, rInt)
		}

		val1 := -1
		val2 := -1

		tree := &utils.BTS{}

		for _, x := range input {
			tree.Insert(x)
		}

		for _, x := range input {
			val := 2020 - x

			foundVal := tree.Find(val)
			if foundVal != nil {
				val1 = x
				val2 = foundVal.Value
				break
			}
		}

		if val1 == -1 || val2 == -1 {
			log.Fatalln("no matches found")
		}

		log.Printf("Answer is:\n%d", val1 * val2)
	})
}