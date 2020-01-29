## Purpose of this module
This module teaches how to use the format sub-command and the standalone program to format `Go` programs.
It also touches a bit on the design decisions for formatting of `Go` programs.

## Philosophy of Go regarding formatting
TODO

## Verify locally
To test this module locally:

- Open up a terminal at the project root
- Run command `cd test` to change to the test directory
- Run command `go test -run Module2` to run all tests for module 2, or 
- Run command `go test -v -run Module2` to run all tests for module 2 with verbose information 

## Task 1: Format `Go` source code using a standalone program
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

Open a file named `module2.txt` and write the `gofmt` command to format a source file named `common.go`.


## Task 2: Use `go fmt` with proper flags
`go fmt` takes following input:
- (`stdin`)[https://en.wikipedia.org/wiki/Standard_streams#Standard_input_(stdin)]: when without an explicit file or path
- file: when given a file
- all .go files under a directory: when given a directory

Run the command `go help fmt` and examine the provided information.

What command do we use to achieve the same result with that in Task 1?
This time we use `go fmt`.

Write the complete command (equivalent of the answer in Task 1), in the `module2.txt` file.



## Extra help
Here is the command to get more information on `gofmt`: 
- `go help fmt`

Here is the command to get more information on `go fmt`:
- `go doc cmd/gofmt`
