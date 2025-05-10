package main

import "fmt"

func main() {
	// 1. Basic for loop (similar to C, Java, JavaScript)
	// In Go, there is no 'while' keyword - we use 'for' instead
	fmt.Println("1. Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// 2. While-style loop in Go
	// Other languages: while (condition) { ... }
	// Go: for condition { ... }
	fmt.Println("\n2. While-style loop:")
	count := 0
	for count < 3 {
		fmt.Printf("count = %d\n", count)
		count++
	}

	// 3. Infinite loop
	// Other languages: while(true) or for(;;)
	// Go: for { ... }
	fmt.Println("\n3. Infinite loop (commented out to prevent actual infinite loop):")
	/*
		for {
			fmt.Println("This would run forever")
		}
	*/

	// 4. Range loop (similar to Python's for-in or JavaScript's for-of)
	// Iterating over arrays, slices, maps, or strings
	fmt.Println("\n4. Range loop over slice:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// 5. Range loop over map (similar to Python's dict.items())
	fmt.Println("\n5. Range loop over map:")
	person := map[string]string{
		"name": "John",
		"age":  "30",
	}
	for key, value := range person {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	// 6. Range loop over string (iterates over runes, not bytes)
	fmt.Println("\n6. Range loop over string:")
	for i, char := range "Hello" {
		fmt.Printf("index: %d, character: %c\n", i, char)
	}

	// 7. Skip index or value using blank identifier (_)
	fmt.Println("\n7. Range loop skipping index:")
	for _, value := range numbers {
		fmt.Printf("value: %d\n", value)
	}

	// 8. Break and continue (similar to other languages)
	fmt.Println("\n8. Break and continue example:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip iteration when i is 2
		}
		if i == 4 {
			break // Exit loop when i is 4
		}
		fmt.Printf("i = %d\n", i)
	}
}