package main

import (
	"math/rand/v2"

	"maze_solving/maze"
	"maze_solving/missionaries"
	"maze_solving/search"
)

func main() {
	runMazes()
	runMissionaries()
}

func runMazes() {
	println("---- Maze ----")

	m := maze.MakeMaze(
		10,
		10,
		maze.Location{Row: 0, Column: 0},
		maze.Location{Row: 9, Column: 9},
		0.2,
		rand.New(rand.NewPCG(43, 1024)),
	)
	println(m.String())

	println("---- DFS ----")

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

	println("---- BFS ----")

	node, err = search.Bfs(
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

	println("---- A* ----")

	node, err = search.Astar(
		m.Start,
		m.GoalTest,
		m.Successors,
		m.ManhattanDistanceFromGoal,
	)

	if err != nil {
		println(err.Error())
	} else {
		m.Mark(node.Path())
		println(m.String())
		m.Clear(node.Path())
	}
}

func runMissionaries() {
	println("---- Missionaries ----")

	start := missionaries.NewMCState(3, 3, true)

	println("---- BFS ----")

	node, err := search.Bfs[missionaries.MCState](
		start,
		missionaries.GoalTest,
		missionaries.Successors,
	)

	if err != nil {
		println(err.Error())
	} else {
		missionaries.DisplaySolution(node.Path())
	}
}
