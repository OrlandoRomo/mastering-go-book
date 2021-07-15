package main

import "fmt"

// Node struct
type Node struct {
	Next  *Node
	Value interface{}
}

// LinkedList struct
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// New creates a new instance of *LinkedList
func NewLinkedList() *LinkedList {
	emptynode := &Node{
		Next:  nil,
		Value: nil,
	}
	return &LinkedList{
		Head:   emptynode,
		Tail:   emptynode,
		Length: 0,
	}
}

// Append appends a new node at the end of the linked list
func (l *LinkedList) Append(value interface{}) *LinkedList {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
		l.Length++
		return l
	}

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length++
	return l
}

//Prepend appends a new node at the beginning of the linked list
func (l *LinkedList) Prepend(value interface{}) *LinkedList {

	if l.Head == nil {
		l.Append(value)
		return l
	}

	newNode := &Node{
		Value: value,
		Next:  l.Head,
	}
	l.Head = newNode
	l.Length++
	return l
}

//InsertAtIndex inserts a new node given an index.
//if index is greater than the length of the *LinkedList will be calle l.Append() function
func (l *LinkedList) InsertAtIndex(value interface{}, index int) *LinkedList {
	counter := 0
	pre := l.Head
	for counter < index-1 {
		if pre.Next == nil {
			l.Append(value)
			return l
		}
		pre = pre.Next
		counter++
	}
	newNode := &Node{
		Value: value,
		Next:  pre.Next,
	}
	pre.Next = newNode
	l.Length++
	return l
}

// RemoveHead removes the head of the linked list.
// If the linked list does not have a head it will just return the current pointer of *LinkedList
func (l *LinkedList) RemoveHead() *LinkedList {
	head := l.Head
	if head.Value == nil {
		return l
	}
	head = head.Next
	l.Head = head
	l.Length--
	return l
}

// RemoveTail removes the tail of the linked list.
// If the linked list does not have tail it will just return the current pointer of *LinkedList
func (l *LinkedList) RemoveTail() *LinkedList {
	counter := 0
	node := l.Head

	if l.Length == 0 {
		l.RemoveHead()
		l.Tail = nil
		return l
	}

	for counter < l.Length-1 {
		counter++
		node = node.Next
	}
	if node.Value == nil {
		return l
	}
	node.Next = nil
	l.Tail = node
	l.Length--
	return l
}

// RemoveAtIndex removes a node passing an index
// if index is greater than the length of linked list will delete the tail
func (l *LinkedList) RemoveAtIndex(index int) *LinkedList {
	counter := 0
	pre := l.Head
	for counter < index-1 {
		if pre.Next == nil {
			l.RemoveTail()
			return l
		}
		pre = pre.Next
		counter++
	}
	after := pre.Next
	pre.Next = after.Next
	l.Length--
	return l
}

// PrintAll prints all nodes of the linked list
func (l *LinkedList) PrintAll() {
	node := l.Head
	for {
		fmt.Printf("%v --> ", node.Value)
		if node.Next == nil {
			fmt.Println()
			return
		}
		node = node.Next
	}
}

// PrintReverse prints in revers single linked list
func (l *LinkedList) PrintReverse() {
	copy := *l
	first := copy.Head
	copy.Tail = copy.Head
	second := first.Next
	for second != nil {
		tmp := second.Next
		second.Next = first
		first = second
		second = tmp
	}
	copy.Head.Next = nil
	copy.Head = first
	copy.PrintAll()
}

func main() {
	linkedList := NewLinkedList()
	linkedList.Append(10)
	linkedList.Append(20)
	linkedList.Append(30)
	linkedList.Append(40)
	linkedList.PrintAll()
	linkedList.Prepend(50)
	linkedList.Prepend(60)
	linkedList.PrintAll()
	linkedList.InsertAtIndex(35, 7)
	linkedList.PrintAll()
	linkedList.InsertAtIndex(100, 0)
	linkedList.PrintAll()
	linkedList.InsertAtIndex(205, 45)
	linkedList.PrintAll()
	linkedList.RemoveAtIndex(3)
	linkedList.PrintAll()
	linkedList.RemoveHead()
	linkedList.PrintAll()
	linkedList.RemoveTail()
	linkedList.PrintAll()
	linkedList.RemoveAtIndex(54)
	linkedList.PrintAll()
	linkedList.InsertAtIndex(97, 5)
	linkedList.PrintAll()
	linkedList.Append(97)
	linkedList.PrintAll()
	linkedList.Append(90)
	linkedList.PrintAll()
	linkedList.RemoveAtIndex(2)
	linkedList.PrintAll()
	linkedList.PrintReverse()
}
