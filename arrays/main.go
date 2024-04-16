package main

import "fmt"

func main() {
	/* Arrays */
	// First we create two arrays and print them
	fmt.Println()
	fmt.Println("** printing an array of strings... **")
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Println()
	fmt.Println("** printing an array of ints... **")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/* Loop over an array */
	fmt.Println()
	fmt.Println("** looping over an array... **")

	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d", primes[i])
	}

	/* Loop with 'range' */
	fmt.Println()
	fmt.Println("** loop over array using 'range'... **")
	for _, num := range primes {
		fmt.Println(num)
	}
}
