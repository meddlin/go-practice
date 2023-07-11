# Using the io package

This example has a *slightly* more advanced way to open and read parts of a file, using a ReadSeeker
and making use of `defer` to close the file once we're done with it.

## Defer

"Defer is used to ensure that a function call is performed later in a program's execution..."

When `defer` is called, it sets the "deferred" function to be executed at the end of the enclosing
function.

- Ref: [https://gobyexample.com/defer](https://gobyexample.com/defer)

- Defer example: [https://stackoverflow.com/questions/20602131/io-writeseeker-and-io-readseeker-from-byte-or-file](https://stackoverflow.com/questions/20602131/io-writeseeker-and-io-readseeker-from-byte-or-file)


## io.ReadSeeker

ReadSeeker combines the Reader and Seeker interfaces, allowing a file (or stream) to be read
and searched at the same time.

Ref: [https://golang.hotexamples.com/examples/io/ReadSeeker/Seek/golang-readseeker-seek-method-examples.html](https://golang.hotexamples.com/examples/io/ReadSeeker/Seek/golang-readseeker-seek-method-examples.html)