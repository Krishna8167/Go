/*
Interfaces are collection of mathod signatures .
A type "implements" an interface if it has all of the methods of the given interfaces defined on it

type shape interface  {
area() float64
perimeter() float64
}
// for rectangle shape
type rect struct {
width,height float64
}

type (r rect) area() float64 {
return r.width * r.height
}

type (r rect) perimeter() float64 {
return 2*r.width + 2*r.height
}

// for circle shape

type circle struct {
radius float64
}

func(c circle) area() float64 {
return math.Pi * c.radius * c.radius
}

func(c circle) perimeter() float64 {
return 2*math.Pi * c.radius
}

// Interfaces are implemented implicity

there is no "implements" keyword
you may add method to the type and in the process be unknowingly implementing
various interfaces, and that's okay.

//Multiple Interfaces

a type can implement any number of interfaces in GO .
for ex :, interfaces{}, is always implemented by every type because it has no requirements.

type expense interface {
cost() float64
}

type printer interface {
print()
}

type email struct {
isSubscibred bool
body string
}

func print(p,printer) {
p.print()
}

func test(e expense, p printer) {
fmt.Println("Printing with cost: $%.2f ...\n", e.cost())
p.print()
}


// Name your Interface Argument

type copier interface {
Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

// type assertions in Go

when working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value.
You can cast an interface to its underlying type using a type assertion.package startgo1

type shape interface {
area() float64
}
 type circle struct {
 radius float64
 }

 // "c" is the new circle cast from "s"
 // which is an instance of a shape.
 // "ok" is a bool that is true if s was a circle
 // or false if s isn't a circle
 c,ok := s.(circle)  ......................................................................................................
*/

/*
	 Package main

	 import ("fmt")

	 func getExpenseReport(e expense) (string, float64) {
			em , ok = e.(email)
			if ok {
				return em.toaddress, em.cost()
		}
		s, ok = e.(sms)
		if ok {
				return s.toaddress, s.cost()
		}
		return "",0.0
	 }

	 func (e email) cost() float64 {
		if !e.isSubscibed {
			return float64(len(e.body)) * .05
		}
		return float64(len(e.body)) * .01
	 }
	  func (s sms) cost() float64 {
		if !e.isSubscibed {
			return float64(len(s.body)) * .01
		}
		return float64(len(s.body)) * .03
	 }

	 type expense interface {
		cost() float64
	 }

	 type email struct {
		isSubscribed bool
		body string
		toaddress string
	 }

	 type sms struct {
		isSubscribed bool
		body string
		tophonenumber string
	 }

// Type Switches
a type switch is similar to a regular switch statement, but the cases specify types instead of values.

	func printNumericValues(num interface{}) {
		switch v := num.(type){
	      case int :
		   fmt.Printf("%T\n",v)
		  case string :
		   fmt.Printf("%T\n",v)
		  default :
		   fmt.Printf("%T\n",v)
		}
	}
		func main() {
		printNumericValues(1)  // int

		printNumericValues("1") // string

		printNumericValues(struct{}{}) // struct
	}

	// using switch case for the last program
	 func getExpenseReport(e expense) (string, float64) {
		switch v := e.(type){
		      case email:
		      return v.toaddress, v.cost()
			  case sms:
		      return v.to[honenumber], v.cost()
			  default:
		      return "", 0.0
			  }
	 }
*/
package main
