package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("EX6 - Making use of io package")

	f2, err := os.Open(".\\data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer closeFile(f2) // <- this will be executed at the end of the enclosing function
	// DoSomething(f, f2)
	DoSomething(f2)
}

// func DoSomething(source io.ReadSeeker, destination io.WriteSeeker) {
// 	io.Copy(destination, source)
// }

func DoSomething(source io.Reader) {
	fmt.Print("Copying from source to destination...")
	io.Copy(os.Stdout, source)
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// it's good to check for errors when closing a file,
	// even though we deferred it earlier.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
