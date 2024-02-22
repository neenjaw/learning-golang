package eight_queens

import (
	"constraint_solver/csp"

	"github.com/adam-lavrik/go-imath/ix"
)

type EightQueensConstraint struct {
	columns []int
}

func (e EightQueensConstraint) Variables() []int {
	return e.columns
}

func (e EightQueensConstraint) Satisfied(assignment map[int]int) bool {
	for q1c, q1r := range assignment {
		for q2c := q1c + 1; q2c < len(e.columns); q2c++ {
			if q2r, ok := assignment[q2c]; ok {
				if q1r == q2r || ix.Abs(q1r-q2r) == ix.Abs(q2c-q1c) {
					return false
				}
			}
		}
	}
	return true
}

func EightQueensCSP() *csp.CSPSolver[int, int] {
	variables := make([]int, 8)
	for i := range variables {
		variables[i] = i
	}
	domains := make(map[int][]int)
	for _, variable := range variables {
		domains[variable] = make([]int, 8)
		for i := range domains[variable] {
			domains[variable][i] = i
		}
	}
	csp, _ := csp.NewCSPSolver(variables, domains)
	csp.AddConstraint(EightQueensConstraint{variables})

	return csp
}
