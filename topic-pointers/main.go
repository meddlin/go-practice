package main

import (
	"fmt"
)

func main() {
	/*
	 * The code example below is taken direclty from:
	 * https://www.educative.io/answers/what-are-pointers-in-golang
	 */

	// TODO - Add an exmaple that highlights the double (*) operator
	// from this StackOverflow answer:
	// https://stackoverflow.com/questions/29020523/how-to-cast-a-pointer-back-to-a-value-in-golang/29020534#29020534

	// a normal variable  whose address the pointer will store
	var intData = 20

	//declaration of a pointer
	var intPointer *int

	//intPointer now points towards intData
	intPointer = &intData

	fmt.Println("what intData stores:", intData)
	fmt.Println("address of intData:", &intData)
	fmt.Println("what intPointer stores:", intPointer)

	//this updates the value of intData using dereferncing operator
	*intPointer = 30

	fmt.Println("what intData now stores:", intData)
}
