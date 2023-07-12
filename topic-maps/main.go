package main

import (
	"fmt"
)

/*
 * Determine if a given string exists in a given map
 * NOTE - This only works for maps with the defined type (map[string]int)
 */
func existsInMap(lookup string, myMap map[string]int) bool {
	_, ok := myMap[lookup]
	return ok
}

func main() {
	fmt.Println("Topic - Maps")

	// Let's create a map of strings to ints
	// We'll use this to count the number of characters in a word
	// We'll use the word as the key and the count as the value
	wordCount := make(map[string]int) // We'll start with an empty map
	wordCount["hello"] = len("hello")
	wordCount["world"] = len("world")

	fmt.Println("is 'hello' in map? ", existsInMap("hello", wordCount))
	fmt.Println("is 'new string' in map? ", existsInMap("new string", wordCount))
}
