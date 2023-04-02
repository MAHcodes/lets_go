package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func main() {
	min := 1
	max := 10

	fmt.Printf("Guess a number between %d and %d: ", min, max)

	reader := bufio.NewReader(os.Stdin)

	number := generateRandomNumber(min, max)

	var guess int

	for guess != number {
		inputNumber, err := reader.ReadString('\n')
		inputNumber = strings.TrimSpace(inputNumber)
		guess, err = strconv.Atoi(inputNumber)
		if err != nil {
			log.Fatalln(err)
		}
		if guess < number {
			fmt.Printf("%d is too low, try again: ", guess)
		} else if guess > number {
			fmt.Printf("%d in too high, try again: ", guess)
		} else {
			fmt.Printf("Correct! %d is the number\n", guess)
		}
	}
}
