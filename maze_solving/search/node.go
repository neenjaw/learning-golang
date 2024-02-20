package search

import "slices"

type Node[T any] struct {
	state     T
	parent    *Node[T]
	cost      float64
	heuristic float64
}

func NewNode[T any](state T, parent *Node[T]) *Node[T] {
	return &Node[T]{
		state:  state,
		parent: parent,
	}
}

func NewNodeWithCost[T any](state T, parent *Node[T], cost float64, heuristic float64) *Node[T] {
	return &Node[T]{
		state:     state,
		parent:    parent,
		cost:      cost,
		heuristic: heuristic,
	}
}

func (n *Node[T]) CompareTo(other *Node[T]) int {
	return int(n.cost + n.heuristic - other.cost - other.heuristic)
}

func (n *Node[T]) Path() []T {
	path := make([]T, 0)
	for n != nil {
		path = append(path, n.state)
		n = n.parent
	}
	slices.Reverse(path)
	return path
}
