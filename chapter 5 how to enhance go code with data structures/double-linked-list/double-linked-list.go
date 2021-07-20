package main

import "fmt"

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		root = &Node{v, nil, nil}
		return 0
	}
	if v == t.Value {
		fmt.Println("node already exists!")
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, t, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> empty list")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> empty list")
		return
	}
	tmp := t
	for t != nil {
		tmp = t
		t = t.Next
	}
	for tmp.Previous != nil {
		fmt.Printf("%d -> ", tmp.Value)
		tmp = tmp.Previous
	}
	fmt.Printf("%d -> ", tmp.Value)
	fmt.Println()
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> empty list")
		return 0
	}
	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookUp(t *Node, v int) bool {
	if root == nil {
		return false
	}

	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookUp(t.Next, v)

}

func main() {
	fmt.Println()
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 2)
	traverse(root)
	addNode(root, 3)
	addNode(root, 4)
	addNode(root, 5)
	addNode(root, 6)
	addNode(root, 7)
	traverse(root)
	fmt.Println(size(root))
	addNode(root, 8)
	traverse(root)
	reverse(root)

}
