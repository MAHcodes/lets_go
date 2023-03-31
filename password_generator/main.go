package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	digits := "0123456789"
	letters := "abcdefghijklmnopqrstuvwxyz"
	symbols := "*!@#$?"

	all := digits + letters + symbols

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var password string

	for i := 0; i < length; i++ {
		random_num := r.Intn(len(all))
		password += all[random_num : random_num+1]
	}
	return password
}

func main() {
	pass := generatePassword(10)
	fmt.Println(pass)
}
