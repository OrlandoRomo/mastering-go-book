package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type Stack struct {
	Top    *Node
	Length uint
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Peek() *Node {
	// remove references
	top := new(Node)
	*top = *s.Top
	top.Next = nil
	return top
}

func (s *Stack) Push(value interface{}) *Stack {
	node := &Node{value, nil}
	if s.IsEmpty() {
		s.Top = node
		s.Length++
		return s
	}
	node.Next = s.Top
	s.Top = node
	s.Length++
	return s
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}
	top := s.Top
	next := top.Next
	top.Next = nil
	s.Top = next
	s.Length--
	return top
}

func (s *Stack) IsEmpty() bool {
	return s.Length == 0
}

func (s *Stack) Print() {
	node := s.Top
	for {
		if node == nil {
			break
		}
		fmt.Printf("--- %v ---\n", node.Value)
		node = node.Next
	}
	fmt.Printf("\nLength: %v\n", s.Length)
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	n := s.Pop()
	fmt.Printf("I removed %d\n", n.Value)
	s.Print()
	top := s.Peek()
	fmt.Printf("Top %v\n", top.Value)
	s.Print()
	s.Pop()
	s.Pop()
	s.Print()
	if s.IsEmpty() {
		fmt.Println("The stack is empty alv")
	}
}
