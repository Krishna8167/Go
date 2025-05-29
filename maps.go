package main

/*
Maps are similar to javascript Objects, python dictionaries, and ruby hashes.
Maps are a data structure that provides key-.value mapping.

the zero value of a map is nil.

We can create a map by using a literal or by using the make() function

ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24

ages = map[string]int{
"John": 37,
"Mary": 21,
}

the len() function works on a map, it returns the total number of key/value pairs.

ages = map[string]int{
"John": 37,
"Mary": 21,
}
fmt.Println(len(ages)) // 2


// MUTATION IN MAPS

INSERT AND ELEMENT
m[key] = elem

GET AN ELEMENT
elem = m[key]

DELETE AN ELEMENT
delete(m, key)

CHECK IF A KEY EXISTS
elem, ok := m[key]

// NOTE ON PASSING MAP
like slices, maps are also passed by reference into functions. This means that when a map is passed into the function we write,
we can make changes to the original, we don't have a copy.

// nested maps

*/
