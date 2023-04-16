package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d started\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("worker %d ended\n", id)
}

func main() {
	print("normal")

	var wg sync.WaitGroup

	go print("goroutine")

	for i := 0; i <= 3; i++ {

		wg.Add(1)
		i := i

		go func() {
			worker(i)
			defer wg.Done()
		}()
	}

	wg.Wait()
}
