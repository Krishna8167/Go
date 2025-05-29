// Structs in Go
/*
To represent structured data
often convenient to group different types of variables together.

Ex: represent a car we could do the following

type car struct {
Make string
Model string
height int
Width int
}

this creates a new struct type called car . all cars have a Make, Model,Height and Width

In Go, you will often use a struct to represent information that you would have used a dictionary
for  in python, or an object literal for in javascript.package startgo

*/

//Ex 1. programme for struct structure

package main

import "fmt"

type messageToSend struct {
	phoneNumber uint64
	message     string
}

func test(n messageToSend) {
	fmt.Printf("Sending message: '%s'to: %v\n", n.message, n.phoneNumber)
	fmt.Println("===============================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so exited to have you",
	})
}
