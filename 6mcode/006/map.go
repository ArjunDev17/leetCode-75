package main

import "fmt"

// Simple Go map examples: creation, insertion, lookup, deletion, iteration.
func main() {
	// 1) Create a map using a literal
	ages := map[string]int{
		"Arjun": 30,
		"kabbu": 25,
	}

	fmt.Println("Initial map:", ages)

	// 2) Insert / update
	ages["kab"] = 28
	ages["kabbu"] = 26 // update existing key
	fmt.Println("After add/update:", ages)

	// 3) Lookup with "comma ok" idiom to check existence
	if age, ok := ages["Eve"]; ok {
		fmt.Printf("Eve is %d\n", age)
	} else {
		fmt.Println("Eve not found")
	}

	// 4) Delete a key
	delete(ages, "Arjun")
	fmt.Println("After delete:", ages)

	// 5) Iterate over map (order is unspecified)
	fmt.Println("Iterate:")
	for name, a := range ages {
		fmt.Printf("  %s -> %d\n", name, a)
	}

	// 6) Length of map
	fmt.Println("Length:", len(ages))

	// 7) nil vs made map
	var m map[string]int // nil map, cannot assign to it
	fmt.Println("nil map?", m == nil)
	// m["x"] = 1 // would panic: assignment to entry in nil map
	m = make(map[string]int)
	m["x"] = 1
	fmt.Println("made map:", m)

	// 8) Map with slice values (common pattern)
	groups := map[string][]string{}
	groups["admins"] = []string{"Arjun", "kab"}
	// appending to a slice value for a key
	groups["users"] = append(groups["users"], "kabbu")
	fmt.Println("Groups:", groups)
}
