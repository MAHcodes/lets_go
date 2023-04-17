package main

import (
	"fmt"
	"time"
)

func main() {
	const requestsCoutn = 5
	requests := make(chan int, requestsCoutn)

	for i := 0; i < requestsCoutn; i++ {
		requests <- i
	}
  close(requests)

	limiter := time.Tick(time.Second)

  for req := range requests {
    <-limiter
		fmt.Println(req)
	}
}
