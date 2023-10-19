package main

import (
	"sort"
)

// priorityQueue is a data structure used for storing
// positions and always fetching the item with the lowest priority number.
type priorityQueue struct {
	// TODO: Don't use an array for this
	// Use a binary tree
	items []pqItem
}

type pqItem struct {
	priority int
	pos      position
}

func (pq *priorityQueue) put(pos position, priority int) {
	pq.items = append(pq.items, pqItem{priority: priority, pos: pos})
	sort.Slice(pq.items, func(i, j int) bool {
		return pq.items[i].priority < pq.items[j].priority
	})
}

// get returns the item with the lowest priority number.
// call empty() first to ensure this queue is not empty, or
// it will panic.
func (pq *priorityQueue) get() position {
	if len(pq.items) == 0 {
		panic("Cannot get from empty priorityQueue")
	}

	pos := pq.items[0].pos
	pq.items = pq.items[1:]

	return pos
}

func (pq *priorityQueue) empty() bool {
	return len(pq.items) == 0
}
