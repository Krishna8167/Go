// variables for Go .

/*
bool

string

int int8 int16 int32 int 64 //whole numbers
uint uint8 uint16 uint32 uint64 //positive whole number

byte  (alais for uint8)
rune (alais for int32 {represented a unicode code point})

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)
*/

/*
constants
const keyword , and declared just by variables
constants can't use the ':=' for short declaration
constants can be character , string, boolean , or numeric
values.
as name implies , the value of constant can't be changes
after it has been declared.
*/

/*
Conditionals

if statement in go don't use parentheses arount the condition.
and preetymuch like any other programming language.

the initial statement of an if block

an if conditional can have an "initial" statement.
the variable(s) created in the initial statement
are only defined within the scope of the "if" body

if initial_statment; CONDITION {
}

Example:
length := getlength(email)
if length<1 {
fmt.Println("Email is invalid")
}

instead

if length := getlength(email) ;length < 1 {
fmt.Println("Email is invalid")
}
*/
package main

import "fmt"

/* Declaring a variable
func main() {
	var username string
	var smsSendingLimit float64
	var costperSMS float64
	var hasPermission bool

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costperSMS,
		hasPermission,
		username,
	)
}
*/
// short variable declaration
func main() {
	congrats := "happy birthday!" // declared a empty string
	fmt.Println(congrats)
	penniesPertext := 3
	fmt.Printf("the type of penniesPertext is %T\n", penniesPertext)
	// the above line meant for knowing the format .

	// declaring multiple values on the same line .
	averageOpenrates, displayMessage := .23, "is the average open rate of your message"
	fmt.Println(averageOpenrates, displayMessage)

	// some types can be converted the following ways
	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
	// constants
	const premiumPlanName = "Premium Plan"
	const BasicPlanName = "Basic Plan"

	fmt.Println("plans:", premiumPlanName)
	fmt.Println("plans:", BasicPlanName)
	// Constants can be computed so long as the compuataion can happen at compile time
	const secondInMinute = 60
	const minuteaInhour = 60
	const secondsInhour = secondInMinute * minuteaInhour

	fmt.Println("number of seconds in an hour: ", secondsInhour)

	//Formatting Strings in Go
	/* Go follows the printf tradition form the c language.
	fmt.Printf - prints a formatted string to standard out
	fmt.Sprintf - Returns the formatted string

	Examples :
	*/
	const name = "Krishna"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent",
		name, openRate)
	fmt.Println(msg)
	// some conditionals code .

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("message Sent")
	} else {
		fmt.Println("Message not sent")
	}
}
