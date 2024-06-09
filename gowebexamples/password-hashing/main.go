package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"

	hash, _ := HashPassword(password)

	fmt.Printf("Password: %s\n", password)
	fmt.Printf("    Hash: %s\n", hash)

	match := CompareHashAndPassword(hash, password)

	fmt.Printf("   Match: %t\n", match)
}
