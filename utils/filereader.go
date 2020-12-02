package utils

import (
	"io/ioutil"
	"log"
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