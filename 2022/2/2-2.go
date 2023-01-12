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

	// Outcome
	"X": 1, // Lose
	"Y": 2, // Draw
	"Z": 3, // Win
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
		you := 0
		outcomeNeeded := Scores[string(matchup[1])]

		// Place ring at opponents shape if not already
		if r.Prev().Value.(int) == op {
			r = r.Prev()
		} else if r.Next().Value.(int) == op {
			r = r.Next()
		}

		if outcomeNeeded == 1 { // Lose
			you = r.Prev().Value.(int)
			score += 0 + you
		} else if outcomeNeeded == 2 { // Draw
			you = r.Value.(int)
			score += 3 + you
		} else if outcomeNeeded == 3 { // Win
			you = r.Next().Value.(int)
			score += 6 + you
		}
		fmt.Printf("Opponent: %d, You: %d, Outcome Needed: %s, Points total: %d\n", op, you, outcomeNeeded, score)
	}

	fmt.Println(score)
}
