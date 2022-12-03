package main

import (
	"bufio"
	"fmt"
	"os"
)

type Shape struct {
	name   string // The name of the hand shape
	points int    // Poins of the handshape
	beats  string
	loses  string
}

var Shapes = map[string]Shape{
	"ROCK":     {name: "ROCK", points: 1, beats: "SCISSORS", loses: "PAPER"},
	"PAPER":    {name: "PAPER", points: 2, beats: "ROCK", loses: "SCISSORS"},
	"SCISSORS": {name: "SCISSORS", points: 3, beats: "PAPER", loses: "ROCK"},
}

var opp_scoring = map[string]Shape{
	"A": Shapes["ROCK"],
	"B": Shapes["PAPER"],
	"C": Shapes["SCISSORS"],
}
var my_scoring = map[string]Shape{
	"X": Shapes["ROCK"],
	"Y": Shapes["PAPER"],
	"Z": Shapes["SCISSORS"],
}

var outcome_scoring = map[string]string{
	"X": "LOOSE",
	"Y": "DRAW",
	"Z": "WIN",
}

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	score_p1 := 0
	score_p2 := 0
	for scanner.Scan() {
		round := scanner.Text()
		opp_shape := opp_scoring[string(round[0])]
		my_shape := my_scoring[string(round[2])]
		outcome := outcome_scoring[string(round[2])]

		score_p1 += calculate_points_p1(my_shape, opp_shape)
		score_p2 += calculate_points_p2(opp_shape, outcome)
	}

	fmt.Println("total points Part 1:", score_p1)
	fmt.Println("total points Part 2:", score_p2)
}

func calculate_points_p1(me, opp Shape) int {
	if me.beats == opp.name {
		return 6 + me.points
	} else if me == opp {
		return 3 + me.points
	} else {
		return me.points
	}
}

func calculate_points_p2(opp Shape, outcome string) int {
	switch outcome {
	case "LOOSE":
		return Shapes[opp.beats].points + 0
	case "DRAW":
		return Shapes[opp.name].points + 3
	default:
		return Shapes[opp.loses].points + 6
	}
}
