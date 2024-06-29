package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)

    // time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg *sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

        i := i
 
        go worker(wg, i)
     
    }

    wg.Wait()

}