package main

import (
	"fmt"
)

const size = 20

func NewTopNPriorityQueue() *topNPriorityQueue {
	q := &topNPriorityQueue{}

	// 20 new priority-0 elements
	for i := 0; i < size; i++ {
		q.queue[i] = &element{}
	}
	return q
}

type element struct {
	key      []byte
	priority int
}

// topNPriorityQueue is a modified version of the common priority queue
// who stores only the top 20 elements and ignores the others
type topNPriorityQueue struct {
	queue [size]*element
}

// every Add keeps the array sorted
func (q *topNPriorityQueue) Add(key []byte, priority int) {
	if priority <= q.queue[0].priority {
		return
	}

	q.queue[0].key = key
	q.queue[0].priority = priority

	// single bubble to make it emerge in the right position
	for i := 0; i < size-1; i++ {
		if q.queue[i].priority > q.queue[i+1].priority {
			q.queue[i], q.queue[i+1] = q.queue[i+1], q.queue[i]
		} else {
			break //saves some cycles
		}
	}
}

func (q *topNPriorityQueue) Print() {
	for i := size - 1; i >= 0; i-- {
		fmt.Printf("%7d %s\n", q.queue[i].priority, string(q.queue[i].key))
	}
}
