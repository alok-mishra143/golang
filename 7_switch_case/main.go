package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Basic switch statement
	fmt.Println("=== Basic Switch Example ===")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	case "Wednesday", "Thursday": // Multiple expressions in a case
		fmt.Println("It's mid-week!")
	default:
		fmt.Println("It's another day")
	}

	// Example 2: Switch with initialization
	fmt.Println("\n=== Switch with Initialization ===")
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday")
	}

	// Example 3: Switch without expression (like if-else)
	fmt.Println("\n=== Switch without Expression ===")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	// Example 4: Fallthrough example
	fmt.Println("\n=== Fallthrough Example ===")
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough // Continues to the next case
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Greater than three")
	}

	// Example 5: Type switch
	fmt.Println("\n=== Type Switch Example ===")
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %v\n", v)
	case string:
		fmt.Printf("String: %v\n", v)
	case bool:
		fmt.Printf("Boolean: %v\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	fmt.Println("\n=== Type Switch Example Type Assertion ===")


	// Example 6: Type Function
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("this is int")
		case bool:
			fmt.Println("this is bool")
		case float64:
			fmt.Println("this is float")
		case string:
			fmt.Println("this is string")
		}
	}

	whoAmI(5)
}

/*
Key points about switch statements in Go:

1. Basic Switch:
   - No need for 'break' statements (unlike C/C++)
   - Cases are evaluated from top to bottom
   - Only the first matching case is executed

2. Switch with Initialization:
   - Can include a simple statement before the expression
   - Scope of the initialized variable is limited to the switch block

3. Switch without Expression:
   - Works like if-else statements
   - Cases must be boolean expressions
   - Useful for complex conditions

4. Fallthrough:
   - Use 'fallthrough' to execute the next case
   - Unlike C/C++, it doesn't automatically fall through
   - Must be the last statement in a case

5. Type Switch:
   - Used to discover the type of an interface value
   - Syntax: switch v := i.(type) { ... }
   - Each case specifies a type

6. Multiple Cases:
   - Can have multiple expressions in a single case
   - Example: case "Monday", "Tuesday":

7. Default Case:
   - Optional but recommended
   - Handles unmatched cases
   - Can be placed anywhere in the switch
*/
