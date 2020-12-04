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

				fields := strings.Split(p, " ")
				isValid := true
				for _, f := range fields {
					kv := strings.Split(f, ":")

					if kv[0] == "byr" {
						byr := toInt(kv[1])

						if byr < 1920 || byr > 2002 {
							isValid = false
							break
						}
					}
					if kv[0] == "iyr" {
						iyr := toInt(kv[1])

						if iyr < 2010 || iyr > 2020 {
							isValid = false
							break
						}
					}
					if kv[0] == "eyr" {
						eyr := toInt(kv[1])

						if eyr < 2020 || eyr > 2030 {
							isValid = false
							break
						}
					}
					if kv[0] == "hgt" {
						hgt := toInt(kv[1])

						if strings.Contains(kv[1], "cm") {
							if hgt < 150 || hgt > 193 {
								isValid = false
								break
							}
						} else {
							if hgt < 59 || hgt > 76 {
								isValid = false
								break
							}
						}
					}
					if kv[0] == "hcl" {
						re := regexp.MustCompile("^#[a-f0-9]{6}$")
						if !re.MatchString(kv[1]) {
							isValid = false
							break
						}
					}

					if kv[0] == "ecl" {
						if
							kv[1] != "amb" &&
							kv[1] != "blu" &&
							kv[1] != "brn" &&
							kv[1] != "gry" &&
							kv[1] != "grn" &&
							kv[1] != "hzl" &&
							kv[1] != "oth" {
							isValid = false
							break
						}
					}
					if kv[0] == "pid" {
						re := regexp.MustCompile("^\\d{9}$")
						if !re.MatchString(kv[1]) {
							isValid = false
							break
						}
					}
				}

				if isValid {
					validPassports += 1
				}
			}
		}

		fmt.Printf("Valid passports: %d\n", validPassports)
	})
}

func toInt(s string) int {
	re := regexp.MustCompile("[^0-9]+")
	s = re.ReplaceAllString(s, "")

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int: %v", err)
	}

	return i
}
