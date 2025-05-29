package main

/*
//Arrays are fixed-size groups of variables of the same types
the types [n]T is an array of type T

To declare a array of 20 integers:
var myInts [10]int
or to declare an initialized literal:
primes :=[6]int{2,3,5,7,11,13}

// Slices in Go ,
99 out of 100 times you will use a slice instead of an array when working with ordered lists.
Arrays are fixed in size. once you make an array like [10]int you can't add an 11th element

Slice are dynamically-sized, flexible view of the elements of an array.

Slices always have an underlying array, though it isn't always specified explicitly,
To explicitly create a slice on top of an array we can do :

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

Syntax is :(lower Index is inclusive and higher index is exclusive)

Arrayname[lowIndex:HighIndex]
Arrayname[lowIndex:]
Arrayname[:HighIndex]
Arrayname[:]

// Slices hold "references" to an underlying array, and if you assign one slice to another , both refer to the same array.
// slices references arrays
// A function that only has access to a slice can modify the underlying array


// Make -> most of the time we don't need to think about with the underlying array of the slice.
We can create a new slice using the make function:

mySlice := make([]int, 5, 10) // func make([]T, len, cap)

mySlice := make([]int, 5) // the capacity argument is usually and defaults to the length
Slices created with make will be filled with zero value of the type.
 if we  want to create a slice with a specific set of values, we can use a slice literal:

 mySlice := []string{"I","love", "go"}

 note that the array brackets do not have a 3 in them .If they did, you'd have an array instead of a slice.

 // the length of a slice is the number of elements in the underlying array, counting from the first element in the underlying array,
 counting from the first element in the slice. It is accessed using the build-in len() function:

 fmt.Println(len(mySlice))  // 3

 // the capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. it
 is accessed using the built-in cap() function // don't worry about it , if you are not hyper-otimising the code.

fmt.Println(cap(mySlice))  // 3

// len() and cap() , when slice is nil , it will return zero // this are safe functions

// Variadic Functions:
many functions especially those in standard library, can take an arbitrary number of final arguments .
this is accomplished by using "..." suntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func sum(nums ...int) int{
//nums is just a slice
for i := 0; i<len(nums); i++ {
num := nums[i]
}
}

func main() {
 totals := sum(1,2,3)
 fmt.Println(total)
 // prints "6"
 }


 the fimiliar fmt.Println() and fmt.Sprintf() are variadic.

 // spread Operator -> allows us to pass a slice into a variadic function. the Op consists
 of three dots following the slice in the function call.

 names := string{"bob","sue","alice"}
printStrings(names...)

func printStrings(strings ...string) {
for i := 0 ; i < lens(strings); i++ {
fmt.Println(strings[i])
}
}

//Append -> the built-in append function is used to dynamically add elements to a slice :

func append(slice []Type, elems ...Type) []Type

if the underlying array is not large enough, append() will create a new underlying array and point the slice to it.

notice that the append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)

// SLices of Slices
Slices can hold other slices, effectively creating a matrix, or a 2D Slice.

rows := [][]int{}


// Range
Go provides syntatic sugar to iterate easiky over elements of a slice:

for INDEX, ELEMENT := range SLICE {
}

ex:

fruits := []string{"apple","banana","grape"}
for i, fruit := range fruits {
 fmt.Println(i,fruit)
}
// 0 apple
// 1 banana
// 2 grape

*/
