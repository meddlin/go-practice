# Context

[https://go.dev/blog/context-and-structs](https://go.dev/blog/context-and-structs)

What is a context?


## Understanding goroutines, channels


### Goroutines 

“A goroutine is a lightweight thread of execution”. Goroutines are also lighter than 
a full thread. -- p.agnihotry.com, Ref 1





### Channels

Channels are used when you want to pass in results or errors or any other kind of information from one goroutine to another. (p.agnihotry.com, Ref 1) Channels also have types. You can have a channel specifically for `int` or errors.

Syntax

send something to a channel

```go
ch <- 1
```

receive something from a channel

```go
myVar := <- ch // receives from a channel and stores value in `myVar`
```

Create channels using the `make()` function.

```go
ch := make(chan int)
```


--- 

References:
- Ref 1: [https://p.agnihotry.com/post/understanding_the_context_package_in_golang/](https://p.agnihotry.com/post/understanding_the_context_package_in_golang/)