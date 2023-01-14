// Code here from https://www.reddit.com/r/adventofcode/comments/zb865p/comment/iz4qzyp/?utm_source=share&utm_medium=web2x&context=3
//
// Running on no sleep so decided to just learn from this one.
package main

import (
	"fmt"
	"os"
	"strings"
)

func charToInt(c rune) int {
	if int(c) > 90 {
		// lower case
		return int(c) - 96
	} else {
		// upper case
		return int(c) - 38
	}
}

func partOne(input []byte) {
	rucksack := strings.Split(strings.TrimSpace(string(input)), "\n")
	partOneTotal := 0
	for _, s := range rucksack {
		sideOne := s[len(s)/2:]
		sideTwo := s[:len(s)/2]
		for _, c := range sideOne {
			if strings.Contains(sideTwo, string(c)) {
				partOneTotal += charToInt(c)
				break
			}
		}
	}
	fmt.Println(partOneTotal)
}

func partTwo(input []byte) {
	rucksack := strings.Split(strings.TrimSpace(string(input)), "\n")
	m := make(map[rune][3]int)
	partTwoTotal := 0
	groupPointer := 0
	for _, s := range rucksack {
		for _, c := range s {
			if val, ok := m[c]; ok {
				val[groupPointer] = 1
				m[c] = val
				if val[0] == 1 && val[1] == 1 && val[2] == 1 {
					partTwoTotal += charToInt(c)
					break
				}
			} else {
				m[c] = [3]int{0, 0, 0}
				val := m[c]
				val[groupPointer] = 1
				m[c] = val
			}
		}
		groupPointer += 1
		// reset every 3 strings
		if groupPointer == 3 {
			groupPointer = 0
			m = make(map[rune][3]int)
		}
	}
	fmt.Println(partTwoTotal)
}

func main() {
	input, _ := os.ReadFile("2022/3/input.txt")
	partOne(input)
	partTwo(input)
}
