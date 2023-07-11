package main

import (
	"flag"
	"fmt"
)

/*
 * Docs - https://pkg.go.dev/flag
 */

var myFlag = flag.String("myFlag", "default", "myFlag description")

func main() {
	flag.Parse()

	fmt.Println("EX4 - Files")
	fmt.Println("myFlag: ", *myFlag)
}
