package main

import (
	"fmt"
	"time"
)

func main() {
	var birthYear int

	fmt.Println("Welcome to the Age Calculator!")
	fmt.Print("Enter your birth year: ")
	fmt.Scanln(&birthYear)

	currentYear := time.Now().Year()
	age := currentYear - birthYear

	if age < 0 {
		fmt.Println("You entered a future year ")
	} else {
		fmt.Printf("You are approximately %d years old.\n", age)
		fmt.Println("Current year is:", currentYear)
	}
}
