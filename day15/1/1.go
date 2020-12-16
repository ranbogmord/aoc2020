package main

import (
	"aoc2020/utils"
	"fmt"
	"strings"
)

func addToMem(mem map[int][]int, num, idx int) map[int][]int {
	if _, ok := mem[num]; !ok || len(mem[num]) < 0 {
		mem[num] = append(mem[num], idx)
	} else {
		if len(mem[num]) < 2 {
			mem[num] = append(mem[num], idx)
		} else {
			mem[num] = append(mem[num][1:], idx)
		}
	}

	return mem
}

func doRun(input []string, name string, limit int) {
	lastNum := -1
	mem := make(map[int][]int)

	// init mem and nums
	for idx, val := range input {
		v := utils.ToInt(val)
		lastNum = v
		mem = addToMem(mem, v, idx)
	}

	// do the run
	for i := len(input); i < limit; i++ {
		if _, ok := mem[lastNum]; !ok || lastNum == -1 {
			mem = addToMem(mem, 0, i)
			lastNum = 0
		} else {
			num := utils.SliceMax(mem[lastNum]) - utils.SliceMin(mem[lastNum])
			mem = addToMem(mem, num, i)
			lastNum = num
		}
	}

	fmt.Println(name, lastNum)
}

func main() {
	utils.Time(func() {
		//input := strings.Split("0,3,6", ",")
		//input := strings.Split("1,3,2", ",")
		//input := strings.Split("3,1,2", ",")
		input := strings.Split("1,20,8,12,0,14", ",")

		doRun(input, "Part 1", 2020)
		doRun(input, "Part 2", 30000000)
	})
}