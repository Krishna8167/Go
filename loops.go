package main

/* the basic loop in go is written in standard c-like syntax:
for INITIAL; CONDITION, AFTER {
// DO SOMETHING..
 }

Ex :
FOR i=0; i< 10; i++ {
fmt.Println(i)
 }
// prints 0 through 9

Ex 2: Return cost of batch message send
func bulkSend(numMessage int) float64 {
	 totalCost := 0.0
	for i=0; i<=numMessage ;i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
}

// Omitting conditions from a for loop in GO
loops in go can omit section of a for loop.

for example :

for INITIAL; ; AFTER {                          // CONDITION OMITTED.
//DO SOMETHING..
}

EX 3:

func maxMessage(thresh float64) int {
totalCost := 0.0
for i = 0; ; i++ {
 totalCost += 1.0 + (0.01 * float64(i))
 if totalCost > thresh {
 return i
    }
 }
}

//There is no while loopmin go , most programming language have a concept of a while loop.because go allows for the
omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.

for CONDITION {
//DO SOME STUFF WHILE CONDITION IS TRUE
}

for example:

plantHeight := 1
for plantHeight < 5 {
fmt.Println("still growing! current height:", plantHeight)
plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")

//Go supports standard modulo operator:
  ->   7%3  //1

also logical AND , OR operator
  ->   true && false     // false
  ->   true || false     // true

  EX 4 :
  Complete the fizzbuzz function that prints the number 1 to 100 inclusive each on their
  own line, but substitute multiples of 3 for the text fizz and multiples of 5 for buzz.
  for multiples of 3 AND 5 print instead fizzbuzz.


import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}

// continue keyword stops the current iteration of a loop and continues to the next iteration.
continue is a powerful way to use the "guard clause" pattern within loops.

for i := 0; i < 10; i++ {
if i % 2 ==0 {
continue
}
}

*/
