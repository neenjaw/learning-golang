package search

import (
	"errors"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
)

func Dfs[T any](initial T, goalTest func(T) bool, successors func(T) []T) (*Node[T], error) {
	frontier := arraystack.New()
	frontier.Push(NewNode(initial, nil))

	explored := hashset.New()
	explored.Add(initial)

	for frontier.Size() > 0 {
		currentMaybeNode, ok := frontier.Pop()
		if !ok {
			return nil, errors.New("unable to pop from frontier stack")
		}

		currentNode := currentMaybeNode.(*Node[T])
		currentState := currentNode.state

		if goalTest(currentState) {
			return currentNode, nil
		}

		for _, child := range successors(currentState) {
			if explored.Contains(child) {
				continue
			}

			explored.Add(child)
			frontier.Push(NewNode(child, currentNode))
		}
	}

	return nil, nil
}
