package main

import (
	"aoc2020/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	Min int
	Max int
	Char string
	Word string
}

func parseInput(row string) (int, int, string, string) {
	re := regexp.MustCompile("(\\d+)-(\\d+)\\s(\\w):\\s(.*)")
	matches := re.FindStringSubmatch(row)
	if len(matches) == 0 {
		log.Fatalf("failed to parse row: '%s'", row)
	}

	min, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatalf("failed to convert to int: %v", err)
	}

	max, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatalf("failed to convert to int: %v", err)
	}

	char := matches[3]
	word := matches[4]

	return min, max, char, word
}

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		valid := 0
		invalid := 0

		for _, val := range input {
			min, max, char, word := parseInput(val)

			instances := strings.Count(word, char)
			if instances >= min && instances <= max {
				valid += 1
			} else {
				invalid += 1
			}
		}

		fmt.Printf("Valid: %d, Invalid: %d\n", valid, invalid)
	})
}
