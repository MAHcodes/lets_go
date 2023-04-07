package main

import (
	"fmt"
	"crypto/rand"
)

func generatePassword(length int) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	symbols := "*!@#$?"
	all := digits + lower + upper + symbols

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	for i, b := range bytes {
		bytes[i] = all[int(b)%len(all)]
	}

	return string(bytes)
}

func main() {
	pass := generatePassword(32)
	fmt.Println(pass)
}
