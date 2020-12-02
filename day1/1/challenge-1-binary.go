package main

import (
	"aoc2020/utils"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type bstnode struct {
	value int
	left *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst) reset() {
	b.root = nil
}

func (b *bst) insert(value int) {
	b.insertRec(b.root, value)
}

func (b *bst) insertRec(node *bstnode, value int) *bstnode {
	if b.root == nil {
		b.root = &bstnode{
			value: value,
		}
		return b.root
	}
	if node == nil {
		return &bstnode{
			value: value,
		}
	}
	if value <= node.value {
		node.left = b.insertRec(node.left, value)
	}
	if value > node.value {
		node.right = b.insertRec(node.right, value)
	}

	return node
}

func (b *bst) find(value int) *bstnode {
	return b.findRec(b.root, value)
}

func (b *bst) findRec(node *bstnode, value int) *bstnode {
	if node == nil {
		return nil
	}

	if node.value == value {
		return node
	}

	if value < node.value {
		return b.findRec(node.left, value)
	} else {
		return b.findRec(node.right, value)
	}
}

func readInput() []int {
	data, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	dataStr := string(data[:])
	rows := strings.Split(dataStr, "\n")

	var rowsInt []int
	for _, row := range rows {
		rInt, err := strconv.Atoi(row)
		if err != nil {
			log.Fatalf("failed to parse number as int: %v", err)
		}
		rowsInt = append(rowsInt, rInt)
	}

	return rowsInt
}

func main() {
	utils.Time(func () {
		input := readInput()

		val1 := -1
		val2 := -1

		tree := &bst{}

		for _, x := range input {
			tree.insert(x)
		}

		for _, x := range input {
			val := 2020 - x

			foundVal := tree.find(val)
			if foundVal != nil {
				val1 = x
				val2 = foundVal.value
				break
			}
		}

		if val1 == -1 || val2 == -1 {
			log.Fatalln("no matches found")
		}

		log.Printf("Answer is:\n%d", val1 * val2)
	})
}