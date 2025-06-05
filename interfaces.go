package main

import (
	"fmt"
	"math"
)

/*
📘 THEORY: INTERFACES IN GO

An **interface** in Go is a type that defines a set of method signatures. Any type that implements all the methods in an interface is said to "satisfy" that interface — *implicitly*, without using an `implements` keyword.

✅ Benefits:
- Polymorphism without inheritance
- Flexible APIs
- Decoupling dependencies

✅ Core Concepts:
- Implicit implementation
- Multiple interface implementation
- Type assertions
- Type switches
*/

// ===============================
// 1. Basic Interface: Shape
// ===============================
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func interfaceExample() {
	var s shape = rect{10, 5}
	fmt.Println("Area of rectangle:", s.area())
	fmt.Println("Perimeter of rectangle:", s.perimeter())

	s = circle{7}
	fmt.Println("Area of circle:", s.area())
	fmt.Println("Perimeter of circle:", s.perimeter())
}

// ===============================
// 2. Multiple Interfaces
// ===============================
type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.03
	}
	return float64(len(s.body)) * 0.01
}

func (e email) print() {
	fmt.Println("Email to:", e.toAddress, "Body:", e.body)
}

func (s sms) print() {
	fmt.Println("SMS to:", s.toPhoneNumber, "Body:", s.body)
}

func multiInterfaceExample() {
	e := email{true, "Hello there!", "test@example.com"}
	s := sms{false, "Verify your code", "+9188888888"}

	test(e, e)
	test(s, s)
}

func test(e expense, p printer) {
	fmt.Printf("Cost: $%.2f\n", e.cost())
	p.print()
}

// ===============================
// 3. Type Assertions
// ===============================
func getExpenseReport(e expense) (string, float64) {
	if em, ok := e.(email); ok {
		return em.toAddress, em.cost()
	}
	if s, ok := e.(sms); ok {
		return s.toPhoneNumber, s.cost()
	}
	return "", 0.0
}

// ===============================
// 4. Type Switch
// ===============================
func getExpenseReportSwitch(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

func printNumericValues(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Println("It's an int:", v)
	case string:
		fmt.Println("It's a string:", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// ===============================
// MAIN
// ===============================
func main() {
	fmt.Println("=== Interface Shape ===")
	interfaceExample()

	fmt.Println("\n=== Multiple Interfaces ===")
	multiInterfaceExample()

	fmt.Println("\n=== Type Assertion Report ===")
	addr, cost := getExpenseReport(email{false, "Testing assertion", "dev@example.com"})
	fmt.Println("Sent to:", addr, "Cost:", cost)

	fmt.Println("\n=== Type Switch Report ===")
	addr, cost = getExpenseReportSwitch(sms{false, "OTP Code", "+9177777777"})
	fmt.Println("Sent to:", addr, "Cost:", cost)

	fmt.Println("\n=== Type Switch on Values ===")
	printNumericValues(123)
	printNumericValues("abc")
	printNumericValues(true)
}
