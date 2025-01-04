# Go Map Access Without Initialization

This repository demonstrates a common error in Go: accessing a map element before initializing the map.  This leads to a runtime panic.

## The Bug

The `bug.go` file contains code that attempts to assign a value to a map element without first initializing the map.  This results in a runtime panic because the program is trying to access a non-existent map.

## The Solution

The `bugSolution.go` file shows the corrected code.  The map is initialized using `make()` before any elements are accessed. This prevents the runtime panic.

## How to Reproduce

1. Clone the repository.
2. Navigate to the directory.
3. Run the command `go run bug.go` to see the panic.
4. Run the command `go run bugSolution.go` to see the corrected code running without a panic.