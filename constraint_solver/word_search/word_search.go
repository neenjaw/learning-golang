// Description: This package enables the creation of word search puzzles.
//
// The word search puzzle is a grid of random letters with words placed in the grid.
// The words can be placed horizontally, vertically, or diagonally, but they do not overlap.
// The placement of the words in the grid are determined by a constraint satisfaction problem (CSP) solver.

package word_search

import (
	"constraint_solver/csp"
	"math/rand"

	"github.com/emirpasic/gods/sets/hashset"
)

type WordSearchConstraint struct {
	words []string
}

func (w WordSearchConstraint) Variables() []string {
	return w.words
}

// Satisfied checks if the assignment is satisfied
// by ensuring that there are no overlapping locations
// for the words in the grid
func (w WordSearchConstraint) Satisfied(assignment map[string][]GridLocation) bool {
	allLocations := []GridLocation{}
	uniqueLocations := hashset.New()

	for _, locations := range assignment {
		for _, location := range locations {
			allLocations = append(allLocations, location)
			uniqueLocations.Add(location)
		}
	}

	return len(allLocations) == uniqueLocations.Size()
}

func WordSearchCSP(
	grid *WordGrid,
	wordList []string,
	random *rand.Rand,
) (*csp.CSPSolver[string, []GridLocation], error) {
	domains := make(map[string][][]GridLocation)
	for _, word := range wordList {
		domains[word] = grid.generateDomain(word)
	}

	csp, err := csp.NewCSPSolver(
		wordList,
		domains,
	)
	if err != nil {
		return nil, err
	}

	constraint := WordSearchConstraint{wordList}
	csp.AddConstraint(constraint)

	return csp, nil
}
