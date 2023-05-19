package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("password123")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err)
}
