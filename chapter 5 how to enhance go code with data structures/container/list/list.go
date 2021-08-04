package main

import (
	"container/list"
	"fmt"
)

func printList(list *list.List) {
	for t := list.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
	for t := list.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	l := list.New()
	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	l.PushFront(3)
	l.InsertBefore(4, e1)
	l.InsertBefore(5, e2)
	l.Remove(e2)
	l.Remove(e2)
	l.InsertAfter(5, e2)
	l.PushBackList(l)

	printList(l)

	l.Init()
	for i := 0; i < 20; i++ {
		l.PushFront(i)
	}
	fmt.Println()
	printList(l)
}
