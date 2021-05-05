package main

import "fmt"

type Coordinate struct {
	x, y int
}

func main() {
	// 19. Maps
	// A map maps a key to a value; just like objects in Javascript
	// The zero value of a map is nil; it has no keys, neither any key can be added
	// the make function returns a map for given type initialized and ready to use
	vertexMap := make(map[string]Coordinate) // a map with key of type string and value of type Coordinate
	vertexMap["P1"] = Coordinate{1, 2}
	vertexMap["P2"] = Coordinate{1, 2}
	vertexMap["P3"] = Coordinate{2, 1}

	if vertexMap["P1"] == vertexMap["P2"] {
		fmt.Println("Both P1 and P2 are same coordinates")
	}

	if vertexMap["P1"] != vertexMap["P3"] {
		fmt.Println("P1 and P3 are different coordinates")
	}

	// 20. Map literals
	// map literals are like struct literals, but the keys are required
	var verticesMapLiterals = map[string]Coordinate{
		"P1": Coordinate{3, 1}, // 21. type Coordinate at this level is optional as map defines the type Coordinate at top level already, example shown below
		"P2": {1, 3},
		"P3": {3, 1},
	}
	fmt.Println()
	if verticesMapLiterals["P1"] == verticesMapLiterals["P2"] {
		fmt.Println("Both P1 and P2 are same coordinates")
	} else {
		fmt.Println("P1 and P2 are different coordinates")
	}

	if verticesMapLiterals["P1"] != verticesMapLiterals["P3"] {
		fmt.Println("P1 and P3 are different coordinates")
	} else {
		fmt.Println("Both P1 and P3 are same coordinates")
	}

	// 22. Mutating Maps
	var randomMap = make(map[int]string)
	randomMap[1] = "Amit" // insert or update
	randomMap[2] = "delete me"

	fmt.Println()
	fmt.Println("Random map : ", randomMap)

	name := randomMap[1] // retrieve
	fmt.Println("Value for key 1 in random map : ", name)

	delete(randomMap, 2) // delete element from map with delete(map,key)
	fmt.Println("Random map after deleting element for key 2 : ", randomMap)

	value, ok := randomMap[5]                 // test the presence of some key, if ok is true value corresponds to the value assigned for given key else default zero value is returned
	fmt.Println("Value and ok = ", value, ok) // zero value for string is ""
}
