package main

import (
	"constraint_solver/eight_queens"
	"constraint_solver/map_coloring"
	"constraint_solver/word_search"
	"fmt"
)

func main() {
	runMapColoring()
	runEightQueens()
	runWordSearch()
}

func runMapColoring() {
	fmt.Println("--- Map coloring problem ---")
	mapColoringCSP := map_coloring.MapColoringCSP()
	solution := mapColoringCSP.BacktrackingSearch(make(map[string]string))

	if solution == nil {
		println("No solution found")
	} else {
		for k, v := range solution {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
}

func runEightQueens() {
	fmt.Println("--- Eight queens problem ---")
	eightQueensCSP := eight_queens.EightQueensCSP()
	solution := eightQueensCSP.BacktrackingSearch(make(map[int]int))

	if solution == nil {
		println("No solution found")
	} else {
		for rank := 0; rank < 8; rank++ {
			for file := 0; file < 8; file++ {
				if q, ok := solution[rank]; ok && q == file {
					print("Q ")
				} else {
					print("* ")
				}
			}
			println()
		}
	}
}

func runWordSearch() {
	fmt.Println("--- Word search problem ---")

	// random := rand.New(rand.NewSource(0))
	grid, err := word_search.NewWordGrid(9, 9, nil)
	if err != nil {
		println(err)
		return
	}

	words := []string{
		"MATTHEW",
		"JOE",
		"MARY",
		"SARAH",
		"SALLY",
	}

	wordSearchCSP, _ := word_search.WordSearchCSP(
		grid,
		words,
		nil,
	)
	solution := wordSearchCSP.BacktrackingSearch(make(map[string][]word_search.GridLocation))

	if solution == nil {
		println("No solution found")
	} else {
		for word, locations := range solution {
			grid.Mark(word, locations)
		}
		println(grid.String())
	}
}
