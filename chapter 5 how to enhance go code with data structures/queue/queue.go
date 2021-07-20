package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type Queue struct {
	Head   *Node
	Rear   *Node
	Length uint
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Peek() *Node {
	if q.Length == 0 {
		return &Node{}
	}
	return q.Head
}

func (q *Queue) Enqueue(value interface{}) *Queue {
	node := &Node{value, nil}
	if q.Length == 0 {
		q.Head = node
		q.Rear = node
		q.Length++
		return q
	}

	last := q.Rear
	node.Next = last
	q.Rear = node
	q.Length++
	return q
}

func (q *Queue) Dequeue() *Queue {
	if q.Length == 0 {
		return q
	}
	last := q.Rear
	for last != nil {
		if last.Next == q.Head {
			last.Next = nil
			q.Head = last
			break
		}
		last = last.Next
	}
	q.Length--
	return q
}

func (q *Queue) Print() {
	node := q.Rear
	for {
		fmt.Printf("%v --> ", node.Value)
		if node.Next == nil {
			fmt.Println()
			return
		}
		node = node.Next
	}
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
}
