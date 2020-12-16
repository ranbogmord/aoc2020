package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func applyMask(mask, number string) string {
	maskSlice := strings.Split(mask, "")
	numSlice := strings.Split(number, "")
	for idx, char := range maskSlice {
		if char != "X" {
			numSlice[idx] = char
		}
	}

	return strings.Join(numSlice, "")
}

func setMemory(mem map[int]int64, mask string, index, value int) map[int]int64 {
	numString := fmt.Sprintf("%036b", value)
	numString = applyMask(mask, numString)
	newVal, err := strconv.ParseInt(numString, 2, 64)
	if err != nil {
		log.Fatalf("failed to cast binary string to int: %v", err)
	}

	mem[index] = newVal

	return mem
}

func main() {
	utils.Time(func() {
		//input := utils.ReadInput("day14/input-test.txt")
		input := utils.ReadInput("day14/input.txt")

		mask := ""
		mem := make(map[int]int64)
		re := regexp.MustCompile("mem\\[(\\d+)\\]\\s+=\\s+(\\d+)")

		for _, row := range input {
			if row[:4] == "mask" {
				mask = row[7:]
			} else if row[:3] == "mem" {
				matches := re.FindStringSubmatch(row)
				addr := utils.ToInt(matches[1])
				val := utils.ToInt(matches[2])

				mem = setMemory(mem, mask, addr, val)
			}
		}

		sum := int64(0)
		for _, val := range mem {
			sum += val
		}

		fmt.Println("Sum:", sum)
	})
}