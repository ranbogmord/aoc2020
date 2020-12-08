package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		acc := 0
		seenIdx := make(map[int]bool)
		idx := 0

		for {
			instr := input[idx]
			if _, ok := seenIdx[idx]; ok {
				fmt.Println("Double hit")
				fmt.Println("idx =", idx)
				fmt.Println("instr =", instr)
				fmt.Println("acc =", acc)
				break
			}

			inst, val, err := parseInstruction(instr)
			if err != nil {
				log.Fatalf("error parsing instruction: %v", err)
			}

			seenIdx[idx] = true
			if inst == "acc" {
				acc += val
				idx += 1
			} else if inst == "jmp" {
				idx += val
			} else {
				idx += 1
			}
		}
	})
}

func parseInstruction(inst string) (string, int, error) {
	parts := strings.Split(inst, " ")
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, err
	}

	return parts[0], val, nil
}
