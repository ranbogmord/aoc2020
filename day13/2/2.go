package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInput("day13/input-test.txt")
		//input := utils.ReadInput("day13/input.txt")

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
		sort.Ints(busses)

		fmt.Println(busses)

		for i := utils.SliceMax(busses); i < math.MaxInt32; i++ {

			if i % busses[0] == 0 { // if first bus mods to 0 we're at the start of the sequence
				lastBusTs := i
				lastFound := 0

				busloop:
				for j := i + 1; j < math.MaxInt32; j++ {
					for idx, bus := range busses {
						if idx == 0 {
							continue
						}
						if j % bus == 0 {
							if idx != lastFound + 1 {
								break busloop
							} else {
								lastFound = idx

								if lastFound == len(busses) - 1 {
									fmt.Println("Found sequence, TS:", lastBusTs)
								}

								break
							}
						}
					}
				}
			}
		}
	})
}