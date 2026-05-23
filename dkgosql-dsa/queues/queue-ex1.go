package main

import (
	"fmt"
	"os"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	fmt.Println("Size:", len(q.items), cap(q.items))
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
func (q Queue) LenNew() {
	fmt.Println("Inside:", len(q.items))
}

func (q *Queue) AddNew(v int) {
	q.items = append(q.items, v)
}

func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("List item before:", queue.items)
	queue.AddNew(99)
	fmt.Println("List item after:", queue.items)

	queue.LenNew()

	os.Exit(0)
	fmt.Println("Size:", queue.Size())
	item, err := queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued item:", item)
	}
	item, err = queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued item:", item)
	}
	fmt.Println("Size:", queue.Size())
}
