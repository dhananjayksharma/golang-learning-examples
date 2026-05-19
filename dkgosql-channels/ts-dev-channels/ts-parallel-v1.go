package main

import (
	"fmt"
	"sync"
	"time"
)

var maxCapacity = 2

func doParallelTask(targets []string) []error {
	var maxCapParallel = make(chan struct{}, maxCapacity)
	var errList []error
	var wg sync.WaitGroup
	for i, v := range targets {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println("Task number: ", i)
			maxCapParallel <- struct{}{}
			defer func() { <-maxCapParallel }()
			err := execute(i, v)
			errList = append(errList, err)

		}()

	}
	wg.Wait()
	return errList
}

func execute(i int, task string) error {
	var err error
	if i%2 == 0 {
		err = fmt.Errorf("error even number find %d\n", i)
	} else {
		err = fmt.Errorf("error odd number find %d\n", i)
	}
	time.Sleep(2 * time.Second)
	// fmt.Println("Task is started:", task)

	fmt.Println("Task is done:", task)
	return err
}

func main() {
	fileList := []string{"file-8-apr-2026.json", "file-28-apr-2026.json", "file-20-apr-2026.json", "file-16-apr-2026.json", "file-12-apr-2026.json", "file-13-apr-2026.json", "file-14-apr-2026.json", "file-1-apr-2026.json", "file-11-apr-2026.json", "file-19-apr-2026.json"}
	errlist := doParallelTask(fileList)
	fmt.Println("error list:\n", errlist)
}
