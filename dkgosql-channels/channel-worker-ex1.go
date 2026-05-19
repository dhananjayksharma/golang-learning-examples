package main

import (
	"fmt"
	"sync"
)

func workerPool(jobs chan string, isDoneChannel *sync.WaitGroup, jobId string) {
	defer isDoneChannel.Done()
	for v := range jobs {
		fmt.Printf("Job:#%s -> Jobs details: %s\n", jobId, v)
	}

}
func main() {
	fmt.Println("starting now")
	jobsChannel := make(chan string)
	var isDoneChannel sync.WaitGroup
	isDoneChannel.Add(5)

	go workerPool(jobsChannel, &isDoneChannel, "workerpool-four")
	go workerPool(jobsChannel, &isDoneChannel, "workerpool-five")
	go workerPool(jobsChannel, &isDoneChannel, "workerpool-one")
	go workerPool(jobsChannel, &isDoneChannel, "workerpool-two")
	go workerPool(jobsChannel, &isDoneChannel, "workerpool-three")
	jobAllocationCnt := 0
	for i := 1; i <= 2; i++ {
		jobsChannel <- fmt.Sprintf("job id: %d", i)
		jobAllocationCnt++
	}

	close(jobsChannel)
	isDoneChannel.Wait()
	fmt.Printf("Job Allocation done Count: %d\n", jobAllocationCnt)
	fmt.Println("ending... now")
}
