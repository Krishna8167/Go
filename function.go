/*
   Go Functions – Syntax, Concepts & Examples

  This file demonstrates function concepts in Go, useful for backend, cloud-native, and DevOps engineering.

  ------------------------------------
   Function Signatures:
    A function's signature defines the parameter types and return type.

    Syntax:
      func sub(x int, y int) int  // returns an int after accepting two int arguments

    If multiple parameters share a type:
      func add(x, y int) int

  ------------------------------------
   Go vs C Function Declaration Syntax:
    C:   int (*fp){int (*ff)(int x, int y), int b}
    Go:  func(func(int, int) int, int) int

    Go is significantly cleaner and more readable.

  ------------------------------------
   Pass-by-Value:
    Go passes arguments **by value** by default. Functions receive a **copy**, not a reference.

    Example:
*/

package main

import (
	"errors"
	"fmt"
)

// Demonstrating pass-by-value
func incrementSends(current, toAdd int) int {
	current += toAdd
	return current
}

// ------------------------------------
//  String Concatenation Function

func concat(s1, s2 string) string {
	return s1 + s2
}

// ------------------------------------
//  Ignoring Multiple Return Values

func getNames() (string, string) {
	return "Krishna", "Sharma"
}

// ------------------------------------
//  Named Return Values (Naked Return)

func getCoordinates() (x, y int) {
	// x and y are automatically initialized to 0
	return // Naked return returns x, y
}

// ------------------------------------
//  Early Return (Guard Clause)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, nil
}

// ------------------------------------
//  Main Function to Demonstrate All Concepts

func main() {
	// Pass-by-Value Example
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages.")

	// String Concatenation
	fmt.Println(concat("Lane", " Happy Birthday!"))
	fmt.Println(concat("Elon", " — Hope that Tesla thing works out!"))
	fmt.Println(concat("Go ", "is Fantastic!"))

	// Ignoring multiple return values
	firstName, _ := getNames()
	fmt.Println("Welcome to the party,", firstName)

	// Named Return Values
	x, y := getCoordinates()
	fmt.Printf("Coordinates: (%d, %d)\n", x, y)

	// Early Return / Error Handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// Attempt divide by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}
}
