package main

import (
	"fmt"
	"strconv"
)

func strToInt(s string) int {
	// var s string = "123"
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return i
}

func main() {
	fmt.Println("Topic - Type Conversions")

	// Let's convert a string to an int
	s := "123"
	i := strToInt(s)

	fmt.Printf("s (our string) is %s\n", s)
	fmt.Printf("i (our int) is %d\n", i)
}
