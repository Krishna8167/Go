package main

import "fmt"

// =====================================
// LOOPS IN GO
// =====================================

/*
The basic loop in Go is the `for` loop, written in a C-like syntax:

	for INIT; CONDITION; AFTER {
		// body
	}
*/

func standardLoop() {
	fmt.Println("Standard Loop:")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("\n")
}

// =====================================
// BULK MESSAGE COST CALCULATION
// =====================================

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i <= numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i)) // increasing cost with each message
	}
	return totalCost
}

// =====================================
// OMITTING CONDITION (INFINITE UNTIL BREAK)
// =====================================

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

// =====================================
// WHILE-STYLE LOOP IN GO
// =====================================

func simulatePlantGrowth() {
	height := 1
	for height < 5 {
		fmt.Println("Still growing! Current height:", height)
		height++
	}
	fmt.Println("Plant has grown to", height, "inches.\n")
}

// =====================================
// FIZZBUZZ PROBLEM
// =====================================

func fizzBuzz() {
	fmt.Println("FizzBuzz:")
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
	fmt.Println()
}

// =====================================
// CONTINUE STATEMENT
// =====================================

func skipEvens() {
	fmt.Println("Odd numbers from 0 to 10:")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println(i)
	}
}

// =====================================
// MAIN FUNCTION
// =====================================

func main() {
	standardLoop()

	fmt.Println("Total Cost for 10 Messages:", bulkSend(10))
	fmt.Println("Max messages before reaching ₹15.0:", maxMessages(15.0), "\n")

	simulatePlantGrowth()
	fizzBuzz()
	skipEvens()
}
