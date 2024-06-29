package main

import (
	"fmt"
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

func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
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
