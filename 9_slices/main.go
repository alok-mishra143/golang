package main

import (
	"fmt"
)

func main() {
	// 1. Creating Slices
	fmt.Println("=== Creating Slices ===")
	
	// Using make() to create a slice with length and capacity
	numbers := make([]int, 5, 10) // length=5, capacity=10
	fmt.Printf("numbers: %v, length: %d, capacity: %d\n", numbers, len(numbers), cap(numbers))

	// Slice literal
	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("fruits: %v, length: %d, capacity: %d\n", fruits, len(fruits), cap(fruits))

	// Empty slice
	var emptySlice []int
	fmt.Printf("emptySlice: %v, is nil: %v\n", emptySlice, emptySlice == nil)

	// 2. Basic Slice Operations
	fmt.Println("\n=== Basic Slice Operations ===")
	
	// Accessing elements
	fmt.Printf("First fruit: %s\n", fruits[0])
	fmt.Printf("Last fruit: %s\n", fruits[len(fruits)-1])

	// Slicing
	fmt.Printf("First two fruits: %v\n", fruits[:2])
	fmt.Printf("Last two fruits: %v\n", fruits[1:])

	// 3. Slice Manipulation
	fmt.Println("\n=== Slice Manipulation ===")
	
	// Appending elements
	fruits = append(fruits, "grape")
	fmt.Printf("After append: %v\n", fruits)

	// Appending multiple elements
	fruits = append(fruits, "mango", "kiwi")
	fmt.Printf("After multiple append: %v\n", fruits)

	// Appending one slice to another
	moreFruits := []string{"pear", "plum"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("After appending slice: %v\n", fruits)

	// 4. Common Slice Patterns
	fmt.Println("\n=== Common Slice Patterns ===")
	
	// Copying slices
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)
	fmt.Printf("Copied numbers: %v\n", numbersCopy)

	// Filtering elements
	var evenNumbers []int
	for _, num := range []int{1, 2, 3, 4, 5, 6} {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	// 5. Slice Internals and Best Practices
	fmt.Println("\n=== Slice Internals and Best Practices ===")
	
	// Demonstrating slice capacity growth
	var growingSlice []int
	for i := 0; i < 10; i++ {
		growingSlice = append(growingSlice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(growingSlice), cap(growingSlice))
	}

	// Pre-allocating slice capacity for better performance
	preAllocated := make([]int, 0, 1000)
	fmt.Printf("Pre-allocated slice - Length: %d, Capacity: %d\n", len(preAllocated), cap(preAllocated))

	// Demonstrating slice sharing
	fmt.Println("\n=== Slice Sharing ===")
	original := []int{1, 2, 3, 4, 5}
	shared := original[1:4]
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Shared slice: %v\n", shared)
	
	// Modifying shared slice affects original
	shared[0] = 99
	fmt.Printf("After modifying shared slice - Original: %v\n", original)
	fmt.Printf("After modifying shared slice - Shared: %v\n", shared)
}

/*
Key Points about Slices in Go:

1. Slices are dynamic arrays that can grow and shrink
2. A slice has three components:
   - Pointer to the underlying array
   - Length (number of elements)
   - Capacity (maximum number of elements)

3. Important slice operations:
   - make() - creates a slice with specified length and capacity
   - append() - adds elements to a slice
   - copy() - copies elements between slices
   - len() - returns the length
   - cap() - returns the capacity

4. Best Practices:
   - Pre-allocate capacity when you know the size
   - Be careful with slice sharing (slicing creates views of the same array)
   - Use append() for dynamic growth
   - Consider using make() when you need specific capacity

5. Common Gotchas:
   - Slicing creates a view of the original array
   - Modifying a shared slice affects the original
   - nil slices are valid and have length 0
   - Empty slices ([]int{}) are not nil
*/
