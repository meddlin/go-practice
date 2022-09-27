package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// csvReader()

	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n') // ReadString will block until the delimiter is entered
		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")
		fmt.Println("your input was: " + input)
		// fmt.Printf("length: %v", len(input))

		log.Printf("length: %v", len(input))

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		// our custom message
		// if strings.Compare("hi", input) == 0 {
		// 	fmt.Println("Hello!")
		// }
		if len(input) == 3 { // somehow
			fmt.Println("short message")
			break
		}
	}

	return
}

func csvReader() {
	// 1. open file
	recordFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	// 2. initialize reader
	reader := csv.NewReader(recordFile)

	// 3. read all records
	records, _ := reader.ReadAll()

	// 4. iterate through records
	for i := 0; i < len(records); i++ {
		// fmt.Println(records[i])
		fmt.Println(records[i][0])
		fmt.Println(records[i][1])
	}
	// fmt.Println(records)
}
