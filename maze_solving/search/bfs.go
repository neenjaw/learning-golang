package search

import (
	"errors"

	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func Bfs[T any](initial T, goalTest func(T) bool, successors func(T) []T) (*Node[T], error) {
	frontier := arrayqueue.New()
	frontier.Enqueue(NewNode(initial, nil))

	explored := hashset.New()
	explored.Add(initial)

	for frontier.Size() > 0 {
		currentMaybeNode, ok := frontier.Dequeue()
		if !ok {
			return nil, errors.New("unable to dequeue from frontier queue")
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
			frontier.Enqueue(NewNode(child, currentNode))
		}
	}

	return nil, nil
}
