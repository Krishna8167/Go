/*
  Go Variables, Types, Constants, and Conditionals – Reference

   Data Types in Go:
  - Boolean:             bool
  - String:              string
  - Integers:            int, int8, int16, int32, int64
  - Unsigned Integers:   uint, uint8, uint16, uint32, uint64
  - Aliases:             byte (alias for uint8), rune (alias for int32, used for Unicode)
  - Floating Points:     float32, float64
  - Complex Numbers:     complex64, complex128 (rarely used)

   Constants:
  - Declared using the `const` keyword.
  - Cannot be changed after declaration.
  - Cannot use short declaration syntax `:=`.
  - Types supported: string, character, boolean, numeric.

   Conditionals:
  - No parentheses required around conditions in `if` statements.
  - An optional "initial" statement can be used to declare variables within the scope of the `if`.

    Example:
    if length := getLength(email); length < 1 {
        fmt.Println("Email is invalid")
    }
*/

package main

import "fmt"

func main() {
	// Declaring variables using 'var'
	var username string
	var smsSendingLimit float64
	var costPerSMS float64
	var hasPermission bool

	fmt.Printf("Limit: %f | Cost/SMS: %f | Permission: %v | Username: %q\n",
		smsSendingLimit, costPerSMS, hasPermission, username)

	// Declaring and initializing using shorthand ':='
	congrats := "Happy Birthday!"
	fmt.Println(congrats)

	penniesPerText := 3
	fmt.Printf("The type of 'penniesPerText' is %T\n", penniesPerText)

	// Declaring multiple variables in a single line
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your message"
	fmt.Println(averageOpenRate, displayMessage)

	// Type conversion
	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years")

	// Declaring constants
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("Plans:", premiumPlanName, "and", basicPlanName)

	// Constant expressions (must be resolvable at compile time)
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("Number of seconds in an hour:", secondsInHour)

	// Formatting strings using fmt.Sprintf
	const name = "Krishna"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f%%", name, openRate)
	fmt.Println(msg)

	// Example of conditional logic
	messageLen := 10
	maxMessageLen := 20

	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent.")
	} else {
		fmt.Println("Message not sent.")
	}
}
