# Dynamic Array

The code for this structure is taken from this NeetCode 
problem: [https://neetcode.io/problems/dynamicArray](https://neetcode.io/problems/dynamicArray)

We need to create a "dynamic array" structure by hand.

## First Problem - Static arrays

Arrays in Go are static by design, and difficult to work with when trying to create *any* sort
of dynamically set length of an array.

If an array is:

```go
var arr = [3]int{1,2,3}
```

then, this is syntactically wrong:

```go
var arr = [n]int{}
```

So, we're going to solve this with slices.

Ref: [https://go.dev/blog/slices-intro](https://go.dev/blog/slices-intro)

The following code creates an array (ok, it's a slice with an underlying array), of `n-length`.

```go
arr := make([]int, n)
```

## First Problem - Pt. 2: Structs & Methods

Perhaps, we can create this without `slices`. 

If instead,

- we create a struct
- hold a length and capacity