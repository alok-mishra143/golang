package main

import "fmt"

func main() {
	// * Constants in Go
	// 1. Basic constant declaration
	const PI = 3.14159
	fmt.Println("PI:", PI)

	// 2. Multiple constants declaration
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
		Sunday    = 7
	)
	fmt.Println("Monday:", Monday)

	// 3. Using iota for sequential constants
	const (
		Red   = iota // 0
		Green        // 1
		Blue         // 2
	)
	fmt.Println("Colors:", Red, Green, Blue)

	// 4. Constants with type declaration
	const (
		MaxInt    int     = 9223372036854775807
		MinFloat  float64 = -1.797693134862315708145274237317043567981e+308
		ComplexNum complex128 = 1 + 2i
	)
	fmt.Println("MaxInt:", MaxInt)
	fmt.Println("MinFloat:", MinFloat)
	fmt.Println("ComplexNum:", ComplexNum)

	// 5. Constants with expressions
	const (
		SecondsPerMinute = 60
		MinutesPerHour   = 60
		SecondsPerHour   = SecondsPerMinute * MinutesPerHour
	)
	fmt.Println("Seconds in an hour:", SecondsPerHour)

	// Note: Constants must be initialized at declaration time
	// Note: Constants cannot be modified after declaration
	// Note: Constants can be declared at package level or inside functions
}