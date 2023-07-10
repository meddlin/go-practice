A collection of practicing exercises I walked through while learning Go.

## Contents

- Example
  - This is the first "exercise"/example. (It's the simplest.)
- ex2
  - 2nd example. Shows using/calling code from a separate module.

## Notes

For the exercises I'm walking through, it comes from here the...

Go Tutorial: [https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started)

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