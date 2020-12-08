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

func parseInput(rows []string) map[string]string {
	re := regexp.MustCompile("([a-z\\s]+)\\sbags\\scontain\\s([a-z1-9\\s,]+)\\sbags?")
	re2 := regexp.MustCompile("\\s?bags?,?")
	result := make(map[string]string)
	for _, row := range rows {
		matches := re.FindStringSubmatch(row)
		contain := re2.ReplaceAllString(matches[2], ",")
		result[matches[1]] = contain
	}

	return result
}

func findBagsRec(rows map[string]string, search string) int {
	if rows[search] == "no other" {
		return 0
	}

	sum := 0
	childStr := rows[search]
	re := regexp.MustCompile("(\\d+)\\s(.+)")
	for _, child := range strings.Split(childStr, ",") {
		child = strings.Trim(child, " ")
		matches := re.FindStringSubmatch(child)

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalf("failed to parse int: %v", err)
		}

		sum += count * (findBagsRec(rows, matches[2]) + 1)
	}

	return sum
}

func main() {
	utils.Time(func() {
		searchWord := "shiny gold"
		input := utils.ReadInput(os.Args[1])
		parsed := parseInput(input)

		sum := findBagsRec(parsed, searchWord)

		fmt.Println(sum)
	})
}