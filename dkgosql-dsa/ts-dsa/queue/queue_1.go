package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue -> add element to end
func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

// Dequeue -> remove from front
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

// Peek -> front element
func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

// Size
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Queue:", q.items)

	val, _ := q.Dequeue()
	fmt.Println("Dequeued:", val)

	fmt.Println("Queue after dequeue:", q.items)
}