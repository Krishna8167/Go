package main

import (
	"fmt"
)

// =====================================
// MAPS IN GO
// =====================================

/*
Maps are Go's built-in data structure for key-value storage.
They are similar to:
- JavaScript Objects
- Python Dictionaries
- Ruby Hashes

Declaration Syntax:

	ages := make(map[string]int) // using make
	ages["John"] = 37

	// using map literal
	ages := map[string]int{
		"John": 37,
		"Mary": 21,
	}
*/

// create and print map
func basicMapExample() {
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // overwrites previous value

	fmt.Println("Ages:", ages)
	fmt.Println("Total keys:", len(ages)) // Output: 2
}

// check existence and retrieve value
func checkKey() {
	ages := map[string]int{
		"John": 37,
		"Mary": 21,
	}

	// retrieve and check key existence
	age, ok := ages["John"]
	if ok {
		fmt.Println("John is", age, "years old.")
	} else {
		fmt.Println("John not found.")
	}

	// key that doesn't exist
	_, exists := ages["Alice"]
	fmt.Println("Is Alice present?", exists)
}

// delete key from map
func deleteKeyExample() {
	ages := map[string]int{
		"John": 37,
		"Mary": 21,
	}
	delete(ages, "Mary")
	fmt.Println("After deleting Mary:", ages)
}

// =====================================
// MAPS ARE PASSED BY REFERENCE
// =====================================

func updateMap(m map[string]string) {
	m["status"] = "active"
}

func referenceExample() {
	user := map[string]string{
		"name": "Alice",
	}
	updateMap(user)
	fmt.Println("Updated user map:", user)
}

// =====================================
// NESTED MAPS
// =====================================

func nestedMaps() {
	users := map[string]map[string]string{
		"alice": {
			"email": "alice@example.com",
			"role":  "admin",
		},
		"bob": {
			"email": "bob@example.com",
			"role":  "member",
		},
	}

	for username, info := range users {
		fmt.Println("Username:", username)
		for key, val := range info {
			fmt.Printf("   %s: %s\n", key, val)
		}
	}
}

// =====================================
// LOOPING OVER A MAP
// =====================================

func loopMap() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Carol": 92,
	}

	fmt.Println("Student Scores:")
	for name, score := range scores {
		fmt.Printf("%s scored %d\n", name, score)
	}
}

// =====================================
// MAIN
// =====================================

func main() {
	fmt.Println("=== Basic Map Example ===")
	basicMapExample()

	fmt.Println("\n=== Check Key Existence ===")
	checkKey()

	fmt.Println("\n=== Delete Key Example ===")
	deleteKeyExample()

	fmt.Println("\n=== Reference Behavior of Maps ===")
	referenceExample()

	fmt.Println("\n=== Nested Maps ===")
	nestedMaps()

	fmt.Println("\n=== Looping Over a Map ===")
	loopMap()
}
