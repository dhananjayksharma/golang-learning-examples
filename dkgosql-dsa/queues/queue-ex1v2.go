package main

import "fmt"

type Queue struct {
	items []interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		items: make([]interface{}, size),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
	fmt.Println("Size len,cap=", len(q.items), cap(q.items))
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := NewQueue(3)

	queue.Enqueue("Apple")
	queue.Enqueue("Banana")
	queue.Enqueue("Cherry")

	for !queue.IsEmpty() {
		item := queue.Dequeue()
		fmt.Println(item)
	}
}
