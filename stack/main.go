package main

import "fmt"

type Stack []string

// check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// push a new value on the stack
func (s *Stack) Push(str string) {
	// Appends the new value to the end of the Stack
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // get index of top-most element
		element := (*s)[index] // index into slice and get the element
		*s = (*s)[:index]      // remove element, by slicing it off
		return element, true
	}
}

func main() {
	var stack Stack

	stack.Push("this")
	stack.Push("is")
	stack.Push("a")
	stack.Push("...stack")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
