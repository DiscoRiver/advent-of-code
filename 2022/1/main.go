package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Elves struct {
	Population []Elf
	CalorieMax int
}

type Elf struct {
	Items         []Item
	TotalCalories int
}

type Item struct {
	Calories int
}

func (e *Elves) FindHighCalorieElvesTotal(depth int) int {
	var ar []int
	for i := range e.Population {
		if len(ar) == 0 {
			ar = append(ar, e.Population[i].TotalCalories)
		} else {
			for j := range ar {
				if e.Population[i].TotalCalories > ar[j] {
					fmt.Println(ar)
					var shiftAfter []int
					copy(shiftAfter, ar[j:])
					fmt.Println(shiftAfter)

					var newArray []int
					newArray = append(newArray, ar[:j]...)
					newArray = append(newArray, e.Population[i].TotalCalories)
					newArray = append(newArray, shiftAfter...)
					copy(ar, newArray)
					break
				}
			}
		}
	}
	fmt.Println(ar)

	total := 0
	for i := range ar {
		total += ar[i]
	}
	return total
}

func (e *Elf) AddItems(b [][]byte) {
	totalCalories := 0
	for i := range b {
		calories, _ := strconv.Atoi(string(b[i]))
		totalCalories += calories
		e.Items = append(e.Items, Item{Calories: calories})
	}
	e.TotalCalories = totalCalories
}

func main() {
	input, err := os.ReadFile("2022/1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	Elves := &Elves{}

	lines := bytes.Split(input, []byte("\n"))
	lastBlankIndex := 0
	for i := range lines {
		if string(lines[i]) == "" {
			Elf := Elf{}
			Elf.AddItems(lines[lastBlankIndex+1 : i])
			Elves.Population = append(Elves.Population, Elf)
			lastBlankIndex = i

			if Elf.TotalCalories > Elves.CalorieMax {
				Elves.CalorieMax = Elf.TotalCalories
			}
		}
	}

	fmt.Printf("%d", Elves.FindHighCalorieElvesTotal(3))
}
