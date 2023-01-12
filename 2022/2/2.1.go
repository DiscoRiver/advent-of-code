package main

import (
	"bytes"
	"container/ring"
	"fmt"
	"os"
)

var Scores map[string]int = map[string]int{
	// Opponent
	"A": 1,
	"B": 2,
	"C": 3,

	// You
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	input, err := os.ReadFile("2022/2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := bytes.Split(input, []byte("\n"))

	// Roshambo has a circle outcome: S -cuts-> P -covers-> R -smashes-> S -cuts-> P ... and so on.
	r := ring.New(3)
	for i := 0; i < 3; i++ {
		r.Value = i + 1
		r = r.Next()
	}

	score := 0
	for i := range lines {
		matchup := bytes.Split(lines[i], []byte(" "))

		op := Scores[string(matchup[0])]
		you := Scores[string(matchup[1])]

		// Place ring at opponents shape if not already
		if r.Prev().Value.(int) == op {
			r = r.Prev()
		} else if r.Next().Value.(int) == op {
			r = r.Next()
		}

		outcome := ""
		if you == r.Value { // Draw
			score += 3 + you
		} else if you == r.Prev().Value.(int) { // Lose
			score += 0 + you
		} else if you == r.Next().Value.(int) { // Win
			score += 6 + you
		}
		fmt.Printf("Opponent: %d, You: %d, Current Ring Value: %d\n --> Outcome: %s, Points total: %d\n", op, you, r.Value, outcome, score)
	}

	fmt.Println("Final score:", score)
}
