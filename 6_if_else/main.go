package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic if statement
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with initialization statement
	// This is a common Go idiom where you can initialize a variable in the if statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// If with multiple conditions
	temperature := 25
	humidity := 80
	if temperature > 30 && humidity > 70 {
		fmt.Println("It's hot and humid")
	} else if temperature > 30 {
		fmt.Println("It's hot but not humid")
	} else if humidity > 70 {
		fmt.Println("It's humid but not hot")
	} else {
		fmt.Println("Weather is pleasant")
	}

	// If with short circuit evaluation
	// Go uses short-circuit evaluation for logical operators
	name := "John"
	if len(name) > 0 && name[0] == 'J' {
		fmt.Println("Name starts with 'J'")
	}

	// If with error checking (common Go idiom)
	currentTime := time.Now()
	if hour := currentTime.Hour(); hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}

	// Nested if statements
	userAge := 25
	hasLicense := true
	if userAge >= 18 {
		if hasLicense {
			fmt.Println("You can drive")
		} else {
			fmt.Println("You need a license to drive")
		}
	} else {
		fmt.Println("You are too young to drive")
	}

	// If with type assertion (common in Go)
	var value interface{} = "Hello"
	if str, ok := value.(string); ok {
		fmt.Printf("Value is a string: %s\n", str)
	} else {
		fmt.Println("Value is not a string")
	}
}
