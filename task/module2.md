## Purpose of this module
This module teaches how to use the format sub-command and the standalone program to format `Go` programs.
It also touches a bit on the design decisions for formatting of `Go` programs.

## Go's approach regarding formatting
Consistent format makes reading the source easier, both as the source file and the diff in code reviews.

Here's an excerpt of [Formatting](https://golang.org/doc/effective_go.html#formatting) section in [Effective Go](https://golang.org/doc/effective_go.html):
```
With Go we take an unusual approach and let the machine take care of most formatting issues. 
The `gofmt` program and `go fmt` reads a Go program and emits the source in a standard style of
indentation and vertical alignment, retaining and if necessary reformatting comments.
```

## Verify locally
To test this module locally:

- Open up a terminal at the project root
- Run command `cd test` to change to the test directory
- Run command `go test -run Module2` to run all tests for module 2, or 
- Run command `go test -v -run Module2` to run all tests for module 2 with verbose information 


## Task 1: Use `go fmt` with proper flags
`go fmt` takes following input:
- file: when given a file
- all .go files under a directory: when given a directory
- [`stdin`](https://en.wikipedia.org/wiki/Standard_streams#Standard_input_(stdin)): when there is no explicit file or path

Run the command `go help fmt` and examine the provided information.

What command do we use to achieve the same result with that in Task 1?
This time we use `go fmt`.

Write the complete command to format the provided `module2_content.go` file.


What changes are made between the before and after version of `module2_content.go`?

Bonus questions:
- Are there changes that can be made by, but not made?
- What are some limitations of the `go fmt` command?


## Task 2: Format `Go` source code using a standalone program
Note that there are two tools when using formatting in `Go`:
- `gofmt` is a standalone program (typically available as `bin/gofmt`)
- `go fmt` calls `go` command with a sub-command `fmt`

We focus on the standalone program, `gofmt` in this task.

### source code file
`gofmt` formats Go programs. 
Usage is as follows:
```
    gofmt [flags] [path ...]
```
here, `path` can be a file or a directory.

What flags should be used to achieve the exact same effect as Task 1.

Write the `gofmt` command to format a source file named `module2_code.go`.



## Extra help
Here is the command to get more information on `gofmt`: 
- `go help fmt`

Here is the command to get more information on `go fmt`:
- `go doc cmd/gofmt`
