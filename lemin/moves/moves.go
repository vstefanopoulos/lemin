package moves

import (
	"fmt"
	"lemin/lemin/colony"
)

var MovesPerAnt [][]string

// Each ant takes the shortest path. The length of each path is increased as more ants walk it
// The outer slice of the return represents each ant and the inner slice the rooms that it moves to
func Make(numAnts int, paths [][]*colony.Room) {
	pathData := createPD(paths)

	MovesPerAnt = make([][]string, numAnts)

	for i := 0; i < numAnts; i++ {
		minIndex := bestRated(pathData)

		current := pathData[minIndex]
		antMoves := make([]string, current.rating-len(current.path))
		for _, room := range current.path {
			antMoves = append(antMoves, fmt.Sprintf("L%d-%s", i+1, room.Name))
		}
		MovesPerAnt[i] = antMoves
		pathData[minIndex].rating++
	}
}

// Returns the index of the best rated path in pathList
func bestRated(paths []pathData) (minIndex int) {
	minIndex = 0
	for idx, p := range paths {
		if paths[minIndex].rating >= p.rating {
			minIndex = idx
		}
	}
	return
}
