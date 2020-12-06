package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
	"strings"
)

type Group struct {
	Rows []string
	Chars map[string]int
}

func (g *Group) AddRow(row string) {
	if g.Chars == nil {
		g.Chars = make(map[string]int)
	}
	g.Rows = append(g.Rows, row)

	for _, char := range strings.Split(row, "") {
		if _, ok := g.Chars[char]; !ok {
			g.Chars[char] = 1
		} else {
			g.Chars[char] += 1
		}
	}
}

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		var rows []Group
		group := Group{}
		group.Chars = make(map[string]int)

		for _, in := range input {
			if in == "" {
				rows = append(rows, group)
				group = Group{}
				continue
			}

			group.AddRow(in)
		}

		sum := 0
		for _, group := range rows {
			target := len(group.Rows)

			for _, count := range group.Chars {
				if count == target {
					sum += 1
				}
			}
		}

		fmt.Println("Sum is:", sum)
	})
}
