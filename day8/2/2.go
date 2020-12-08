package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInstruction(inst string) (string, int, error) {
	parts := strings.Split(inst, " ")
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, err
	}

	return parts[0], val, nil
}

func checkInstructions(instructions []string) (bool, int) {
	acc := 0
	seenIdx := make(map[int]bool)
	idx := 0

	for {
		if _, ok := seenIdx[idx]; ok {
			return false, acc
		}

		if idx >= len(instructions) {
			return true, acc
		}

		instr := instructions[idx]
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
}

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		for idx, row := range input {
			if strings.Contains(row, "jmp") || strings.Contains(row, "nop") {
				newInput := utils.ReadInput(os.Args[1])
				if strings.Contains(row, "jmp") {
					newInput[idx] = strings.Replace(row, "jmp", "nop", -1)
				} else if strings.Contains(row, "nop") {
					newInput[idx] = strings.Replace(row, "nop", "jmp", -1)
				}

				done, acc := checkInstructions(newInput)
				if done {
					fmt.Println("Acc =", acc)
				}
			}
		}
	})
}
