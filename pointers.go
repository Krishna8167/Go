package main

import (
	"errors"
	"fmt"
)

// =====================================
// POINTERS IN GO
// =====================================

/*
A **variable** is a named location in memory that stores a value.

Example:
	x := 42 // x is a variable storing the value 42
*/

/*
A **pointer** is a variable that stores the *memory address* of another variable.

The `*` syntax is used to define a pointer type.

	var p *int // p is a pointer to an int

The `&` operator is used to get the address of a variable.

	x := 10
	ptr := &x // ptr stores the address of x
*/

func pointerBasics() {
	x := 100
	var ptr *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value stored in pointer ptr:", ptr)
	fmt.Println("Value pointed to by ptr:", *ptr) // Dereferencing
}

/*
Pointers allow us to directly modify the value stored at a memory location.
*/

func modifyValueThroughPointer(p *int) {
	*p = *p * 2
}

/*
Pointers can be `nil` (i.e., they point to nothing). Dereferencing a `nil` pointer causes a runtime panic.
So we must check before accessing it.
*/

func safeAccess(s *string) error {
	if s == nil {
		return errors.New("cannot access value: pointer is nil")
	}
	fmt.Println("Value:", *s)
	return nil
}

/*
POINTER RECEIVERS

Methods can use value or pointer receivers.

- Value Receiver: operates on a copy (doesn't affect the original)
- Pointer Receiver: operates on the original value (can modify it)
*/

type Counter struct {
	value int
}

// Value receiver (does NOT modify the original)
func (c Counter) IncrementByVal() {
	c.value++
}

// Pointer receiver (modifies the original)
func (c *Counter) IncrementByPtr() {
	c.value++
}

// This works even if you call pointer receiver on a value; Go handles it automatically

func receiverDemo() {
	c := Counter{value: 5}

	c.IncrementByVal()
	fmt.Println("After IncrementByVal:", c.value) // Output: 5 (unchanged)

	c.IncrementByPtr()
	fmt.Println("After IncrementByPtr:", c.value) // Output: 6
}

// ===========================
// MAIN FUNCTION
// ===========================

func main() {
	fmt.Println("1. Pointer Basics")
	pointerBasics()

	fmt.Println("\n2. Modify Through Pointer")
	num := 50
	modifyValueThroughPointer(&num)
	fmt.Println("Modified value:", num) // Output: 100

	fmt.Println("\n3. Nil Pointer Check")
	var str *string = nil
	err := safeAccess(str)
	if err != nil {
		fmt.Println("Error:", err)
	}

	msg := "hello"
	err = safeAccess(&msg)

	fmt.Println("\n4. Pointer Receivers vs Value Receivers")
	receiverDemo()
}
