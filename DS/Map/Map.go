package main

import (
	"fmt"
	"maps"
)

func main() {
	// init map
	mp := map[string]int{}

	// assigning value
	mp["A"] = 1
	mp["C"] = 2
	mp["E"] = 5

	fmt.Println("A", mp["A"])
	fmt.Println("E", mp["E"])

	// modifying value
	mp["A"]++

	fmt.Println("A", mp["A"])
	fmt.Println("C", mp["C"])

	mp["C"] = 21

	fmt.Println("C", mp["C"])

	// loop through map
	for key, value := range mp {
		fmt.Printf("mp[\"%s\"] = %d\n", key, value)
	}

	// get value with existance check
	v, ok := mp["A"]
	fmt.Println(v, ok)

	v, ok = mp["B"]
	fmt.Println(v, ok)

	// delete from map
	delete(mp, "A")

	v, ok = mp["A"]
	fmt.Println(v, ok)

	// emptying map
	clear(mp)
	v, ok = mp["C"]
	fmt.Println(v, ok)
	v, ok = mp["E"]
	fmt.Println(v, ok)

	// init map with make
	m := make(map[string]int, 1) // type, initial size
	m["Hello"] = 5
	m["World"] = 10

	fmt.Println(m)

	// comparing maps
	n := map[string]int{
		"Hello": 5,
		"World": 10,
	}

	fmt.Println(n)

	fmt.Println(maps.Equal(m, n))
}
