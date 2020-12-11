package main

import (
	"aoc2020/utils"
	"fmt"
	"os"
)

func main() {
	utils.Time(func() {
		input := utils.ReadInputInts(os.Args[1])

		tree := utils.BTS{}

		for _, val := range input {
			tree.Insert(val)
		}

		adapterDiffCount := make(map[int]int)
		adapterDiffCount[1] = 0
		adapterDiffCount[2] = 0
		adapterDiffCount[3] = 1 // Output can always handle +3 jolts

		var adapterList []int

		for len(input) != len(adapterList) {
			lastAdapter := 0
			if len(adapterList) > 0 {
				lastAdapter = adapterList[len(adapterList) - 1]
			}

			for i := 1; i < 4; i++ {
				val := tree.Find(lastAdapter + i)
				if val != nil {
					adapterList = append(adapterList, lastAdapter + i)
					adapterDiffCount[i] += 1
					break
				}
			}
		}

		fmt.Println("Differences product", adapterDiffCount[1] * adapterDiffCount[3])
	})
}
