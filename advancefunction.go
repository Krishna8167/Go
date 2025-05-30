package main

/*
// function as data
For example :

func add(x, y int) int {
return x+y
}

func mul(x, y int) int {
return x*y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arthimetic func(int, int) int) int {
return arithematic(arithmetic(a,b), c)
}

func main() {
fmt.Println(aggregate(2,3,4, add))
// prints 9
fmt.Println(aggregate(2,3,4, mul))
// prints 24
}

// ORDER FUNCTIONS ?
dynamically creating functions and passing them around as variables adds unecessary complexity.
most of the time, this is right

however, when functions as values make a lot of sense. some of these includes
HTTP API HANDLERS
PUB/SUB HANDLERS
ONCLICK CALLBACKS

NORMAL functions are first class functions, but the above aggregator function is a higher-Order function.

// Currying
Function currying is the practice of writing a function that takes a function(or functions) as input
and returns a new function.

for ex:
func main() {
squareFunc := selfMath(multiply)
doubleFunc := selfMath(add)

fmt.Println(squareFunc(5))
// prints 25

fmt.Println(squareFunc(5))
// prints 10
}

func multiply(x, y int) int {
return x*y
}

func add(x, y int) int {
return x +y
}

func selMath(mathFunc func(int, int) int) func (int) int {
return func(x int) int {
return mathFunc(x, x)
}
}
// the defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically
just before its enclosing function returns

the deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding
function returns.

deferred function are typically used to close database connections, file handlers and the like.

// defer example -----(please give some)


//  */
