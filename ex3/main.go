package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readBasic() {
	/* Doing it this way, we pull the whole file into memory */
	dat, err := os.ReadFile(".\\data.csv")
	check(err)
	fmt.Print(string(dat))
}

func readBytes() {
	f, err := os.Open(".\\data.csv")
	check(err)

	// Let's read 5 bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Let's read 10 bytes
	// NOTE - this will read the next 10 bytes, from where we left off
	b2 := make([]byte, 10)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes: %s\n", n2, string(b2[:n2]))

	f.Close()
}

func readWithSeek() {
	f, err := os.Open(".\\data.csv")
	check(err)

	// o2, err := f.Seek(6, 0)
	o2, err := f.Seek(0, 0)
	check(err)
	b2 := make([]byte, 10)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	f.Close()
}

func main() {
	fmt.Println("EX3 - Files")

	f, err := os.Open(".\\data.csv")
	check(err)

	// This "rewinds" the file to the beginning
	// _, err = f.Seek(0, 0)
	// check(err)

	r4 := bufio.NewReader(f)
	size := 15
	b4, err := r4.Peek(size)
	check(err)
	fmt.Printf("%d bytes: %s\n", size, string(b4))

	f.Close()
}
