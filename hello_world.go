package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// $ go run hello-world.go
// hello world

// Sometimes we’ll want to build our programs into binaries. We can do this using go build.
// $ go build hello-world.go
// $ ls
// hello-world    hello-world.go

// We can then execute the built binary directly.
// $ ./hello-world
// hello world
