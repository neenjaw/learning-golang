package main

import (
	"math/rand/v2"

	"maze_solving/maze"
	"maze_solving/search"
)

func main() {
	maze.Randomizer = *rand.New(rand.NewPCG(43, 1024))
	m := maze.MakeMaze(
		10,
		10,
		maze.Location{Row: 0, Column: 0},
		maze.Location{Row: 9, Column: 9},
		0.2,
	)
	println(m.String())

	node, err := search.Dfs(
		m.Start,
		m.GoalTest,
		m.Successors,
	)

	if err != nil {
		println(err.Error())
	} else {
		m.Mark(node.Path())
		println(m.String())
		m.Clear(node.Path())
	}
}
