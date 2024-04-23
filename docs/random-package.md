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