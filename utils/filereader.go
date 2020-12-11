package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInput(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	dataStr := string(data[:])
	rows := strings.Split(dataStr, "\n")

	return rows
}

func ReadInputInts(path string) []int {
	input := ReadInput(path)
	var inputs []int
	for _, val := range input {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("failed to parse int %v", err)
		}

		inputs = append(inputs, i)
	}

	return inputs
}

func ReadInputMatrix(path string) [][]string {
	var inputs [][]string
	input := ReadInput(path)

	for _, row := range input {
		inputs = append(inputs, strings.Split(row, ""))
	}

	return inputs
}