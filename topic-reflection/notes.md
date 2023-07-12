# Reflection

- Go docs: [https://pkg.go.dev/reflect](https://pkg.go.dev/reflect)
- Explanation, "Golangbot": [https://golangbot.com/reflection/](https://golangbot.com/reflection/)

In Go (at least for me while I was debugging using Delve), there isn't a quick and easy "type" operator
available for direct use, like there is with JavaScript or Visual Studio's "Immediate Window" for C#. So,
we rely on reflection if we need to determine a variable's type at runtime.