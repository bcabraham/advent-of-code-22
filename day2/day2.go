package day2

import (
	"advent-of-code-22/lib"
	"fmt"
	"log"
	"strings"
)

/*
A for Rock
B for Paper
C for Scissors

X for Rock
Y for Paper
Z for Scissors

Rock defeats Scissors
Scissors defeats Paper
Paper defeats Rock
If both players choose the same shape, the round instead ends in a draw.

The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

Your total score is the sum of your scores for each round.
*/

var (
	playerOneMap = map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	playerTwoMap = map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	choiceScoresMap = map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}
	resultScoresMap = map[string]int{
		"WIN":  6,
		"DRAW": 3,
		"LOSE": 0,
	}

	playerTwoResults = map[string]string{
		"X": "LOSE",
		"Y": "DRAW",
		"Z": "WIN",
	}
)

func Run() {
	rpsStrategy, err := lib.ReadLines("day2", "rock-paper-scissors.txt")
	lib.HandleError(err)

	WithPlayerChoice(rpsStrategy)
	WithPlayerOutcome(rpsStrategy)
}

func WithPlayerOutcome(rpsStrategy []string) {
	totalScore := 0
	fmt.Println("WithPlayerOutcome...")
	for i, round := range rpsStrategy {
		choices := strings.Split(round, " ")

		if len(choices) != 2 {
			log.Fatalf("bad data: %+v -> %+v", round, choices)
		}

		p1Choice := playerOneMap[choices[0]]
		p2Outcome := playerTwoResults[choices[1]]
		p2Choice := getPlayerChoice(p1Choice, p2Outcome)

		result := getRoundResult(p1Choice, p2Choice)
		score := calcRoundScore(p2Choice, result)
		totalScore += score
		fmt.Printf("Round %d: %s - %s = %s (%s) | %d\n", i+1, p1Choice, p2Choice, result, p2Outcome, score)
	}

	fmt.Printf("Total score: %d\n", totalScore)
}

func WithPlayerChoice(rpsStrategy []string) {
	totalScore := 0
	fmt.Println("WithPlayerChoice...")
	for i, round := range rpsStrategy {
		choices := strings.Split(round, " ")

		if len(choices) != 2 {
			log.Fatalf("bad data: %+v -> %+v", round, choices)
		}

		p1Choice := playerOneMap[choices[0]]
		p2Choice := playerTwoMap[choices[1]]

		result := getRoundResult(p1Choice, p2Choice)
		score := calcRoundScore(p2Choice, result)
		totalScore += score
		fmt.Printf("Round %d: %s - %s = %s | %d\n", i+1, p1Choice, p2Choice, result, score)
	}

	fmt.Printf("Total score: %d\n", totalScore)
}

func getRoundResult(p1, p2 string) string {
	if p1 == p2 {
		return "DRAW"
	}

	switch p1 {
	case "Rock":
		if p2 == "Scissors" {
			return "LOSE"
		}
	case "Paper":
		if p2 == "Rock" {
			return "LOSE"
		}
	case "Scissors":
		if p2 == "Paper" {
			return "LOSE"
		}
	}

	return "WIN"
}

// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func calcRoundScore(p2 string, result string) int {
	return choiceScoresMap[p2] + resultScoresMap[result]
}

func getPlayerChoice(p1 string, result string) string {
	if result == "WIN" {
		switch p1 {
		case "Rock":
			return "Paper"
		case "Paper":
			return "Scissors"
		case "Scissors":
			return "Rock"
		}
	} else if result == "LOSE" {
		switch p1 {
		case "Rock":
			return "Scissors"
		case "Paper":
			return "Rock"
		case "Scissors":
			return "Paper"
		}
	}

	return p1
}
