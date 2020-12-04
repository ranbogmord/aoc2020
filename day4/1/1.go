package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInput(os.Args[1])

		var passports []string
		validPassports := 0
		currentPassport := ""

		for _, row := range input {
			if row == "" {
				currentPassport = strings.Trim(currentPassport, " ")
				passports = append(passports, currentPassport)
				currentPassport = ""
			}

			currentPassport += " " + row
		}

		for _, p := range passports {
			if
				strings.Contains(p, "byr:") &&
				strings.Contains(p, "iyr:") &&
				strings.Contains(p, "eyr:") &&
				strings.Contains(p, "hgt:") &&
				strings.Contains(p, "hcl:") &&
				strings.Contains(p, "ecl:") &&
				strings.Contains(p, "pid:")  {
				validPassports += 1
				fmt.Println(p)
			}

		}

		fmt.Printf("Valid passports: %d\n", validPassports)
	})
}
