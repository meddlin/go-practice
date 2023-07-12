package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exec 7 - Range")

	// We can loop over a slice of ints like this
	nums := []int{1, 2, 3, 4, 5}
	for idx, num := range nums {
		fmt.Printf("idx: %d -- num: %d\n", idx, num)
	}

	// Or loop over a key, value map like this
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// TODO - Write an example showing range over channels
	// Using this example: https://gobyexample.com/range-over-channels
}
