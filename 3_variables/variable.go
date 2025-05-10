package main

import "fmt"

func main() {
	// Variable Declaration in Go
	// There are three ways to declare variables in Go:

	// 1. Explicit type declaration
	// Syntax: var variableName type = value
	var var1 string = "hello golang 1"    // string type
	var intVar int = 42                   // integer type
	var floatVar float64 = 3.14           // float64 type
	var boolVar bool = true               // boolean type

	// 2. Type inference (Go automatically detects the type)
	// Syntax: var variableName = value
	var var2 = "hello golang 2"           // type inferred as string
	var num = 100                         // type inferred as int
	var pi = 3.14159                      // type inferred as float64

	// 3. Short variable declaration (inside functions only)
	// Syntax: variableName := value
	var3 := "hello golang 3"              // type inferred as string
	isValid := true                       // type inferred as bool

	// Common Go Data Types:
	// - string: for text
	// - int, int8, int16, int32, int64: for integers
	// - uint, uint8, uint16, uint32, uint64: for unsigned integers
	// - float32, float64: for floating-point numbers
	// - bool: for true/false values
	// - byte: alias for uint8
	// - rune: alias for int32 (represents Unicode code points)
	// - complex64, complex128: for complex numbers

	// Zero Values in Go:
	// - int: 0
	// - float: 0.0
	// - bool: false
	// - string: "" (empty string)
	// - pointer, function, interface, slice, channel, map: nil

	fmt.Println("String variables:", var1, var2, var3)
	fmt.Println("Numeric variables:", intVar, floatVar, num, pi)
	fmt.Println("Boolean variable:", boolVar, isValid)

	// Note: var1 string 
	// This explicit type declaration is useful when:
	// 1. You want to declare a variable but initialize it later
	// 2. You need to specify a type different from the default
	// 3. You want to make the code more readable and explicit
}