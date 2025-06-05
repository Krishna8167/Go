package main

import (
	"fmt"
)

/*
📘 THEORY: FUNCTIONS IN GO

Go treats functions as *first-class citizens*. This means:
- You can assign functions to variables
- Pass them as arguments to other functions
- Return them from functions

Functions are *values* — they can be stored, manipulated, and passed around like data.
Go also supports powerful features like:
- Higher-order functions
- Currying (returning functions)
- Closures (functions that capture outer scope variables)
- Defer (delayed execution)
- Anonymous functions (functions without names)
*/

// ===============================
// 1. HIGHER-ORDER FUNCTIONS
// ===============================

/*
A higher-order function is one that either:
- Takes a function as an argument, or
- Returns a function as its result
*/

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// aggregate applies a given math function to three inputs
func aggregate(a, b, c int, operation func(int, int) int) int {
	return operation(operation(a, b), c)
}

func higherOrderExample() {
	fmt.Println("Sum:", aggregate(2, 3, 4, add))     // 9
	fmt.Println("Product:", aggregate(2, 3, 4, mul)) // 24
}

// ===============================
// 2. FUNCTION CURRYING
// ===============================

/*
Currying is the process of returning a function from a function,
often with some pre-bound arguments or logic.

Useful for:
- Configurable behavior
- Partial application
- Code reuse
*/

func multiply(x, y int) int {
	return x * y
}

func addAgain(x, y int) int {
	return x + y
}

// curry-style helper
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func curryingExample() {
	square := selfMath(multiply)
	double := selfMath(addAgain)

	fmt.Println("Square of 5:", square(5)) // 25
	fmt.Println("Double of 5:", double(5)) // 10
}

// ===============================
// 3. DEFER
// ===============================

/*
`defer` is used to schedule a function call to run *after* the current function returns.

Use-cases:
- Cleanup (e.g., close files, release locks)
- Tracing and logging
- Recovering from panics

Multiple `defer` calls run in LIFO (Last-In-First-Out) order.
*/

func deferExample() {
	defer fmt.Println("This prints last.")
	fmt.Println("This prints first.")
}

func multiDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("main logic")
}

// ===============================
// 4. CLOSURES
// ===============================

/*
A closure is a function that "remembers" and can manipulate variables from outside its own body.

Closures are powerful for:
- Encapsulating state
- Creating factories and counters
- Writing concurrent code
*/

func closureExample() {
	counter := 0
	increment := func() {
		counter++
		fmt.Println("Counter:", counter)
	}

	increment() // 1
	increment() // 2
}

// closure returning a function
func counterClosure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func closureReturningFunction() {
	next := counterClosure()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
}

// ===============================
// 5. ANONYMOUS FUNCTIONS
// ===============================

/*
Anonymous functions have no name.

Useful when:
- You need a quick function once
- You want to immediately invoke a function
- You're passing logic as a parameter
*/

func anonymousFunctionExample() {
	// Immediately Invoked Function Expression (IIFE)
	func(msg string) {
		fmt.Println("Hello,", msg)
	}("World")

	// Function assigned to a variable
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square of 6 is:", square(6))
}

// ===============================
// MAIN FUNCTION
// ===============================

func main() {
	fmt.Println("=== Higher-Order Functions ===")
	higherOrderExample()

	fmt.Println("\n=== Currying ===")
	curryingExample()

	fmt.Println("\n=== Defer Keyword ===")
	deferExample()

	fmt.Println("\n=== Multiple Defers ===")
	multiDefer()

	fmt.Println("\n=== Closures ===")
	closureExample()

	fmt.Println("\n=== Closure Returning Function ===")
	closureReturningFunction()

	fmt.Println("\n=== Anonymous Functions ===")
	anonymousFunctionExample()
}
