/*
  Go vs Other Languages – Performance Overview

  1. Execution & Compilation Speed:
     - Go outperforms many interpreted languages such as JavaScript, Python, Ruby, and PHP in both execution and compilation speed.
     - Go code generally runs faster than interpreted languages and compiles significantly quicker than many compiled languages like C++ and Rust.

  2. Runtime Performance:
     - Go includes a lightweight runtime that handles memory management, garbage collection, and concurrency.
     - Although the Go runtime adds some overhead, the trade-off is improved safety and developer productivity.

  3. Comparison with Rust:
     - Rust still outperforms Go in terms of raw execution speed due to its zero-cost abstractions and lack of a garbage collector.
     - However, Go's simplicity and fast compile times make it a pragmatic choice for many backend and DevOps applications.

  4. Memory Management:
     - Go includes efficient garbage collection and runtime memory management, making it a great choice for long-running services.
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
