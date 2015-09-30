// Based on the Go team's blog post:
// https://blog.golang.org/go-maps-in-action

package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")

	var m map[string]int

	// A map is a reference type, like pointers or slices.
	m = make(map[string]int)

	// Sets key "route" to the value of 66
	m["route"] = 66

	/* Retrieves the value stored uner the key "route" and assigns it to a new variable i.
	Nonexistant keys return the value's zero value.*/
	i := m["route"]

	// Removes an entry from the map, doesn't return anything and will do nothing if specified key doesn't exist.
	delete(m, "route")

	// Two value assignment tests for the existence of a key. `ok` is true if exists, false otherwise.
	i, ok := m["route"]

	// Maps can be iterated using the range keyword:
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Map literal
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
	}

	// Initialize an empty map:
	m = map[string]int{}
}
