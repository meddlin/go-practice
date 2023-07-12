# Pointers

An explanation and some examples of working with pointers.

- Go docs: [https://go.dev/tour/moretypes/1](https://go.dev/tour/moretypes/1)
- Go by Example: [https://gobyexample.com/pointers](https://gobyexample.com/pointers)
- Code explanation: [https://www.educative.io/answers/what-are-pointers-in-golang](https://www.educative.io/answers/what-are-pointers-in-golang)
- Helpful StackOverflow answer: [https://stackoverflow.com/a/29020534/1670148](https://stackoverflow.com/a/29020534/1670148)

---

This answer from StackOverflow was the key for me to understand how to actually *get at the value*
of a freaking pointer. 

> "Just like in C, & means "address of" and assigns a pointer which is marked as *T. But * also means "value of", or dereferencing the pointer, and assigns the value.
> 
> This is a bit off-topic, but IMO this dual use of * is partly what gets newbies in C and friends confused about pointers"

[https://stackoverflow.com/a/29020534/1670148](https://stackoverflow.com/a/29020534/1670148)

---