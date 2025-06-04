package main

import (
	"fmt"
)

// ==================================
// ARRAYS AND SLICES IN GO
// ==================================

/*
1. ARRAYS

Arrays are fixed-size sequences of elements of the same type.
Syntax:
	var myInts [10]int

Array literals:
	primes := [6]int{2, 3, 5, 7, 11, 13}
*/

func arrayExample() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", a)
}

/*
2. SLICES

Slices are dynamically-sized and more commonly used than arrays.
A slice references an underlying array.

Syntax to slice an array:
	slice := arr[low:high] // low is inclusive, high is exclusive
	slice := arr[:3]       // first 3 elements
	slice := arr[2:]       // from index 2 to end
	slice := arr[:]        // full slice
*/

func sliceFromArray() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	slice := primes[1:4]
	fmt.Println("Slice from array:", slice) // Output: [3 5 7]
}

/*
Slices are references to arrays. Modifying a slice will affect the original array.
*/

func modifySlice() {
	arr := [4]int{10, 20, 30, 40}
	slice := arr[:2]
	slice[0] = 999
	fmt.Println("Modified array:", arr)
}

/*
3. MAKE FUNCTION

To create a slice without explicitly creating an array:
	slice := make([]int, length, capacity)
	slice := make([]int, 5) // capacity defaults to length
*/

func makeSlice() {
	s := make([]int, 3)
	fmt.Println("Slice from make:", s) // [0 0 0]
	fmt.Println("Length:", len(s), "Cap:", cap(s))
}

/*
4. SLICE LITERAL

Directly declare and initialize:
	s := []string{"Go", "is", "awesome"}
*/

func sliceLiteral() {
	mySlice := []string{"I", "love", "Go"}
	fmt.Println("Slice literal:", mySlice)
}

/*
5. LEN() AND CAP()

	len(slice) gives current number of elements
	cap(slice) gives capacity from starting index to end of underlying array
*/

func sliceLenCap() {
	s := []int{1, 2, 3}
	fmt.Println("Length:", len(s))   // 3
	fmt.Println("Capacity:", cap(s)) // 3
}

/*
6. VARIADIC FUNCTIONS

Functions that take a variable number of arguments.

	func myFunc(args ...int)
	The args are received as a slice.
*/

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/*
7. SPREAD OPERATOR

Used to pass a slice to a variadic function:
	fn(slice...)
*/

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

/*
8. APPEND FUNCTION

Adds elements to a slice.

	slice = append(slice, element)
	slice = append(slice, moreElements...)
*/

func appendExample() {
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	fmt.Println("After append:", s)

	t := []int{6, 7}
	s = append(s, t...)
	fmt.Println("After appending another slice:", s)
}

/*
9. SLICES OF SLICES

A slice can hold other slices (2D matrix):
	var matrix [][]int
*/

func matrixExample() {
	matrix := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println("2D Slice (matrix):", matrix)
}

/*
10. RANGE

A cleaner way to iterate over slices:
	for i, val := range slice
*/

func rangeExample() {
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}
}

// ===========================
// MAIN FUNCTION
// ===========================

func main() {
	fmt.Println("1. Array Example")
	arrayExample()

	fmt.Println("\n2. Slice From Array")
	sliceFromArray()

	fmt.Println("\n3. Modify Slice")
	modifySlice()

	fmt.Println("\n4. Make Slice")
	makeSlice()

	fmt.Println("\n5. Slice Literal")
	sliceLiteral()

	fmt.Println("\n6. Slice Length and Capacity")
	sliceLenCap()

	fmt.Println("\n7. Variadic Function Example")
	fmt.Println("Sum:", sum(1, 2, 3, 4))

	fmt.Println("\n8. Spread Operator Example")
	names := []string{"Alice", "Bob", "Charlie"}
	printNames(names...)

	fmt.Println("\n9. Append Example")
	appendExample()

	fmt.Println("\n10. Slices of Slices (Matrix)")
	matrixExample()

	fmt.Println("\n11. Range Loop Example")
	rangeExample()
}
