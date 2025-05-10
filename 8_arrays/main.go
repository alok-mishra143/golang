package main

import (
	"fmt"
	"strings"
)

func main() {
	// Basic array declaration and initialization
	numbers := [5]int{10, 20, 30, 40, 50}
	
	// Different ways to print arrays
	fmt.Println("Method 1 - Direct print:", numbers)
	fmt.Printf("Method 2 - Using Printf: %v\n", numbers)
	fmt.Printf("Method 3 - Using Printf with type: %#v\n", numbers)
	fmt.Printf("Method 4 - Custom format: [%s]\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ", "), "[]"))

	// Array insertion methods
	// Method 1: Direct assignment (if you know the index)
	numbers[2] = 25  // Replaces element at index 2
	fmt.Println("\nAfter direct insertion:", numbers)

	// Method 2: Shift and insert (manual way)
	// To insert at index 1, we need to shift elements
	for i := len(numbers) - 1; i > 1; i-- {
		numbers[i] = numbers[i-1]
	}
	numbers[1] = 15
	fmt.Println("After shift and insert:", numbers)

	// Common array operations
	fmt.Println("\nArray Operations:")
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	
	// Array comparison
	numbers2 := [5]int{10, 15, 25, 30, 40}
	fmt.Println("Are arrays equal?", numbers == numbers2)

	// Array iteration methods
	fmt.Println("\nIteration Methods:")
	// Method 1: Using range
	fmt.Println("Using range:")
	for i, v := range numbers {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	// Method 2: Traditional for loop
	fmt.Println("\nUsing for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	// Method 3: Using range with index only
	fmt.Println("\nUsing range with index only:")
	for i := range numbers {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	// Array copying
	numbers3 := numbers  // Creates a copy
	numbers3[0] = 100
	fmt.Println("\nOriginal array:", numbers)
	fmt.Println("Copied array:", numbers3)
}
