package main

import "fmt"

type Queue struct {
	items []string
}

func (q *Queue) Enqueue(i string) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() string {
	q.IsEmpty()

	deleted := q.items[0]
	q.items = q.items[1:]
	return deleted
}

func (q *Queue) IsEmpty() {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
	}
}

func (q *Queue) showQueue() {
	q.IsEmpty()
	for _, v := range q.items {
		fmt.Println("Queue:", v)
	}
}
func main() {
	q := Queue{}
	q.showQueue()
	q.Enqueue("apple")
	q.Enqueue("orange")
	q.showQueue()
	deleted := q.Dequeue()
	fmt.Println("Deleted:", deleted)
	deleted = q.Dequeue()
	fmt.Println("Deleted:", deleted)
	deleted = q.Dequeue()
	fmt.Println("Deleted:", deleted)
	q.showQueue()
}
