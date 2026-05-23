package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) InsertAtBeginning(num int) {
	// fmt.Printf("InsertAtBeginning LinkedList: %#v\n\n", l)
	newNode := &Node{value: num}

	if l.head == nil {
		l.head = newNode
		return
	}

	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) InsertAtEnd(num int) {
	// fmt.Printf("InsertAtEnd LinkedList: %#v\n\n", l)
	newNode := &Node{value: num}

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

func (l *LinkedList) Delete(num int) {
	if l.head == nil {
		return
	}

	if l.head.value == num {
		l.head = l.head.next
		return
	}
	temp := l.head
	for temp.next != nil && temp.next.value != num {
		temp = temp.next
	}

	if temp.next != nil {
		temp.next = temp.next.next
	}
}

func (l *LinkedList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.value, " -> ")
		temp = temp.next
	}

	fmt.Println("nil")
}

func main() {
	fmt.Println("Linkedlist start")

	list := &LinkedList{}

	list.InsertAtEnd(97)
	list.InsertAtBeginning(24)
	list.InsertAtBeginning(21)
	list.InsertAtEnd(96)
	list.InsertAtBeginning(11)
	// list.InsertAtBeginning(10)
	// list.InsertAtBeginning(5)
	// list.InsertAtEnd(20)
	list.Delete(96)
	// fmt.Printf("List LinkedList: %#v\n\n", list.head.value)
	list.Print()
	fmt.Println("Linkedlist end")
}
