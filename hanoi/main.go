package main

import "fmt"

type Stack []int

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, true
}

type HanoiGame struct {
	TowerA    Stack
	TowerB    Stack
	TowerC    Stack
	num_disks int
}

func NewHanoiGame(num_disks int) HanoiGame {
	game := HanoiGame{
		TowerA:    Stack{},
		TowerB:    Stack{},
		TowerC:    Stack{},
		num_disks: num_disks,
	}

	for i := 0; i < num_disks; i++ {
		game.TowerA.Push(i + 1)
	}

	return game
}

func (h *HanoiGame) Move(source, target, auxiliary *Stack, n int) {
	if n == 1 {
		value, _ := source.Pop()
		target.Push(value)
		return
	}

	h.Move(source, auxiliary, target, n-1)
	h.Move(source, target, auxiliary, 1)
	h.Move(auxiliary, target, source, n-1)
}

func main() {
	game := NewHanoiGame(3)

	fmt.Println("Initial State")
	fmt.Println("Tower A:", game.TowerA)
	fmt.Println("Tower B:", game.TowerB)
	fmt.Println("Tower C:", game.TowerC)

	game.Move(&game.TowerA, &game.TowerC, &game.TowerB, game.num_disks)

	fmt.Println("Final State")
	fmt.Println("Tower A:", game.TowerA)
	fmt.Println("Tower B:", game.TowerB)
	fmt.Println("Tower C:", game.TowerC)
}
