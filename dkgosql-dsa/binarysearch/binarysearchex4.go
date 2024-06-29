package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//	2
//
// Search
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}

var count int = 0

func main() {
	fmt.Println("main start")

	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(99)
	tree.Insert(150)
	tree.Insert(125)
	tree.Insert(75)
	tree.Insert(78)
	tree.Insert(66)
	tree.Insert(10)
	fmt.Println(tree)
	fmt.Println(tree.Search(150))
	fmt.Println("Search count:", count)
}
