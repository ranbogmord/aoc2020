package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func SliceContainsString(slice []string, str string) bool {
	sliceString := strings.Join(slice, "|")
	return strings.Contains(sliceString, str)
}

func createAddresses(address []string) [][]string {
	if !SliceContainsString(address, "X") {
		return [][]string{address}
	}

	var ret [][]string
	for idx, char := range address {
		if char == "X" {
			ones := append([]string{}, address...)
			zeroes := append([]string{}, address...)
			ones[idx] = "1"
			zeroes[idx] = "0"

			onesResult := createAddresses(ones)
			zeroesResult := createAddresses(zeroes)

			for _, val := range onesResult {
				ret = append(ret, val)
			}

			for _, val := range zeroesResult {
				ret = append(ret, val)
			}

			break
		}
	}

	return ret
}

func applyMask(mask, number string) []int64 {
	maskSlice := strings.Split(mask, "")
	numSlice := strings.Split(number, "")

	for idx, char := range maskSlice {
		if char != "0" {
			numSlice[idx] = char
		}
	}

	var indices []int64
	addresses := createAddresses(numSlice)
	for _, addr := range addresses {
		a := strings.Join(addr, "")
		newVal, err := strconv.ParseInt(a, 2, 64)
		if err != nil {
			log.Fatalf("failed to cast binary string to int: %v", err)
		}

		indices = append(indices, newVal)
	}

	return indices
}

func setMemory(mem map[int64]int64, mask string, index, value int) map[int64]int64 {
	numString := fmt.Sprintf("%036b", index)
	indices := applyMask(mask, numString)

	for _, val := range indices {
		mem[val] = int64(value)
	}

	return mem
}

func main() {
	utils.Time(func() {
		//input := utils.ReadInput("day14/input-test2.txt")
		input := utils.ReadInput("day14/input.txt")

		mask := ""
		mem := make(map[int64]int64)
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