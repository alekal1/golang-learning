package points

import "strings"

/**
Team's match results are recorded in a collection of strings.
Each match is represented by a string in the format "x:y",
where x is our team's score and y is our opponents score.

Points are awarded for each match as follows:

if x > y: 3 points (win)
if x < y: 0 points (loss)
if x = y: 1 point (tie)

Write a function that takes this collection
and returns the number of points our team (x) got in the championship by the rules given above
*/
func Points(games []string) int {
	var score int
	for _, game := range games {
		score += calculateWinner(game)
	}
	return score
}

func calculateWinner(game string) int {
	splitedScore := strings.Split(game, ":")
	if splitedScore[0] > splitedScore[1] {
		return 3
	}

	if splitedScore[0] == splitedScore[1] {
		return 1
	}

	return 0
}
