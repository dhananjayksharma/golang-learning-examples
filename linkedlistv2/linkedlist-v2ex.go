package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtBeginning(input int) {
	newNode := &Node{data: input}

	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) InsertAtEnd(input int) {
	newNode := &Node{data: input}
	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (l *LinkedList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}

	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(5)
	list.InsertAtEnd(20)
	list.InsertAtEnd(9)
	list.InsertAtBeginning(15)

	list.Print() // 5 -> 10 -> 20 -> nil
}
