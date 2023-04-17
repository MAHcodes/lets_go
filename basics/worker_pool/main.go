package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finish job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const jobsCount = 5
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	for i := 0; i < 5; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < jobsCount; i++ {
		<-results
	}
}
