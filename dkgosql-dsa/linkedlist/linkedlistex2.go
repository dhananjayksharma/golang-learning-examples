package main

import "fmt"

type Node struct {
	Data int
	next *Node
}

type OrderList struct {
	Head   *Node
	Length int
}

// add ...
func (ol *OrderList) add(id int) {
	newNode := &Node{Data: id}
	if ol.Head == nil {
		ol.Head = newNode
		ol.Length = 1
		return
	}

	curr := ol.Head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
	ol.Length += 1
}

// append ...
func (ol *OrderList) append(pos, value int) {
	newNode := &Node{Data: value}
	// adding into head position
	if pos == 1 {
		// head[data:1975]
		curr := ol.Head
		ol.Head = newNode
		ol.Head.next = curr
		ol.Length += 1
		return
	}

	curr := ol.Head
	counter := 1
	prev := ol.Head
	for curr != nil {
		// head[ prev:[data:1965], <---> curr:[data:1970], [data:1975],
		if counter == pos {
			newNode.next = curr
			prev.next = newNode
			ol.Length += 1
		}
		prev = curr
		curr = curr.next
		counter++
	}

	// adding as tell if not matched any positions or not exist
	ol.add(value)
}

func (ol *OrderList) print() {
	curr := ol.Head
	for curr != nil {
		fmt.Printf("OrderList : %d\n", curr.Data)
		curr = curr.next
	}
}

// remove ...
func (ol *OrderList) remove(value int) {
	if ol.Head == nil {
		return
	}

	if ol.Head.Data == value {
		ol.Length = ol.Length - 1
		ol.Head = ol.Head.next
		return
	}

	curr := ol.Head
	for curr.next != nil && curr.next.Data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
	ol.Length -= 1
}

func main() {
	orderList := &OrderList{}
	orderList.add(1975)
	orderList.add(1980)
	orderList.add(1985)
	orderList.add(1995)
	orderList.add(2000)
	// fmt.Printf("length:%d\n", orderList.Length)
	// orderList.remove(1985)
	// orderList.remove(2000)
	// fmt.Printf("length:%d\n", orderList.Length)

	// Append at pos 1
	// orderList.append(1, 1970)
	// orderList.append(1, 1965)
	// orderList.append(8, 1990)
	// orderList.add(1995)
	// orderList.remove(1975)
	orderList.print()
}
