package queue

import (
	"fmt"
	"strings"
)

// NODE
type DQueueNode struct {
	data int
	next *DQueueNode
}

func newDQueueNode(data int) *DQueueNode {
	return &DQueueNode{
		data: data,
		next: nil,
	}
}

func (n DQueueNode) String() string {
	return fmt.Sprintf("%d -> ", n.data)
}

// DYNAMIC QUEUE
type DQueue struct {
	first *DQueueNode
	last  *DQueueNode
}

// -- BASIC OPERATIONS
func (q DQueue) IsEmpty() bool {
	return q.first == nil && q.last == nil
}

func (q DQueue) Size() int {
	res := 0
	for tmp := q.first; tmp != nil; tmp = tmp.next {
		res++
	}
	return res
}

func (q DQueue) String() string {
	builder := strings.Builder{}
	builder.WriteString("first -> ")
	for tmp := q.first; tmp != nil; tmp = tmp.next {
		builder.WriteString(tmp.String())
	}
	builder.WriteString("last")
	return builder.String()
}

// -- UPDATE OPERATIONS
func (q *DQueue) Enqueue(val int) *DQueueNode {
	newNode := newDQueueNode(val)
	if q.last == nil {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	return newNode
}

func (q *DQueue) Dequeue() *DQueueNode {
	if q.first == nil { // empty queue
		return nil
	}
	if q.first == q.last { // queue with one element
		tmp := q.first
		q.first, q.last = nil, nil
		return tmp
	}

	dequeued := q.first
	q.first = q.first.next
	dequeued.next = nil
	return dequeued
}

// -- ACCESS OPERATIONS
func (q DQueue) First() *DQueueNode {
	return q.first
}

func (q DQueue) Last() *DQueueNode {
	return q.last
}
