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

	DoSeekerStuff(f2)
	DoSomething(f2)
}

// func DoSomething(source io.ReadSeeker, destination io.WriteSeeker) {
// 	io.Copy(destination, source)
// }

func DoSomething(source io.Reader) {
	fmt.Print("Copying from source to destination...")
	io.Copy(os.Stdout, source)
}

/*
 *
 * ReadSeeker combines the Reader and Seeker interfaces
 * - https://golang.hotexamples.com/examples/io/ReadSeeker/Seek/golang-readseeker-seek-method-examples.html
 */
func DoSeekerStuff(source io.ReadSeeker) {
	source.Seek(5, io.SeekStart)
	io.Copy(os.Stdout, source)
}

/*
 * This function comes from here:
 * https://stackoverflow.com/questions/20602131/io-writeseeker-and-io-readseeker-from-byte-or-file
 */
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
