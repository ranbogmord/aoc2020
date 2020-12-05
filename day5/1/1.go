package main

import (
	"aoc2020/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		maxId := 0

		for _, code := range input {
			min := 0.0
			max := 127.0

			rMin := 0.0
			rMax := 7.0

			for _, char := range strings.Split(code, "") {
				if char == "F" {
					max = math.Floor(max - ((max - min) / 2))
				} else if char == "B" {
					min = math.Ceil(max - ((max - min) / 2))
				} else if char == "L" {
					rMax = math.Floor(rMax - ((rMax - rMin) / 2))
				} else if char == "R" {
					rMin = math.Ceil(rMax - ((rMax - rMin) / 2))
				}
			}

			// fmt.Println(min, rMin, (min * 8) + rMin)
			id := int((min * 8) + rMin)
			if id > maxId {
				maxId = id
			}
		}

		fmt.Println("Max ID: ", maxId)
	})
}
