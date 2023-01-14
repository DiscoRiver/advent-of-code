package main

import (
	"bytes"
	"fmt"
	"os"
)

var alpha = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	input, err := os.ReadFile("2022/3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(input, []byte("\n"))

	var totalPrio int
	for i := range lines {
		firstPack := lines[i][:(len(lines[i]))/2]
		secondPack := lines[i][(len(lines[i]))/2 : len(lines[i])]

		common := map[byte]struct{}{}
		for a := range firstPack {
			for b := range secondPack {
				if firstPack[a] == secondPack[b] {
					common[firstPack[a]] = struct{}{}
				}
			}
		}

		stepPrio := 0
		for k := range common {
			for d := range alpha {
				if k == alpha[d] {
					stepPrio += d + 1
				}
			}
		}
		totalPrio += stepPrio
		fmt.Printf("Value: %s (%d), Split: (%d)(%d), Common values: %v, Priority: %d, Total Priority: %d\n", lines[i], len(lines[i]), len(firstPack), len(secondPack), common, stepPrio, totalPrio)
	}
	fmt.Println(totalPrio)
}
