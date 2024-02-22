package main

import (
	"constraint_solver/eight_queens"
	"constraint_solver/map_coloring"
	"fmt"
	"slices"
)

func main() {
	runMapColoring()
	runEightQueens()

	a := []int{1, 2, 3, 4, 5}
	b := slices.Clone(a)
	fmt.Printf("%v %p\n%v %p\n%v\n", a, &a, b, &b, &a == &b)

	d := slices.DeleteFunc(b, func(i int) bool {
		return true
	})
	fmt.Printf("%v %p\n%v %p\n%v\n", b, &b, d, &d, &b == &d)
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
