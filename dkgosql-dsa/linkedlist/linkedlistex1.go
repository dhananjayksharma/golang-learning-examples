package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) insert(data int) {
	newNode := Node{data: data, next: nil}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else if l.tail == nil {
		l.tail = &newNode
		l.head.next = l.tail
	} else {
		l.tail.next = &newNode
		l.tail = &newNode
	}
}

func (l *List) Print() {
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		fmt.Println(currentNode.data)
	}
}
func main() {
	linkedlist := &List{}
	linkedlist.insert(1)
	linkedlist.insert(10)
	linkedlist.insert(12)
	linkedlist.Print()
}
