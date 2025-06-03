package main

/*
what is concurrency ?
concurrency is the ability to perform multiple tasks at the same time. Typically, our code is executed one
line at a time , one after the other.This is called sequential execution.

Synchronous Execution
------> ------>
task 1 task 2

Concurrent Execution
task 1
------>
------>
task 2

if the computer is we're running our code on has multiple cores, we can even execute multiple tasks at exactly the same time.
if we're running on a single core, a single code executes code at almost the same time by swutching between tasks very quickly.

either way, the code we write looks the same in Go and takes advantages of whatever resources are available.

how does concurrency works in Go?
Go was designed to be concurrent, which is a trait fairly unique to Go. It excels at performing many tasks simultaneously
safely using a simple syntax.

there isn't a popular programming language in existence where spawning concurrent execution is quite as elegant,

Concurrency is as simple as using a go keyword when calling a function:

go doSomething()

-> doSomething() will be executed concurrently with the rest of the function. the go keyword is used to spawn a new goroutine.

(use concurrency in a program..)

 Channels :
 are typed, thread-safe queue. channels allows different goroutines to communicate with each other.

 create a channel
 like maps and slices, channels must be created before use. they also use the same make keyword :

 ch := make(chan int)

 Send data to a Channel

 ch <- 69
 the <- operator is called the channel operator. Data flows in the direction of the arrow. this operator will
 block until another goroutine is ready to receive the value.

 Receive data from a channel :

 v := <-ch
 this reads and removes a value from the channels and saves it into the variable v. this operation will block until there is a value in the
 channel to be read.

 Blocking and Deadlocks :

 A deadlock is when a group of goroutines are all blocking so none of them can continue. this is common
 bug that you need to watch out for in concurrent programming.

 (example..)

 empty structs are often used as tokens in go programs. in this context, a token is unary value. in other words, we don't care
 what is passed through the channel. we care when and if it is passed.

 we can block and wait until something is sent on a channel using the following syntax.

 <-ch
 this will block until it pops a single item off the chsnnel, then continue, discarding the item.

 (Example..)

 BUffered Channels :
 channels are optionally be buffered.

 create a channel witha buffer
 we can provide a buffer length as soon as the second argument to make() to create a buffered channel:

 ch := make(chan int, 100)
 Sending on a buffered channel only blocks when the buffer is full.
 receiving blocks only when the buffer is empty.

 Closing Channels in go
 ch := make(chan int)
 // do some stuff with channel

 close(ch)

 checking if a channel is closed
 Similar to the ok value when accessing data ina map, receivers can check the ok value when receiving from a channel
 to test if a channel was closed.

 v, ok := <-ch
ok is false , if the channel is empty and closed.

Don't Send on a closed channel
Sending on a closed channel will cause a panic. A painc on the main goroutine will cause the entire program to crash,
and a painc in any other goroutine will cause that goroutine to crash.

closing isn't necessary. There's nothing wrong with leaving channel open, they'll still be gaebage collected if they're unused.
you should close channels to indicate explicity to a receiver that nothing else is going to come across.

Range
similar to slices and maps, channels can be ranged over.

for item := range ch {
// item is the next value received from the channel
}
will receive values over the channel(blocking at each iterationif nothing new is there) and will exit only when the channel is closed.

Select statement :
Sometimes we have a single goroutine listening to multiple channels and want to process
data in the order it comes through each channel.

A select statement used to listen to multiple channels at the same time. it is similar to a switch statement but for channels.

select {
case i, ok := <-chInts:
	fmt.Println(i)
case s, ok := <-chStrings:
	fmt.Println(s)
	}

Select Default Case :
the default case in a select statement executes immediately if no other channel has a value ready. Adefault case stops the
select statement from  blocking.

select {
case v := <-ch:
	// use v
default:
	//receiving from ch would block
	// so do something else
}

Tickers :
time.Tick() is astandard library function that returns a channel that sends a value on a given interval.
time.After() sends a value once after the duration has passed.
time.Sleep() blocks the current goroutine for the specified amount of time.

Read-Only channels :
A channels can be marked as read-only by casting it from a chan to a <-chan type. for example:

func main() {
ch := make(chan int)
readCh(ch)
}

func readCh(ch <-chan int) {
//ch can only be read from
// in this function
}

Write-Only channels :
the same goes for write-only channels, but the arrow's position moves.

func writeCh(ch <-chan int) {
// ch can  only write to
// this function
}

What happens when you read from a nil channel -> the receiver will block forever
What happens when you send to a closed channel -> panic


*/
