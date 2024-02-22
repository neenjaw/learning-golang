package word_search

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type GridLocation struct {
	row, column int
}

const (
	alphabetSize = 26
	firstLetter  = 'A'
)

type WordGrid struct {
	rows, columns int
	grid          [][]rune
}

func NewWordGrid(rows, columns int, random *rand.Rand) (*WordGrid, error) {
	if random == nil {
		random = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	if rows < 1 || columns < 1 {
		return nil, errors.New("invalid grid size")
	}

	grid := make([][]rune, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]rune, columns)
		for column := 0; column < columns; column++ {
			// Fill the grid with random letters
			letter := rune(firstLetter + random.Intn(alphabetSize))
			grid[row][column] = letter
		}
	}

	return &WordGrid{rows, columns, grid}, nil
}

func (w *WordGrid) Mark(word string, locations []GridLocation) {
	for i, r := range word {
		location := locations[i]
		w.grid[location.row][location.column] = r
	}
}

func (w *WordGrid) String() string {
	var result strings.Builder
	for row := 0; row < w.rows; row++ {
		for column := 0; column < w.columns; column++ {
			result.WriteRune(w.grid[row][column])
		}
		result.WriteRune('\n')
	}
	return result.String()
}
