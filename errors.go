package main

import (
	"errors"
	"fmt"
	"strconv"
)

// =====================
// ERROR HANDLING IN GO
// =====================

/*
In Go, errors are represented by the built-in `error` type. It is an interface:

	type error interface {
		Error() string
	}

Functions that might fail typically return an `error` as their last return value.
If the error is `nil`, the operation was successful; otherwise, it failed.
*/

// =======================
// Standard Error Example
// =======================

func convertStringToInt(str string) {
	i, err := strconv.Atoi(str) // Atoi converts string to int
	if err != nil {
		fmt.Println("Couldn't convert:", err)
		return
	}
	fmt.Println("Conversion successful:", i)
}

// =========================
// Custom Error Definitions
// =========================

// Custom error type implementing the error interface
type userError struct {
	name string
}

// Error method makes userError satisfy the error interface
func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

// Dummy function to simulate whether user can receive SMS
func canSendToUser(userName string) bool {
	return userName != "bannedUser"
}

// Function demonstrating return of a custom error
func sendSMS(msg, userName string) error {
	if !canSendToUser(userName) {
		return userError{name: userName}
	}
	fmt.Println("SMS sent to", userName, ":", msg)
	return nil
}

// ===========================
// Using errors.New Example
// ===========================

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("division by zero is not allowed")
	}
	return x / y, nil
}

// ===========================
// MAIN FUNCTION
// ===========================

func main() {
	fmt.Println("1. Standard Error Example:")
	convertStringToInt("42b") // invalid input to trigger error
	fmt.Println()

	fmt.Println("2. Custom Error Example:")
	err := sendSMS("Hello!", "bannedUser")
	if err != nil {
		fmt.Println("SMS error:", err)
	}
	fmt.Println()

	fmt.Println("3. errors.New Example:")
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}
