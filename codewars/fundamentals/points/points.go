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

const (
	SCORE_DELIMETER      = ":"
	TEAM_POINT_INDEX     = 0
	OPPONENT_POINT_INDEX = 1

	WIN_POINT  = 3
	TIE_POINT  = 1
	LOSE_POINT = 0
)

func Points(games []string) int {
	var score int
	for _, game := range games {
		score += calculateWinner(game)
	}
	return score
}

func calculateWinner(game string) int {
	splitedScore := strings.Split(game, SCORE_DELIMETER)
	x, y := splitedScore[TEAM_POINT_INDEX], splitedScore[OPPONENT_POINT_INDEX]

	if x > y {
		return WIN_POINT
	}

	if x == y {
		return TIE_POINT
	}

	return LOSE_POINT
}
