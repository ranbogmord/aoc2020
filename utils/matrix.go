package utils

import "fmt"

func DeepCopyMatrix(in [][]string) [][]string {
	out := make([][]string, len(in))

	for row, _ := range in {
		out[row] = append([]string{}, in[row]...)
	}

	return out
}

func PrintMatrix(matrix [][]string) {
	for row, _ := range matrix {
		for col, _ := range matrix[row] {
			fmt.Print(matrix[row][col])
		}
		fmt.Println("")
	}
	fmt.Println("-----")
}