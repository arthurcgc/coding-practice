package main

import "math"

func numberOfMatches(n int) int {
	matches := 0
	for currTeams := n; currTeams > 1; currTeams = currTeams / 2 {
		matches += int(math.Floor(float64(currTeams / 2)))
		if currTeams%2 != 0 {
			currTeams++
		}
	}
	return matches
}
