package main

import (
	"fmt"
	"reflect"
)

func main() {
	numberData := 42
	strData := "Hello, World!"

	fmt.Println("numberData is type: ", reflect.TypeOf(numberData))
	fmt.Println("strData is type: ", reflect.TypeOf(strData))
}
