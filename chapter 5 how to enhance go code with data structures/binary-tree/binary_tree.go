package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

type Trunk struct {
	Previous *Trunk
	Str      string
}

func NewTrunk(prev *Trunk, str string) *Trunk {
	return &Trunk{
		Previous: prev,
		Str:      str,
	}
}

func ShowTrunks(p *Trunk) {
	if p == nil {
		return
	}
	ShowTrunks(p.Previous)
	fmt.Print(p.Str)
}

func NewTree(k int) *Tree {
	// var t *Tree
	// rand.Seed(time.Now().Unix())
	// for i := 0; i < 2*k; i++ {
	// 	tmp := rand.Intn(k * 2)
	// 	t = Insert(t, tmp)
	// }
	return &Tree{nil, k, nil}
}

func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}

func Traverse(t *Tree, prev *Trunk, isLeft bool) {
	if t == nil {
		return
	}
	prevStr := "    "
	trunk := NewTrunk(prev, prevStr)
	Traverse(t.Right, trunk, true)

	if prev == nil {
		trunk.Str = "-- "
	} else if isLeft {
		trunk.Str = " .--- "
		prevStr = "   |"
	} else {
		trunk.Str = " `-- "
		prev.Str = prevStr
	}
	ShowTrunks(trunk)
	fmt.Println(t.Value)
	if prev != nil {
		prev.Str = prevStr
	}
	trunk.Str = "   |"
	Traverse(t.Left, trunk, false)
}

func main() {
	t := NewTree(5)
	t = Insert(t, 5)
	t = Insert(t, 6)
	t = Insert(t, 4)
	t = Insert(t, 2)
	t = Insert(t, 1)
	t = Insert(t, 7)
	t = Insert(t, 8)
	t = Insert(t, 2)
	Traverse(t, nil, false)
}
