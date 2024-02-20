package search

import (
	"errors"

	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func comparator[T any](a, b interface{}) int {
	return a.(*Node[T]).CompareTo(b.(*Node[T]))
}

func Astar[T any](
	initial T,
	goalTest func(T) bool,
	successors func(T) []T,
	heuristic func(T) float64,
) (*Node[T], error) {
	frontier := priorityqueue.NewWith(comparator[T])
	frontier.Enqueue(NewNodeWithCost(initial, nil, 0, heuristic(initial)))

	explored := hashmap.New()
	explored.Put(initial, 0.0)

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
			newCost := currentNode.cost + 1
			if cost, ok := explored.Get(child); ok && cost.(float64) <= newCost {
				continue
			}

			explored.Put(child, newCost)
			frontier.Enqueue(
				NewNodeWithCost(
					child,
					currentNode,
					newCost,
					heuristic(child),
				),
			)
		}
	}

	return nil, nil
}
