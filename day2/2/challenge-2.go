package main

import (
	"aoc2020/utils"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Password struct {
	Min int
	Max int
	Char string
	Word string
}

func (p Password) IsValid() bool {
	minChar := string(p.Word[p.Min - 1])
	maxChar := string(p.Word[p.Max - 1])

	if minChar == p.Char && maxChar != p.Char {
		return true
	} else if minChar != p.Char && maxChar == p.Char {
		return true
	} else {
		return false
	}
}

func parseInput(row string) *Password {
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

	return &Password{
		Min: min,
		Max: max,
		Char: char,
		Word: word,
	}
}

func main() {
	utils.Time(func () {
		input := utils.ReadInput(os.Args[1])

		valid := 0
		invalid := 0

		for _, val := range input {
			p := parseInput(val)

			if p.IsValid() {
				valid += 1
			} else {
				invalid += 1
			}
		}

		log.Printf("Valid: %d, Invalid: %d", valid, invalid)
	})
}
