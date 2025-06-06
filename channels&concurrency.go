package main

import (
	"fmt"
	"time"
)

/*
📘 THEORY: CONCURRENCY IN GO

Concurrency is the ability to execute multiple tasks seemingly at the same time. Go makes this easy and safe with:
- Goroutines
- Channels
- Select statements
- Tickers, timers
- Buffered & unbuffered communication
- Safe communication and synchronization

Goroutines are lightweight threads managed by the Go runtime.
*/

// ===============================
// 1. GOROUTINE BASICS
// ===============================

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func basicGoroutine() {
	go sayHello()                     // executed concurrently
	time.Sleep(time.Millisecond * 10) // Give goroutine time to complete
	fmt.Println("Main routine done")
}

// ===============================
// 2. CHANNEL BASICS
// ===============================

func channelExample() {
	ch := make(chan string)
	go func() {
		ch <- "Data sent from goroutine"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)
}

// ===============================
// 3. DEADLOCK DEMO
// ===============================

func deadlockExample() {
	// Uncommenting this will result in a deadlock
	// ch := make(chan int)
	// ch <- 1 // No receiver
	// fmt.Println(<-ch)
}

// ===============================
// 4. BUFFERED CHANNELS
// ===============================

func bufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// ===============================
// 5. CHANNEL CLOSE AND RANGE
// ===============================

func rangeOverChannel() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)

	for item := range ch {
		fmt.Println("Received:", item)
	}
}

// ===============================
// 6. SELECT STATEMENT
// ===============================

func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "one" }()
	go func() { ch2 <- "two" }()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	}
}

// ===============================
// 7. SELECT DEFAULT CASE
// ===============================

func selectWithDefault() {
	ch := make(chan int)
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received, skipping")
	}
}

// ===============================
// 8. TICKERS AND TIME
// ===============================

func tickerExample() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			fmt.Println("Ticker stopped")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}

// ===============================
// 9. READ-ONLY AND WRITE-ONLY CHANNELS
// ===============================

func readOnly(ch <-chan int) {
	fmt.Println("Reading:", <-ch)
}

func writeOnly(ch chan<- int) {
	ch <- 999
}

func readWriteDemo() {
	ch := make(chan int)
	go writeOnly(ch)
	readOnly(ch)
}

// ===============================
// MAIN
// ===============================

func main() {
	fmt.Println("\n=== Basic Goroutine ===")
	basicGoroutine()

	fmt.Println("\n=== Channel Example ===")
	channelExample()

	fmt.Println("\n=== Buffered Channel ===")
	bufferedChannel()

	fmt.Println("\n=== Range Over Channel ===")
	rangeOverChannel()

	fmt.Println("\n=== Select Statement ===")
	selectExample()

	fmt.Println("\n=== Select Default ===")
	selectWithDefault()

	fmt.Println("\n=== Ticker Example ===")
	tickerExample()

	fmt.Println("\n=== Read/Write Only Channels ===")
	readWriteDemo()
}

/*
📌 Summary:
- `go` keyword starts a goroutine.
- `chan` is used to send and receive data between goroutines.
- Channels can be buffered or unbuffered.
- Use `select` to handle multiple channel operations.
- Use `defer`, `close()`, and `range` for clean communication.
- Mark channels as read-only or write-only to enforce safe usage.
*/
