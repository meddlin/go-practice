A collection of practicing exercises I walked through while learning Go.

## Contents

- Example
  - This is the first "exercise"/example. (It's the simplest.)
- ex2
  - 2nd example. Shows using/calling code from a separate module.

## Notes

### Tutorial Used for Exercises 

For the exercises I'm walking through, it comes from here the...

Go Tutorial: [https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started)

### Directories and Configuring Modules

Organizing work for this repo is deisgned to be "one exercise per directory". However, this means we need to configure this to be "one module per directory". This effectively means we have a single repo with multiple modules. *To work with multiple modules, we need to make use of a Go workspace.*
- [https://github.com/golang/tools/blob/master/gopls/doc/workspace.md](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)

Another note on multiple directories...
The `go.work` file needs to be updated for every new directory we add, for each of the modules.

When adding the `ex2` exercise, it required two subdirectories: `greetings` and `hello`. Simply adding the `ex2` directory to the `go.work` didn't allow everything to work together.

This wasn't enough: 

```
use (
	./ex2
	./example
)
```

It had to tbe this:

```
use (
	./ex2/greetings
	./ex2/hello
	./example
)
```


## Language Notes

### Slices

- Go slices: [https://go.dev/blog/slices-intro](https://go.dev/blog/slices-intro)

### Maps

Maps are Go's associative data type. They are commonly used to store key-value pairs.

Setting them up, they follow this pattern:

```go
make(map[key-type]val-type)

/*
 * A map where the key is a string, and the value is a string
 * -- similar to JSON -- would look like this.
*/
make(map[string]string)
```

- Resources
  - "Go by example": [https://gobyexample.com/maps](https://gobyexample.com/maps)
  - go.dev: [https://go.dev/blog/maps](https://go.dev/blog/maps)


### Seeding the random package

It's mentioned, but not well documented on the "random greeting" exercise, *you have to seed the random number generator* before it will return random numbers.

You can seed it in the `init()` function of a module like below. **And** every module has an inherit `init()` function that is called when the code executes. If we specify it, then we can setup variables, etc. before everything else runs--*a great place to set a randomizer seed*.

```go
func init() {
	rand.Seed(time.Now().UnixNano())
}
```

- Exercise: [https://go.dev/doc/tutorial/random-greeting](https://go.dev/doc/tutorial/random-greeting)
- StackOverflow post: [https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator](https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator)