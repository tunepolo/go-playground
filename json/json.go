package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person is a personal information
type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

func main() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "Golang test",
	}

	b, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
