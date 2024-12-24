// channels are a way to communicate between goroutines. They allow goroutines to share data by sending and receiving values of a specific type.
//
//	Think of a channel as a pipeline through which data can flow from one goroutine to another.
package main

import (
	"fmt"
	"time"
)

// func main() {
// 	//unbuffered channel
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 42
// 	}()
// 	fmt.Println(<-ch)

// 	//buffered channel
// 	ch1 := make(chan int, 3)
// 	ch1 <- 42
// 	ch1 <- 43
// 	ch1 <- 44
// 	fmt.Println(<-ch1)

// }

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(1 * time.Second) // Simulating work
		results <- job * 2
	}
}

func main() {
	num_Jobs := 5
	jobs := make(chan int, num_Jobs)
	results := make(chan int, num_Jobs)

	//start workers
	for i := 1; i <= 5; i++ {
		go worker(i, jobs, results)
	}

	//send jobs
	for j := 1; j <= num_Jobs; j++ {
		jobs <- j
	}
	close(jobs)

	//get results
	for r := 1; r <= num_Jobs; r++ {
		fmt.Println("Results:", <-results)
	}
}
