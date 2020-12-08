package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseInput(rows []string) map[string]string {
	re := regexp.MustCompile("([a-z\\s]+)\\sbags\\scontain\\s([a-z1-9\\s,]+)\\sbags?")
	re2 := regexp.MustCompile("\\s?bags?,?")
	result := make(map[string]string)
	for _, row := range rows {
		matches := re.FindStringSubmatch(row)
		contain := re2.ReplaceAllString(matches[2], ",")
		//contain := strings.Replace(matches[2], "bags", "", -1)
		//contain = strings.Replace(contain, "bag", "", -1)
		result[matches[1]] = contain
	}

	return result
}

func findBagsRec(rows map[string]string, search string) string {
	matches := ""
	for bag, children := range rows {
		if strings.Contains(children, search) {
			matches += fmt.Sprintf("%s,", findBagsRec(rows, bag))
		}
	}

	matches += fmt.Sprintf("%s,", search)
	return matches
}

func main() {
	utils.Time(func() {
		searchWord := "shiny gold"
		input := utils.ReadInput(os.Args[1])
		parsed := parseInput(input)

		found := findBagsRec(parsed, searchWord)
		uniqueFound := make(map[string]bool)
		for _, f := range strings.Split(found, ",") {
			if f != "" && f != searchWord {
				uniqueFound[f] = true
			}
		}
		fmt.Println(uniqueFound)
		fmt.Println(len(uniqueFound))
	})
}