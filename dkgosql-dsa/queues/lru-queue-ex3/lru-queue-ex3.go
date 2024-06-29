package main

import (
	"fmt"
)

type Queue struct {
	items    []int
	hitcache map[int]bool
}

func (q *Queue) Hit(item int) int {
	index := -1
	if q.IsEmpty() {
		return index
	}
	if _, ok := q.hitcache[item]; ok {
		index = q.Get(item)
	}
	return index
}
func (q *Queue) Get(item int) int {
	index := -1
	for i, v := range q.items {
		if item == v {
			return i
		}
	}
	return index
}
func (q *Queue) Enqueue(item int) {
	fmt.Printf("q.items:%v\n", q.items)
	queueLength := len(q.items)
	fmt.Printf("q.items length:%d\n", queueLength)
	hitPos := q.Hit(item)
	if hitPos == -1 {
		if queueLength == MAXCAP {
			fmt.Println("in deque")
			deletedItem, _ := q.Dequeue()
			delete(q.hitcache, deletedItem)
		}
		q.hitcache[item] = true
		q.items = append(q.items, item)
		fmt.Println("Size:", queueLength, cap(q.items))
	} else {
		fmt.Println("ELSE hit item, hitPos:", item, hitPos)
		leastItems := q.items[:hitPos]
		recentItems := q.items[hitPos+1 : queueLength]
		q.items = append(leastItems, recentItems...)
		q.Enqueue(item)
		fmt.Printf("leastItems:%v\n", leastItems)
		fmt.Printf("recentItems:%v\n", recentItems)
		fmt.Printf("q.items:%v\n", q.items)
		// if queueLength == MAXCAP {
		// 	fmt.Println("hit:", item)
		// 	deletedItem, _ := q.Dequeue()
		// 	delete(q.hitcache, deletedItem)
		// 	q.hitcache[item] = true
		// 	q.items = append(q.items, item)
		// }
	}
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
func (q *Queue) Display() {
	// for _, v := range q.items {
	// 	fmt.Printf("%d", v)
	// }
	fmt.Printf("Queue items:%v\n", q.items)
	fmt.Printf("Queue hitcache:%v\n", q.hitcache)
}

const MAXCAP = 4

func main() {
	queue := Queue{items: make([]int, 0, MAXCAP), hitcache: make(map[int]bool)}
	queue.Enqueue(10)
	queue.Enqueue(20)
	// queue.Enqueue(30)
	// queue.Enqueue(20)
	// queue.Enqueue(10)
	// queue.Enqueue(30)
	// queue.Enqueue(20)
	// queue.Enqueue(40)
	// queue.Enqueue(50)
	// queue.Enqueue(60)
	// queue.Enqueue(40)
	// queue.Enqueue(20)
	fmt.Printf("\n")
	queue.Display()
	fmt.Printf("\n")
}
