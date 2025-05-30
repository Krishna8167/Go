package main

/*
a variable is a named memory location in memory that stores a value. we can manipulate the value of a variable by assigning a new
value to it or by performing operations on it.

We can manipulate the value of a variable by assigning a new value to it or by performing operations on it.
when we assign a value in a specific location in memory.

x := 42

// "x" is the name of the memory. that location is stor

// A pointer is a variable that stores the memory address of another variable. this means
that a pointer "points to" the location of where the data is stored NOT the actual data itself.

the  * syntax defines a pointer

var p *int

the & operator generates a pointer to its operand.

myString := "hello"
myStrungPtr = &myString

Why are pointers useful ?

pointers allows us to manipulate data in memory directly, without making copies or duplicating data. This can make programs
more efficient and allows us to do things that would be difficult or impossible without them.

// Nil pointers,
Pointer can be dangerous

if a pointer points to nothing (the zero value of the pointer type) then dereferncing it will cause a runtime error (a panic)
that crashes the program. Generally speaking, whenever you're dealing with pointers you should check it's nil before trying
to dereference it.

func removePanic(message *string) error {
if message == nil {
return errors.New("invalid")
}
}

// Pointer Receivers
a receiver type on a method can be a pointer.const
Method with pointer receivers can modify the value to which the receiver points. Since
methods often need to modify their, pointer receivers are more common than value receivers.

more widely used in go is a pointer receiver than a value receiver.

MEthods with pointer receivers son't requires that a pointer is used to call the method. the pointer
will automatically be derived from the value.

---(example)



*/
