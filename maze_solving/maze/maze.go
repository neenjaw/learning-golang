package maze

import (
	"math/rand/v2"
	"strings"
)

type Cell rune

const (
	Empty   Cell = ' '
	Blocked Cell = 'X'
	Start   Cell = 'S'
	Goal    Cell = 'G'
	Path    Cell = '*'
)

type Location struct {
	Row, Column int
}

func (ml Location) equals(other Location) bool {
	return ml.Row == other.Row && ml.Column == other.Column
}

type Maze struct {
	grid          [][]Cell
	Start, Goal   Location
	rows, columns int
}

var Randomizer rand.Rand

func MakeMaze(rows, columns int, start, goal Location, sparseness float64) *Maze {
	maze := &Maze{
		grid:    make([][]Cell, rows),
		Start:   start,
		Goal:    goal,
		rows:    rows,
		columns: columns,
	}

	for i := range maze.grid {
		maze.grid[i] = make([]Cell, columns)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			maze.grid[r][c] = Empty
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if Randomizer.Float64() < sparseness {
				maze.grid[r][c] = Blocked
			}
		}
	}

	maze.grid[start.Row][start.Column] = Start
	maze.grid[goal.Row][goal.Column] = Goal

	return maze
}

func (maze *Maze) String() string {
	var builder strings.Builder
	for r := 0; r < maze.rows; r++ {
		for c := 0; c < maze.columns; c++ {
			builder.WriteRune(rune(maze.grid[r][c]))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (maze *Maze) GoalTest(ml Location) bool {
	return ml.equals(maze.Goal)
}

func (maze *Maze) Successors(ml Location) []Location {
	successors := make([]Location, 0, 4)
	if ml.Row+1 < maze.rows && maze.grid[ml.Row+1][ml.Column] != Blocked {
		successors = append(successors, Location{ml.Row + 1, ml.Column})
	}
	if ml.Row-1 >= 0 && maze.grid[ml.Row-1][ml.Column] != Blocked {
		successors = append(successors, Location{ml.Row - 1, ml.Column})
	}
	if ml.Column+1 < maze.columns && maze.grid[ml.Row][ml.Column+1] != Blocked {
		successors = append(successors, Location{ml.Row, ml.Column + 1})
	}
	if ml.Column-1 >= 0 && maze.grid[ml.Row][ml.Column-1] != Blocked {
		successors = append(successors, Location{ml.Row, ml.Column - 1})
	}
	return successors
}

func (maze *Maze) EuclideanDistanceFromGoal(a Location) float64 {
	dx := a.Row - maze.Goal.Row
	dy := a.Column - maze.Goal.Column
	return float64(dx*dx + dy*dy)
}

func (maze *Maze) ManhattanDistanceFromGoal(a Location) float64 {
	dx := a.Row - maze.Goal.Row
	dy := a.Column - maze.Goal.Column
	return float64(dx + dy)
}

func (maze *Maze) Mark(path []Location) {
	for _, loc := range path {
		maze.grid[loc.Row][loc.Column] = Path
	}
	maze.grid[maze.Start.Row][maze.Start.Column] = Start
	maze.grid[maze.Goal.Row][maze.Goal.Column] = Goal
}

func (maze *Maze) Clear(path []Location) {
	for _, loc := range path {
		maze.grid[loc.Row][loc.Column] = Empty
	}
	maze.grid[maze.Start.Row][maze.Start.Column] = Start
	maze.grid[maze.Goal.Row][maze.Goal.Column] = Goal
}
