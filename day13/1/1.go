package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	utils.Time(func() {
		//input := utils.ReadInput("day13/input-test.txt")
		input := utils.ReadInput("day13/input.txt")

		earliestTs, err := strconv.Atoi(input[0])
		if err != nil {
			log.Fatalf("failed to parse int: %v", err)
		}

		originalTs := earliestTs
		earliestTs -= 1

		bussesStr := utils.SliceFilterString(strings.Split(input[1], ","), func(thing string) bool {
			return thing != "x"
		})

		var busses []int
		for _, item := range bussesStr {
			i, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("failed to parse int: %v", err)
			}

			busses = append(busses, i)
		}

		foundBus := false
		busNum := -1
		for !foundBus {
			earliestTs += 1
			for _, bus := range busses {
				if earliestTs % bus == 0 {
					foundBus = true
					busNum = bus
					break
				}
			}
		}

		fmt.Println("ID:", (earliestTs - originalTs) * busNum)
	})
}