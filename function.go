// Functions.
// func sub(x int,y int) int     -> is a function signature .
package main

//import "fmt"

/*  function to concatinate two string
func concat(s1 string, s2 string) string {
	return s1 + s2
}
func main() {
	fmt.Println(concat("Lane", " happey birthday!"))
	fmt.Println(concat("Elon", " hope that tesla think works out"))
	fmt.Println(concat("Go ", "is Fantastic!"))
}
*/

// if you have multiple paramater you can also write as
/*   func add(x, y int) int       instead of ->           func(x int, y int) int
if they are not in order they need to be defined separately.
*/

// Declaration Syntax
/* it is less confusing than c

int (*fp){int (*ff)(int x, int y), int b}   -> c

func(func(int, int) int , int) int            -> Go
*/

/* Variables in Go are passed by value (except for a few data types).
   "Pass by value" means that when a variable is passed into the function, that function
   receives a copy of the variable . the function is unable to mutate the caller's data.
*/
/*   pass by value example -> when resign the functional return to the variable it gave the desired result.
func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}
*/

/*  ignoring multiple return values

A function can return a value that the caller doesn't care about .
We can explicity ignore variables by using an underscore:_

func getPoint() (x int, y int) int {
return 3,4
}

ignore y value
x, _ := getPoint()

*/
/* Example for ignoring variables .
func main() {
	firstName, _ := getnames()
	fmt.Println("Welcome to party,", firstName)
}

func getnames() (string, string) {
	return "Krishna", "Sharma"
}

*/

/* Named Return Values 

A return statement without the arguments returns the named return values.
this is known as a "naked" return . Naked return statements should be used only in short functions. 
they can harm readability in longer functions. 

func getcoords() (x, y int) {
// x and y are initialized with zero values 

return // automatically returns x and y
}

is same as 

func getcoords() (int, int) {
var x int
var y int
return x, y
}
*/

/*Early return Statements.

Go supports early return from a function, this a powerful feature that can clean up code, especially
when used as guard clauses.
Guard clauses levearges the ability to return early from a function(or continue theough a loop)
to make nested conditionals one-dimensional.
instead of using ifelse chains, we return early from the function at the end of each conditional block. 

func divide(dividend, divisor int) (int, error} {
if divisor == 0 {
     return 0, errors.New("can't divide by zero")
   }
return dividend/divisor, nil
}