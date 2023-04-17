package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second)
		messages <- "1 sec"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		messages <- "2 sec"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case msg2 := <-messages:
			fmt.Println(msg2)
		}
	}
}
